package app

import "fmt"

func Start() {
	fmt.Println("========================================")
	fmt.Println("        ABABIL NETWORK")
	fmt.Println("========================================")
	fmt.Println("Version :", Version)
	fmt.Println("Chain ID:", ChainID)
	fmt.Println("Coin     :", Denom)
	fmt.Println("Binary   :", BinaryName)
	fmt.Println("========================================")
}
