+++
title = 'gorp と squirrel で CRUD'
tags = ['go', 'mysql']
date = '2016-01-17T15:50:42+09:00'
+++

[これ]({{< ref "/blog/2016/01/15/january.md" >}}) の 2 つ目のやつの API サーバーを Go で書くので、DB 周りどうしようかなあと考えながら [gorp](https://github.com/go-gorp/gorp) ＋ [squirrel](https://github.com/Masterminds/squirrel) でサンプルコードを書いたのでメモ。

<!--more-->

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/Masterminds/squirrel"
    _ "github.com/go-sql-driver/mysql"

    "gopkg.in/gorp.v1"
)

type DbMap struct {
    *gorp.DbMap
}

func newDbMap() (*DbMap, error) {
    db, err := sql.Open("mysql", "root:@/test?loc=Local&parseTime=true")
    if err != nil {
        return nil, err
    }

    dbm := &gorp.DbMap{
        Db: db,
        Dialect: gorp.MySQLDialect{
            Engine:   "InnoDB",
            Encoding: "UTF8",
        },
    }

    return &DbMap{dbm}, nil
}

func (dbMap *DbMap) initDB() error {
    dbMap.addUserTable()
    if err := dbMap.DropTablesIfExists(); err != nil {
        return err
    }
    if err := dbMap.CreateTablesIfNotExists(); err != nil {
        return err
    }
    if err := dbMap.TruncateTables(); err != nil {
        return err
    }
    return nil
}

func (dbMap *DbMap) addUserTable() {
    userTableMap := dbMap.AddTableWithName(User{}, "user").SetKeys(true, "ID")
    userTableMap.ColMap("Name").SetNotNull(true)
    userTableMap.ColMap("Age").SetNotNull(true)
    userTableMap.ColMap("UpdatedAt").SetNotNull(true)
    userTableMap.ColMap("CreatedAt").SetNotNull(true)
    return
}

type User struct {
    ID        int64     `db:"id"`
    Name      string    `db:"name"`
    Age       int32     `db:"age"`
    UpdatedAt time.Time `db:"updated_at"`
    CreatedAt time.Time `db:"created_at"`
}

func newUser(name string, age int32) *User {
    return &User{
        Name: name,
        Age:  age,
    }
}

func (u *User) PreInsert(s gorp.SqlExecutor) error {
    now := time.Now()
    u.UpdatedAt = now
    u.CreatedAt = now
    return nil
}

func (u *User) PreUpdate(s gorp.SqlExecutor) error {
    u.UpdatedAt = time.Now()
    return nil
}

func (u *User) PrintValues() {
    fmt.Println(fmt.Sprintf(
        "id: %d, name: %s, age: %d, created_at: %s, updated_at: %s",
        u.ID, u.Name, u.Age, u.CreatedAt, u.UpdatedAt))
    return
}

func getUserNum(dbMap *DbMap) (int64, error) {
    sql, _, err := squirrel.Select("COUNT(*)").From("user").ToSql()
    if err != nil {
        return 0, err
    }

    return dbMap.SelectInt(sql)
}

func getUsers(dbMap *DbMap) ([]*User, error) {
    sql, _, err := squirrel.Select("*").From("user").ToSql()
    if err != nil {
        return nil, err
    }

    var users []*User
    if _, err := dbMap.Select(&users, sql); err != nil {
        return nil, err
    }

    return users, err
}

func getUserByID(dbMap *DbMap, id int64) (*User, error) {
    sql, _, err := squirrel.Select("*").From("user").Where("id = ?").ToSql()
    if err != nil {
        return nil, err
    }

    var u User
    if err := dbMap.SelectOne(&u, sql, id); err != nil {
        return nil, err
    }

    return &u, nil
}

func main() {
    dbMap, err := newDbMap()
    checkError(err)
    defer dbMap.Db.Close()

    err = dbMap.initDB()
    checkError(err)

    // create
    u1 := newUser("m0t0k1ch1", 27)
    u2 := newUser("m0t0k1ch2", 28)
    err = dbMap.Insert(u1, u2)
    checkError(err)

    // read - user num
    userNum, err := getUserNum(dbMap)
    checkError(err)
    fmt.Println("userNum:", userNum)
    fmt.Println("---")

    // read - each user
    users, err := getUsers(dbMap)
    checkError(err)
    for _, u := range users {
        u.PrintValues()
    }
    fmt.Println("---")

    time.Sleep(3 * time.Second)

    // update
    u := users[0]
    u.Age = 29
    _, err = dbMap.Update(u)
    checkError(err)

    // read - by id
    u, err = getUserByID(dbMap, u.ID)
    checkError(err)
    u.PrintValues()
    fmt.Println("---")

    // delete
    _, err = dbMap.Delete(u)
    checkError(err)

    // read - each user
    users, err = getUsers(dbMap)
    checkError(err)
    for _, u := range users {
        u.PrintValues()
    }

    return
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
    return
}
```

出力はこんな感じ。

```txt
userNum: 2
---
id: 1, name: m0t0k1ch1, age: 27, created_at: 2016-01-17 16:16:44 +0900 JST, updated_at: 2016-01-17 16:16:44 +0900 JST
id: 2, name: m0t0k1ch2, age: 28, created_at: 2016-01-17 16:16:44 +0900 JST, updated_at: 2016-01-17 16:16:44 +0900 JST
---
id: 1, name: m0t0k1ch1, age: 29, created_at: 2016-01-17 16:16:44 +0900 JST, updated_at: 2016-01-17 16:16:47 +0900 JST
---
id: 2, name: m0t0k1ch2, age: 28, created_at: 2016-01-17 16:16:44 +0900 JST, updated_at: 2016-01-17 16:16:44 +0900 JST
```
