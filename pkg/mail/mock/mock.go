package mock

import (
	"fmt"
	"strings"
)

type mock int

func New() *mock {
	var m mock = 0
	return &m
}

func (m *mock) Mail(from, subject, body string, to, cc []string) error {
	fmt.Printf("FROM: %s\n TO: %s\n CC: %s\n SUBJECT: %s\n BODY: %s\n", from, strings.Join(to, ","), strings.Join(cc, ","), subject, body)
	return nil
}
