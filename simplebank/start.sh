 #!/bin/sh 

#Alpine linux is very light weight linux distribution design to make docker images as small as possible so it does not include bash shell to save the space so we have to run our script on the sh shell originally command line use bash shell to run the commands

#!/bin/sh = "use whatever small shell the system provides (ash in Alpine)."

set -e 

echo "run db migration"
/THEPROJECT/migrate -path /THEPROJECT/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
