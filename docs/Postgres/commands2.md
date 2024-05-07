## 1) Connect to PostgreSQL database

The following command [connects to a database](https://www.postgresqltutorial.com/postgresql-jdbc/connecting-to-postgresql-database/) under a specific user. After pressing `Enter` PostgreSQL will ask for the password of the user.

```
psql -d database -U  user -W
Code language: SQL (Structured Query Language) (sql)
```

For example, to connect to `dvdrental` database under `postgres` user, you use the following command:

```
C:\Program Files\PostgreSQL\9.5\bin>psql -d dvdrental -U postgres -W
Password for user postgres:
dvdrental=#Code language: SQL (Structured Query Language) (sql)
```

If you want to connect to a database that resides on another host, you add the -h option as follows:

```
psql -h host -d database -U user -WCode language: SQL (Structured Query Language) (sql)
```

In case you want to use SSL mode for the connection, just specify it as shown in the following command:

```
psql -U user -h host "dbname=db sslmode=require"Code language: SQL (Structured Query Language) (sql)
```

## 2) Switch connection to a new database

Once you are connected to a database, you can switch the connection to a new database under a user specified by `user`. The previous connection will be closed. If you omit the `user` parameter, the current `user` is assumed.

```
\c dbname username
```

The following command connects to `dvdrental` database under `postgres` user:

```php
postgres=# \c dvdrental
You are now connected to database "dvdrental" as user "postgres".
dvdrental
```

## 3) List available databases

To [list all databases](https://www.postgresqltutorial.com/postgresql-show-databases/) in the current PostgreSQL database server, you use `\l` command:

```sql
\l
```

## 4) List available tables

To [list all tables](https://www.postgresqltutorial.com/postgresql-show-tables/) in the current database, you use `\dt` command:

```sql
\dt
```

Note that this command shows the only table in the currently connected database.

## 5) Describe a table

To [describe a table](https://www.postgresqltutorial.com/postgresql-describe-table/) such as a column, type, modifiers of columns, etc., you use the following command:

```sql
\d table_name
```

## 6) List available schema

To list all [schemas](https://www.postgresqltutorial.com/postgresql-schema/) of the currently connected database, you use the `\dn` command.

```sql
\dn
```

## 7) List available functions

To list available functions in the current database, you use the `\df` command.

```sql
\df
```

## 8) List available views

To list available [views](https://www.postgresqltutorial.com/postgresql-views/) in the current database, you use the `\dv` command.

```sql
\dv
```

## 9) List users and their roles

To list all users and their assign roles, you use `\du` command:

```
\du
```

## 10) Execute the previous command

To retrieve the current version of PostgreSQL server, you use the `version()` function as follows:

```
SELECT version();
```

Now, you want to save time typing the previous command again, you can use `\g` command to execute the previous command:

```
\g
```

psql executes the previous command again, which is the [SELECT statement](https://www.postgresqltutorial.com/postgresql-select/),.

## 11) Command history

To display command history, you use the `\s` command.

```
\s
```

If you want to save the command history to a file, you need to specify the file name followed the `\s` command as follows:

```
\s filename
```

## 12) Execute psql commands from a file

In case you want to execute psql commands from a file, you use `\i` command as follows:

```
\i filename
```

## 13) Get help on psql commands

To know all available psql commands, you use the `\?` command.

```
\?
```

To get help on specific PostgreSQL statement, you use the `\h` command.

For example, if you want to know detailed information on [ALTER TABLE](https://www.postgresqltutorial.com/postgresql-alter-table/) statement, you use the following command:

```
\h ALTER TABLE
```

## 14) Turn on query execution time

To turn on query execution time, you use the `\timing` command.

```
dvdrental=# \timing
Timing is on.
dvdrental=# select count(*) from film;
 count
-------
  1000
(1 row)

Time: 1.495 ms
dvdrental=#
```

You use the same command `\timing` to turn it off.

```
dvdrental=# \timing
Timing is off.
dvdrental=#
```