Create Database EWE;

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
  path   text ,
  tag    varchar (40),
  pTime  varchar (100)
)

insert into Manager values('wyf','150031','wyf')

insert into ManagerVerificationcode values('wyf','150031')


--为rayRun数据库加载数据
create table rayrun(
 id varchar(11)  primary key,
 data text
)

--RayRun iri 模型
insert into iridata values('PM','2011','10','11','')
insert into iridata values('AM','2011','10','11','')
create table iridata(
    id varchar(11) primary key,
    year varchar(20),
    month varchar(20),
    day   varchar(20),
    data text
)


//新开一个表描述feedback
create table feedback(
     account varchar(11),
     time    varchar(50),
     message text
)


