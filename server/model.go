package server

type SubscribeRequest struct {
	Address string `json:"address"`
}

type GetTransactionRequest struct {
	Address string `json:"address"`
}
