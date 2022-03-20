# 14 Join - Union - Agregasi - Subquery - Function (DBMS)
## Resume

Dalam materi ini, mempelajari:
  1. Join
  2. Union dan Agregasi
  3. Subquery
  
## Join
Join merupakan sebuah klausa untuk mengkombinasikan record dari dua tabel atau lebih. Join pada SQL ada 3 yaitu inner, lef, dan right join. Inner join akan mengembalikan baris-baris dari 2 tabel atau lebih yang memnuhi syarat. Sedangkan left join dan right join akan mengembalikan seluruh baris pada tabel sebelah kiri (left join) atau sebelah kanan (right join) yang dikenai kondisi ON dan hanya baris tabel sebelah kanan (left join) dan sebelah kiri (left join) yang memenuhi kondisi. 
## Union dan Agregasi
Hal yang perlu diperhatikan dari union adalah jumlah field yang dikeluarkan/dipanggil harus sama. Sedangkan untuk agregasi merupakan fungsi yang beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal. Fungsi agregasi sendiri yaitu min, max, sum, avg, count, dan having. 
## Subquery
Subquery atau inner query adalah query di dalam query SQL lain. Subquery ini digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil.