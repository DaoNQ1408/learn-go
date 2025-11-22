package models

type User struct {
	ID       int    `json: "id"`
	Mail     string `json: "mail"`
	Password string `json: "-"`
	Fullname string `json: "fullname"`
}

type Brand struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

type Product struct {
	ID      int    `json: "id"`
	Name    string `json: "name"`
	Price   int    `json: "price"`
	BrandID int    `json: "brand_id"`
	Brand   *Brand `json: "brand,omitempty"`
}
