# Connection to the *gator* database

```
psql "postgres://postgres:@localhost:5432/gator"
```
passw : `postgres`

# Running up migration 
From  `sql/schema` directory :

```
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

# Connecting to the server

```
sudo -u postgres psql
```