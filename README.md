# sqlgen

English | [中文](README_cn.md)

[![Go](https://github.com/anqiansong/sqlgen/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/anqiansong/sqlgen/actions/workflows/go.yml)

sqlgen is a tool to generate **bun**, **gorm**, **sql**, **sqlx** and **xorm** sql code from SQL file which is inspired by 
- [go-zero](https://github.com/zeromicro/go-zero)
- [goctl](https://github.com/zeromicro/go-zero/tree/master/tools/goctl)
- [sqlc](https://github.com/kyleconroy/sqlc).

# Installation

```bash
go install github.com/anqiansong/sqlgen@latest
```

# Example

See [example](https://github.com/anqiansong/sqlgen/tree/main/example)

# Queries rule
## 1. Function Name
You can define a function via `fn` keyword in line comment, for example:

```sql
-- fn: my_func
SELECT * FROM user;
```

it will be generated as:

```go
func (m *UserModel) my_func (...) {
    ...
}
```

## 2. Get One Record
The expression `limit 1` must be explicitly defined if you want to get only one record, for example:

```sql
-- fn: FindOne
select * from user where id = ? limit 1;
```

## 3. Marker or Values?
For arguments of SQL, you can use `?` or explicitly values to mark them, in sqlgen, the arguments will be converted into variables, for example, the following query are equivalent:

> NOTES: It does not apply to rule 2

```sql
-- fn: FineLimit
select * from user where id = ?;

-- fn: FineLimit
select * from user where id = 1;

```

## 4. SQL Function
sqlgen supports aggregate function queries in sql, other than that, other functions are not supported so far. All the aggregate function query expressions must contain AS expression, for example:

```sql
-- fn: CountAll
select count(*) as count from user;
```

For most query cases, you can see [example.sql](https://github.com/anqiansong/sqlgen/blob/main/example/example.sql) for details.

# How it works
1. Create a SQL file
2. Write your SQL code in the SQL file
3. Run `sqlgen` to generate code

# Notes
1. Only support MYSQL code generation.
3. Do not support multiple tables in one SQL file.
4. Do not support join query.