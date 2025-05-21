package mailer

import (
	"context"
	"time"

	"mailer-service/internal/model"

	mailersend "github.com/mailersend/mailersend-go"
)

type Mailer struct {
	client *mailersend.Mailersend
}

func NewMailer(client *mailersend.Mailersend) *Mailer {
	return &Mailer{client: client}
}

func (m *Mailer) Send(ctx context.Context, customer model.Customer) error {
	ctxC, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  "SSN Motors",
		Email: "test-vz9dlemrqv14kj50.mlsender.net", // должен быть валидный, привязанный к домену в MailerSend
	}

	recipients := []mailersend.Recipient{
		{Email: customer.Email},
	}

	message := m.client.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject("Welcome!")
	message.SetHTML("Welcome to our awesome platform!")
	message.SetText("Welcome to our awesome platform!")

	_, err := m.client.Email.Send(ctxC, message)
	return err
}
