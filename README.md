
## V2版本更新说明

V2版本，结构较比V1版本有很大的改动 全部转化成 controller、dao、service 模式，更加符合企业开发

由于整合上传oss和上传到本地，需要在 conf 中进行配置 UploadModel 字段，上传到 oss 则配置 oss，上传本地则配置 local


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


```shell 
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mall_db            |
| mall_db_test       |
| mysql              |
| performance_schema |
| sys                |
| todolist           |
+--------------------+
7 rows in set (0.22 sec)

mysql> use mall_db;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> show tables;
+-------------------+
| Tables_in_mall_db |
+-------------------+
| address           |
| admin             |
| carousel          |
| cart              |
| category          |
| favorite          |
| notice            |
| order             |
| product           |
| product_img       |
| relation          |
| skill_product     |
| skill_product2_mq |
| user              |
+-------------------+
14 rows in set (0.00 sec)

mysql> select * from user;
+----+---------------------+---------------------+------------+-----------+-------+--------------------------------------------------------------+-----------+--------+-----------------------------------------------------+--------------------------+
| id | created_at          | updated_at          | deleted_at | user_name | email | password_digest                                              | nick_name | status | avatar                                              | money                    |
+----+---------------------+---------------------+------------+-----------+-------+--------------------------------------------------------------+-----------+--------+-----------------------------------------------------+--------------------------+
|  1 | 2023-06-19 07:12:55 | 2023-06-19 07:12:55 | NULL       | xi        |       | $2a$12$bSkd66FwPau0vEP0141ZruOh9p9rkjMIjgZ2ncCp21er7ZmSe./vC | ??        | active | http://127.0.0.1:3000/static/imgs/avatar/avatar.png | Zt3uB50jiObAwOoJ/kLpaQ== |
|  2 | 2023-07-20 03:21:01 | 2023-07-20 03:21:01 | NULL       | x         |       | $2a$12$3a76sUYwl8TbdILrCd8sku22567S8D6.J2/.Xo5prf39GVrLp58.i | ??        | active | http://127.0.0.1:3000/static/imgs/avatar/avatar.png | Zt3uB50jiObAwOoJ/kLpaQ== |
|  3 | 2023-07-20 09:11:23 | 2023-07-20 09:11:23 | NULL       | test      |       | $2a$12$u9CYbv7rYD4BnDjdc7bRKeBVui21quziudVGlTUTL5soMvn9rXMc. | ??        | active | http://127.0.0.1:3000/static/imgs/avatar/avatar.png | Zt3uB50jiObAwOoJ/kLpaQ== |
|  4 | 2023-07-28 07:04:48 | 2023-07-28 07:04:48 | NULL       | ??        |       | $2a$12$IeLjHH1sQ1D5kDBHE6622uHLXKbfXSdX4Pro3BKHqJ7PG8m/lddFm | ??        | active | avatar.JPG                                          | wVlreio71PEArDlotVjrJw== |
|  5 | 2023-08-13 07:10:05 | 2023-08-13 07:10:05 | NULL       | test1     |       | $2a$12$NrV9rB1BwdeskRx0GFyLN..3s86GTzVeCchxXaREwQBBxUclPKWuq | ??        | active | avatar.JPG                                          | wVlreio71PEArDlotVjrJw== |
+----+---------------------+---------------------+------------+-----------+-------+--------------------------------------------------------------+-----------+--------+-----------------------------------------------------+--------------------------+
5 rows in set (0.01 sec)
```