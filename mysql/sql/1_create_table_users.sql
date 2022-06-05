use order_development;

create table if not exists users
(
  id      int unsigned not null primary key auto_increment,
  name    varchar(128) not null,
  memo    varchar(128) not null,
  api_key varchar(256) not null,
  unique uniq_api_key (api_key)
);
