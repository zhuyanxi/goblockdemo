package block

// BlockChain :
type BlockChain struct {
	Blocks []*Block
}

// AddBlock :
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// newGenesisBlock :
func newGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain :
func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{newGenesisBlock()}}
}
