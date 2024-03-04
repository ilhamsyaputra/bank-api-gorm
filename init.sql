-- buat database
create database bank;

-- install uuid module
create extension if not exists "uuid-ossp";

-- buat table
-- buat table
create table nasabah (
	no_nasabah varchar(30) not null,
	nama varchar(100) not null,
	nik varchar(55) not null,
	no_hp varchar(14) not null,
	pin text not null,
	kode_cabang varchar(5) not null,
	tanggal_registrasi timestamp default now(),
	
	primary key(no_nasabah)
);

create table enums (
    scope varchar(20) not null,
	value varchar(5) unique not null,
	description varchar(200) not null,
	
	constraint enums_pkey
		primary key(scope, value)
);

create table rekening (
	no_rekening varchar(30) not null,
	no_nasabah varchar(30) not null,
	saldo numeric(38, 2) not null,
	
	primary key(no_rekening),
	constraint no_nasabah
		foreign key(no_nasabah)
			references nasabah(no_nasabah)
);

create table transaksi (
	id uuid not null,
	no_rekening_asal varchar(30) not null,
	no_rekening_tujuan varchar(30) not null,
	tipe_transaksi varchar(5) not null,
	nominal numeric(38, 2) not null,
	waktu_transaksi timestamp default now(),
	
	primary key(id),
	
	constraint fk_no_rekening_asal
		foreign key(no_rekening_asal)
			references rekening(no_rekening),
			
	constraint fk_tipe
		foreign key(tipe_transaksi)
			references enums(value)
);

create table counter (
	name varchar(50),
	value int,
	
	primary key(name)
);

insert into enums (scope, value, description) values 
	('tipe_transaksi', 'D', 'Tarik'),
	('tipe_transaksi', 'C', 'Tabung');

insert into counter (name, value) values
	('No Nasabah', 0),
	('No Rekening', 0);