
# run this comman on terminal to make shart.sh executable
# chmod +x start.sh

# this file will be run by /bin/sh because we are using alpine image
# !/bin/sh

# to make sure that the script will exit immediately if a command returns a non-zero status
set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"

# takes all parameters passed to the script and run it
exec "$@"