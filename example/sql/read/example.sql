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

-- example1: find one by primary key
-- if you want to find one result, you have to explicitly declare limit 1 statement.
-- fn: FindOne
select *
from user
where id = ? and name in (?,?,?) limit 1;

-- example2: find one by unique key
-- fn: FindByName
select *
from user
where name = ? limit 1;

-- example3: find part of fields by primary key
-- fn: FindOnePart
select id, name, nickname
from user
where id = ? limit 1;

-- example4: find part of fields by unique key
-- fn: FindByNamePart
select id, name, nickname
from user
where name = ? limit 1;

-- example5: find all
-- fn: FindAll
select *
from user;

-- example6: find all count, if call function, you must use AS keyword to alias result.
-- fn: FindAllCount
select count(*) AS count
from user;

-- example7: find all part of fields
-- fn: FindAllPart
select id, name, nickname
from user;

-- example8: find all part of fields count, if call function, you must use AS keyword to alias result.
-- fn: FindAllPartCount
select count(id) AS count
from user;

-- example9: find one by name and password
-- fn: FindOneByNameAndPassword
select *
from user
where name = ?
  and password = ? limit 1;

-- example10: list user by primary key, group by name
-- fn: ListUserByNameAsc
select *
from user
where id > ?
group by name;

-- example11: list user by primary key, group by name asc, having count(type) > ?
-- having clause must be a alias, do not use function expression, for example:
-- select * from user where id > ? group by name asc having count(type) > ?; in this
-- statement, count(type) is a function expression, it will not work, you can use
-- select *,count(type) AS typeCount from user where id > ? group by name asc having typeCount > ?; instead.
-- fn: ListUserByNameAscHavingCountTypeGt
select *, count(type) AS typeCount
from user
where id > ?
group by name
having typeCount > ?;

-- example13: list user by primary key, group by name desc, having count(type) > ?, order by id desc
-- fn: ListUserByNameDescHavingCountTypeGtOrderByIdDesc
select *, count(type) AS typeCount
from user
where id > ?
group by name
having typeCount > ?
order by id desc;

-- example14: list user by primary key, group by name desc, having count(type) > ?, order by id desc, limit 10
-- fn: ListUserByNameDescHavingCountTypeGtOrderByIdDescLimit10
select *, count(type) AS typeCount
from user
where id > ?
group by name
having typeCount > ?
order by id desc limit 10;

-- example15: list user by primary key, group by name desc, having count(type) > ?, order by id desc, limit 10, 10
-- fn: ListUserByNameDescHavingCountTypeGtOrderByIdDescLimit10Offset10
select *, count(type) AS typeCount
from user
where id > ?
group by name
having typeCount > ?
order by id desc limit 10, 10;

-- example16: find one by name like
-- fn: FindOneByNameLike
select *
from user
where name like ? limit 1;

-- example17: find all by name not like
-- fn: FindAllByNameNotLike
select *
from user
where name not like ?;

-- example18: find all by id in
-- fn: FindAllByIdIn
select *
from user
where id in (?);

-- example19: find all by id not in
-- fn: FindAllByIdNotIn
select *
from user
where id not in (?);

-- example20: find all by id between
-- fn: FindAllByIdBetween
select *
from user
where id between ? and ?;

-- example21: find all by id not between
-- fn: FindAllByIdNotBetween
select *
from user
where id not between ? and ?;

-- example22: find all by id greater than or equal to
-- fn: FindAllByIdGte
select *
from user
where id >= ?;

-- example23: find all by id less than or equal to
-- fn: FindAllByIdLte
select *
from user
where id <= ?;

-- example24: find all by id not equal to
-- fn: FindAllByIdNeq
select *
from user
where id != ?;

-- example25: find all by id in, or, not in
-- fn: FindAllByIdInOrNotIn
select *
from user
where id in (?)
   or id not in (?);

-- example26: complex query
-- fn: ComplexQuery
select *
from user
where id > ?
  and id < ?
  and id != ?
  and id in (?)
  and id not in (?)
  and id between ? and ?
  and id not between ? and ?
  and id >= ?
  and id <= ?
  and id != ?
  and name like ?
  and name not like ?
  and name in (?)
  and name not in (?)
  and name between ?
  and ? and name not between ? and ?
  and name >= ?
  and name <= ?
  and name != ?;
