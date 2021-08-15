ARG go_version=1.16


# Dvenelopment environment
FROM golang:${go_version} AS development

ENV GO111MODULE=on
WORKDIR /backend

RUN apt-get update

COPY ./backend/go.mod ./
COPY ./backend/go.sum ./
RUN go mod download

RUN go build -o /go/bin/air github.com/cosmtrek/air
RUN go build -o /go/bin/sql-migrate github.com/rubenv/sql-migrate

COPY ./backend ./

EXPOSE 8000

ENTRYPOINT ["air", "-c", ".air.toml"]


# Builder
FROM golang:${go_version} AS builder

ENV GO111MODULE=on
WORKDIR ${WORKDIR}

RUN apt-get update

COPY ./backend/go.mod ./
COPY ./backend/go.sum ./
RUN go mod download

COPY ./backend ./

RUN go build -o ./bin/golang-nextjs-template


# Production environment
FROM gcr.io/distroless/base:debug AS production

COPY --from=builder /backend/bin/build /bin/golang-nextjs-template

ENTRYPOINT ["/bin/golang-nextjs-template"]
