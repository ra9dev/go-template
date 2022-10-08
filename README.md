# go-template

Template for productive high-tech creators

## Setup

- Run `make environment`
- Replace all occuriences of `go-template` to `your-service`

## Features

### Development

- Shared git hooks: on [commit](./scripts/pre-commit.sh) and on [push](./scripts/pre-push.sh) 🪝
- Friendly [graceful shutdown](./pkg/shutdown/global.go) that can be used in any part of your code 🤳
- [Smart fixer](https://github.com/incu6us/goimports-reviser) for your imports, keeping it within 3 blocks 🗄

### Delivery

- [Multi-command](https://github.com/spf13/cobra) support 🤾🏼‍♀️ 🤾🏼 🤾🏼‍♂️
- Extensive multi-env [configuration](https://github.com/spf13/viper) via [config.yaml](./config/config.yaml), environment variables, flags 💽
- Multi-port api server for: `http, admin_http, grpc` 🎏
- Swagger spec [generation](https://github.com/swaggo/swag) (available at [Admin HTTP](./internal/api/http/admin/router.go)) 😎
- Minimal Docker image ~ 25MB 🐳

### Database

- [Database](./docker-compose.yml) for local development ([postgres](./.ra9/make/db.make) by default) 💾
- [Migrations engine](https://github.com/golang-migrate/migrate) with predefined [make scripts](./.ra9/make/db.make) 🎼

### Site Reliability Engineering

- [Lightweight logger](https://github.com/uber-go/zap) ✉️
- Tracing via [Jaeger](https://www.jaegertracing.io/) and [OpenTelemetry](https://opentelemetry.io).
View your traces at [Jaeger UI](http://localhost:16686/) 🔎

## To Be Done
- SRE best practices support: traced logger, traced transport, metrics, etc.
- Protocols support:
  - GRPC
    - automated proto dependencies fetch
    - swagger-like proto contracts availability
    - multi-transport handlers
  - QUIC
- Dynamic configuration via etcd/consul/etc
