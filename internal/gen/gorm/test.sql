--  用户表 --
CREATE TABLE `user`
(
    `id`          bigint(10) unsigned NOT NULL AUTO_INCREMENT primary key,
    `user`        varchar(50)                             NOT NULL DEFAULT '' COMMENT '用户',
    `name`        varchar(255) COLLATE utf8mb4_general_ci NULL COMMENT '用户\t名称',
    `password`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户\n密码',
    `mobile`      varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `gender`      char(5) COLLATE utf8mb4_general_ci      NOT NULL COMMENT '男｜女｜未公\r开',
    `nickname`    varchar(255) COLLATE utf8mb4_general_ci          DEFAULT '' COMMENT '用户昵称',
    `type`        tinyint(1) COLLATE utf8mb4_general_ci DEFAULT 0 COMMENT '用户类型',
    `create_time` timestamp NULL,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `name_index` (`name`),
    UNIQUE KEY `user_index` (`user`),
    UNIQUE KEY `type_index` (`type`),
    UNIQUE KEY `mobile_index` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表' COLLATE=utf8mb4_general_ci;

-- fn: Case1
select * from user where id = 1 limit 1;
-- fn: Case2
select * from user where id = ? limit 1;
-- fn: Case3
select * from user where name = '' limit 1;
-- fn: Case4
select * from user where name = 'foo' limit 1;
-- fn: Case5
select * from user where name = ? limit 1;
-- fn: Case6
select id,name from user where id = 1 limit 1;
-- fn: Case7
select id,name from user where id > 1 limit 10;
-- fn: Case8
select id,name from user where id > 1 limit 1,10;
-- fn: Case9
select id,name from user where id > 1 group by name;
-- fn: Case10
select id,name from user where id > 1 group by name having id > 1;
-- fn: Case11
select id,name from user where id > 1 group by name having id > 1 order by id desc;
-- fn: Case12
select id,name from user where id > 1 group by name having id > 1 order by id asc;
-- fn: Case13
select id,name from user where id > 1 group by name having id > 1 order by id desc limit 1,10;
-- fn: Case14
select count(1) AS count from user;
-- fn: Case15
select count(*) AS count from user;
-- fn: Case16
select count(id) AS count from user;
-- fn: Case17
select count(id) AS count from user where id > ?;
-- fn: Case18
select max(id) AS maxID from user;
-- fn: Case19
select max(id) AS maxID from user where id > ?;
-- fn: Case20
select avg(id) AS avgID from user;
-- fn: Case21
select id,name,count(id) AS count from user;