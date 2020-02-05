package helloworld

type AdminService struct {
}

func (s *AdminService) HelloAdmin() string {
	return "I am admin"
}
