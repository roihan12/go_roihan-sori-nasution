# (18) ORM and Code Structure (MVC)

## Resume Materi

### ORM (Object Relational Mapping)

Dalam ilmu komputer adalah teknik pemrograman untuk mengonversi data antara tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek.ORM (Object Relation Mapping) merupakan teknik yang merubah suatu table menjadi sebuah object yang nantinya mudah untuk digunakan. Object yang dibuat memiliki property yang sama dengan field — field yang ada pada table tersebut. ORM memungkinkan kita melakukan query dan memanipulasi data di database menggunakan object oriented.

### Kelebihan ORM

- Mengurangi query yang berulang
- Secara otomatis mengambil data menjadi siap digunakan
- obyek
- Cara sederhana jika ingin menyaring data
- Sebelum menyimpannya di database Beberapa memiliki permintaan cache fitur untuk melakukan chache query (case di beberapa ORM)

### Kelemahan ORM

- Menambah layer dalam kode dan membuat biaya overhead proses
- Memuat data hubungan yang tidak diperlukan
- Query kompleks yang rumit bisa membuatnya lama untuk ditulis ORM ( > 10 tabel bergabung )
- Fungsi SQL khusus yang terkait dengan satu vendor mungkintidak didukung atau tidak ada fungsi khusus

### Gorm

GORM merupakan fungsi ORM yang merupakan RDBMS. GORM dapat dipasangkan dengan framework echo yang sudah dibuat sebelumnya.fungsi GORM akan memungkinkan kita melakukan auto migration.

Cara Install

```
go get -u gorm.io/gorm
```

### Database Migration

Data migration penting dalam industri karena ketika melakukan update structure pada database secara simpel maka migrate akan melakukan tracking siapa yang menambahkan atau apa yang ditambahkan di versi yang baru.

- pembaruan basis data lebih simple
- rollback basis data menjadi lebih simpel
- dapat melacak perubahan pada struktur database
- Selalu kompatibel dengan aplikasi jika perubahan versi

### MVC

MVC adalah sebuah pola desain arsitektur dalam sistem yang terdiri dari model, view, controllers. Artinya nanti projek kita dipisah - pisah menjadi bagian- bagian tersebut sehingga lebih mudah dibaca. Model itu seperti struct dimana model dari struktur data yang akan dimasukkan ke database. controller adalah isi dari logic bisnis nya kita, view adalah bagian awal seperti main.
