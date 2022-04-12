# 22 Middleware
## Resume

Dalam materi ini, mempelajari:
  1. Middleware 
  2. Auth Middleware
  3. Basic & JWT Middleware
  
## Middleware
Middleware merupakan entitas yang menghubungkan kedalam proses request atau respons pada suatu server. Middleware memungkinkan untuk mengimplementasikan bermacam fungsionalitas disekitar tangkapan HTTP request ke server dan mengeluarkan respons

## Auth Middleware
Auth Middleware merupakan salah satu jenis middleware yang dibahas pada materi kali ini. Auth middleware ini secara singkatnya berisi authentikasi yang berfungsi untuk user identification dan mengamankan informasi serta data yang tersimpan. Auth Middleware ini dibagi menjadi 2 pada bahasan kali ini yaitu Basic Auth dan JWT Auth.

## Basic & JWT Middleware
Basic Auth sendiri merupakan teknik autentikasi request dari http, metode ini membutuhkan nama dan password untuk dimasukkan kedalam request header. Basic auth ini sendiri bisa dibilang cukup mudah diakses pihak lain dikarenakan metode encode/decode nya yang hanya menggunakan base64. Selanjutnya ada JWT yang merupakan singkatan dari JSON Web Token yang fungsinya sama dengan basic auth, perbedaan dari keduanya adalah pada JWT ini menggunakan secret key yang harus diinputkan ketika ingin mengakses fungsi yang memerlukan authentication.


