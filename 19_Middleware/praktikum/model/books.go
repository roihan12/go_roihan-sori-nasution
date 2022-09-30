package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	JudulBuku   string `json:"judul_buku" form:"judul_buku"`
	Kategori    string `json:"kategori" form:"kategori"`
	Pengarang   string `json:"pengarang" form:"pengarang"`
	Penerbit    string `json:"penerbit" form:"penerbit"`
	TahunTerbit string `json:"tahun_terbit" form:"tahun_terbit"`
	Isbn        uint   `json:"isbn" form:"isbn"`
}
