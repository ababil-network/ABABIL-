package module

import "fmt"

type AppModule struct{}

func (AppModule) Name() string {
	return "ababil"
}

func (AppModule) RegisterServices() {
	fmt.Println("ABABIL module services registered")
}
