package main

import (
	"kongyixueyuan.com/publicChain/part8-proof-of-work/BLC"
	"fmt"
)

func main()  {

	//data string,height int64,prevBlockHash []byte
	block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)

	proofOfWork := BLC.NewProofOfWork(block)

	fmt.Printf("%v",proofOfWork.IsValid())

}
