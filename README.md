# go-template

[![ci](https://github.com/ra9dev/go-template/actions/workflows/cicd.yaml/badge.svg)](https://github.com/ra9dev/go-template/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ra9dev/go-template)](https://goreportcard.com/report/github.com/ra9dev/go-template)
[![Documentation](https://godoc.org/github.com/ra9dev/go-template?status.svg)](https://pkg.go.dev/mod/github.com/ra9dev/go-template)

Template for productive high-tech creators

## Setup

- Replace all occuriences of `go-template` to `your-service`
- Run `make environment`

## Features

### Development

- Shared git hooks: on [commit](./scripts/pre-commit.sh) and on [push](./scripts/pre-push.sh) 🪝
- Friendly [graceful shutdown](https://github.com/ra9dev/shutdown) that can be used in any part of your code 🤳
- [Smart fixer](https://github.com/incu6us/goimports-reviser) for your imports, keeping it within 3 blocks 🗄

### Delivery

- [Multi-command](https://github.com/spf13/cobra) support 🤾🏼‍♀️ 🤾🏼 🤾🏼‍♂️
- Extensive multi-env [configuration](https://github.com/spf13/viper) via [config.yaml](./config/config.yaml), environment variables, flags 💽
- Multi-port api server for: `http, admin_http, grpc` 🎏
- Swagger spec [generation](https://github.com/swaggo/swag) (available at [Admin HTTP](./internal/api/http/admin.go)) 😎
- Minimal Docker image ~ 25MB 🐳

### Database

- [Database](./docker-compose.yml) for local development ([postgres](./.ra9/make/db.make) by default) 💾
- [Migrations engine](https://github.com/golang-migrate/migrate) with predefined [make scripts](./.ra9/make/db.make) 🎼

### Site Reliability Engineering

- [Traced logger](./pkg/sre/log) ✉️
- [Traced transport](./pkg/sre/tracing/transport) 🛞
- Tracing via [Jaeger](https://www.jaegertracing.io/) and [OpenTelemetry](https://opentelemetry.io).
View your traces at [Jaeger UI](http://localhost:16686/) 🔎

## To Be Done
- SRE best practices support: profiling, metrics, etc.
- Protocols support:
  - GRPC
    - automated proto dependencies fetch
    - swagger-like proto contracts availability
  - QUIC
  - multi-transport handlers
- Dynamic configuration via etcd/consul/etc
