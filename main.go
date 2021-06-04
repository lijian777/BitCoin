package main

import (
	"fmt"
	"github.com/lijian777/BitCoin/block"
)


//启动
func main() {
	//var bc=bc.NewBlock(1, []byte("aaa"), []byte("vvvv"))
	//fmt.Printf("frist bolck: %v\n",bc)
	//初始化区块链bc
	bc := block.CreateBlockChainWithGenesisBlock()
	//上链
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,
		bc.Blocks[len(bc.Blocks)-1].Hash, []byte("lily send 10 btc to bob"))
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,
		bc.Blocks[len(bc.Blocks)-1].Hash, []byte("mike send 10 btc to bob"))
	for _, b := range bc.Blocks {
		fmt.Printf("prev hash: %x,current hash:%x \n", b.PrevBlockHash, b.Hash)
	}
}
