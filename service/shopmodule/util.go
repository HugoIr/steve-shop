package shopmodule

import (
	"errors"
	"fmt"

	m "github.com/HugoIr/steve-shop/service/model"
)

func SanitizeInsert(param m.ProductRequest) (m.ProductRequest, error) {
	if param.Name == "" {
		return param, errors.New("name cannot be empty")
	}
	if param.Description == "" {
		return param, errors.New("description cannot be empty")
	}
	if param.Price < 0 {
		return param, errors.New("invalid rating range")
	}
	if param.Discount < 0 {
		return param, errors.New("invalid rating range")
	}
	if param.Stock < 0 {
		return param, errors.New("invalid rating range")
	}

	return param, nil
}

func BuildQuery(id int64, param m.ProductRequest) (finalQuery string, fieldValues []interface{}) {
	var fieldQuery string
	fieldValues = make([]interface{}, 0)

	var i = 1
	if param.Name != "" {
		fieldQuery += fmt.Sprintf("name=$%d,", i)
		fieldValues = append(fieldValues, param.Name)
		i++
	}
	if param.Description != "" {
		fieldQuery += fmt.Sprintf("description=$%d,", i)
		fieldValues = append(fieldValues, param.Description)
		i++
	}
	if param.Price != 0 {
		fieldQuery += fmt.Sprintf("price=$%d,", i)
		fieldValues = append(fieldValues, param.Price)
		i++
	}
	if param.Discount != 0 {
		fieldQuery += fmt.Sprintf("discount=$%d,", i)
		fieldValues = append(fieldValues, param.Discount)
		i++
	}
	if param.Stock != 0 {
		fieldQuery += fmt.Sprintf("stock=$%d,", i)
		fieldValues = append(fieldValues, param.Stock)
		i++
	}

	i++

	finalQuery = fmt.Sprintf(updateProductQuery, fieldQuery[:len(fieldQuery)-1], id)

	return
}
