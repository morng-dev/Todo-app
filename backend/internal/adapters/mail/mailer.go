package mail

import (
	"context"
	"fmt"
	"morng-dev/internal/core/domain/ports/services"

	mail "github.com/wneessen/go-mail"
)

type SmtpMailer struct {
	Host     string
	Username string
	Password string
	From     string
	AppURL   string
	Port     int
}

func NewSMTPMailer(host string, port int, username, password, from, appURL string) services.Mailer {
	return &SmtpMailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
		AppURL:   appURL,
	}
}

func (s *SmtpMailer) SendResetPassword(ctx context.Context, to, resetToken string) error {
	resetURL := fmt.Sprintf("%s/reset-password?token=%s", s.AppURL, resetToken)

	msg := mail.NewMsg()
	if err := msg.From(s.From); err != nil {

		return err

	}
	if err := msg.To(to); err != nil {

		return err

	}
	msg.Subject("รีเซ็ตรหัสผ่านของคุณ")
	msg.SetBodyString(mail.TypeTextHTML, fmt.Sprintf(`

		<h2>รีเซ็ตรหัสผ่าน</h2>

		<p>เราได้รับคำขอรีเซ็ตรหัสผ่านสำหรับบัญชีของคุณ</p>

		<p><a href="%s">กดที่นี่เพื่อตั้งรหัสผ่านใหม่</a></p>

		<p>ลิงก์นี้จะหมดอายุใน 15 นาที</p>

		<p>หากคุณไม่ได้เป็นผู้ร้องขอ กรุณาเพิกเฉยต่ออีเมลนี้</p>

	`, resetURL))
	client, err := mail.NewClient(

		s.Host,

		mail.WithPort(s.Port),

		mail.WithSMTPAuth(mail.SMTPAuthPlain),

		mail.WithUsername(s.Username),

		mail.WithPassword(s.Password),

		mail.WithTLSPortPolicy(mail.TLSMandatory),
	)

	if err != nil {

		return err

	}
	return client.DialAndSendWithContext(ctx, msg)
}
