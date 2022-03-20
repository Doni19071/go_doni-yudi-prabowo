# 16 MongoDB  Advanced Query - Array - Aggregation
## Resume

Dalam materi ini, mempelajari:
  1. Advanced Query
  2. Array
  3. Aggregation
  
## Advanced Query
Pada mongoDB kita dapat memberikan math logic pada code dengan beberapa cara diantaranya $ne yang artinya not equal, $gt yang artinya greater than, $lt yang artinya less tha, atau jika ingin ada sama dengan (equal) tinggal ditambah menjadi $gte dan $lte. Dan untuk logic ada $or, $and, $in, $nin, atau $not. Adapula evaluator berupa $regex.

## Array
Untuk array pada mongoDB ini terdapat beberapa fungsi. Untuk fungsi find kita dapat menggunakan $all yang akan menampilkan semua array yang memiliki kondisi tersebut, $size yang akan menampilkan data dengan length sesuai size yang ditulis, dan $slice. Untuk  fungsi update ada $push untuk menambahkan detail data kedalam data yang sudah ada dan $pop untuk menghapus detail data yang dipilih.

## Aggregation
Sama halnya pada materi sebelumnya, agregasi disini berfungsi untuk menggabungkan data. Fungsi-fungsi agregasi diantaranya $lookup:join, $match, $project, $cond, $group, $unwind, dan $sort.