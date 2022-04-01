create database if not exists order-development character set utf8mb4 collate utf8mb4_bin;

use order-development;

drop table if exists users;
create table if not exists users
(
  id      int unsigned not null primary key auto_increment,
  name    varchar(128) not null,
  api_key varchar(256) not null,
  unique uniq_api_key (api_key)
) character set utf8mb4 collate utf8mb4_bin;

insert into users (name, api_key)
values ('test-user-1', 'secure-api-key-1'),
       ('test-user-2', 'secure-api-key-2'),
       ('test-user-3', 'secure-api-key-3'),
       ('test-user-4', 'secure-api-key-4'),
       ('test-user-5', 'secure-api-key-5'),
       ('test-user-6', 'secure-api-key-6'),
       ('test-user-7', 'secure-api-key-7'),
       ('test-user-8', 'secure-api-key-8'),
       ('test-user-9', 'secure-api-key-9'),
       ('test-user-10', 'secure-api-key-10');