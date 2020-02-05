package helloworld

type Service struct {
}

func (s *Service) HelloWorld(name string) string {
	if name == "" {
		return "Hello world"
	}
	return "Hello " + name
}
