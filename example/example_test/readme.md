# example_test

## before test
1. started docker
2. run a mysql container which dsn is `root:mysqlpw@(localhost:55000)`
3. new a schema `test`
4. create a table use the following sql
```sql
CREATE TABLE `user`
(
    `id`          bigint(10) unsigned NOT NULL AUTO_INCREMENT primary key,
    `name`        varchar(255) COLLATE utf8mb4_general_ci NULL COMMENT 'The username',
    `password`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The \n user password',
    `mobile`      varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
    `gender`      char(10) COLLATE utf8mb4_general_ci      NOT NULL COMMENT 'gender,male|female|unknown',
    `nickname`    varchar(255) COLLATE utf8mb4_general_ci          DEFAULT '' COMMENT 'The nickname',
    `type`        tinyint(1) COLLATE utf8mb4_general_ci DEFAULT 0 COMMENT 'The user type, 0:normal,1:vip, for test golang keyword',
    `create_at` timestamp NULL,
    `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `name_index` (`name`),
    UNIQUE KEY `type_index` (`type`),
    UNIQUE KEY `mobile_index` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'user table' COLLATE=utf8mb4_general_ci; 
```
5. clean the test data and set auto_increment to `1`
6. run the test


