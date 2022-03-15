create database alta_online_shop;
use alta_online_shop;

	create table product(
	id int(10) primary key,
    product_type_id int(10),
    operator_id int(10),
    nama varchar(225),
    status smallint,
    updated_at timestamp,
    created_at timestamp
    );
    
    create table product_type(
    id int(10) primary key,
    nama varchar(225),
    updated_at timestamp,
    created_at timestamp
    );
    
    create table product_deskripsi(
    deskripsi TEXT,
    id INT(10) primary key,
    updated_at timestamp,
    created_at timestamp
    );
    
    create table operators(
    id int(10) primary key,
    nama varchar(225),
    updated_at timestamp,
    created_at timestamp
    );
    
    create table payment_method(
    id int(10) primary key,
    nama varchar(225),
    status smallint,
    updated_at timestamp,
    created_at timestamp
    );
    
    create table transaksi(
    id int(10) primary key,
    user_id int(10),
    status varchar(10),
    total_qty int(11),
    total_price numeric(21,2),
    updated_at timestamp,
    created_at timestamp
    );
    
    create table detail_transaksi(
    transaksi_id int(10),
    product_id int(10),
    status varchar(10),
    qty int(10),
    price numeric(20,2),
    updated_at timestamp,
    created_at timestamp
    );
    
    create table user(
    id int(10) primary key,
    nama varchar(25),
    alamat text,
    tanggal_lahir date,
    status varchar(15),
	gender enum('laki-laki', 'perempuan'),
    created_at datetime,
    pdated_at date
    );
  
    create table kurir(
    field_id int(10) primary key,
    nama varchar(20),
    created_at date,
    updated_at datetime
    );
    alter table kurir add column ongkos_dasar int(10);
    
    alter table kurir rename to shipping;
    
    drop table shipping;
    
    