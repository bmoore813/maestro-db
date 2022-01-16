FROM golang:alpine as build

LABEL MAINTAINER="MAESTRO-TEAM"
WORKDIR /go/src/github.com/bmoore813/meastro-db

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/meastro-db
RUN apk update \
    && apk upgrade \
    && apk add --no-cache git \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true

COPY rds-combined-ca-bundle.pem /rds-combined-ca-bundle.pem

## Build Stage II
FROM scratch

COPY --from=build /go/src/github.com/bmoore813/meastro-db/rds-combined-ca-bundle.pem /rds-combined-ca-bundle.pem 
COPY --from=build /bin/meastro-db /bin/meastro-db
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


EXPOSE 50051 1313

CMD [ "/bin/bin/meastro-db" ]