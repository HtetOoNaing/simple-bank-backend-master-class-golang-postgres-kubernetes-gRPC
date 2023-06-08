# Build stage
FROM golang:1.20.4-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# add curl to download migrate and unzip it
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
# copy from builder - the unzipped migrate to current migrate
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
# copy all migration files under db/migration to current migration
COPY db/migration ./migration

EXPOSE 8080
# if CMD is used with ENTRYPOINT, CMD is just a second parameter of ENTRYPOINT
# passed /app/main parameters to start.sh shell
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]