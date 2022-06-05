use order_development;

create table if not exists authors
(
  id      int unsigned not null primary key auto_increment,
  name    varchar(128) not null
);
