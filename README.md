# Golang Nextjs Template

This is a template for a web application that uses golang for the backend and nextjs for the frontend.

## Components

### Web Server

[nginx](https://github.com/nginx/nginx)

### Application Server

- backend
  - golang
    - framework: [Gin](https://github.com/gin-gonic/gin)
    - database library: [sqlx](https://github.com/jmoiron/sqlx)
    - sql migration tool: [goose](https://github.com/pressly/goose)
- frontend
  - typescript
    - framework: [Nextjs](https://github.com/vercel/next.js)

## Install

### prepare

- docker
- docker-compose

### clone

```shell
git clone https://github.com/ITK13201/golang-nextjs-template.git
cd golang-nextjs-template
```

### init environment variables and create log files

```shell
./scripts/environment/init.sh
```

### build

```shell
docker-compose build
```

## Usage

### start container

```shell
docker-compose up
```

### stop container

```shell
docker-compose down
```


## References

For the design of the clean architecture, I referred to the following article.

- [Clean ArchitectureでAPI Serverを構築してみる - Qiita](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
