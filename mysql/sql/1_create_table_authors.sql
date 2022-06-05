use order_development;

create table if not exists authors
(
  id      int unsigned not null primary key auto_increment,
  name    varchar(128) not null
);

insert into `authors` (`name`) VALUES ('Dustin Boswell');
insert into `authors` (`name`) VALUES ('増田 亨');
insert into `authors` (`name`) VALUES ('ミック');
