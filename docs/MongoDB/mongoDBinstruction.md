## Using Docker to Deploy a MongoDB Cluster/Replica Set

MongoDB is a [general-purpose](https://www.mongodb.com/why-use-mongodb) database that was built with the web in mind. Amongst other things, it offers high availability when used in [clusters](https://www.mongodb.com/basics/clusters), also called [replica sets](https://www.mongodb.com/basics/replication). A replica set is a group of MongoDB servers, called nodes, containing an identical copy of the data. If one of the servers fails, the other two will pick up the load while the crashed one restarts.

![A diagram showing a client application with read and write access to a primary node. Arrows are showing that the data from the primary node is asynchronously replicated in the secondary nodes.](https://webimages.mongodb.com/_com_assets/cms/kv9owcafupia6mnsi-image2.png?auto=format%252Ccompress)

When you create a MongoDB instance on [MongoDB Atlas](https://www.mongodb.com/cloud/atlas), the database-as-a-service offering by MongoDB, a cluster is automatically created for you, ensuring that you have the best possible experience.

However, if you need to experiment with MongoDB clusters, you can use [Docker](https://www.mongodb.com/compatibility/docker) to create a cluster on your personal computer. This tutorial will provide you with the necessary instructions to create your own MongoDB Docker cluster.

## Benefits of Docker Deployment

There are many different ways to create a MongoDB cluster. The easiest way is by using [MongoDB Atlas](https://www.mongodb.com/cloud/atlas/register), the database-as-a-service offering by MongoDB. With just a few clicks, you can create your free cluster, hosted in the cloud.

If you want to experiment with clusters to understand better how they work, or if you need a development environment with a cluster, you can install MongoDB on your computer and [deploy a replica set](https://docs.mongodb.com/manual/replication/).

If you do not want to install MongoDB on your laptop, you can use Docker to run your cluster. Docker is an application to launch containers, packages that contain applications along with all the required dependencies necessary to run them.

Using Docker is an alternative way to get started with replica sets without a local installation of MongoDB. It is a convenient way to ensure that everyone on your team runs the same cluster configuration.

The same technique would be used to deploy a highly available cluster on [Kubernetes](https://www.mongodb.com/presentations/creating-highlyavailable-mongodb-microservices-with-docker-containers-and-kubernetes).

## Instructions

The steps to create a docker cluster are as follows.

1. Create a Docker network.
2. Start three instances of MongoDB.
3. Initiate the Replica Set.

Once you have a MongoDB cluster up and running, you will be able to experiment with it.

### Prerequisites

For this tutorial, you will need to have a runtime for containers. You can find the installation instructions for your operating system on the [Docker official website](https://www.docker.com/get-started).

### Create a Docker Network

The first step is to create a Docker network. This network will let each of your containers running in this network see each other. To create a network, run the `docker network create` command.

```
docker network create mongoCluster
```

The `mongoCluster` parameter here is the network’s name; you can pick one that is appropriate for your setup.

After executing the command, you should see the id of the network you just created.

Note that this command only needs to run once. If you restart your containers afterwards, you won’t need to recreate this network.

### Start MongoDB Instances

You are now ready to start your first container with MongoDB. To start the container, use the `docker run` command.

```
docker run -d --rm -p 27017:27017 --name mongo1 --network mongoCluster mongo:5 mongod --replSet myReplicaSet --bind_ip localhost,mongo1
```

In here, you tell docker to start a container with the following parameters:

- `-d` indicates that this container should run in detached mode (in the background).
- `-p` indicates the port mapping. Any incoming request on port 27017 on your machine will be redirected to port 27017 in the container.
- `--name` indicates the name of the container. This will become the hostname of this machine.
- `--network` indicates which Docker network to use. All containers in the same network can see each other.
- `mongo:5` is the image that will be used by Docker. This image is the MongoDB Community server version 5 ([maintained by Docker](https://hub.docker.com/_/mongo)). You could also use a MongoDB Enterprise [custom image](https://docs.mongodb.com/manual/tutorial/install-mongodb-enterprise-with-docker/).

The rest of this instruction is the command that will be executed once the container is started. This command creates a new `mongod` instance ready for a replica set.

If the command was successfully executed, you should see a long hexadecimal string representing the container id. Start two other containers. You will need to use a different name and a different port for those two.

```
docker run -d --rm -p 27018:27017 --name mongo2 --network mongoCluster mongo:5 mongod --replSet myReplicaSet --bind_ip localhost,mongo2
 
docker run -d --rm -p 27019:27017 --name mongo3 --network mongoCluster mongo:5 mongod --replSet myReplicaSet --bind_ip localhost,mongo3
```

You now have three containers running MongoDB. You can use `docker ps` to validate that they are running.

![img](https://webimages.mongodb.com/_com_assets/cms/kv9parnf8tw36l94x-image3.png?auto=format%252Ccompress)

### Initiate the Replica Set

The next step is to create the actual replica set with the three members. To do so, you will need to use the [MongoDB Shell](https://docs.mongodb.com/mongodb-shell/). This CLI (command-line interface) tool is available with the default MongoDB installation or [installed independently](https://www.mongodb.com/try/download/shell). However, if you don’t have the tool installed on your laptop, it is possible to use `mongosh` available inside containers with the `docker exec` command.

```
docker exec -it mongo1 mongosh --eval "rs.initiate({
 _id: \"myReplicaSet\",
 members: [
   {_id: 0, host: \"mongo1\"},
   {_id: 1, host: \"mongo2\"},
   {_id: 2, host: \"mongo3\"}
 ]
})"
```

This command tells Docker to run the `mongosh` tool inside the container named `mongo1`. `mongosh` will then try to evaluate the `rs.initiate()` command to initiate the replica set.

As part of the configuration object that is passed to `rs.initiate()`, you will need to specify the name of the replica set (`myReplicaSet`, in this case), along with the list of `members` that will be part of the replica set. The hostnames for the containers are the names of the containers as specified by the `--name` parameter in the `docker run` command. You can learn more about the possible options for configuring a replica set from the [documentation](https://docs.mongodb.com/manual/reference/replica-configuration/#std-label-replica-set-configuration-document).

If the command was successfully executed, you should see a message from the `mongosh` CLI indicating the version numbers of MongoDB and Mongosh, followed by a message indicating:

```
{ ok: 1 }
```

### Test and Verify the Replica Set

You should now have a running replica set. If you want to verify that everything was configured correctly, you can use the `mongosh` CLI tool to evaluate the `rs.status()` instruction. This will provide you with the status of your replica set, including the list of members.

```
docker exec -it mongo1 mongosh --eval "rs.status()"
```

You can also connect to your cluster using [MongoDB Compass](https://www.mongodb.com/products/compass) to [create a database](https://www.mongodb.com/basics/create-database) and [add some documents](https://docs.mongodb.com/compass/current/documents/insert/). Note that the data is created inside the container storage and will be destroyed when the containers are removed from the host system. To verify that your replica set is working, you can try stopping one of the containers with docker stop and try to read from your database again.

```
docker stop mongo1
```

The data will still be there. You can see that the cluster is still running by using `rs.status()` on the `mongo2` container.

```
docker exec -it mongo2 mongosh --eval "rs.status()"
```

You will still see the replica set, but notice that the first member is now down, and one of the other two members has been [elected](https://docs.mongodb.com/manual/core/replica-set-elections/) as the primary node. If you start your container `mongo1` again, you will be able to see it back in the replica set, but as a secondary member.