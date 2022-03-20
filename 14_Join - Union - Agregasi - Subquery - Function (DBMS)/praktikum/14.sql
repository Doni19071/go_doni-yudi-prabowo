insert into operators (id, nama, updated_at, created_at) values ('001', 'Operator 1', '2022-03-16 11:58:00','2022-03-16 11:58:00');
insert into operators (id, nama, updated_at, created_at) values ('002', 'Operator 2', '2022-03-16 11:59:00','2022-03-16 11:59:00');
insert into operators (id, nama, updated_at, created_at) values ('003', 'Operator 3', '2022-03-16 11:59:30','2022-03-16 11:59:30');
insert into operators (id, nama, updated_at, created_at) values ('004', 'Operator 4', '2022-03-16 12:00:30','2022-03-16 12:00:30');
insert into operators (id, nama, updated_at, created_at) values ('005', 'Operator 5', '2022-03-16 12:01:30','2022-03-16 12:01:30');

insert into product_type (id, nama, updated_at, created_at) values ('001', 'Type 1', '2022-03-16 12:05:30','2022-03-16 12:05:30');
insert into product_type (id, nama, updated_at, created_at) values ('002', 'Type 2', '2022-03-16 12:07:30','2022-03-16 12:07:30');
insert into product_type (id, nama, updated_at, created_at) values ('003', 'Type 3', '2022-03-16 12:10:30','2022-03-16 12:10:30');

insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('1', '1', '3', 'Produk 1', '10', '2022-03-16 12:30:00','2022-03-16 12:30:00', 'P001');
insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('2', '1', '3', 'Produk 2', '10', '2022-03-16 12:31:00','2022-03-16 12:31:00', 'P002');
insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('3', '2', '1', 'Produk 3', '10', '2022-03-16 12:33:00','2022-03-16 12:33:00', 'P003');
insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('4', '2', '1', 'Produk 4', '10', '2022-03-16 12:35:00','2022-03-16 12:35:00', 'P004');
insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('5', '2', '1', 'Produk 5', '10', '2022-03-16 12:35:00','2022-03-16 12:35:00', 'P005');
insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('6', '3', '4', 'Produk 6', '10', '2022-03-16 12:35:00','2022-03-16 12:35:00', 'P006');
insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('7', '3', '4', 'Produk 7', '10', '2022-03-16 12:37:00','2022-03-16 12:37:00', 'P007');
insert into product (id, product_type_id, operator_id, nama, status, updated_at, created_at, code) values ('8', '3', '4', 'Produk 8', '10', '2022-03-16 12:37:00','2022-03-16 12:37:00', 'P008');

insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('1', 'Produk 1 ini kualitas oke','2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('2', 'Produk 2 ini kualitas bagus','2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('3', 'Produk 3 terjamin','2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('4', 'Produk 4 sangat terjamin','2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('5', 'Produk 5 mantap','2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('6', 'Produk 6 sangat berkualitas','2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('7', 'Produk 7 kualitas nomer 1','2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into product_deskripsi (id, deskripsi, updated_at, created_at) values ('8', 'Produk 8 paling laris','2022-03-16 12:38:00','2022-03-16 12:38:00');

insert into payment_method (id, nama, status, updated_at, created_at) values ('1', 'Transfer Bank', '5', '2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into payment_method (id, nama, status, updated_at, created_at) values ('2', 'Tunai', '1', '2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into payment_method (id, nama, status, updated_at, created_at) values ('3', 'QR Code', '4', '2022-03-16 12:38:00','2022-03-16 12:38:00');

insert into user (id, tanggal_lahir, status, gender, created_at, updated_at) values ('1', '2000-01-01', 'aktif', 'laki-laki', '2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into user (id, tanggal_lahir, status, gender, created_at, updated_at) values ('2', '2001-04-11', 'aktif', 'laki-laki', '2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into user (id, tanggal_lahir, status, gender, created_at, updated_at) values ('3', '2000-06-29', 'aktif', 'perempuan', '2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into user (id, tanggal_lahir, status, gender, created_at, updated_at) values ('4', '2000-02-28', 'aktif', 'laki-laki', '2022-03-16 12:38:00','2022-03-16 12:38:00');
insert into user (id, tanggal_lahir, status, gender, created_at, updated_at) values ('5', '2000-11-15', 'aktif', 'perempuan', '2022-03-16 12:38:00','2022-03-16 12:38:00');

insert into transaksi (id, user_id, status, total_qty, total_price, updated_at, created_at) values ('1', '1', 'terkirim', '3', '30000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into transaksi (id, user_id, status, total_qty, total_price, updated_at, created_at) values ('2', '1', 'terkirim', '3', '30000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into transaksi (id, user_id, status, total_qty, total_price, updated_at, created_at) values ('3', '1', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into transaksi (id, user_id, status, total_qty, total_price, updated_at, created_at) values ('4', '2', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into transaksi (id, user_id, status, total_qty, total_price, updated_at, created_at) values ('5', '2', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into transaksi (id, user_id, status, total_qty, total_price, updated_at, created_at) values ('6', '2', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into transaksi (id, user_id, status, total_qty, total_price, updated_at, created_at) values ('7', '3', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('8', '3', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('9','3', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('10','4', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('11', '4', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('12','4', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('13','5', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('14','5', 'terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('15','5','terkirim', '3', '32000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');

insert into detail_transaksi (transaksi_id, product_id, status, qty, price, updated_at, created_at) values ('1','1', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('1','2', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('1','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into detail_transaksi (transaksi_id, product_id, status, qty, price, updated_at, created_at) values ('2','2', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('2','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('2','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'), ('3','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('3','1', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('3','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('4','7', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('4','5', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('4','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('5','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('5','1', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('5','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('6','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('6','2', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('6','4', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('7','7', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('7','5', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('7','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('8','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('8','1', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('8','2', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');
insert into detail_transaksi (transaksi_id, product_id, status, qty, price, updated_at, created_at) values ('9','7', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('9','6', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('9','5', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'), ('10','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('10','2', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('10','4', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('11','7', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('11','1', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('11','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('12','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('12','2', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('12','6', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('13','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('13','5', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('13','4', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('14','7', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('14','5', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('14','3', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('15','8', 'terkirim', '1', '12000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('15','1', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00'),('15','6', 'terkirim', '1', '10000', '2022-03-16 11:45:00', '2022-03-16 11:45:00');

select id from user where gender = 'laki-laki';
select id from user where id = '3';
select sum(gender) as 'Total Perempuan' from user where gender ='perempuan';
select id, product_type_id, operator_id, nama, status from product where product_type_id between 1 and 2;
update user set nama = 'Produk Dummy' where id = '1';
update detail_transaksi set qty = '3' where product_id = '1';
delete from product where id = '1';
delete from product where product_type_id = '1';


