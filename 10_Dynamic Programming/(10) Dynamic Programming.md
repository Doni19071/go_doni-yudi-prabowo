# 10 Dynamic Programming
## Resume

Dalam materi ini, mempelajari:
  1. Dynamic Programming
  2. Karakteristik 
  3. Metode
  
## Dynamic Programming
Merupakan suatu algoritma yang digunakan untuk mendapatkan solusi optimal dengan cara memecah masalah tersebut kedalam bentuk yang lebih kecil atau lebih simple (sub problem) dan solusi optimal yang didapat merupakan solusi-solusi dari sub problem itu sendiri.
## Karakteristik 
Karakteristik dari dynamic programming yang pertama yaitu overlapping subproblems atau tiap kita mencari solusi optimal dari suatu masalah maka sama saja dengan kita mencari solusi subproblems dari masalah yang sama beberapa kali. Yang kedua yaitu optimal substructure property atau tiap masalah punya propertu substruktur yang optimal jika semua solusi optimal dapat dibangun dari solusi optimal tiap-tiap subproblem .
## Metode 
Metode pertama yang dapat digunakan yaitu Top-down with Memoization, pada metode ini kita akan melakukan pendekatan dengan cara menyelesaikan masalah yang besar dengan metode recursive mencari solusi-solusi dari subproblemnya atau dapat dikatakan urutan langkahnya adalah dari atas atau nilai yang dituju (n) bergerak ke bawah (0). Metode yang kedua yaitu Bottom-up with Tabulation, metode ini berlawanan dengan metode sebelumnya yaitu secara garis besar langkahnya adalah dari bawah (0) bergerak ke nilai yang dituju (n). Dengan cara menyelesaikan subproblem sebelumnya maka kita juga akan mendapat solusi dari masalah yang nilainya kita cari.