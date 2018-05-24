package block

const dbName = "blockchain"

// BlockChain :
type BlockChain struct {
	tipHash []byte
}

// BlockChainDB :
type BlockChainDB struct {
	BCMap map[string]BlockChain
}

// newGenesisBlock :
func newGenesisBlock() *Block {
	return NewBlock(0, "The First Block", []byte{}) //6929, 22949
}

// NewBlockChain :
func NewBlockChain() *BlockChain {
	var tipHash []byte
	gene := newGenesisBlock()
	tipHash = gene.Hash
	bc := BlockChain{tipHash: tipHash}
	return &bc
}
