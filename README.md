# Gator

*Gator* is a Go command-line RSS reader. It is a multi-user CLI application. It stores users, feeds, follows, and scraped posts in PostgreSQL local server and reads its local CLI state from a config file in your home directory.

## Requirements

- Go 1.25 or newer
- PostgreSQL database

## Setting up the PostgresSQL database 

The next steps assume you have installed PostgreSQL on your machine :

1. Start the postgres server in the background. 
    
    - mac : ` brew services start postgresql@15`

    - linux : ` sudo service postgresql start`
2. Connect to the server by entering the `psql` shell :

    - Mac : `psql postgres`
    - Linux : `sudo -u postgres psql`

    You should see a promp that looks like 

    ```bash
    postgres=#
    ```

3. Create new database
    ```sql 
    CREATE DATABASE gator;
    ```
4. Connect to the database
    ```sql
    \c gator
    ```
5. Set the user password (Linux / WSL only)
    ```sql
    ALTER USER postgres PASSWORD 'postgres'
    ```
You can type `exit` to leave the shell.

## Updating the database

After sucessfully creating the blank database, we need to update it to the version that this program uses. From the root directory of this codebase, run the shell script inside the `sql/schema/up.sh` file that will do so, with the default configuration setup (username is postgres, password also postgres).

It uses a connection string (a URL with the information needed to connect to the database), of format :

```
protocol://username:password@host:port/database
```

The default port number is `5432`, modify the username and password or the port number if needed.


Test your connection string by running `psql`, for example 
```bash
psql "postgres://postgres:postgres@localhost:5432/gator"
```

## Setting up the configuration file

We are using a single JSON file to keep track of who is currently logged in and the connection credentials for the PostgreSQL database.

1. Create a config file in your home directory, `~/.gatorconfig.json` with the following content : 
``` 
{
  "db_url": "protocol://username:password@host:port/database?sslmode=disable"
}
```
Replace the protocal with `postgres`, the username, password, host, port, database name to what you configured, else by default use :

```
{
"db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
``` 


## Build

To add the executatble binary to the list of programs with go, run 

` go install ` that will let you run it from anywhere as 

```
Gator [argument]
```

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

## Usage

- Register a user
- Add some feeds
- From a separate terminal, call the aggregate command and set a reasonable frequency (every 30min for instance)
- Browse the articles using the `browse [arg]` command

## Examples

```bash
Gator register alice
Gator login alice
Gator addfeed "Hacker News" https://news.ycombinator.com/rss
Gator agg 30min
Gator browse 10
```

## Notes

- The aggregator command will skip duplicate posts when PostgreSQL returns unique-constraint error `23505`.
- The browse command works with plain `int` values at the CLI boundary.