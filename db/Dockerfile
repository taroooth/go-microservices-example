# 使用するDockerイメージ
FROM mysql:5.7

# タイムゾーンを東京にする
ENV TZ Asia/Tokyo

# MySQLのConfigファイルをイメージ内にコピー
COPY ./config/db.cnf /etc/mysql/conf.d/db.cnf

# docker runに実行
CMD ["mysqld"]

# ポートを開ける
EXPOSE 3306
