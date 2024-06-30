package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
    pb "github.com/perfectbleu/commons/api"
)

const (
    httpAddr = ":8080"
    orderServiceAddr = "localhost:2000"
)

func main() {
    conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to dial server %v", err)
    }
    defer conn.Close()

    log.Println("Dialing orders service at", orderServiceAddr)
    c := pb.NewOrderServiceClient(conn)

    mux := http.NewServeMux() 
    handler := NewHandler(c)
    handler.registerRoutes(mux)

    log.Printf("🚀 Starting HTTP server at http://localhost/%s", httpAddr)

    if err := http.ListenAndServe(httpAddr, mux); err != nil {
        log.Fatal("Failed to start HTTP server")
    }
}

