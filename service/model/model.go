package model

type ProductResponse struct {
	ID          int64  `json:"product_id,omitempty" db:"id"`
	Name        string `json:"product_name,omitempty" db:"name"`
	Description string `json:"product_description,omitempty" db:"description"`
	Price       int    `json:"product_price,omitempty" db:"price"`
	Discount    int    `json:"product_discount,omitempty" db:"discount"`
	Stock       int    `json:"product_stock,omitempty" db:"stock"`
}

type ProductRequest struct {
	Name        string `json:"product_name"`
	Description string `json:"product_description"`
	Price       int    `json:"product_price"`
	Discount    int    `json:"product_discount"`
	Stock       int    `json:"product_stock"`
}
