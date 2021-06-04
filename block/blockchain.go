package block

import (
	"github.com/boltdb/bolt"
	"log"

)
//数据库名称
const dbName = "block.db"

//表名称
const blockTableName = "blocks"

type BlockChain struct {
	Blocks []*Block		//区块的切片
	DB *bolt.DB 			//数据库对象
	Tip []byte 				//保存最新区块哈希
}

//初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//保存最新区块的哈希
	var blockHash []byte

	//1.创建或者打开一个数据库
	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Panicf("create db [%s] failed %v\n",dbName,err)
	}
	//2.创建桶
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//没找到桶
			b,err := tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panicf("create bucket[%s] failed %v\n",blockTableName,err)
			}

			//生成创世区块
			genesisBlock := CreateGenesisBlock([]byte("init blockchain"))
			//存储
			err=b.Put(genesisBlock.Hash,genesisBlock.Serialize())
			if nil!=err{
				log.Printf("insert the genesis block failed %v\n",err)
			}
			blockHash = genesisBlock.Hash
			//存储最新区块的哈希
			b.Put([]byte("1"),genesisBlock.Hash)
		}
		return err
	})
	return &BlockChain{DB:db,Tip: blockHash}
}

//添加区块到区块链中
func (bc *BlockChain) AddBlock(height int64, prevBlockHash []byte, data []byte) {
	newBlock := NewBlock(height, prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

