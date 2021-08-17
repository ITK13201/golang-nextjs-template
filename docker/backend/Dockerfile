ARG go_version=1.16


# Dvenelopment environment
FROM golang:${go_version} AS development

ARG air_version=v1.27.3
ARG goose_version=v3.0.1
ENV GO111MODULE=on
WORKDIR /backend

RUN apt-get update

COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download

RUN go get -u github.com/cosmtrek/air@${air_version} && \
    go build -o /go/bin/air github.com/cosmtrek/air && \
    go get -u github.com/pressly/goose/v3/cmd/goose@${goose_version} && \
    go build -o /go/bin/goose github.com/pressly/goose/v3/cmd/goose

COPY ./backend .
COPY ./scripts/backend /usr/local/bin

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
