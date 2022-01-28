package BLC

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//区块基本结构与功能管理文件

//实现一个最基本的区块结构
type Block struct {
	Timestamp     int64  //区块时间戳
	Hash          []byte //当前区块哈希
	PrevBlockHash []byte //前区块哈希
	Height        int64  //区块高度
	Data          []byte //交易数据
}

//新建区块
func NewBlock(height int64, prevblockhash []byte, data []byte) *Block {
	var block Block
	block = Block{
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: prevblockhash,
		Height:        height,
		Data:          data,
	}
	block.SetHash()
	return &block
}

//计算区块哈希
func (b *Block) SetHash() {
	//调用sha256实现哈希生成
	//将int64->hash
	HeightBytes := IntTopHex(b.Height)
	TimeStampBytes := IntTopHex(b.Timestamp)
	blockbytes := bytes.Join([][]byte{
		HeightBytes,
		TimeStampBytes,
		b.PrevBlockHash,
		b.Data,
	}, []byte{})
	hash := sha256.Sum256(blockbytes)
	b.Hash = hash[:]
}

//创建创世区块
func CreateGenesisBlock(data []byte) *Block {

	return NewBlock(0, nil, data)
}
