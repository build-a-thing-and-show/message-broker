package services

// HelloWorldService defines the service interface.
type HelloWorldService interface {
	SayHello(name string) string
}

// helloWorldService is the concrete implementation.
type helloWorldService struct{}

// NewHelloWorldService creates a new HelloWorldService instance.
// Constructor function
func NewHelloWorldService() HelloWorldService {
	// This is a pointer to the above declared struct
	return &helloWorldService{}
}

// SayHello returns a greeting message.
func (s *helloWorldService) SayHello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
