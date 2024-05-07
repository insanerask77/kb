# Kubernetes KIND Cheat Sheet

![Oleg Sucharevich](https://miro.medium.com/v2/resize:fill:88:88/0*_q3aLNEu_zF23lBe.)

![ITNEXT](https://miro.medium.com/v2/resize:fill:48:48/1*yAqDFIFA5F_NXalOJKz4TA.png)

[Oleg Sucharevich](https://medium.com/@olegsucharevich?source=post_page-----2605da77984--------------------------------)

·

[Follow](https://medium.com/m/signin?actionUrl=https%3A%2F%2Fmedium.com%2F_%2Fsubscribe%2Fuser%2Feb8896063d4d&operation=register&redirect=https%3A%2F%2Fitnext.io%2Fkubernetes-kind-cheat-shee-2605da77984&user=Oleg+Sucharevich&userId=eb8896063d4d&source=post_page-eb8896063d4d----2605da77984---------------------post_header-----------)

Published in

ITNEXT

·

2 min read

·

Oct 29, 2021



44









![img](https://miro.medium.com/v2/resize:fit:700/1*pDiLy3VezAqPZMYbbvODpA.jpeg)

> [kind](https://sigs.k8s.io/kind) is a tool for running local Kubernetes clusters using Docker container “nodes”.
> kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.

# Autocompletion

## Bash

```
# Setup the current shell
source <(kind completion bash)########
## OR ##
######### Update permanently
echo “source <(kind completion bash) >> ~/.bashrc”
```

## ZSH

```
source <(kind completion zsh)
```

# Basic

Create kind cluster named cncf-cheat-sheet

```
kind create cluster — name cncf-cheat-sheet
```

Create cluster and wait for all the components to be ready

```
kind create cluster — wait 2m
```

Get running clusters

```
kind get clusters
```

Delete kind cluster named cncf-cheat-sheet

```
kind delete cluster — name cncf-cheat-sheet
```

# Advanced Configuration

Use kind.yaml config file for more advanced use cases

## Ports

[Info](https://kind.sigs.k8s.io/docs/user/configuration/#extra-port-mappings)
Map port 80 from the cluster control plane to the host.

```
cat <<EOF | kind create cluster — name cncf-cheat-sheet — config -
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
 extraPortMappings:
 — containerPort: 80
 hostPort: 80
 protocol: TCP
EOF
```

## Mount Directories

[Info](https://kind.sigs.k8s.io/docs/user/configuration/#extra-mounts)
Mount current directory into clusters control plane located at `/app`

*NOTE: MacOS users: make sure to share resources in docker-for-mac preferences*

```
cat <<EOF | kind create cluster — name cncf-cheat-sheet — config -
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
 extraMounts:
 — hostPath: .
 containerPath: /app
EOF
```

## Add Local Registry

[Info](https://kind.sigs.k8s.io/docs/user/local-registry/)

Step 1: Create local registry

```
docker run -d — restart=always -p 127.0.0.1:5000:5000 — name cncf-cheat-sheet-registry registry:2
```

Step 2: Create cluster

```
cat <<EOF | kind create cluster — name cncf-cheat-sheet — config -
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
 [plugins.”io.containerd.grpc.v1.cri”.registry.mirrors.”localhost:5000"]
 endpoint = [“http://cncf-cheat-sheet-registry:5000"]
nodes:
- role: control-plane
EOF
```

Step 3: Connect registry with created network

```
docker network connect kind cncf-cheat-sheet-registry
```

Step 4: Update cluster about new registry

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
 name: local-registry-hosting
 namespace: kube-public
data:
 localRegistryHosting.v1: |
 host: “localhost:5000”
EOF
```

## Multiple Workers

[Info](https://kind.sigs.k8s.io/docs/user/configuration/#nodes)
The default configuration will create cluster with one node (control-plane).

```
cat <<EOF | kind create cluster — name cncf-cheat-sheet — config -
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
EOF
```