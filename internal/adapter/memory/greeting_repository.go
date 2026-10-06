package memory

// GreetingRepository is an in-memory driven adapter for greeting.Repository.
type GreetingRepository struct{}

func NewGreetingRepository() *GreetingRepository {
	return &GreetingRepository{}
}

func (GreetingRepository) DefaultMessage() string {
	return "Hello, %s!"
}
