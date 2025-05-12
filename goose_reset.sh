export $(grep DB_CONNECTION_STRING .env | xargs)
goose -dir sql/schema postgres "$DB_CONNECTION_STRING" reset