package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/api/note"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/closer"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/config"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	noteRepository "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
	noteService "github.com/atlasir0/Chat_service/Auth_chat/internal/service/note"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool         *pgxpool.Pool
	noteRepository repository.UserRepository

	noteService service.NoteService

	noteImpl *note.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) GetPgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.GetPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) repository.UserRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.GetPgPool(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}

func (s *serviceProvider) GetNoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.GetNoteService(ctx))
	}

	return s.noteImpl
}
