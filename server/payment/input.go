package payment

type InputCreatePayment struct {
	ProductId int `json:"product_id" binding:"required"`
}

type InputUpdatePayment struct {
	Name  string `json:"name", omitempty:`
	Price string `json:"price", omitempty`
}
