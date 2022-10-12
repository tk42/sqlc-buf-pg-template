package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	bufbuild "github.com/tk42/sqlc-buf-pg-template/gen/proto/golang/github.com/tk42/sqlc-buf-pg-template"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	log.Print("server is starting...")
	client, err := sql.Open(
		"pgx",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			"db", "5432", "postgres", "postgres", "e8a48653851e28c69d0506508fb27fc5"))
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}
	defer client.Close()

	svc := NewServiceServer(client)
	server := grpc.NewServer()

	reflection.Register(server) // Failed to list services: server does not support the reflection API
	bufbuild.RegisterPetStoreServiceServer(server, svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
