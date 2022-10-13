package shop

import "github.com/HugoIr/steve-shop/service/shopmodule"

type InsertProductResponse struct {
	ID int64 `json:"id"`
}

type Handler struct {
	shop *shopmodule.Module
}

func NewProductHandler(p *shopmodule.Module) *Handler {
	return &Handler{
		shop: p,
	}
}
