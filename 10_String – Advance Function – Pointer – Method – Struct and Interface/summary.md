# (10) String – Advance Function – Pointer – Method – Struct and Interface

## Resume Materi

### Advance Function

Definisi fungsi sendiri adalah sekumpulan blok kode yang dibungkus dengan nama tertentu. Cara membuat fungsi yaitu dengan menuliskan keyword func , diikuti setelahnya nama fungsi, kurung yang berisikan parameter, dan kurung kurawal untuk membungkus blok kode.Parameter sendiri adalah variabel yang disisipkan pada saat pemanggilan fungsi. Sebuah fungsi bisa dirancang tidak mengembalikan nilai balik (void), atau bisa mengembalikan suatu nilai. Fungsi yang memiliki nilai kembalian, harus ditentukan tipe data nilai baliknya pada saat deklarasi.

Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel. Dengan menerapkan konsep tersebut, kita bisa membuat
fungsi di dalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi. Closure merupakan anonymous function atau fungsi tanpa nama. Biasa dimanfaatkan untuk membungkus suatu proses yang hanya
dipakai sekali atau dipakai pada blok tertentu saja.

### Pointer

Pointer adalah reference atau alamat memori.pointer adalah variabel yang menyimpan alamat memori dari variabel lain.
Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.

Variabel bertipe pointer ditandai dengan adanya tanda asterisk ( \* ). Nilai default variabel pointer adalah nil (kosong).

Ada dua hal penting yang perlu diketahui mengenai pointer:

Mengakses alamat variabel ke pointer caranya dengan menambahkan tanda ampersand ( & ). Metode ini disebut dengan referencing.

Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk ( \* ) . Metode ini disebut dengan dereferencing.

### Struct

Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu.

### Method

Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.

### Interface

Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu. Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah nil
