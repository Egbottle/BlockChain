package main

import (
	"Project/BlockChain/P1-bkc/BLC"
	"fmt"
)

func main() {
	//block := BLC.NewBlock(1, nil, []byte("the first testing"))
	//fmt.Printf("the first block : %v\n", block)
	//gblock := BLC.CreateGenesisBlock([]byte("yanzihan"))
	//fmt.Println(gblock)
	var bc = BLC.CreateBlockChain()
	//fmt.Printf("blockchain%v\n", bc.Blocks[0])*/
	var i int64
	for i = 1; i <= 4; i++ {
		bc.AddBlock(i, bc.Blocks[i-1].PrevBlockHash, []byte("qukuai"))
	}
	fmt.Printf("blockchain%x\n", bc.Blocks)
}
