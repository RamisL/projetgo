package payment

type InputPayment struct {
	ProductId int `json:"product_id" binding:"required"`
}
