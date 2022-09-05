# (9) Clean Code

## Resume Materi

### Clean Code

Clean Code adalah istilah untuk Kode yang mudah dibaca, dipahami dan diubah oleh programmer. Sehingga programmer lain pun dapat dengan mudah membaca atau memodifikasi code tersebut. Alasan mengapa menggunakan clean code ada 3 yaitu Work Collaboration, Feature Development dan Faster Development

### Karakteristik Clean Code

Berikut beberapa karakteristik dari clean code : 1. Penamaan Mudah dipahami 2. Mudah dieja dan dicari 3. Singkat namum mendeskripsikan konteks 4. Konsisten 5. Hindari penambahan konteks yang tidak perlu 6. Komentar 7. Good function 8. Gunakan konvensi 9. Formatting

### Clean Code Principle

#### Kiss (Keep it So Simple)

Prinsip KISS menyatakan bahwa sebagian besar sistem bekerja paling baik jika dibuat sederhana daripada dibuat rumit; oleh karena itu, kesederhanaan harus menjadi tujuan utama dalam desain, dan kerumitan yang tidak perlu harus dihindari.

Kode sederhana memiliki manfaat sebagai berikut:

- sedikit waktu untuk menulis
- lebih sedikit kemungkinan bug
- lebih mudah dipahami, di-debug, dan dimodifikasi
- Lakukan hal paling sederhana yang mungkin bisa berhasil.

#### Don't Repeat Yourself (DRY)

Duplikasi kode dapat membuat kode sangat sulit untuk dipelihara. Setiap perubahan logika dapat membuat kode rentan terhadap bug atau dapat membuat perubahan kode menjadi sulit. Ini dapat diperbaiki dengan melakukan penggunaan kembali kode (Prinsip DRY). Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.

#### Refactoring

Refactoring adalah proses restrukturisasi kode yang dibuat, dengan mengubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor.
