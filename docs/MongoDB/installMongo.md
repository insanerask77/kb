# Instalar Mongo en Ubuntu 22.04

## Install  MongoDB 6.0 on Ubuntu 22.04|20.04|18.04

The following steps will guide to accomplish the installation of MongoDB 6.0 on Ubuntu 20.04|18.04.

#### MongoBD Community & Enterprise Editions

Both the free Community edition and the Enterprise edition of MongoDB are available as a part of the Mongo Enterprise Advanced subscription.
On-disk encryption, LDAP and Kerberos support, auditing, and other features have been introduced to the Enterprise edition.

### Step 1: Perform System Update

Before we can proceed with our installation, let’s update Ubuntu 20.04|18.04 system and install the required packages:

```
sudo apt update
sudo apt install wget curl gnupg2 software-properties-common apt-transport-https ca-certificates lsb-release
```

### Step 2: Import the public key

Run the following command to import the MongoDB public GPG Key:

```
curl -fsSL https://www.mongodb.org/static/pgp/server-6.0.asc|sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/mongodb-6.gpg
```

### Step 3: Configure MongoDB Repo

Using the following commands, add the repository to your system right away:

**Ubuntu 22.04:**

```
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list
```

**Ubuntu 20.04**

```
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list
```

**Ubuntu 18.04**

```
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/6.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list
```

### Step 4: Install MongoDB 6.0 on Ubuntu 22.04|20.04|18.04

Once the appropriate repository has been added, use the following command to install MongoDB 6.0 on Ubuntu.

***Ubuntu 22.04:\***

```
wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1f-1ubuntu2.16_amd64.deb
sudo dpkg -i ./libssl1.1_1.1.1f-1ubuntu2.16_amd64.deb
sudo apt update
sudo apt install mongodb-org
```

***Ubuntu 20.04 / 18.04\***:

```
sudo apt update
sudo apt install mongodb-org
```

Dependency tree:

```
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following package was automatically installed and is no longer required:
  libfwupdplugin1
Use 'sudo apt autoremove' to remove it.
The following additional packages will be installed:
  mongodb-database-tools mongodb-mongosh mongodb-org-database mongodb-org-database-tools-extra mongodb-org-mongos mongodb-org-server mongodb-org-shell mongodb-org-tools
The following NEW packages will be installed:
  mongodb-database-tools mongodb-mongosh mongodb-org mongodb-org-database mongodb-org-database-tools-extra mongodb-org-mongos mongodb-org-server mongodb-org-shell mongodb-org-tools
0 upgraded, 9 newly installed, 0 to remove and 4 not upgraded.
Need to get 134 MB of archives.
After this operation, 458 MB of additional disk space will be used.
Do you want to continue? [Y/n] y
```

After successful installation, start and enable MongoDB:

```
sudo systemctl enable --now mongod
```

Check MongoDB Status:

```
$ systemctl status mongod
● mongod.service - MongoDB Database Server
     Loaded: loaded (/lib/systemd/system/mongod.service; enabled; vendor preset: enabled)
     Active: active (running) since Fri 2022-07-29 04:11:46 EAT; 8s ago
       Docs: https://docs.mongodb.org/manual
   Main PID: 430451 (mongod)
     Memory: 61.6M
     CGroup: /system.slice/mongod.service
             └─430451 /usr/bin/mongod --config /etc/mongod.conf

Jul 29 04:11:46 frank-PC systemd[1]: Started MongoDB Database Server.
```

Check MongoDB installed version:

```
$ mongod --version
db version v6.0.0
Build Info: {
    "version": "6.0.0",
    "gitVersion": "e61bf27c2f6a83fed36e5a13c008a32d563babe2",
    "openSSLVersion": "OpenSSL 1.1.1f  31 Mar 2020",
    "modules": [],
    "allocator": "tcmalloc",
    "environment": {
        "distmod": "ubuntu2004",
        "distarch": "x86_64",
        "target_arch": "x86_64"
    }
}
```

### Step 5: Configure MongoDB 6.0

The configuration file for MongoDB is located at `/etc/mongod.conf`. All the needed configurations, including the database path, logs directory, etc., can be made in this file. Some of the MongoDB setups are listed below.

