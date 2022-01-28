package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//设置目标难度值（前两位为零）
const targetBit = 16

//共识算法管理文件
type pow struct {
	//需检验的目标区块
	block *Block
	//目标难度的哈希值（用大数存储）
	//用大数存储的目的：1.可以存储范围很大的数据 2.有专用方法进行位左移和右移
	target *big.Int
}

func CreateNewPOW(block *Block) *pow {
	//将target初始为一
	target := big.NewInt(1)
	//左移的目的是为了制定目标难度的哈希值
	//一个字节是8bit，哈希值是32位的字节数组，前两位为零，得左移（256-16）bit
	target = target.Lsh(target, 256-targetBit)
	return &pow{
		block,
		target,
	}
}

//执行pow，比较哈希
//返回哈希值，以及碰撞次数
func (pow *pow) Run() ([]byte, int64) {
	//碰撞次数
	var nonce int64 = 0
	var hashInt *big.Int
	hashInt = big.NewInt(0)
	var hash [32]byte //生成的哈希值
	//无限循环，生成符合条件的哈希值
	for {
		//生成准备数据
		databytes := pow.PrepareData(nonce)
		hash = sha256.Sum256(databytes)
		hashInt.SetBytes(hash[:])
		if pow.target.Cmp(hashInt) == 1 {
			//找到了符合条件的哈希，循环终止

			break
		}
		nonce++
	}
	fmt.Printf("\n打印碰撞次数：%d\n", nonce)
	return hash[:], nonce
}

//生成准备数据
func (pow *pow) PrepareData(nonce int64) []byte {
	var data []byte
	pow.block.Nonce = nonce
	//拼接区块属性
	HeightBytes := IntTopHex(pow.block.Height)
	TimeStampBytes := IntTopHex(pow.block.Timestamp)
	NonceBytes := IntTopHex(pow.block.Nonce)
	data = bytes.Join([][]byte{
		HeightBytes,
		TimeStampBytes,
		pow.block.PrevBlockHash,
		pow.block.Data,
		NonceBytes,
		IntTopHex(targetBit),
	}, []byte{})
	return data
}
