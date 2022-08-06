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

-- example1: update by primary key
-- fn: Update
update user set name = ?, password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where id = ?;

-- example2: update by unique key
-- fn: UpdateByName
update user set password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where name = ?;

-- example3: update part columns by primary key
-- fn: UpdatePart
update user set name = ?, nickname = ? where id = ?;

-- example4: update part columns by unique key
-- fn: UpdatePartByName
update user set name = ?, nickname = ? where name = ?;

-- example5: update name limit ?
-- fn: UpdateNameLimit
update user set name = ? where id > ? limit ?;

-- example6: update name limit ? order by id desc
-- fn: UpdateNameLimitOrder
update user set name = ? where id > ? order by id desc limit ?;

