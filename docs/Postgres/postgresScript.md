```bash
#!/usr/bin/env bash

# Create a txt file having these separate words db_variables.txt: 
# user_name user_password db_host db_port db_name db_schema
# as separate columns. Call this script in bash shell as ./db_repopulate.sh $(cat db_variables.txt)
user_name=$1
user_password=$2
db_host=$3
db_port=$4
db_name=$5
db_schema=$6

# Creates database with desired parameters.
PGPASSWORD=$user_password psql -U $user_name -h $db_host -p $db_port -c "CREATE DATABASE $db_name WITH ENCODING 'UTF8';"

if [[ "$db_schema" != "public" ]]
then
    db_command_1="DROP SCHEMA IF EXISTS public CASCADE; DROP SCHEMA IF EXISTS $db_schema CASCADE;"
    db_command_2="CREATE SCHEMA IF NOT EXISTS public AUTHORIZATION $user_name; CREATE SCHEMA IF NOT EXISTS $db_schema AUTHORIZATION $user_name"
else
    db_command_1="DROP SCHEMA IF EXISTS public CASCADE;"
    db_command_2="CREATE SCHEMA IF NOT EXISTS public AUTHORIZATION $user_name;"
fi

PGPASSWORD=$user_password psql -U $user_name -h $db_host -p $db_port -d $db_name -c "$db_command_1$db_command_2"

# This populates our database.

npm run sequelize-config

npx sequelize-cli db:migrate

cd sequelize/module2_sequelize_utils
node timezoneUTC.js

cd ../..
npx sequelize-cli db:seed:all

cd sequelize
node executeUtils.js

```