#### Enable Password Authentication on MongoDB 6.0

Users must enter a password in order to log in and read and edit databases when this is enabled.

To do this, uncomment the following line:

```
security:
  authorization: enabled
```

#### Enable Remote Access on MongoDB 6.0

MongoDB is typically configured to only allow local access. You must modify the code below to include the server IP or hostname as preferred if you want to access it remotely.

```
# network interfaces
net:
  port: 27017
  bindIp: 127.0.0.1  # Enter 0.0.0.0,:: to bind to all IPv4 and IPv6 addresses or, alternatively, use the net.bindIpAll setting.
```

To bind to all IPv4 and IPv6 address you’ll set:

```
 bindIp: 0.0.0.0
```

Save the file after making the required changes, then restart the service:



```
sudo systemctl restart mongod
```

Allowing the port past the firewall will allow remote access to the service:

```
sudo ufw allow 27017
```

### Step 6: Change MongoDB 6.0 default Data Path

The path to MongoDB’s data storage is /var/lib/mongodb by default. However, if another path is selected, this can be adjusted. Stop the currently operating instance to achieve this:

```
sudo systemctl stop mongod.service
```

Create the needed path and give it the appropriate rights to store data:

```
sudo mkdir -p /newdata/mongo
sudo chown -R mongodb:mongodb  /newdata/mongo
```

The former path’s contents should be copied to the new directory:



```
sudo rsync -av /var/lib/mongodb  /newdata/mongo
```

Then rename the old directory to a backup file:

```
sudo mv /var/lib/mongodb /var/lib/mongodb.bak
```

Now, create a symbolic link to the new directory:

```
sudo ln -s /newdata/mongo /var/lib/mongodb
```

Restart MongoDB Service:

```
sudo systemctl daemon-reload
sudo systemctl start mongod
```

### Step 7: Using MongoDB 6.0 Database

You can use the following command to reach the MongoDB shell:



```
$ mongosh
```

Here is the output:

```
Current Mongosh Log ID:	62e33c2dfed670cec73bbf7f
Connecting to:		mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.5.3
Using MongoDB:		6.0.0
Using Mongosh:		1.5.3

For mongosh info see: https://docs.mongodb.com/mongodb-shell/


To help improve our products, anonymous usage data is collected and sent to MongoDB periodically (https://www.mongodb.com/legal/privacy-policy).
You can opt-out by running the disableTelemetry() command.

Warning: Found ~/.mongorc.js, but not ~/.mongoshrc.js. ~/.mongorc.js will not be loaded.
  You may want to copy or rename ~/.mongorc.js to ~/.mongoshrc.js.
test> 
```

#### Create a User and Add a Role in MongoDB

We’ll make a user named ***mongdbuser\*** and give them the admin roles for the purposes of this guide.
The desired user can, however, be created with this command.



```
use admin
db.createUser(
{
user: "mongdbuser",
pwd: passwordPrompt(), // or cleartext password
roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
}
)
```

Give the user the proper password, then use the command listed below to close the shell.

```
> exit
bye
```

Log in to the shell with the user credentials to see if the user has been created.

```
$ mongosh -u mongdbuser -p --authenticationDatabase admin
```

#### Create a database in MongoDB

You can check the available databases:



```
> show dbs
admin   132.00 KiB
config   12.00 KiB
local    72.00 KiB
```

It’s simple to create a database with MongoDB; all you have to do is switch to a blank database. ***mymongodb\***, for instance:

```
> use mongotestdb
```

#### Create a collection in MongoDB

In SQL databases, a collection is comparable to a table. Here, a table can be created in the appropriate database using the commands listed below:

```
db.employeedetails.insertOne(
   {F_Name: "John",
    L_NAME: "Doe",
    ID_NO: "23245",
     AGE: "25",
     TEL: "63365467666"
   }
)
```

Once created, use the following command to view the collections:



```
> show collections
employeedetails
```

## Conclusion

Our tutorial on how to install MongoDB 6.0 on Ubuntu 20.04|18.04 has come to a conclusion. We have also learned how to create databases, collections, and users. We sincerely hope you found this information helpful.