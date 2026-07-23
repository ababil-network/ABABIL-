package module

import "fmt"

type AppModule struct{}

func (AppModule) Name() string {
	return "ababil"
}

func (AppModule) Register() {
	fmt.Println("ABABIL module registered")
}
