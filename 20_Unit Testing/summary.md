# (20) Unit Testing

## Resume Materi

### Software Testing

Menurut standar ANSI , software testing adalah sebuah proses menganalisa sistem / software untuk mendeteksi perbedaan antara kondisi fitur sistem saat ini dan fitur yang diinginkan oleh stakeholder / po / pembuat/ product owner requirement ). Tujuan dilakukannya testing adalah unutk mengetahui apakah sistem yang telah dibuat sudah bejalan baik dan tidak muncul sebuah bugs dalam sistem.

### Tujuan Testing

- Preventing regression ( menegah regresi ) : software regression adalah kondisi dimana software yang sebelumnya sudah berjalan baik , menjadi salah dikarenakan adanya perubahan atau penambahan fitur.

- Confidence level in refactoring = > refactoring adalah proses mengubah sistem tanpa mengubah fungsionalitas akhir dari sistem tersebut. Dengan testing yang telah dibuat sebelumnya kita bisa tahu efek dari refactoring itu merubah fungsi yang sudah ada atau tidak

- Better code design : ketika developers sering berlatih membuat test untuk fitur yang mereka buat , mereka akan membuat ethod / function mereka sekecil dan se to the point yang mereka bisa buat , karena mereka akan familiar dengan method atau function mana yang dapat di buat testnya .

- Code documentation : untuk mengetahui secara detail fitur yang dikerjakan seseorang , selain kita dapat melihat langsung dari code nya , kita dapat melihat testing nya , bagaimana seharusnya fitur tersebut bekerja.

- Good scaling scaling adalah proses pengembangan sistem menjadi lebih besar , karena testing menghindari Regression , maka proses development akan lebih cepat karena proses validasi jika ada error pun lebih cepat diketahui

### Level Of Testing

- UI Test : melakukan test secara end to end melalui User Interface sistem yang kita buat

- integration Test: Saat kita melakukan testing dengan cara melakukan hit pada endpoint sebuah api sistem.
  Pengujian integrasi adalah fase dalam pengujian perangkat lunak di mana modul perangkat lunak individu
  digabungkan dan diuji sebagai suatu kelompok . Didalam sebuah endpoint pasti ada beberapa modul / function yang digunakan untuk melakukan proses dari data yang diterima endpoint tersebut sampai mengembalikan response data terhadap client.

- Unit Test : melakukan test pada sebuah method / function tertentu pada sebuah sistem.

#### Framework

framework untuk testing biasanya sudah tersedia tools dan struktur file yang dibutuhkan untuk menjalankan testing dengan efficient. beberapa framework untuk testing yaitu Pada python ada framework pytest atau nose, Pada js ada mocha dan jest, pada go ada go testify

### 2 Pattern testing

2 usual pattern :

1. Centralize your test file inside tests folder.
2. Save test file together with production file.

Keuntungan kita centralize didalam test folder , directory nya lebih clean karena script untuk melakukan test terpisah dengan code yang ingin kita test . Sehingga meminimalisir script test mempengaruhi code

### Test File

Didalam test file biasanya terdapat beberapa test suites. Test suites adalah kumpulan dari beberapa test case. Test cases adalah kumpulan kondisi / scenario yang akan kita jalankan untuk melakukan test ( biasanya berisi input , proses dan output ) . Bagian output inilah yang akan kita pastikan apakah sesuai yang kita inginkan atau tidak .
Test fixture adalah suatu proses / script untuk memastikan bahwa environment yang dipakai untuk melakukan test tetap konsisten , sehingga saat kita melakukan test ulang akan menghasilkan hasil yang sama.

### Yang Tidak perlu dilakukan testing

- 3rd party modules
- Database
- 3rd party api or other external system
- Object class that you have test in other place

### Coverage

Test coverage adalah suatu alat ukur untuk menunjukan source code program bagian mana saja yang di executed saat kita menjalankan suatu test . Test coverage ini penting karena dengan coverage dapat mengukur apakah test yang buat cukup atau tidak untuk megetes logic - logic yang ada pada sistem yang buat.
