package mailgun

import (
	"context"
	"errors"
	"testing"

	"github.com/darkrockmountain/gomail"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/stretchr/testify/assert"
)

// Mock implementations for Mailgun
type mockMailgun struct{}

func (m *mockMailgun) NewMessage(from, subject, text string, to ...string) *mailgun.Message {
	return mailgun.NewMailgun("", "").NewMessage(from, subject, text, to...)
}

func (m *mockMailgun) Send(ctx context.Context, message *mailgun.Message) (string, string, error) {
	return "", "", nil
}

type mockMailgunWithError struct{}

func (m *mockMailgunWithError) NewMessage(from, subject, text string, to ...string) *mailgun.Message {
	return mailgun.NewMailgun("", "").NewMessage(from, subject, text, to...)
}

func (m *mockMailgunWithError) Send(ctx context.Context, message *mailgun.Message) (string, string, error) {
	return "", "", errors.New("send error")
}

func TestNewMailgunEmailSender(t *testing.T) {
	domain := "example.com"
	apiKey := "fake-api-key"
	emailSender, err := NewMailgunEmailSender(domain, apiKey)
	assert.NoError(t, err)
	assert.NotNil(t, emailSender)
	assert.NotNil(t, emailSender.mailgunClient)
}

func TestMailgunEmailSender_SendEmail(t *testing.T) {
	emailSender := &mailgunEmailSender{
		mailgunClient: &mockMailgun{},
	}

	message := gomail.EmailMessage{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Email",
		Text:    "This is a test email.",
		HTML:    "<p>This is a test email.</p>",
		CC:      []string{"cc@example.com"},
		BCC:     []string{"bcc@example.com"},
		ReplyTo: "replyto@example.com",
		Attachments: []gomail.Attachment{
			{
				Filename: "test.txt",
				Content:  []byte("This is a test attachment."),
			},
		},
	}

	err := emailSender.SendEmail(message)
	assert.NoError(t, err)
}

func TestMailgunEmailSender_SendEmailWithSendError(t *testing.T) {
	emailSender := &mailgunEmailSender{
		mailgunClient: &mockMailgunWithError{},
	}

	message := gomail.EmailMessage{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Email",
		Text:    "This is a test email.",
	}

	err := emailSender.SendEmail(message)
	assert.Error(t, err)
	assert.Equal(t, "failed to send email: send error", err.Error())
}

func TestMailgunEmailSender_SendEmailWithEmptyFields(t *testing.T) {
	emailSender := &mailgunEmailSender{
		mailgunClient: &mockMailgun{},
	}

	message := gomail.EmailMessage{
		From:    "sender@example.com",
		To:      []string{},
		Subject: "",
		Text:    "",
	}

	err := emailSender.SendEmail(message)
	assert.NoError(t, err)
}
