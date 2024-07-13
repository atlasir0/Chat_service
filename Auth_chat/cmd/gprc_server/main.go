package main

import (
	"context"
	"log"
	"net"

	//	desc "github.com/atlasir0/Chat_service/Auth_chat/grpc/pkg/note_v1"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/config"
	//	noteAPI "github.com/atlasir0/Chat_service/Auth_chat/internal/api/note"
	noteRepository "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note"
	noteService "github.com/atlasir0/Chat_service/Auth_chat/internal/service/note"
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

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.GRPCAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	noteRepo := noteRepository.NewRepository(pool)
	noteSrv := noteService.NewService(noteRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, noteAPI.NewImplementation(noteSrv))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
