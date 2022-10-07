# gin-mall
基于 gin+gorm+redis+mysql 读写分离的电商网站(高并发秒杀系统)，包括 JWT 鉴权，CORS跨域，AES 对称加密，引入ELK体系，使用docker容器化部署



Docker 安装 MySQL

``` s

➜  ~ docker search mysql
NAME                            DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                           MySQL is a widely used, open-source relation…   13234     [OK]



➜  ~ docker pull mysql:latest
latest: Pulling from library/mysql
Digest: sha256:e9027fe4d91c0153429607251656806cc784e914937271037f7738bd5b8e7709
Status: Image is up to date for mysql:latest
docker.io/library/mysql:latest


➜  ~ docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=密码 mysql
8d227cb57844817048571ee06efe21bc206c9399ddddc1ac741be0026f38603c


➜  ~ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS          PORTS                               NAMES
8d227cb57844   mysql                  "docker-entrypoint.s…"   45 seconds ago   Up 42 seconds   0.0.0.0:3306->3306/tcp, 33060/tcp   mysql








➜  ~ docker exec -it mysql bash
root@8e6ae0287092:/# mysql -uroot -p
Enter password:

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 17
Server version: 8.0.27 MySQL Community Server - GPL

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
mysql>



use mysql;

ALTER USER 'root'@'%' IDENTIFIED BY '密码' PASSWORD EXPIRE NEVER;

ALTER USER 'root'@'localhost' IDENTIFIED BY '密码';

```