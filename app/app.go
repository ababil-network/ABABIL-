package app

import "fmt"

func Start() {
	fmt.Println("=================================")
	fmt.Println("        ABABIL NETWORK")
	fmt.Println("=================================")
	fmt.Println("Version :", AppVersion)
	fmt.Println("Chain ID:", ChainID)
	fmt.Println("Coin    :", DisplayDenom)
	fmt.Println("Binary  :", BinaryName)
	fmt.Println("=================================")
}
