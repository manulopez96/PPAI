#!/bin/bash

# sudo apt-get update
# sudo apt-get install postgresql postgresql-contrib

# sudo systemctl start postgresql
# sudo systemctl enable postgresql


# sudo -u postgres psql 
# DROP DATABASE ppai_db;
# ALTER USER postgres WITH PASSWORD '1234';


DB_NAME="ppai_db"
DB_USER="postgres"
DB_HOST="localhost"
DB_PORT="5432"
DB_PASS="1234"

PGPASSWORD="$DB_PASS" psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -c "CREATE DATABASE $DB_NAME;"

go mod tidy

# firefox --class "WebApp-TEST9224" \
#         --name "WebApp-TEST9224" \
#         --profile "$HOME/.local/share/ice/firefox/TEST9224" \
#         --no-remote "http://localhost:8080/" &
# disown


firefox --kiosk "http://localhost:8080/" & disown


go build -o ./ppai ./cmd/main.go
# ./ppai 
go run ./cmd/main.go

# pkill firefox
pkill -f "localhost:8080"

