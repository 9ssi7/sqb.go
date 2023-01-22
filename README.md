<p align="center"><br><img src="https://avatars.githubusercontent.com/u/76786120?v=4" width="128" height="128" style="border-radius: 50px;" /></p>
<h3 align="center">sqb.go</h3>
<p align="center">
  Sql Query Builder for Go
</p>

### What is it?

sqb.go is a sql query builder for Go. It is inspired
by [ssibrahimbas/query typescript package](https://github.com/ssibrahimbas/query).

### Motivation

Generally, adding parameters to a query is a tedious task. You have to write a lot of code to add parameters to a query.
This package aims to make it easier to add parameters to a query.

### Installation

```bash
go get gitlab.com/ssibrahimbas/sqb.go
```

### Usage

```go
package main

import (
	"fmt"

	"gitlab.com/ssibrahimbas/sqb.go"
)

func main() {
	query := sqb_go.QB.Table("my_table").GetAll()
	fmt.Println(query) // SELECT * FROM my_table
}
```

### Contributing

Contributions are always welcome!

### License

[MIT](https://choosealicense.com/licenses/mit/)

### Documentation

<docgen-index>

| Functions                                    |
|----------------------------------------------|
| **[`Table(...)`](#table)**                   |
| **[`Select(...)`](#select)**                 |
| **[`GroupConcat(...)`](#groupconcat)**       |
| **[`Least(...)`](#least)**                   |
| **[`Max(...)`](#max)**                       |
| **[`Min(...)`](#min)**                       |
| **[`Sum(...)`](#sum)**                       |
| **[`Count(...)`](#count)**                   |
| **[`Avg(...)`](#avg)**                       |
| **[`InnerJoin(...)`](#innerjoin)**           |
| **[`LeftJoin(...)`](#leftjoin)**             |
| **[`RightJoin(...)`](#rightjoin)**           |
| **[`FullOuterJoin(...)`](#fullouterjoin)**   |
| **[`KeftOuterJoin(...)`](#leftouterjoin)**   |
| **[`RightOuterJoin(...)`](#rightOuterjoin)** |
| **[`Where(...)`](#where)**                   |
| **[`OrWhere(...)`](#orwhere)**               |
| **[`NotWhere(...)`](#notwhere)**             |
| **[`OrNotWhere(...)`](#ornotwhere)**         |
| **[`WhereNull(...)`](#wherenull)**           |
| **[`WhereNotNull(...)`](#wherenotnull)**     |
| **[`Grouped(...)`](#grouped)**               |
| **[`In(...)`](#in)**                         |
| **[`NotIn(...)`](#notin)**                   |
| **[`OrIn(...)`](#orin)**                     |
| **[`OrNotIn(...)`](#ornotin)**               |
| **[`FindInSet(...)`](#findinset)**           |
| **[`NotFindInSet(...)`](#notfindinset)**     |
| **[`OrFindInSet(...)`](#orfindinset)**       |
| **[`OrNotFindInSet(...)`](#ornotfindinset)** |
| **[`Between(...)`](#between)**               |
| **[`NotBetween(...)`](#notbetween)**         |
| **[`OrBetween(...)`](#orbetween)**           |
| **[`OrNotBetween(...)`](#ornotbetween)**     |
| **[`Like(...)`](#like)**                     |
| **[`OrLike(...)`](#orlike)**                 |
| **[`NotLike(...)`](#notlike)**               |
| **[`OrNotLike(...)`](#ornotlike)**           |
| **[`Limit(...)`](#limit)**                   |
| **[`Pagination(...)`](#pagination)**         |
| **[`OrderBy(...)`](#orderby)**               |
| **[`GroupBy(...)`](#groupby)**               |
| **[`Having(...)`](#having)**                 |
| **[`FromQuery(...)`](#fromQuery)**           |
| **[`GetQuery(...)`](#getQuery)**             |
| **[`Get()`](#get)**                          |
| **[`GetAll()`](#getall)**                    |
| **[`Insert(...)`](#insert)**                 |
| **[`InsertMany(...)`](#insertMany)**         |
| **[`Update(...)`](#update)**                 |
| **[`UpdateMany(...)`](#updateMany)**         |
| **[`Delete()`](#delete)**                    |
| **[`Analyze()`](#analyze)**                  |
| **[`Check()`](#check)**                      |
| **[`Checksum()`](#checksum)**                |
| **[`Optimize()`](#optimize)**                |
| **[`Repair()`](#repair)**                    |
| **[`Reset()`](#reset)**                      |

</docgen-index>

<docgen-api>

### Table

▸ **Table**(`tables`: ...string): Sqb

#### Parameters:

| Name     | Type      | Description                   |
|----------|-----------|-------------------------------|
 | `tables` | ...string | You can pass multiple tables. |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").GetAll()
fmt.Println(query) // SELECT * FROM my_table
```

### Select

▸ **Select**(`columns`: ...string): Sqb

#### Parameters:

| Name      | Type      | Description                   |
|-----------|-----------|-------------------------------|
 | `columns` | ...string | You can pass multiple columns. |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("id", "name").GetAll()
fmt.Println(query) // SELECT id, name FROM my_table
```

### GroupConcat

▸ **GroupConcat**(`column`: string): Sqb

#### Parameters:

| Name     | Type   | Description |
|----------|--------|-------------|
 | `column` | string |             |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("id", "name").GroupConcat("name").GetAll()
fmt.Println(query) // SELECT id, name, GROUP_CONCAT(name) FROM my_table
```

### Least

▸ **Least**(`column`: string): Sqb

#### Parameters:

| Name     | Type   | Description |
|----------|--------|-------------|
 | `column` | string |             |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("id", "name").Least("name").GetAll()
fmt.Println(query) // SELECT id, name, LEAST(name) FROM my_table
```

### Max

▸ **Max**(`column`: string): Sqb

#### Parameters:

| Name     | Type   | Description |
|----------|--------|-------------|
 | `column` | string |             |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("id", "name").Max("name").GetAll()
fmt.Println(query) // SELECT id, name, MAX(name) FROM my_table
```

### Min

▸ **Min**(`column`: string): Sqb

#### Parameters:

| Name     | Type   | Description |
|----------|--------|-------------|
 | `column` | string |             |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("id", "name").Min("name").GetAll()
fmt.Println(query) // SELECT id, name, MIN(name) FROM my_table
```

### Sum

▸ **Sum**(`column`: string): Sqb

#### Parameters:

| Name     | Type   | Description |
|----------|--------|-------------|
 | `column` | string |             |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("id", "name").Sum("price").GetAll()
fmt.Println(query) // SELECT id, name, SUM(price) FROM my_table
```

### Count

▸ **Count**(`column`: string): Sqb

#### Parameters:

| Name     | Type   | Description |
|----------|--------|-------------|
 | `column` | string |             |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("name").Count("id").GetAll()
fmt.Println(query) // SELECT name, COUNT(id) FROM my_table
```

### Avg

▸ **Avg**(`column`: string): Sqb

#### Parameters:

| Name     | Type   | Description |
|----------|--------|-------------|
 | `column` | string |             |

#### Returns:

Sqb

#### Example

```go
query := sqb_go.QB.Table("my_table").Select("name").Avg("price").GetAll()
fmt.Println(query) // SELECT name, AVG(price) FROM my_table
```