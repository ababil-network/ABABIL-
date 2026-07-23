package keeper

import "fmt"

type Keeper struct{}

func (k Keeper) Init() {
	fmt.Println("ABABIL Keeper initialized")
}
