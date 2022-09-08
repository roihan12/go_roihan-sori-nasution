package main

type mobil struct {
	kecepatan int
}

// function jalan berfungsi untuk memanggil function tambahKecepatan dan memberikan nilainya
func (m *mobil) jalan() {
	m.tambahKecepatan(10)
}

// function/ method tambahKecepatan berfungsi untuk menambah kecepatan mobil
func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatan += kecepatanBaru
}

func main() {
	mobilCepat := mobil{}
	mobilCepat.jalan()
	mobilCepat.jalan()
	mobilCepat.jalan()

	mobilLamban := mobil{}
	mobilLamban.jalan()
}