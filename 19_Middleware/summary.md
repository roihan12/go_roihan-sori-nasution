# (19) Middleware

## Resume Materi

### Middleware

Middleware adalah sebuah entitas yang dipasang didalam proses server. proses ini nanti ketika client memberikan sebuah request dan ketika server memberikan response nanti akan dipasangkan sebuah layer atau middleware. middleware ini akan berisi fungsi fungsi khusus untuk membantu komunikasi data antara sisi client dan sisi server.
Beberapa Third Party Middleware

- Negroni
- Echo
- Interpose
- Alice

### Implementation Middleware

- LoggingMiddleware: Mencatat log setiap request yang mengenai REST API yang keluar masuk server.

- SessionMiddleware: Menentukan session dari client dan Validasi sesi pengguna dan jaga komunikasi tetap hidup.

- AuthMiddleware : Otentikasi pengguna

- CustomMiddleware : Untuk custom middleware.

### Setup Middleware Echo

#### Echo # Pre ()

Echo # Pre () adalah setup middleware yang akan di executed sebelum routenya masuk ke controller.

- HTTPSRedirect
- HTTPSWWWRedirect
- WWwRedirect
- AddTrailingSlash
- Remove TrailingSlash
- MethodOverride
- Rewrite

#### Echo # Use ()

Echo # Use () adalah setup middleware yang akan di executed setelah mendapatkan akses echo.Context API

- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- academy
- CORS
- Static

### Logger Middleware

Mencatat informasi - informasi yang ada dalam server terutama pada HTTP request
Data log tersebut dapat dilakukan analisis contohnya untuk analisis bisnis, security

### Authentication

- User Identification
- Secure Data and Information

### Basic Authentication

Basic Auth adalah permintaan teknik otentikasi http, metode ini memerlukan informasi nama pengguna dan kata sandi untuk dimasukkan ke dalam header permintaan.

### Jwt (JSON Web Token)

token ini menggunakan JSON (Javascript Object Notation) berbentuk string panjang yang sangat random, lalu token ini memungkinkan kita untuk mengirimkan data yang dapat diverifikasi oleh dua pihak atau lebih.
