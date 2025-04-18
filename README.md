A lightweight Go backend server, Postgres integrated, for aggregating data from RSS feeds.

Generate db models and data access layer

`sqlc generate`

Run migrations

`goose postgres postgres://admin:admin@localhost:5432/rssdb up`

Install dependencies, build and run

`go mod tidy`

`go build`

`./rss-api.exe`