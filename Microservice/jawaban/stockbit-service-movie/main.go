package main

import (
	"log"
	"net"
	"os"
	grpcServer "stockbit-service-movie/transport/grpc"
	pb "stockbit-service-movie/transport/grpc/proto/movies"
	route "stockbit-service-movie/transport/http"
	db "stockbit-service-movie/util/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	conn, err := db.GetSqlxConnection()
	if err != nil {
		log.Fatalf("failed to connect database")
	}
	defer conn.Close()

	go  route.Serve(conn)

	lis, err := net.Listen("tcp", ":"+os.Getenv("port_grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC to port " + os.Getenv("port_grpc"))
	}

	s := grpc.NewServer()

	pb.RegisterMoviesServer(s, grpcServer.NewServer(conn))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
