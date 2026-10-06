package greeting_test

import (
	"errors"
	"testing"

	"hotstar/internal/application/greeting"
	"hotstar/internal/domain"
)

type stubRepo struct {
	msg string
}

func (s stubRepo) DefaultMessage() string { return s.msg }

func TestHello_success(t *testing.T) {
	svc := greeting.NewService(stubRepo{msg: "Hi, %s"})
	g, err := svc.Hello("Sonu")
	if err != nil {
		t.Fatalf("Hello: %v", err)
	}
	if g.Message != "Hi, Sonu" {
		t.Fatalf("got %q", g.Message)
	}
}

func TestHello_emptyName(t *testing.T) {
	svc := greeting.NewService(stubRepo{msg: "Hi, %s"})
	_, err := svc.Hello("  ")
	if !errors.Is(err, domain.ErrEmptyName) {
		t.Fatalf("got err %v", err)
	}
}
