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

RUN go build -o ./bin/go-next-test


# Production environment
FROM gcr.io/distroless/base:debug AS production

COPY --from=builder /backend/bin/build /bin/go-next-test

ENTRYPOINT ["/bin/go-next-test"]