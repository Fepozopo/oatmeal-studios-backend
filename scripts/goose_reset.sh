export $(grep DB_URL .env | xargs)
goose -dir sql/schema postgres "$DB_URL" reset