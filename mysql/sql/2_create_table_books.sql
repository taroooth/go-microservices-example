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

insert into `books` (`title`, `author_id`) VALUES ('The Art of Readable Code', 1);
insert into `books` (`title`, `author_id`) VALUES ('現場で役立つシステム設計の原則', 2);
insert into `books` (`title`, `author_id`) VALUES ('達人に学ぶDB設計 徹底指南書', 3);
