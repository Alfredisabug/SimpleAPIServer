# About This

Simple API Server 

## Env

OS: Ubuntu 18.04
Language: go1.14.6 linux/amd64
database: mysql Ver 14.14 Distrib 5.7.31, for Linux (x86_64) using  EditLine wrapper

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
