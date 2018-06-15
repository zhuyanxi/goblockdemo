package block

// Transaction :
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// TXInput :
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// TXOutput :
type TXOutput struct {
	Value        int
	ScriptPubKey string
}
