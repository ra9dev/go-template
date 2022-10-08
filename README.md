# go-template

Template for productive high-tech creators

### Features

- Multi-command support with https://github.com/spf13/cobra 🤾🏼‍♀️ 🤾🏼 🤾🏼‍♂️
- Shared git hooks: fast linter + goimports on `commit`, full linter + tests + dependencies update on `push`. Run `make environment` for setup 🪝
- Lightweight logger that can be used both locally and globaly with https://github.com/uber-go/zap ✉️
- Extensive application configuration through file (yaml by default), env, flags with https://github.com/spf13/viper 💽
- Friendly `graceful shutdown` that can be used in any part of your code 🤳 
- Database support both for local development and migrations (postgres by default), can be changed easily at `db.make` and `docker-compose.yml`. Migrations engine - https://github.com/golang-migrate/migrate 💾
- Smart goimports linting that keeps your imports within 3 blocks via https://github.com/incu6us/goimports-reviser 🗄
- Multi-port api server for: `http, admin_http, grpc` 🎏 Note: `grpc` is still in progress
- Swagger spec generation with https://github.com/swaggo/swag (Admin HTTP) 😎
- Minimal Docker image ~ 25MB 🐳

### To Be Done
- SRE best practices support: tracing, metrics, etc.
- Protocols support: grpc, quic, etc.
- Dynamic configuration via etcd/consul/etc

### Tracing
[Jaeger](https://www.jaegertracing.io/) open source, end-to-end distributed<br/>
[OpenTelemetry](https://opentelemetry.io/docs/migration/opentracing/)<br/>

### Jaeger UI:
http://localhost:16686/