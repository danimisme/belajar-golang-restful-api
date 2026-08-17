//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

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

// cara memakai set
var fooSet = wire.NewSet(NewFooService, NewFooRepository)
var barSet = wire.NewSet(NewBarService, NewBarRepository)

func InitializeFooBarService() *FooBarService {
	// wire.Build dapat menerima lebih dari 1 parameter
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

var FooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializeFooBar() *FooBar {
	// wire.Struct digunakan agar wire bisa membuat struct
	// parameter keduanya adalah pointer dari struct yang ingin dibuat
	// parameter ketiga adalah field yang ingin diisi
	wire.Build(FooBarSet, wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarWithValue() *FooBar {
	// Kita memasukkan sebuah Value ke dalam wire
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializeReader() io.Reader {
	// contoh bind interface ke value
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}