# (11) Concurrent Programing

## Resume Materi

### Sequential, Parallel, Concurent

- Pemrograman sekuensial berarti bahwa suatu program beroperasi secara berurutan, yang berarti satu instruksi pada suatu waktu dalam urutan linier. Dalam program berurutan, yaitu sebelum tugas baru dimulai, yang sebelumnya harus juga sudah selesai.

- Dalam program paralel, banyak tugas dapat dilakuka dieksekusi secara bersamaan. (membutuhkan mesin multi-core)

- Dalam program Concurent, banyak tugas dapat dieksekusi secara independen dan dapat muncul bersamaan.

### Goroutines

Goroutine adalah fungsi atau metode yang berjalan secara bersamaan (independen) dengan fungsi atau methods lain.Biaya pembuatan Goroutine sangat kecil jika dibandingkan dengan thread . Sebuah thread adalah proses ringan, atau dengan kata lain, sebuah thread adalah unit yang mengeksekusi kode di bawah program.Goroutine sangat ringan, hanya dibutuhkan sekitar 2kB memori saja untuk satu buah goroutine. Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.

Beberapa FEATURE di miliki oleh GO :

- Concurrent execution ( Goroutines )
- Synchronization and messaging ( Channels )
- Multi - way concurrent control ( Select )

### Channel & Select

Channel dalah komunikasi menggunakan objek goroutine yang dapat berkomunikasi satu sama lain.Dalam komunikasinya, sebuah channel difungsikan sebagai pengirim di sebuah goroutine, dan juga sebagai penerima di goroutine lainnya. Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.

Select membuatnya lebih mudah untuk mengontrol data komunikasi melalui satu atau banyak saluran.Select ini mempermudah kontrol komunikasi data lewat satu ataupun banyak channel.

### Race conditions

Race conditions adalah di mana 2 utas mengakses memori sekaligus, salah satunya adalah menulis. Race conditions terjadi karena tidak sinkron akses ke memori bersama.
