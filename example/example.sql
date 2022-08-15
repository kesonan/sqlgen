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
    UNIQUE KEY `mobile_index` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'user table' COLLATE=utf8mb4_general_ci;


-- operation: create
-- note: sqlgen will generate only one create function named Create, so the next insert sql statements will be ignored.

-- test case: insert one.
-- fn: CreateOne
insert into `user` (`name`, `password`, `mobile`, `gender`, `nickname`, `type`, `create_at`, `update_at`) values (?, ?, ?, ?, ?, ?, ?, ?);

-- test case: insert partial columns.
-- fn: CreatePart
insert into `user` (`name`, `password`, `mobile`) values (?, ?, ?);




-- operation: update
-- note:  sqlgen will generate update function whose name via the value of `fn:` in the comment.
-- update statement please see https://dev.mysql.com/doc/refman/8.0/en/update.html.

-- test case: update one.
-- fn: Update
update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ?;

-- case: update one with order by desc clause.
-- fn: UpdateOrderByIdDesc
update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ? order by id desc;

-- test case: update one with order by desc, limit count clause.
-- fn: UpdateOrderByIdDescLimitCount
update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ? order by id desc;


-- operation: read
-- note:  sqlgen will generate update function whose name via the value of `fn:` in the comment.
-- select statement please see https://dev.mysql.com/doc/refman/8.0/en/select.html.

-- test case: find one by primary key.
-- note: the expression `limit 1` is necessary in order to get the first record, otherwise it will return multiple records.
-- fn: FindOne
select * from `user` where `id` = ? limit 1;

-- test case: find one by unique key.
-- note: the expression `limit 1` is necessary in order to get the first record, otherwise it will return multiple records.
-- fn: FindOneByName
select * from `user` where `name` = ? limit 1;


-- test case: find one with group by clause.
-- note: the expression `limit 1` is necessary in order to get the first record, otherwise it will return multiple records.
-- fn: FindOneGroupByName
select * from `user` where `name` = ? group by name limit 1;

-- test case: find one with group by desc, having clause.
-- note: the expression `limit 1` is necessary in order to get the first record, otherwise it will return multiple records.
-- fn: FindOneGroupByNameHavingName
select * from `user` where `name` = ? group by name having name = ? limit 1;

-- test case: find all
-- fn: FindAll
select * from `user`;

-- test case: find limit count, offset 0.
-- note: the expression both `limit ?`(unsupported marker likes `$1`) and `limit 10`(count must be gather than 1) can return multiple records. do not use `limit 1` if you want to read multiple records.
-- fn: FindLimit
select * from `user` where id > ? limit ?;

-- test case: find records, with limit count, offset clause.
-- note: the expression both `limit ?, ?`(unsupported marker likes `$1`) and `limit 10, 10`(count must gather than 1, offset must gather than 0) can return multiple records. do not use `limit ?,1` or `limit 10,1` if you want to read multiple records.
-- fn: FindLimitOffset
select * from `user` limit ?, ?;

-- test case: find records, with group by, limit, offset clause.
-- note: the expression both `limit ?, ?`(unsupported marker likes `$1`) and `limit 10, 10`(count must gather than 1, offset must gather than 0) can return multiple records. do not use `limit ?,1` or `limit 10,1` if you want to read multiple records.
-- fn: FindGroupLimitOffset
select * from `user` where id > ? group by name limit ?, ?;

-- test case: find records, with group by, having, limit, offset clause.
-- note: the expression both `limit ?, ?`(unsupported marker likes `$1`) and `limit 10, 10`(count must gather than 1, offset must gather than 0) can return multiple records. do not use `limit ?,1` or `limit 10,1` if you want to read multiple records.
-- fn: FindGroupHavingLimitOffset
select * from `user` where id > ? group by name having id > ? limit ?, ?;

-- test case: find records, with group by, having, order by asc, limit, offset clause.
-- note: the expression both `limit ?, ?`(unsupported marker likes `$1`) and `limit 10, 10`(count must gather than 1, offset must gather than 0) can return multiple records. do not use `limit ?,1` or `limit 10,1` if you want to read multiple records.
-- fn: FindGroupHavingOrderAscLimitOffset
select * from `user` where id > ? group by name having id > ? order by id limit ?, ?;

-- test case: find records, with group by, having, order by desc, limit, offset clause.
-- note: the expression both `limit ?, ?`(unsupported marker likes `$1`) and `limit 10, 10`(count must gather than 1, offset must gather than 0) can return multiple records. do not use `limit ?,1` or `limit 10,1` if you want to read multiple records.
-- fn: FindGroupHavingOrderDescLimitOffset
select * from `user` where id > ? group by name having id > ? order by id desc limit ?, ?;

-- test case: find partial columns.
-- fn: FindOnePart
select `name`, `password`, `mobile` from `user` where id > ? limit 1;

-- test case: built-in function: count.
-- note: AS expression is necessary if you are using built-in function.
-- fn: FindAllCount
select count(id) AS countID from `user`;

-- test case: built-in function: count.
-- note: AS expression is necessary if you are using built-in function.
-- fn: FindAllCountWhere
select count(id) AS countID from `user` where id > ?;

-- test case: built-in function: max
-- note: AS expression is necessary if you are using built-in function.
-- fn: FindMaxID
select max(id) AS maxID from `user`;

-- test case: built-in function: min
-- note: AS expression is necessary if you are using built-in function.
-- fn: FindMinID
select min(id) AS minID from `user`;

-- test case: built-in function: avg
-- note: AS expression is necessary if you are using built-in function.
-- fn: FindAvgID
select avg(id) AS avgID from `user`;


-- operation: delete
-- note:  sqlgen will generate update function whose name via the value of `fn:` in the comment.
-- select statement please see https://dev.mysql.com/doc/refman/8.0/en/delete.html.

-- test case: delete one by primary key.
-- fn: DeleteOne
delete from `user` where `id` = ?;

-- test case: delete one by unique key.
-- fn: DeleteOneByName
delete from `user` where `name` = ?;

-- test case: delete one with order by asc clause.
-- fn: DeleteOneOrderByIDAsc
delete from `user` where `name` = ? order by id;

-- test case: delete one with order by desc clause.
-- fn: DeleteOneOrderByIDDesc
delete from `user` where `name` = ? order by id desc;

-- test case: delete one with order by desc clause, limit clause.
-- fn: DeleteOneOrderByIDDescLimitCount
delete from `user` where `name` = ? order by id desc limit ?;