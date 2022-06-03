FROM golang:1.18-alpine as builder

WORKDIR /build

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait.sh
RUN chmod +x /wait.sh

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod .
RUN go mod download

COPY config/config.yml /etc/ssr/config.yml

COPY migrations /migrations

COPY . .
RUN go build -o /bin/ssr ./cmd/app/main.go


CMD ["sh","-c","/wait.sh && migrate -path /migrations -database $PG_DSN up && /bin/ssr"]
