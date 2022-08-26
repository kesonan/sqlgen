--  用户表 --
CREATE TABLE `user`
(
    `id`          bigint(10) unsigned NOT NULL AUTO_INCREMENT primary key,
    `name`        varchar(255) COLLATE utf8mb4_general_ci NULL COMMENT '用户\t名称',
    `password`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户\n密码',
    `mobile`      varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `gender`      char(5) COLLATE utf8mb4_general_ci      NOT NULL COMMENT '男｜女｜未公\r开',
    `nickname`    varchar(255) COLLATE utf8mb4_general_ci          DEFAULT '' COMMENT '用户昵称',
    `type`        tinyint(1) COLLATE utf8mb4_general_ci DEFAULT 0 COMMENT '用户类型',
    `create_time` timestamp NULL,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `name_index` (`name`),
    UNIQUE KEY `type_index` (`type`),
    UNIQUE KEY `mobile_index` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表' COLLATE=utf8mb4_general_ci;

-- fn: findLimit
select * from user where id > ? group by name having id > ? order by id desc limit ?,?;

-- fn: case1
select count(mobile) as mobileCount, count(1) as count, id,name from user;

-- fn: case2
select * from user where name in (?) and id between ? and ? or (mobile = ?) and nickname like ?;

-- fn: count
select count(id) as count from user;

-- fn: test
start transaction;
-- fn: foo1
select * from user where id = 1 ;
-- fn: foo2
select * from user where id = 2;
commit;

-- fn: test2
start transaction;
-- fn: foo3
update user set name = ? where id = ?;
-- fn: foo4
update user set nickname = ? where id = ?;
commit ;

-- fn: deleteUser
delete from user;


