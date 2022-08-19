# go-template

Template for productive high-tech creators

### Features

- Multi-command support with https://github.com/spf13/cobra 🤾🏼‍♀️ 🤾🏼 🤾🏼‍♂️
- Shared git hooks: fast linter on `commit`, full linter + tests + dependencies update on `push` 🪝
- Lightweight logger that can be used both locally and globaly with https://github.com/uber-go/zap ✉️
- Extensive application configuration through file (yaml by default), env, flags with https://github.com/spf13/viper 💽
- Friendly `graceful shutdown` that can be used in any part of your code 🤳 
- Database support both for local development and migrations (postgres by default), can be changed easily at `db.make` and `docker-compose.yml`. Migrations engine - https://github.com/golang-migrate/migrate 💾
