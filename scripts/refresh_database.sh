set -e

echo "Refreshing database..."
psql -U $DB_USERNAME -h $DB_HOST --set=ON_ERROR_STOP << EOF
DROP DATABASE IF EXISTS musicplayer;
DROP DATABASE IF EXISTS musicplayer_test;
CREATE DATABASE musicplayer;
CREATE DATABASE musicplayer_test;
EOF

# Run migrations on dev database
goose -dir ./db/migrations postgres "host=$DB_HOST user=$DB_USERNAME dbname=musicplayer sslmode=disable" up
# Run migrations on test database
goose -dir ./db/migrations postgres "host=$DB_HOST user=$DB_USERNAME dbname=musicplayer_test sslmode=disable" up

# run script for fixtures when i add fixtures
