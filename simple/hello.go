package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello
}

type SayHelloImp struct {

}

func (s *SayHelloImp) Hello(name string) string {
	return "Hello " + name
}

func NewSayHelloImp() *SayHelloImp {
	return &SayHelloImp{}
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}
	