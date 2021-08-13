# gauthfun
RBAC system with tenants and patterns.

## Basic Design

![image](https://user-images.githubusercontent.com/6275608/129182060-b0aa5599-d525-44b3-94d0-e8df922ffb1d.png)

## How to use

- Download package

``` bash
go mod tidy
```

- Modiy `config.toml`

``` ini
[service]
port = 8080

[mysql]
user = "{user}"
password = "{password}"
host = "{host}"
port = 3306
db_name = "{db_name}"
``` 

- Start the process

``` bash
go run main.go
```
