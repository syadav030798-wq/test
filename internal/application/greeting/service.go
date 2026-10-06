package greeting

import (
	"fmt"
	"strings"

	"hotstar/internal/domain"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Hello(name string) (domain.Greeting, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return domain.Greeting{}, domain.ErrEmptyName
	}

	template := s.repo.DefaultMessage()
	return domain.Greeting{
		Message: fmt.Sprintf(template, name),
	}, nil
}
