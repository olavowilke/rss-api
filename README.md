### A lightweight Go backend server, Postgres integrated, for aggregating data from RSS Feeds.

Building locally:

With docker installed, run Postgres:
`docker compose up`

With Goose installed, run migrations:
`cd sql/schema`
`goose postgres postgres://admin:admin@localhost:5432/rssdb up`

With Go installed, install dependencies, build and run
`go build`
`./rss-api`

Check the `/requests` package for request examples.