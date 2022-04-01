drop table if exists transactions;
create table if not exists transactions
(
  id          int unsigned not null primary key auto_increment,
  user_id     int unsigned not null,
  amount      int          not null,
  description varchar(256) not null,
  CONSTRAINT fk_transactions_users FOREIGN KEY (user_id) REFERENCES users (id)
) character set utf8mb4 collate utf8mb4_bin;
