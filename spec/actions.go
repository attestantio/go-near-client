package spec

// ReceiptAction represents an action receipt
type ReceiptAction struct {
	Actions             []Action      `json:"actions"`
	GasPrice            string        `json:"gas_price"`
	InputDataIDs        []string      `json:"input_data_ids"`
	IsPromiseYield      bool          `json:"is_promise_yield"`
	OutputDataReceivers []interface{} `json:"output_data_receivers"`
	SignerID            string        `json:"signer_id"`
	SignerPublicKey     string        `json:"signer_public_key"`
}

// ReceiptActionItem represents an individual action within a receipt
type Action struct {
	Delegate     *DelegateAction `json:"Delegate,omitempty"`
	Transfer     *TransferAction `json:"Transfer,omitempty"`
	FunctionCall *FunctionCall   `json:"FunctionCall,omitempty"`
}

// FunctionCall represents a function call action
type FunctionCall struct {
	Args       string `json:"args"`
	Deposit    string `json:"deposit"`
	Gas        int64  `json:"gas"`
	MethodName string `json:"method_name"`
}

// TransferAction represents a transfer action
type TransferAction struct {
	Deposit string `json:"deposit"`
}

// DelegateAction represents a delegate action
type DelegateAction struct {
	DelegateAction DelegateActionData `json:"delegate_action"`
	Signature      string             `json:"signature"`
}

// DelegateActionData represents the data within a delegate action
type DelegateActionData struct {
	Actions        []DelegateActionItem `json:"actions"`
	MaxBlockHeight int64                `json:"max_block_height"`
	Nonce          int64                `json:"nonce"`
	PublicKey      string               `json:"public_key"`
	ReceiverID     string               `json:"receiver_id"`
	SenderID       string               `json:"sender_id"`
}

// DelegateActionItem represents an individual action within a delegate action
type DelegateActionItem struct {
	FunctionCall *FunctionCall `json:"FunctionCall,omitempty"`
}
