package app

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/api/access"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/api/login"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/api/note"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db/pg"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db/transaction"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/closer"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/config"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/rate_limiter"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	"github.com/sony/gobreaker"

	accessRepository "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/access"
	loginRepository "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/login"
	noteRepository "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
	accessService "github.com/atlasir0/Chat_service/Auth_chat/internal/service/access"
	loginService "github.com/atlasir0/Chat_service/Auth_chat/internal/service/login"
	noteService "github.com/atlasir0/Chat_service/Auth_chat/internal/service/note"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	grpcConfig       config.GRPCConfig
	httpConfig       config.HTTPConfig
	swaggerConfig    config.SwaggerConfig
	loggerConfig     config.LoggerConfig
	prometheusConfig config.PrometheusConfig
	rateLimitConfig  config.RateLimitConfig
	breakerConfig    config.BreakerConfig

	dbClient         db.Client
	txManager        db.TxManager
	noteRepository   repository.UserRepository
	loginRepository  repository.LoginRepository
	accessRepository repository.AccessRepository

	noteService   service.UserService
	loginService  service.LoginService
	accessService service.AccessService

	noteImpl   *note.Implementation
	loginImpl  *login.Implementation
	accessImpl *access.Implementation

	rateLimiter    *rate_limiter.TokenBucketLimiter
	circuitBreaker *gobreaker.CircuitBreaker
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetRateLimitConfig() config.RateLimitConfig {
	if s.rateLimitConfig == nil {
		cfg, err := config.NewRateLimitConfig()
		if err != nil {
			log.Fatalf("failed to get rate limit config: %v", err)
		}

		s.rateLimitConfig = cfg
	}

	return s.rateLimitConfig
}
func (s *serviceProvider) GetRateLimiter(ctx context.Context) *rate_limiter.TokenBucketLimiter {
	if s.rateLimiter == nil {
		s.rateLimiter = rate_limiter.NewTokenBucketLimiter(
			ctx,
			s.GetRateLimitConfig().Limit(),
			s.GetRateLimitConfig().Period())
	}

	return s.rateLimiter
}

func (s *serviceProvider) GetBreakerConfig() config.BreakerConfig {
	if s.breakerConfig == nil {
		cfg, err := config.NewBreakerConfig()
		if err != nil {
			log.Fatalf("failed to get circuit breaker config: %v", err)
		}

		s.breakerConfig = cfg
	}

	return s.breakerConfig
}

func (s *serviceProvider) GetBreaker(_ context.Context) *gobreaker.CircuitBreaker {
	if s.circuitBreaker == nil {
		s.circuitBreaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "auth-service-api",
			MaxRequests: uint32(s.GetBreakerConfig().Requests()),
			Interval:    s.GetBreakerConfig().Interval(),
			Timeout:     s.GetBreakerConfig().Timeout(),
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				// >60% of requests failed => open circuit (no new requests allowed)
				return float64(counts.TotalFailures)/float64(counts.Requests) > 0.6
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				log.Fatal("grpc breaker state changed: %s %s -> %s", name, from, to)
				//WTF?
			},
		})
	}

	return s.circuitBreaker
}

func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusConfig == nil {
		cfg, err := config.NewPrometheusConfig()
		if err != nil {
			log.Fatalf("failed to get prometheus config: %s", err.Error())
		}

		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
}

// LoggerConfig - ...
func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) NoteRepository(ctx context.Context) repository.UserRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.DBClient(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) NoteService(ctx context.Context) service.UserService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(
			s.NoteRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.noteService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.NoteService(ctx))
	}

	return s.noteImpl
}

func (s *serviceProvider) LoginRepository(ctx context.Context) repository.LoginRepository {
	if s.loginRepository == nil {
		s.loginRepository = loginRepository.NewRepository(s.DBClient(ctx))
	}

	return s.loginRepository
}

func (s *serviceProvider) LoginService(ctx context.Context) service.LoginService {
	if s.loginService == nil {
		s.loginService = loginService.NewService(
			s.LoginRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.loginService
}

func (s *serviceProvider) LoginImpl(ctx context.Context) *login.Implementation {
	if s.loginImpl == nil {
		s.loginImpl = login.NewImplementation(s.LoginService(ctx))
	}

	return s.loginImpl
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.DBClient(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.AccessRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.accessService
}
func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}
