package model

type BuyProductBody struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
}
