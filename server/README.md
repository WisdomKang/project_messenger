# Server

메신져 서버

- Go
  - Gin
  - Gorm
- AMQP
- RDBS

## docker 실행 명령어

- database

```cmd
docker run --name=my-database --env=POSTGRES_PASSWORD=mypassword -p 5432:5432 -d postgres
```

- Message Queue broker

```cmd
docker run --hostname=my-rabbit -p 15672:15672 -p 5672:5672 -d rabbitmq:3-management
```
