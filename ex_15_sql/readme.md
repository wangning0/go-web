# Go与database

Go与PHP不同的地方是Go官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动，这样的话，只要是按照标准接口开发的代码，以后需要迁移数据库时，不需要任何修改

## Overview

在Go中连接数据库，将会使用`sql.DB`，可以用这个类型创建语句和事务，执行查询和获取结果。

首先要知道的是 `sql.DB`不是数据库的连接，它也不是指定任何指定的数据库或者是schema,它是一个接口和数据库的抽象实现，它可能会像本地文件一样变化，通过网络连接、内存和进程

`sql.DB`在背后做了一些重要的行为

* 通过driver，建立和关闭真正的数据库的连接
* 根据需要管理一个连接池，这可能是提到的各种各样的事情

## Importing a Database Driver

要使用`database/sql`，你会使用包它本身，还需要你想要使用的特定的数据库的驱动程序

不建议直接使用driver包，你的代码应该址使用在`database/sql`中定义的类型，这样的话，会把你的代码分离出来，可以改变底层的数据库而不改变代码

## Accessing the Database

具体的可以看[mysql连接](https://github.com/go-sql-driver/mysql/wiki/Examples)，有几个要注意的点 

sql.Open() does not establish any connections to the database, nor does it validate driver connection parameters. Instead, it simply prepares the database abstraction for later use. The first actual connection to the underlying datastore will be established lazily, when it’s needed for the first time. If you want to check right away that the database is available and accessible (for example, check that you can establish a network connection and log in), use db.Ping() to do that, and remember to check for errors

```
err = db.Ping()
if err != nil {
    // ...
}
```

以及 不要频繁的Open和Close数据库，因为`sql.DB`是设计为长连接的
