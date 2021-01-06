package data

import "time"

type Product struct {
	Id        int
	Name      string
	Price     float32
	SKU       string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		Id:        1,
		Name:      "Computer",
		Price:     99.9,
		SKU:       "XYZ-111",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		Id:        2,
		Name:      "Cellphone",
		Price:     99.9,
		SKU:       "XYZ-111",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
