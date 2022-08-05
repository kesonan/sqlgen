-- 用户表 --
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

-- -- fn: FindOne
-- select * from user where id = ? limit 1;
--
-- -- fn: ListUsers
-- -- select * from user where id > ? limit ?, ?;
-- select * from user where id > ? limit 1, 10;
--
-- -- fn: ListUserCount
-- select count(1) AS count from user where id > ?;
--
-- -- fn: FindByName
-- select * from user where name = ? limit 1;
--
-- -- fn: FindByMobile
-- select * from user where mobile = ? limit 1;
--
-- -- fn: ListUserInRange
-- -- select * from user where create_time > ? and create_time < ? order by create_time desc limit ?,?;
-- select * from user where create_time > ? and create_time < ? order by create_time desc limit 1,10;
--
-- -- fn: ListUserCountInRange
-- select count(1) AS count from user where create_time > ? and create_time < ?;

-- fn: FindMaxID
select max(id) AS maxID from user;

-- -- fn: UpdateUser
-- update user set name = ?,mobile = ?, password = ?, nickname = ?, type = ?, update_time = ? where id = ?;
--
-- -- fn: UpdateByName
-- update user set mobile = ?, password = ?, nickname = ?, type = ?, update_time = ? where name = ?;
--
-- -- fn: CreateOne
-- insert into user (user, name, password, mobile, nickname, type, create_time, update_time) values (?, ?, ?, ?, ?, ?, ?, ?);
--
-- -- fn: DeleteOne
-- delete from user where id = ?;
--
-- -- fn: DeleteByName
-- delete from user where name = ?;