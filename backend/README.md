# Backend

## Usage

migration directory: [./migrations](./migrations)

See the [goose documentation](https://pressly.github.io/goose/) for details.

### migrate sql files

```shell
docker-compose exec backend goose.sh up
```

### rollback sql files

```shell
docker-compose exec backend goose.sh down
```
