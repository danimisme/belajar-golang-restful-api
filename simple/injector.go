//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	) 
	return nil
}

var fooSet = wire.NewSet(NewFooService, NewFooRepository)
var barSet = wire.NewSet(NewBarService, NewBarRepository)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

//  contoh salah
// func InitializeHelloService() *HelloService {
// 	wire.Build(NewSayHelloImp, NewHelloService)
// 	return nil
// }


var helloSet = wire.NewSet(
	NewSayHelloImp,
	// ini artinya: "Ketika ada kebutuhan tipe data interface SayHello, gunakan instance NewSayHelloImp untuk membuat nya"
	wire.Bind(new(SayHello),new(*SayHelloImp)),
)

func InitializeHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}