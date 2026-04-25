# blog-aggregator

A multi-user RSS feed aggregator CLI written in Go. Users can register, follow feeds, and browse the latest posts pulled from those feeds — all from the terminal.

## Prerequisites

You'll need the following installed on your machine:

- **Go** (1.22 or newer) — https://go.dev/dl/
- **PostgreSQL** (15 or newer) — https://www.postgresql.org/download/

## Installation

Install the `blog-aggregator` CLI using `go install`:

```bash
go install github.com/weilok2021/blog-aggregator@latest
```

This places the binary in your Go bin directory (typically `~/go/bin` or `$(go env GOPATH)/bin`). Make sure that directory is on your `$PATH` so you can run `blog-aggregator` from anywhere:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Configuration

`blog-aggregator` reads its config from a JSON file at `~/.gatorconfig.json`. Create the file with the following content, replacing the connection string with your local database details:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

You'll also need to apply the database schema migrations (using [goose](https://github.com/pressly/goose) or the tool of your choice) against the `sql/schema` directory before running the program.

## Usage

The general form is:

```bash
blog-aggregator <command> [args...]
```

### A few commands to get you started

- `blog-aggregator register <name>` — create a new user and log them in
- `blog-aggregator login <name>` — switch to an existing user
- `blog-aggregator users` — list all users (the current one is marked)
- `blog-aggregator addfeed <name> <url>` — add a new RSS feed (also follows it)
- `blog-aggregator feeds` — list every feed in the database
- `blog-aggregator follow <url>` — follow an existing feed
- `blog-aggregator following` — list feeds the current user follows
- `blog-aggregator unfollow <url>` — unfollow a feed
- `blog-aggregator agg <interval>` — start the aggregator loop (e.g. `1m`, `30s`)
- `blog-aggregator browse [limit]` — show recent posts from feeds you follow (default 2)
- `blog-aggregator reset` — wipe the database (use with care!)
