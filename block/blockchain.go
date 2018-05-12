package block

// BlockChain :
type BlockChain struct {
	Blocks []*Block
}

// AddBlock :
func (bc *BlockChain) AddBlock(data string, nouce int64) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Height+1, data, prevBlock.Hash, nouce)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// newGenesisBlock :
func newGenesisBlock() *Block {
	return NewBlock(0, "The First Block", []byte{}, 6929) //6929, 22949
}

// NewBlockchain :
func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{newGenesisBlock()}}
}
