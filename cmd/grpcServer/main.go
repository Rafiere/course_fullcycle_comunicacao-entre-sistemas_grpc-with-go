package main

import (
	"database/sql"
	"github.com/rafiere/course_fullcycle_comunicacao-entre-sistemas_grpc-with-go/internal/database"
	"github.com/rafiere/course_fullcycle_comunicacao-entre-sistemas_grpc-with-go/internal/pb"
	"github.com/rafiere/course_fullcycle_comunicacao-entre-sistemas_grpc-with-go/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	/* Estamos criando o servidor gRPC e registrando o serviço no servidor gRPC. */
	grpcServer := grpc.NewServer()

	/* Utilizaremos um client de gRPC, que é o "evans", e para isso, teremos que
	utilizar mecanismos de reflection. */

	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	/* Após isso, temos que abrir uma conexão TCP para falarmos com o gRPC. */
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
