package main

// type Product struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price"`
// 	Account     Account `json:"account"`
// }

// type Order struct {
// 	ID         int       `json:"id"`
// 	CreatedAT  time.Time `json:"created_at"`
// 	TotalPrice float64   `json:"total_price"`
// 	Products   []Product `json:"products"`
// }

type Account struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Orders []Order `json:"orders"`
}
