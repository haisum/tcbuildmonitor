package plainsmtp

import (
	"gopkg.in/gomail.v2"
)

type plainsmtp struct {
	dialer      *gomail.Dialer
	toWhiteList []string
}

func New(username, password, host string, port int, toWhiteList []string) *plainsmtp {
	return &plainsmtp{
		gomail.NewDialer(host, port, username, password),
		toWhiteList,
	}
}

func (p *plainsmtp) Mail(from, subject, body string, to, cc []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	var toWhiteList []string
	for _, t := range to {
		if In(p.toWhiteList, t) {
			toWhiteList = append(toWhiteList, t)
		}
	}
	m.SetHeader("To", toWhiteList...)
	m.SetHeader("Cc", cc...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	if err := p.dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func In(l []string, s string) bool {
	for _, v := range l {
		if s == v {
			return true
		}
	}
	return false
}
