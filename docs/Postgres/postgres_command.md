Login usuario postgres:

```
$ sudo su - postgres
```

Creación de un usuario:

```
CREATE USER nombre_usuario WITH password '123456'
```

Eliminar usuario:

```
DROP USER nombre_usuario
```

Crear base de datos:

```
CREATE DATABASE nombre_db WITH OWNER nombre_usuario;
```

Eliminar base de datos:

```
DROP DATABASE nombre_db
```

Acceder database con usuario x:

```
psql -U nombre_usuario nombre_db
```

Crear schema:

```
CREATE SCHEMA [IF NOT EXISTS] schema_name;

In this example, we will create a schema for a user (say, Raju). to do show let’s first create a user using the below statement:
```

```
CREATE USER Raju WITH ENCRYPTED PASSWORD 'Postgres123';
```

Crear schema autorizado

```
CREATE SCHEMA AUTHORIZATION Raju;
```

```
CREATE SCHEMA IF NOT EXISTS geeksforgeeks AUTHORIZATION Raju;
```

Obtener ayuda:

```
\h
```

Quit

```
\q
```

Leer comandos desde un archivo:

```
\i input.sql
```

Dump db a un archivo:

```
$ pg_dump -U nombre_usuario nombre_db > db.out
```

Dump todas las bases de datos:

```
$ sudo su - postgres
$ pg_dumpall > /var/lib/pgsql/backups/dumpall.sql
```

Restaurar db:

```
$ sudo su - postgres
$ psql -f /var/lib/pgsql/backups/dumpall.sql mydb
```

También:

```
$ psql -U postgres nombredb < archivo_restauracion.sql
```

List databases:

```
\l
```

List tables in database:

```
\d
```

Describe table:

```
\d table_name
```

Describe table:

```
\d+ table_name
```

Use database_name:

```
\c nombre_db
```

Show users:

```
select * from "pg_user";
# también
\du
```

Escribir las consultas en tu editor favorito:

```
\e
```

Activar/Desactivar ver el tiempo del query:

```
\timing
```

Reset a user password as admin:

```
ALTER USER usertochange WITH password 'new_passwd';
```

Select version

```
SELECT version();
```

Change Database Owner:

```
ALTER DATABASE database_name OWNER TO new_owner;
```

Create a superuser user:

```
ALTER USER mysuper WITH SUPERUSER;
# or even better
ALTER USER mysuper WITH SUPERUSER CREATEDB CREATEROLE INHERIT LOGIN REPLICATION
```

Saber el tamaño usado las tablas en una base de datos:

Ver mas: http://www.niwi.be/2013/02/17/postgresql-database-table-indexes-size/

```postgresql
SELECT pg_size_pretty(pg_database_size('dbname'));
```

Tamaño DB 

```postgresql
SELECT pg_database.datname, pg_size_pretty(pg_database_size(pg_database.datname)) AS size FROM pg_database;
```

