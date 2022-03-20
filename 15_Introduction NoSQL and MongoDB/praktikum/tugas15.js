// buat collection
db.createCollection('product_types')
db.createCollection('operators')
db.createCollection('product_deskripsi')
db.createCollection('product')
db.createCollection('payment_method')
db.createCollection('transaksi')
db.createCollection('users')
db.createCollection('detail_transaksi')

//insert 5 operators
db.operators.insert([
    {id : '1', nama : 'operator1', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' },
    {id : '2', nama : 'operator2', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' },
    {id : '3', nama : 'operator3', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' },
    {id : '4', nama : 'operator4', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' },
    {id : '5', nama : 'operator5', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' }

])

//insert 3 product type
db.product_type.insert([
    {id : '1', nama : 'type 1', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' },
    {id : '2', nama : 'type 2', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' },
    {id : '3', nama : 'type 3', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00' },
])

//insert product
db.product.insert([
    {id : '1', product_type_id : '1', operator_id : '3', code : 'P001', nama : 'Produk 1', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '2', product_type_id : '1', operator_id : '3', code : 'P002', nama : 'Produk 2', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '3', product_type_id : '2', operator_id : '1', code : 'P001', nama : 'Produk 1', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '4', product_type_id : '2', operator_id : '1', code : 'P001', nama : 'Produk 1', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '5', product_type_id : '2', operator_id : '1', code : 'P001', nama : 'Produk 1', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '6', product_type_id : '3', operator_id : '4', code : 'P001', nama : 'Produk 1', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '7', product_type_id : '3', operator_id : '4', code : 'P001', nama : 'Produk 1', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '8', product_type_id : '3', operator_id : '4', code : 'P001', nama : 'Produk 1', status : '10',  created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'}
])

//insert deskripsi produk
db.product_deskripsi.insert([
    {id :'1', deskripsi : 'Produk pertama', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'},
    {id :'2', deskripsi : 'Produk kedua', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'},
    {id :'3', deskripsi : 'Produk ketiga', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'},
    {id :'4', deskripsi : 'Produk keempat', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'},
    {id :'5', deskripsi : 'Produk kelima', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'},
    {id :'6', deskripsi : 'Produk keenam', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'},
    {id :'7', deskripsi : 'Produk ketujuh', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'},
    {id :'8', deskripsi : 'Produk kedelapan', created_at :'2022-03-18 11:58:00', updated_at :'2022-03-18 11:58:00'}
])

//insert payment method
db.payment_method.insert([
    {id : '1', nama : 'Transfer Bank', status : '5', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '2', nama : 'Tunai', status : '1', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '3', nama : 'QR Code', status : '4', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'}
])

//insert users
db.users.insert([
    {id : '1', status : 'aktif', dob : '2000-01-01', gender : 'laki-laki', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '2', status : 'aktif', dob : '2001-01-02', gender : 'laki-laki', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '3', status : 'aktif', dob : '2002-01-03', gender : 'perempuan', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '4', status : 'aktif', dob : '2001-01-04', gender : 'laki-laki', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '5', status : 'aktif', dob : '2000-01-05', gender : 'perempuan', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'}
])

//insert 3 transaksi per user
db.transaksi.insert([
    {id : '1', user_id : '1', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '2', user_id : '1', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '3', user_id : '1', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '4', user_id : '2', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '5', user_id : '2', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '6', user_id : '2', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '7', user_id : '3', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '8', user_id : '3', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '9', user_id : '3', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '10', user_id : '4', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '11', user_id : '4', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '12', user_id : '4', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '13', user_id : '5', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '14', user_id : '5', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {id : '15', user_id : '5', payment_method_id : '1', status : 'terikirim', total_qty : '3', total_price : '30000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'}
])

//insert 3 produk per transaksi
db.detail_transaksi.insert([
    {transaksi_id : '1', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '1', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '1', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '2', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '2', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '2', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '3', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '3', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '3', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '4', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '4', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '4', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '5', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '5', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '5', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '6', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '6', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '6', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '7', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '7', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '7', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '8', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '8', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '8', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '9', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '9', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '9', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '10', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '10', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '10', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '11', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '11', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '11', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '12', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '12', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '12', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '13', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '13', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '13', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '14', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '14', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '14', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '15', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '15', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},
    {transaksi_id : '15', product_id : '1', status : 'terkirim', qty : '1', price : '10000', created_at : '2022-03-18 11:58:00', updated_at : '2022-03-18 11:58:00'},

])

//tampilkan user dengan gender laki-laki
db.users.find({gender : 'laki-laki'})
//tampilkan produk dengan id = 3
db.product.find({id : '3'})
//hitung jumlah user perempuan
db.users.count()
db.users.find({gender : 'perempuan'}).count();
//tampilkan data pelanggan sesuai abjad
db.users.find().sort({nama : 1})
//tampilkan 5 data produk
db.product.find().limit(5)

//ubah data product id 1 dengan nama 'product dummy'
db.product.update({id : '1'}, {nama : 'product dummy'})
//ubah qty = 3 pada detail transaksi dengan id 1
db.detail_transaksi.update({id : '1'}, {qty : '3'})

//delete data pada tabel produk dengan id 1
db.product.remove({id : '1'})
//delete pada tabel produk dengan product type id 1
db.product.remove({product_type_id : '1'})