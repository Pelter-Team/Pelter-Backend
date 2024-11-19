package dto

type TransactionResponse struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	BuyerID   uint   `json:"buyer_id"`
	Amount    uint   `json:"amount"`
	CreatedAt string `json:"created_at"`
}
