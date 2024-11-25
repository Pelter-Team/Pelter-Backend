package dto

type TransactionResponse struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	BuyerID   uint   `json:"buyer_id"`
	SellerID  uint   `json:"seller_id"`
	Amount    uint   `json:"amount"`
	CreatedAt string `json:"created_at"`
}

type TransactionWithProductResponse struct {
	ID         uint    `json:"id"`
	ProductID  uint    `json:"product_id"`
	BuyerID    uint    `json:"buyer_id"`
	SellerID   uint    `json:"seller_id"`
	Amount     uint    `json:"amount"`
	CreatedAt  string  `json:"created_at"`
	Price      float64 `json:"price"`
	IsVerified bool    `json:"is_verified"`
	IsSold     bool    `json:"is_sold"`
}
