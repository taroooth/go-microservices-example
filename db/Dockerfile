FROM mysql:5.7

ENV TZ Asia/Tokyo

COPY ./config/db.cnf /etc/mysql/conf.d/db.cnf

CMD ["mysqld"]

EXPOSE 3306
