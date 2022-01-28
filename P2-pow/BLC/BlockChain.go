package BLC

//区块链基本结构
type Blockchain struct {
	Blocks []*Block //区块的切片
}

//添加区块到区块链
func (bc *Blockchain) AddBlock(height int64, prevblockhash []byte, data []byte) {
	newblock := NewBlock(height, prevblockhash, data)
	bc.Blocks = append(bc.Blocks, newblock)
}

//初始化区块
func CreateBlockChain() *Blockchain {
	//创建创世区块
	block := CreateGenesisBlock([]byte("the GenesisBlock"))
	//把创世区块加入到区块链中
	return &Blockchain{Blocks: []*Block{block}}
}
