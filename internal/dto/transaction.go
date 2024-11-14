package dto

type TransactionRequest struct {
	ProductID uint `json:"product_id"`
	BuyerID   uint `json:"buyer_id"`
}
