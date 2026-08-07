# Connection to the database

in a terminal, type  `psql "postgres://postgres:@localhost:5432/gator" 
passw : 'postgres"

# Running up migration 

'goose postgres "postgres://postgres:postgres@localhost:5432/gator" up'