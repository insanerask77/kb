# Mongo backup and restore

You can also use `gzip` for taking backup of one collection and compressing the backup on the fly:

```js
mongodump --db somedb --collection somecollection --out - | gzip > collectiondump.gz
```

or with a date in the file name:

```js
mongodump --db somedb --collection somecollection --out - | gzip > dump_`date "+%Y-%m-%d"`.gz
```

**Update:**
Backup all collections of a database in a date folder. The files are gziped:

```js
mongodump --db somedb --gzip --out /backups/`date +"%Y-%m-%d"`
```

Or for a single archive:

```js
mongodump --db somedb --gzip --archive > dump_`date "+%Y-%m-%d"`.gz
```

Or when mongodb is running inside docker:

```js
docker exec <CONTAINER> sh -c 'exec mongodump --db somedb --gzip --archive' > dump_`date "+%Y-%m-%d"`.gz
```

Backup/Restore Mongodb with timing.

**Backup:**

```js
sudo mongodump --db db_name --out /path_of_your_backup/`date +"%m-%d-%y"`
```

`--db` argument for databse name

`--out` argument for path of output

**Restore:**

```js
sudo mongorestore --db db_name --drop /path_of_your_backup/01-01-19/db_name/
```

`--drop` argument for drop databse before restore

**Timing:**

You can use **crontab** for timing backup:

```js
sudo crontab -e
```

It opens with editor(e.g. nano)

```js
3 3 * * * mongodump --out /path_of_your_backup/`date +"%m-%d-%y"`
```

----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Mongo dump and restore with uri to local

```js
mongodump --uri "mongodb://USERNAME:PASSWORD@IP_OR_URL:PORT/DB_NAME" --collection COLLECTION_NAME -o LOCAL_URL
```

Omitting --collection COLLECTION_NAME will dump entire DB.

If your database is in the local system. Then you can type the below command. for Linux terminal

```js
mongodump -h SERVER_NAME:PORT -d DATABASE_NAME
```

If database has username and password then you can use below code.

```js
mongodump -h SERVER_NAME:PORT -d DATABASE_NAME -u DATABASE_USER -p PASSWORD
```

This worked very well in my Linux terminal.



Dump command:

```js
mongodump --host localhost:27017 --gzip --db Alex --out ./testSO
```

Restore Command:

```js
mongorestore --host localhost:27017 --gzip --db Alex ./testSO/Alex
```

Works perfectly!

------

### While using archive:

Dump command:

```js
mongodump --host localhost:27017 --archive=dump.gz --gzip --db Alex
```

Restore Command:

```js
mongorestore --host localhost:27017 --gzip --archive=dump.gz --db Alex
```

------

> **Note:-** While using archive you need to stick with the `database name`