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
docker run --hostname=6b6de5ed5819 --mac-address=02:42:ac:11:00:02 --env=POSTGRES_PASSWORD=mypassword --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/lib/postgresql/15/bin --env=GOSU_VERSION=1.16 --env=LANG=en_US.utf8 --env=PG_MAJOR=15 --env=PG_VERSION=15.2-1.pgdg110+1 --env=PGDATA=/var/lib/postgresql/data --volume=/var/lib/postgresql/data -p 5432:5432 --restart=no --runtime=runc -d postgres
```

- Message Queue broker

```cmd
docker run --hostname=my-rabbit --env=RABBITMQ_DEFAULT_PASS=password --env=RABBITMQ_DEFAULT_USER=user --env=PATH=/opt/rabbitmq/sbin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=RABBITMQ_DATA_DIR=/var/lib/rabbitmq --env=RABBITMQ_VERSION=3.11.13 --env=RABBITMQ_PGP_KEY_ID=0x0A9AF2115F4687BD29803A206B73A36E6026DFCA --env=RABBITMQ_HOME=/opt/rabbitmq --env=HOME=/var/lib/rabbitmq --env=LANG=C.UTF-8 --env=LANGUAGE=C.UTF-8 --env=LC_ALL=C.UTF-8 --volume=/var/lib/rabbitmq -p 8080:15672 -p 5672:5672 --restart=no --label='org.opencontainers.image.ref.name=ubuntu' --label='org.opencontainers.image.version=20.04' --runtime=runc -d rabbitmq:3-management
```
