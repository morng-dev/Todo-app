package services

import "context"

type Mailer interface {
	SendResetPassword(ctx context.Context, to, resetToken string) error
}
