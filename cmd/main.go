package main

import (
    "mailer-service/config"
    "mailer-service/internal/handler"
    "mailer-service/internal/mailer"
    pb "mailer-service/internal/pb"
    "github.com/mailersend/mailersend-go"
    "google.golang.org/grpc"
    "log"
    "net"
)

func main() {
    apiKey := config.LoadAPIKey()
    client := mailersend.NewMailersend(apiKey)
    mailService := mailer.NewMailer(client)

    grpcServer := grpc.NewServer()
    emailHandler := handler.NewEmailHandler(mailService)
    pb.RegisterEmailerServiceServer(grpcServer, emailHandler)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    log.Println("Email gRPC server running on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
