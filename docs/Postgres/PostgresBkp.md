##### Backup

```
#let's create a backup from remote postgresql database using pg_dump:
#
# pg_dump -h [host address] -Fc -o -U [database user] <database name> > [dump file]
#
#later it could be restored at the same remote server using:
#
# sudo -u postgres pg_restore -C mydb_backup.dump
#
#Ex:

pg_dump -h 67.8.78.10 -p 5432 -Fc -o -U myuser mydb > mydb_backup.dump

pg_restore -C mydb_backup.dump

PGPASSWORD=peaceful_cartwright pg_dump -h db.carti3r.tk -p 9955 -Fc -O -U postgres alg_database > ./bkp.dump


#complete (all databases and objects)

pg_dumpall -U myuser -h 67.8.78.10 -p 5432 --clean --file=mydb_backup.dump


#restore from pg_dumpall --clean:

psql -f mydb_backup.dump postgres #it doesn't matter which db you select here
```



```sh
docker exec -t <your-postgres-container-id> pg_dumpall -c -U postgres > dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql
```

**Restaurar**

```sh
cat your_dump.sql | docker exec -i <your-postgres-container-id> psql -U myuser
```

```
psql -h mybts-db.cjeketu2j07o.us-east-2.rds.amazonaws.com -U mybtsadmin dev

chaRmandeR.0905

GRANT ALL ON schema public TO mybtsadmin;

GRANT USAGE ON SCHEMA public TO mybtsadmin;

DROP DATABASE mydb WITH (FORCE);

CREATE ROLE platform_prod WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD 'postgres';

CREATE ROLE user_platform_prod WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD 'postgres';
```

```
#! /bin/bash

pg_dump -h my.host.com -p 1234 -U my_user my_database > my_dump.sql

pg_restore -h another.host.com -p 4321 -U another_user -d another_database my_dump.sql
PGPASSWORD=peaceful_cartwright pg_restore -h db -p 5555 -U postgres -d alg_database ./alg_database.sql
```

