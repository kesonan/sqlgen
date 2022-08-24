# sqlgen

[English](README.md) | 中文

[![Go](https://github.com/anqiansong/sqlgen/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/anqiansong/sqlgen/actions/workflows/go.yml)

sqlgen 是一个 SQL 代码生成工具，其支持 **bun**， **gorm**， **sql**， **sqlx**， **xorm** 的代码生成，灵感来自于：

- [go-zero](https://github.com/zeromicro/go-zero)
- [goctl](https://github.com/zeromicro/go-zero/tree/master/tools/goctl)
- [sqlc](https://github.com/kyleconroy/sqlc).


# 安装

```bash
go install github.com/anqiansong/sqlgen@latest
```

# 视频
<iframe width="1512" height="945" src="https://www.youtube.com/embed/Yt5zXerc7Qo" title="sqlgen, a tool to generate gorm,xorm,sqlx,sql,bun code" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>


# 示例

 见 [example](https://github.com/anqiansong/sqlgen/tree/main/example)

#  SQL 查询编写规则
## 1. 函数名称
你可以通过在查询语句上方添加一个单行注释，用 `fn` 关键字来声明一个函数名称，例如：

```sql
-- fn: my_func
SELECT * FROM user;
```

其生成后代码格式为:

```go
func (m *UserModel) my_func (...) {
    ...
}
```

## 2. 查询一条记录
当你只想要查询一条记录的需求时，你必须明确地指定 `limit 1`，sqlgen 通过此表达式来判断当前查询是单记录查询还是多记录查询，例如：

```sql
-- fn: FindOne
select * from user where id = ? limit 1;
```

## 3. 使用 '?' 还是具体值？
在 SQL 查询语句的编写中，你可以用 `?` 来替代一个参数，也可以是具体值，他们最终都会被 sqlgen 转换成一个变量，下列示例中的两个查询是等价的。

> 注意: 此规则不适用于规则 2

```sql
-- fn: FineLimit
select * from user where id = ?;

-- fn: FineLimit
select * from user where id = 1;

```

## 4. SQL 内置函数支持
sqlgen 支持 SQL 内置的聚合函数查询，除此之外的其他函数暂不支持，聚合函数查询的列必须要用 `AS` 来起一个别名，例如：

```sql
-- fn: CountAll
select count(*) as count from user;
```

更多查询示例, 你可以点击 [example.sql](https://github.com/anqiansong/sqlgen/blob/main/example/example.sql) 查看详情.

#  sqlgen 使用步骤
1. 创建一个 SQL 文件
2. 编写 SQL 查询语句，如建表语句、查询语句等
3. 使用 `sqlgen` 工具，生成代码

# 注意
1. 目前只支持 MYSQL 代码生成
3. 不支持多表操作
4. 不支持联表查询