package model

type ItemInput struct {
	Nama       string  `json:"nama"`
	Deskripsi  string  `json:"deskripsi"`
	Stok       int     `json:"stok"`
	Harga      float64 `json:"harga"`
	CategoryID uint    `json:"category_id"`
}
