package handler

import (
    "context"
    "mailer-service/internal/mailer"
    "mailer-service/internal/model"
    pb "mailer-service/internal/pb"
)

type EmailHandler struct {
    pb.UnimplementedEmailerServiceServer
    mailer *mailer.Mailer
}

func NewEmailHandler(m *mailer.Mailer) *EmailHandler {
    return &EmailHandler{mailer: m}
}

func (h *EmailHandler) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
    customer := model.Customer{Email: req.ToEmail}
    err := h.mailer.Send(ctx, customer)
    if err != nil {
        return &pb.SendEmailResponse{Status: "error", Message: err.Error()}, nil
    }
    return &pb.SendEmailResponse{Status: "success", Message: "Email sent!"}, nil
}
