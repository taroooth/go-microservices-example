FROM mysql:8.0

ENV TZ Asia/Tokyo

EXPOSE 3306

COPY ./config/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
