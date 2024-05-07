# Drop databases Mongo

### Example – Drop MongoDB Database

Following is an example where we shall try deleting database named **tutorialkart**.

Refer [MongoDB Create Database](https://www.tutorialkart.com/mongodb/mongodb-create-database/), if you have not already created one.

Open [Mongo Shell](https://www.tutorialkart.com/mongodb/mongo-shell/) and follow the commands in sequence.

```
> show dbs``admin     0.000GB``local     0.000GB``tutorialkart 0.000GB``> use tutorialkart``switched to db tutorialkart``> db.dropDatabase()``{ "dropped" : "tutorialkart", "ok" : 1 }``> show dbs``admin 0.000GB``local 0.000GB``>
```

Following is the explanation for each mongodb command we executed above

1. show dbs there are three databases. We shall delete tutorialkart database in this demonstration.
2. use tutorialkart switched to tutorialkart database.
3. `db.dropDatabase()` drops the database that is currently in use i.e., tutorialkart database.
4. show dbs now there are only two databases, because tutorialkart database is no more present.