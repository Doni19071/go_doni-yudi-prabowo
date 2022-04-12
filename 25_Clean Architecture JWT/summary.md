# 25 Clean Architecture JWT
## Resume

Dalam materi ini, mempelajari:
  1. Clean Architecture 
  2. Clean Architecture JWT
  3. Constraint Clean Architecture
  
## Clean Architecture 
Clean Architecture bisa dibilang pemisahan bagian-bagian atau fungsi menggunakan layer. Arsitektur yang baik adalah pemisahan menggunakan lapisan-lapisan untuk membuat modular, scalable, maintainable, dan testable aplikasi. Sama halnya asrisektur pada distrik perumahan, arsitektur pada code juga sangat berfungsi untuk keberlangsungan project yang dibuat.

## Clean Architecture JWT
JWT atau JSON Web Token berfungsi sebagai auth, pada ilustrasinya seperti kunci yang hanya dimiliki si pemilik rumah untuk masuk ke dalam rumahnya. Jadi jika ingin mengakses suatu fungsi maka perlu melakukan authentikasi terlebih dulu. Sama halnya seperti code lain, code JWT ini juga perlu disusun sedemikian rupa agar code dapat dilihat, dimaintain, dan diubah dengan mudah.

## Constraint Clean Architecture 
Sebelum melakukan design pada clean arsitektur perlu diperhatikan beberapa batasan diantaranya adalah frameworks yang independen, dapat di test tanpa hal-hal lain, UI yang dapat diubah dengan mudah tanpa mengubah sistem, database yang independen, dan independen dari segala macam hal luar.