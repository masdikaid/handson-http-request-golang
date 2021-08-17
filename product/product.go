package product

import "time"

type Product struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	ProductId string    `json:"product_id"`
	Price     int       `json:"price"`
	BuyAt     time.Time `json:"buy_at"`
}

func ProductList() []Product {
	return []Product{
		{Name: "Mie Instant", Category: "Food", ProductId: "0087765512", Price: 3000, BuyAt: time.Now()},
		{Name: "Shampo", Category: "toiletries", ProductId: "0087765512", Price: 30000, BuyAt: time.Now()},
		{Name: "Pasta Gigi", Category: "toiletries", ProductId: "0087765512", Price: 20000, BuyAt: time.Now()},
		{Name: "Daging", Category: "Food", ProductId: "0087765512", Price: 45000, BuyAt: time.Now()},
	}
}
