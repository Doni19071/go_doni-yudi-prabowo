# 11 Concurrent Programming
## Resume

Dalam materi ini, mempelajari:
  1. Sequential, Parallel, and Concurrent
  2. Goroutines
  3. Channel and Select
  
## Sequential, Parallel, and Concurrent
Sequential merupakan sebuah program yang akan menyelesaikan tugasnya sebelum berpindah ke tugas baru , pada parallel program akan dijalankan secara bersama-sama dalam waktu itu juga , sedangkan Concurrent menjalankan banyak program secara sendiri-sendiri. Jikan diibaratkan menjalankan tugas 1, 2, dan 3 maka sequential akan menjalankan tugas 1 dan 2 sebelum menjalankan tugas 3, parallel akan secara bersamaan menjalankan tugas 1,2, dan 3, sedangkan concurrent secara bersamaan menjalankan tugas 1, 2, dan 3 tetapi outputnya berbeda dengan yang parallel jalankan.
## Goroutines
Goroutines atau concurrent execution pada bahasa pemrograman Go, merupakan suatu fungsi atau metode yang dapat menjalankan suatu fungsi atau metode dengan concurrent. Yang perlu diperhatikan ketika menggunakan Goroutines dalam golang yaitu perlu adanya time sleep atau time delay untuk menjalankan fungsinya.
## Channel and Select 
Channel sendiri merupakan objek komunikasi yang digunakan pada goroutines untuk berkomunikasi dengan lainnya. Sedangkan select mempermudah untuk mengontrol data komunikasi dari satu atau banyak channel