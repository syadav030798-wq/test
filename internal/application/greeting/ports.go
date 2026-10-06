package greeting

import "hotstar/internal/domain"

// Repository is a driven port: persistence or external lookup for greetings.
type Repository interface {
	DefaultMessage() string
}

// Service defines driving-side operations exposed to adapters (HTTP, CLI, etc.).
type Service interface {
	Hello(name string) (domain.Greeting, error)
}
