# db-migration

## Create migrations

Create some migrations using migrate CLI. Here is an example:

``` bash
migrate create -ext sql -dir ./migrations -seq create_users_table
export password=y6w4ziZjVN*RW7vj
migrate -path migrations/ -database postgres://postgres:$password@kc-develop.cosacqsyw3db.ap-southeast-1.rds.amazonaws.com/kidscare force 150
```

