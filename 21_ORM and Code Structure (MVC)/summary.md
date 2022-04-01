# 21 ORM and Code Structure (MVC)
## Resume

Dalam materi ini, mempelajari:
  1. ORM 
  2. Kelebihan dan Kekurangan ORM
  3. Code Structure
  
## ORM
ORM atau Object-Relational Map merupakan teknik programming untuk mengubah data antara tipe sistem yang tidak kompatible menggunakan ORM language. Pada golang sendiri juga terdapat ORM yang dikembangkan untuk bahasa Go yang dinamakan GORM, GORM ini memiliki fungsi yang cukup keren yaitu sebagai alat bantuk untuk mempercepat kerja developer. 

## Kelebihan dan Kekurangan ORM
Kelebihan dari ORM ini diantaranya repetitive query yang lebih kecil, secara otomatis mengambil data ke dalam object yang siap digunakan, cara yang simple untuk screening data sebelum disimpan di database, dan beberapa punya fitur cache query yang mempercepat pemrosesan data. Untuk kekurangan dari ORM ini sendiri adalah memperbanyak code yang akan ditulis serta menambah biaya, memuat data hubungan yang tidak perlu, lebih lama dalam penulisan code raw query yang kompleks, dan fungsi SQL yang ditulis belum tentu cocok pada vendor lain. 

## Code Structure
Pada materi ini dilakukan strukturisasi data dengan menggunakan MVC atau Model View Controller. MVC ini populer dan banyak digunakan untuk pengorganisasian data. Konsep dibalik MVC ini sendiri adalah tiap section dari code mempunyai tujuan masing-masing dan tiap tujuan tersebut berbeda-beda


