package main

import (
	"Project/BlockChain/P2-pow/BLC"
	"fmt"
)

func main() {

	var bc = BLC.CreateBlockChain()
	//fmt.Printf("blockchain%v\n", bc.Blocks[0])
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Hash,
		[]byte("Jerry sent 10 btc to Lewis"))
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Hash,
		[]byte("Abbott sent 80 btc to Jerry"))
	fmt.Printf("blockchain%v\n", bc.Blocks[0])
	fmt.Printf("blockchain%v\n", bc.Blocks[1])
	fmt.Printf("blockchain%v\n", bc.Blocks[2])
}
