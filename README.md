# Aggregator

Aggregator is a Go command-line RSS reader. It stores users, feeds, follows, and scraped posts in PostgreSQL and reads its local CLI state from a config file in your home directory.

## Requirements

- Go 1.25 or newer
- PostgreSQL
- A database URL for the app to use

## Setup

1. Create a PostgreSQL database and run the project schema/migrations.
2. Set the database URL in your local config file at `~/.gatorconfig.json`.
3. Run the program with `go run . <command> [args...]`.

## Build

To build the binary from the project root, run:

```bash
go build .
```

That produces a binary named after the current directory. If you want to choose the output name explicitly, use:

```bash
go build -o bin/agg .
```

Then run it with:

```bash
./bin/agg <command> [args...]
```

Example config:

```json
{
  "Db_url": "postgres://user:pass@localhost:5432/aggregator?sslmode=disable"
}
```

You can also store a `Connection_string` value instead of `Db_url`.

## Commands

- `register <name>`: create a new user and set it as the current user
- `login <name>`: select an existing user as the current user
- `reset`: reset the database state
- `users`: list all users
- `agg <duration>`: continuously fetch feeds on an interval, for example `5s` or `1m`
- `addfeed <name> <url>`: create a feed and follow it for the current user
- `feeds`: list feeds with their creator
- `follow <url>`: follow an existing feed by URL
- `following`: list feeds followed by the current user
- `unfollow <url>`: stop following a feed by URL
- `browse [limit]`: show recent posts for the current user, defaulting to 2

## Examples

From the `bin` directory : 
```bash
./bin/agg register alice
./bin/agg login alice
./bin/agg addfeed "Hacker News" https://news.ycombinator.com/rss
./bin/agg agg 5s
./bin/agg browse 10
```

## Notes

- The config file is stored at `~/.gatorconfig.json`.
- The aggregator command will skip duplicate posts when PostgreSQL returns unique-constraint error `23505`.
- The browse command works with plain `int` values at the CLI boundary.