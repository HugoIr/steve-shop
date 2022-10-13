package shipper

import "github.com/HugoIr/steve-shop/service/shippermodule"

type InsertProductResponse struct {
	ID int64 `json:"id"`
}

type Handler struct {
	shipper *shippermodule.Module
}

func NewProductHandler(p *shippermodule.Module) *Handler {
	return &Handler{
		shipper: p,
	}
}
