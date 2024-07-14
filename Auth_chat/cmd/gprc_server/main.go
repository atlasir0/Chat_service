package main

import (
	"context"
	"log"
	"net"

	noteAPI "github.com/atlasir0/Chat_service/Auth_chat/internal/api/note"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/config"
	noteRepository "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note"
	noteService "github.com/atlasir0/Chat_service/Auth_chat/internal/service/note"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Println("Config loaded successfully")

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}
	log.Printf("GRPC Config: %+v", grpcConfig)

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}
	log.Printf("PG Config: %+v", pgConfig)

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", grpcConfig.Address())

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Connected to database successfully")

	noteRepo := noteRepository.NewRepository(pool)
	log.Println("Note repository initialized")

	noteSrv := noteService.NewService(noteRepo)
	log.Println("Note service initialized")

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserServiceServer(s, noteAPI.NewImplementation(noteSrv))
	log.Println("gRPC server registered and reflection enabled")

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
