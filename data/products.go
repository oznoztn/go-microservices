package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"desc, omitempty"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Product) FromJson(r io.Reader) error {
	de := json.NewDecoder(r)
	return de.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	// Marshal yerine direkt encoder (stream) üzerinde çalıştık
	// Marshal bize array döndürüyordu, bu da allocation demek.
	// Bu büyük JSON dokümanlarında sıkıntı oluştururdu.
	// Artık direkt encoder üzerinde çalıştığımızdan allocation yok.
	// Hem daha hızlı. Belki normal kullanımlarda bu göze batmayabilir,
	// fakat özellikle yüksek trafikli mikroservis ortamlarında fark hissettirebilir.
	// Diğer taraftan bu yöntem öyle mikro optimizasyon vs. de değil.
	// Dolayısıyla bu yol varken neden ötekini kullanalım ki?
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// []*Product döndürmek yerine custom Products tipini döndürüyoruz. İkisi de birbiri yerine kullanılabilir.
func GetProducts() Products {
	return productList
}

// AddProduct adds a new product into data store
func AddProduct(p *Product) {
	p.Id = generateNextProductId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.Id = id
	productList[pos] = p

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.Id == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func generateNextProductId() int {
	lp := productList[len(productList)-1]
	return lp.Id + 1
}

var productList = []*Product{
	&Product{
		Id:          1,
		Name:        "Computer",
		Description: "Lorem impsum dolor sit amet",
		Price:       99.9,
		SKU:         "XYZ-111",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
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
