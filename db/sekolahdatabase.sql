create database sekolah;

CREATE TABLE teacher (
id INT NOT NULL,
first_name varchar(40) NOT NULL,
last_name varchar(40) NOT NULL,
email varchar(40) NOT NULL,
PRIMARY KEY (id));

CREATE TABLE student (
id INT NOT NULL,
first_name varchar(40) NOT NULL,
last_name varchar(40) NOT NULL,
email varchar(40) NOT NULL,
PRIMARY KEY (id));

CREATE TABLE subject (
id INT NOT NULL,
subject_name varchar(40) NOT NULL,
PRIMARY KEY (id));

insert into teacher values 
(1, 'Surya', 'Mentari', 'suryamentari@lalala.com'),
(2, 'Dono', 'Pradana', 'donoprada@lalala.com'),
(3, 'Coki', 'Pardede', 'coki666@lalala.com');

insert into student values 
(1, 'Siskaeee', 'Jane', 'siskaenya3@lalala.com'),
(2, 'Tata', 'Pradita', 'tatapr@lalala.com'),
(3, 'Rigen', 'Rispo', 'rigenpopo@lalala.com');

insert into subject values 
(1, 'Matematika'),
(2, 'Biologi'),
(3, 'Fisika'),
(4, 'Kimia');

select * from teacher;
select * from student;
select * from subject;