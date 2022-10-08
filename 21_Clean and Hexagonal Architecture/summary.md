# (21) Clean and Hexagonal Architecture

## Resume Materi

### Software Testing

Clean architecture mempunyai tag line yaitu "the separation of concerns using layers". mission statement nya adalah Arsitektur yang baik adalah pemisahan menggunakan layer untuk membangun modul, terukur, dapat dipelihara, dan aplikasi yang dapat diuji. ini menjadi penting karena integration test akan sulit jika tidak menggunakan itu. Untuk Memperbaikinya maka kita menggunakan Hexagonal Architecture yang dibagi menjadi 3 yaitu API(Rest controller), Domain, dan SPI(Database). Pada dasarnya semuanya akan dibagi berdasarkan dengan concern nya masing masing.

### Kendala sebelum mendesain Clean Architecture adalah :

- Independent of Frameworks : Ini memungkinkan untuk digunakan sebagai kerangka kerja atau
  alat, bukan sebagai keterbatasan kendala.

- Testable : Aturan bisnis dapat mudah diuji.

- Independent of UI : UI bisa berubah dengan mudah, tanpa mengubah sistem .

- Independent of Database : Dapat mengubah database dengan mudah.

- Independent of any external

### Manfaat Clean Architecture

- Struktur standar, jadi mudah untuk mengembangkan dalam proyek,

- Perkembangan lebih cepat dalam jangka panjang

- Menggunakan mock dependensi sehingga mudah dalam dalam tes unit,

- Mudah mengubah penyimpanan memori atau basis data).

### CA Layer

- layer entitas - objek bisnis sebagaimana adanya mencerminkan konsep yang dikelola aplikasi Anda

- Use Case - layer domain - berisi logika bisnis

- Controller - Presentation layer - menyajikan data ke layar dan tangani interaksi pengguna

- Driver - layer data - mengelola data aplikasi misalnya mengambil data dari jaringan, mengelola cache data

### Communicate between domain

- RPC
- Restful
- Messaging
