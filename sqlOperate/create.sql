Create Database EME;

Create Table informationOfUsers(
id   varchar (11) primary key not null,
name varchar(12) not null unique,
sfp  varchar(50) not null,
passWord varchar(12) not null
)

create table verificationcode(
	id varchar(11)  primary key,
	code varchar(16) not null
)

create table datas(
	id varchar(11),
	fileName varchar(100),
	tag      varchar(50),
	time     varchar(70),
	describe text,
	information text
)

create table Manager(
     name varchar(100),
     password varchar(80),
     account varchar(80)
)

create table  ManagerVerificationcode(
  id varchar(11)  primary key,
	code varchar(16) not null
)

create table Notice(
	title  varchar(100),
	content text,
	pTime   varchar(200)
)


create table filesM(
  title  varchar(100),
  path   varchar (100),
  tag    varchar (40),
  pTime  varchar (100),
)


