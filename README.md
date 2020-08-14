# About This

Simple API Server 

## Env

OS: Ubuntu 18.04  
Language: go1.14.6 linux/amd64  
database: mysql Ver 14.14 Distrib 5.7.31, for Linux (x86_64) using  EditLine wrapper  

## How to use?

### DataBase

請在MySQL中新增使用者
使用者：SimpleDBOwner
密碼：Newpassword
並新增以下database及table

database： mydata  
table：data  
schema：
| Field  |  Type | Null  |  Key | Default  |Extra|
|---|---|---|---|---|---|
|  Id | int  | NO  |PRI   |  NULL |auto_increment|
|  DateAdded | datetime  | YES  |   | NULL  ||

table：Location  
schema：
| Field  |  Type | Null  |  Key | Default  |Extra|
|---|---|---|---|---|---|
|  Id | int（11）  | YES  |   |  NULL ||
|  Lat | double  | YES  |   | NULL  ||
|  Long | double  | YES  |   | NULL  ||

### Golang

#### exec

```bash
go build 
```

後會有執行檔，執行就可以了

#### Test

在各功能區塊下有*_test.go可以進行單元測試

## API

```YAML
/data:
    Get:
        description: 
            取得db資料
        responses:
            "200":
            description: Success
    Post:
        description:
            增加db資料
        parameters:
            -ID:   string
                in: body
                description: ID，目前留空
            -Location: object
                in: body
                parameters:
                    -Lat:   float32
                        description: T.B.D
                    -Long: float32
                        description: T.B.D
                description: T.B.D
            -DateAdded: time
                in: body
                description: 時間標記，目前留空
```
