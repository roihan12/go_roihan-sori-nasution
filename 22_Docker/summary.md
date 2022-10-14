# (22) Docker

## Resume Materi

### Docker

Docker adalah layanan yang menyediakan kemampuan untuk mengemas dan menjalankan sebuah aplikasi dalam sebuah lingkungan terisolasi yang disebut dengan container. Dengan adanya isolasi dan keamanan yang memadai memungkinkan untuk menjalankan banyak container di waktu yang bersamaan pada host tertentu. container bukan virtualmesin.
container adalah proses dengan sistem file isolasi.

### Docker Basic

Basics dari doker itu terdiri dari:

- image
- container
- engine
- registry
- control plane.

infrastruktur dari docker ada client yang tempat menuliskan perintah untuk menjalankan docker seperti cli atau docker dekstop selanjutnya perintah tersebut dieksekusi oleh docker daemon untuk memanajemen komponen yang ada didalam docker host yaitu image dan juga container, yang terakhir adalah komponen registry yang merupakan komponen yang digunakan untuk menyimpan image di docker hub.

### Perbedaan Container dan Virtual machines

#### Container

- Abstraction at the app layer
- Containers take up less space than VMs
- (container images are typically tens of MBs in size)
- Handle more applications and require fewer VMs and Operating systems.

#### Virtual Machines

- Abstraction of physical hardware.
- Each VM includes a full copy of an operating system
- Also be slow to boot.

### Perintah - Perintah Docker

- docker pull => Mendownload image dari registry
- docker images => Melihat daftar image yang tersedia
- docker image rm => Menghapus image
- docker container ls -a => Melihat semua daftar container
- docker cotainer ls => Melihat container yang sedang berjalan
- docker container start => Menjalankan container yang telah dibuat
- docker container rm => Menghapus container
- docker container create => Membuat container dari image

### Dockerfile

Docker dapat membuat image secara otomatis dengan membaca instruksi yang ada pada Dockerfile. Dockerfile adalah dokumen teks yang berisi sebuah perintah untuk membangun sebuah image. Dengan menggunakan perintah docker build pengguna dapat mengeksekusi beberapa instruksi baris perintah secara berurutan.

- FROM : Getting image from docker registry
- RUN : Execute bash command when building container
- ENV : Set variable inside container
- ADD :Copy the file with some other process
- COPY :Copy the file
- WORKDIR : Set working file directory
- ENTRYPOINT : Execute command when finish building container
- CMD : Execute command

### Docker Compose

Docker-Compose adalah alat untuk mendefinisikan dan menjalankan satu atau beberapa container yang saling terkait dengan sebuah command. Pada implementasinya dapat menggunakannya dengan membuat sebuah file berekstensi yaml/yml yang di dalamnya terdapat konfigurasi-konfigurasi terhadap service aplikasi yang akan dijalankan.
