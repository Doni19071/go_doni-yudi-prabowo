# 24 Clean and Hexagonal Architecture
## Resume

Dalam materi ini, mempelajari:
  1. Clean Architecture 
  2. Hexagonal Architecture
  3. Constraint Clean Architecture
  
## Clean Architecture 
Clean Architecture bisa dibilang pemisahan bagian-bagian atau fungsi menggunakan layer. Arsitektur yang baik adalah pemisahan menggunakan lapisan-lapisan untuk membuat modular, scalable, maintainable, dan testable aplikasi. Sama halnya asrisektur pada distrik perumahan, arsitektur pada code juga sangat berfungsi untuk keberlangsungan project yang dibuat.

## Hexagonal Architecture
Sama seperti namanya pada architecture ini digambarkan dengan hexagonal atau segienam dimana antar interface yang satu (API : RestController dan RabbitMQReceiver) dengan interface yang lain (SPI : MongoDBClient dan SQL Client) dipisahkan oleh Domain.

## Constraint Clean Architecture 
Sebelum melakukan design pada clean arsitektur perlu diperhatikan beberapa batasan diantaranya adalah frameworks yang independen, dapat di test tanpa hal-hal lain, UI yang dapat diubah dengan mudah tanpa mengubah sistem, database yang independen, dan independen dari segala macam hal luar.