package app

import "fmt"

func Start() {
	fmt.Println("=================================")
	fmt.Println("Starting", Name)
	fmt.Println("Version:", Version)
	fmt.Println("Chain ID:", ChainID)
	fmt.Println("Native Coin:", Denom)
	fmt.Println("=================================")
}
