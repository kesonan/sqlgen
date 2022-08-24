CREATE TABLE `foo`
(
    `id`   bigint(10) unsigned NOT NULL AUTO_INCREMENT primary key,
    `name` varchar(255) COLLATE utf8mb4_general_ci NULL,
        UNIQUE KEY `name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- fn: Count
select count(id) AS count from foo;

-- fn: FindOne
select name, count(id) AS c from foo where id > ? having c > ? limit 1;