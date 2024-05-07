

https://blog.devgenius.io/how-to-deploy-a-mongodb-replicaset-using-docker-compose-a538100db471

# How to deploy a MongoDB replica set using docker-compose

![img](https://miro.medium.com/max/700/1*HTg4qiPylKW_XjC4RSlhRQ.jpeg)

Image by [Tumisu, please consider ☕ Thank you! 🤗](https://pixabay.com/users/tumisu-148124/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1954920) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=1954920)

# Introduction

In this small MongoDB overview, we are going through some important MongoDB concepts related to replica sets. After some concept presentation, we will see some examples in action using docker-compose. We will use as the main source of information the official MongoDB documentation that can be found in the reference section. Without further due let’s get started.

This story is organized as follows:

- **What is a replica set?**
- **Why should you need a replica set?**
- **Oplog**
- **How does replication work in MongoDB?**
- **What is an arbiter in MongoDB replica sets?**
- **Security**
- **Strategies for deploying your replica set**
- **Replication VS Sharding**
- **Getting practical**
- **BONUS**
- **Conclusion**

# What is a replica set?

The concept of a replica set in mongo simply means that instead of having one process of mongo running you will have more than one. This way you can achieve higher availability of your data and redundancy, two important quality attributes when it comes to production-like environments. MongoDB at the time of this story allows up to 50 instances in a replica set.

# Why should you need a replica set?

Many scenarios might convince you to implement a replica set within your environment but here are some of them(not necessarily solved by the Proof of concept here shown):

- Depending on the configuration, assuring higher reads and writes speed(at some degree);
- Having databases “copies” for different purposes(reads, writes, backup, reporting, disaster recovery — at this point a multiple region setup needs to be configured).

In case you are now convinced about what it is and how could you benefit from it let’s mention other interesting features!

# Oplog

While this term feels strange in this sequence of topics an introduction of what is Oplog plays an important role since the replication of MongoDB takes advantage of this concept, therefore essential in understanding the full picture of what replica sets are.

Oplog stands for operations log, and it is a special collection within the replica set, its main goal is to register all the changes made in the different collections. Since this as you might suspect will be used by some nodes, there might be issues in the configuration of its size(to not create bottlenecks), so be sure to check and understand how to fine-tune it to your use case.

# How does replication work in MongoDB?

As already mentioned a replica set has many nodes that together manage a data set. These nodes can keep mirroring the data with different strategies and have different preferences in the way they do it(preference in the replication, read nodes, backup nodes, etc…). When a replica set is setup, one of the nodes becomes PRIMARY by default while others become SECONDARY. This is one of the concerns that our example has in mind. In our scenario, the most beneficial behaviour is to ensure that a defined node, always becomes PRIMARY. The main role of the PRIMARY node is to ensure that all writes will be processed by it. Now let’s imagine you shutdown the PRIMARY node, what happens? In this case, an “election” takes place so that a SECONDARY becomes the new PRIMARY. The election can be triggered by a bunch of reasons but many of the times come down to the same cause, the heartbeat(like a ping endpoint). PRIMARY nodes keep telling other nodes that they are “alive” when this doesn’t happen(the configuration of the intervals for the heartbeat are configurable), the first secondary node that realizes this asks for an election.

The role of the SECONDARY role is to asynchronously replicate the data that is in the PRIMARY node. Now we may wonder, how does the replication work? Well, replication takes place by sending oplog entries. Oplog holds all the changes in the database, this means that instead of just receiving the same “commands” that the PRIMARY node receives, the SECONDARY nodes are fed with oplogs(they are specially designed for this) from the PRIMARY so that they replicate a similar state of their dataset.

# What is an arbiter in MongoDB replica sets?

As simple as it gets now that we know what is an “election”, an arbiter is just a role that is only responsible to participate in elections. It does not help with the data redundancy quality attribute. So why should we want to have an arbiter then? One of the possible answers is, that you don’t have enough votes or have an even number of votes. If you can’t increase the number of secondary instances(imagine cost-wise or resource-wise), you can add one arbiter to do the trick.

# Security

You can configure your replica set to use authorization, in this case, all the communication exchanged by the different instances will be encrypted. In some cases, you might also be required to configure security between your replica set members.

# Strategies for deploying your replica set

For production environments, the recommended approach by MongoDB documentation is 3 member replica set. While MongoDB recommends it is also advisable that you run the different scenarios that you have prepared for your system against your current strategy. If you want to deploy in different regions, this will affect your MongoDB strategy and you will probably want instead of 3 members, perhaps have a fourth that will be in a different region just backing up data. On the other hand, if you have high fault tolerance requirements having a minimum of 5 or 6 members can also be considered.

The bottom line is that you will need to think about this thoroughly. The noticeable good thing is that changing configurations seems pretty straightforward so, considering you have the time to experiment you should be fine by making this an iterative process(while developing).

# Replication vs Sharding

This topic would make enough information to cover in a whole different story, but it is important to sometimes understand this concept which happens to be used together in the same context. Replication is achieved, as explained in this story, by configuring new nodes and providing a defined role to each one of them. Assuming this, let’s think about the following scenario, what happens if you reach the limit of performance(writes and reads) on your PRIMARY node, what can be done to increase it? One of the options that you should consider is sharding, by sharding you will “split your node” and you will make each “shard” take care of some operations. So instead of having a full data set per node, you will have a “chunk” of it. This is much more performant in many cases.

# Getting practical

We will use this section to present the different files needed to run a simple example, be sure to follow the instructions and pay special attention to where the different files should be put for docker-compose mounts them properly.

To achieve a simple 3 PSS node architecture(Primary-Secondary-Secondary) you can use a docker-compose as follows:

<iframe src="https://blog.devgenius.io/media/d6559b10a8f942d4810e9fd26ae34786" allowfullscreen="" frameborder="0" height="1007" width="692" title="docker-compose-replicaset.yml" class="fo aq as ag ce" scrolling="auto" style="box-sizing: inherit; height: 1007px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

Basically with this configuration, we will make sure that the **mongo1** will always be our primary node, for all the nodes to communicate with each other we used a network, called “**mongo-network”**. We especially load all scripts in its directories.

As mentioned throughout this story for you to get always the same primary you will need to choose the instances to have higher priority, this can be achieved using a field with the same name “**priority**”, to apply these changes we are going to use a script that will be mounted in the **mongo1** container(all these scripts should be put inside a new folder called “**scripts**”):

<iframe src="https://blog.devgenius.io/media/75d4ad161d4b84d16d9702865980ede5" allowfullscreen="" frameborder="0" height="787" width="692" title="rs-init.sh" class="fo aq as ag ce" scrolling="auto" style="box-sizing: inherit; height: 787px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

After looking at the scripts you can notice that there are some delays declared, this is used to give time for the containers to start and, most importantly, later to let the election occur so that we have our mongo1 ready to start receiving connections as the PRIMARY node.

Next and to get the mongo ready to accept your inserts and creation of new users and databases we use a small javascript file that can be expanded to suit your needs(also to be added inside the “**scripts**” folder):

<iframe src="https://blog.devgenius.io/media/0c7ab1c899497fb1b16c7af8592d1aee" allowfullscreen="" frameborder="0" height="83" width="692" title="init.js" class="fo aq as ag ce" scrolling="auto" style="box-sizing: inherit; height: 83px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

Finally and to trigger this example we can use a small bash script to glue everything together:

<iframe src="https://blog.devgenius.io/media/f72d7973e167565674892cf80a7dd2d5" allowfullscreen="" frameborder="0" height="347" width="692" title="startReplicaSetEnvironment.sh" class="fo aq as ag ce" scrolling="auto" style="box-sizing: inherit; height: 347px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

This script does a little bit more than just running the docker-compose file, it also starts by cleaning your docker environment and ends by executing the script that we previously mounted into the **mongo1**.

**Note:** Following some discussion in the comments, be aware that executing **line 6** and **line 7** will clean your docker environment, I personally used this script, to help me iterate over the different experiments I did, in case it doesn’t suit your use-case feel free to remove/comment them.

# BONUS

Let us imagine that you instead of having a 3-node architecture, just want to have one node working as a replica set, this can be achieved in a simple docker-compose like this:

<iframe src="https://blog.devgenius.io/media/812388aadc37254b706d75fda003ac08" allowfullscreen="" frameborder="0" height="518" width="692" title="docker-compose-single-node.yml" class="fo aq as ag ce" scrolling="auto" style="box-sizing: inherit; height: 518px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

# Conclusion

In this story, a quick overview of some important concepts of MongoDB was presented, and some ideas about how to choose an architecture and what factors to consider. In the final part, two practical examples show what are replica sets and how to configure them. For more extensive information be sure to check the official MongoDB docs!

Thanks for spending your precious time reading this story, I hope it helped you in some way. Any feedback you have be sure to use the comments below. If you liked this story clap and if you want to keep following my stories feel free to follow as well.