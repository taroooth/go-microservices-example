use order_development;

create table if not exists books
(
  id      int unsigned not null primary key auto_increment,
  title    varchar(128) not null,
  author_id    int unsigned not null,
  CONSTRAINT fk_author_id
    FOREIGN KEY (author_id)
    REFERENCES authors (id)
    ON DELETE RESTRICT ON UPDATE RESTRICT
);
