# k8s-athenz-syncer-the-hard-clean-way

`k8s-athenz-syncer-the-hard-clean-way` [^1] is a Kubernetes controller that syncs Athenz roles into Kubernetes RBAC, just like [Athenz/k8s-athenz-syncer](https://github.com/AthenZ/k8s-athenz-syncer), but in a more manual and educational way.

<!-- TOC -->

- [k8s-athenz-syncer-the-hard-clean-way](#k8s-athenz-syncer-the-hard-clean-way)
  - [Features](#features)
  - [How to run locally](#how-to-run-locally)
    - [For those who want to run easy way](#for-those-who-want-to-run-easy-way)
    - [Run locally](#run-locally)

<!-- /TOC -->

## Features

Operator `k8s-athenz-syncer-the-hard-clean-way` creates the following when you simply create a namespace in your Kubernetes cluster:
- Athenz domain under certain parent domain (e.g., `eks.users`)
- Athenz roles under the created domain, that you can define in the config file
- Kubernetes RBAC Roles that correspond to the created Athenz roles, that you define in the config file

![Demo](./assets/01_create_ns.gif)

Operator `k8s-athenz-syncer-the-hard-clean-way` periodically polls Athenz roles under certain parent domain (e.g., `eks.users`), and syncs the members of the Athenz roles into corresponding Kubernetes RBAC Roles.

![Demo](./assets/02_polling_athenz_roles.gif)

Operator `k8s-athenz-syncer-the-hard-clean-way` makes sure that if you delete members from Athenz roles, the members are also removed from corresponding Kubernetes RBAC Roles.

![Demo](./assets/03_remove_athenz_role_members.gif)

## How to run locally

This operator requires the following:

- Running kubernetes cluster
- Running Athenz Server


### For those who want to run easy way

> [!TIP]
> If you know what you are doing, you can always skip this section go build your own way here: [Run locally](#run-locally)


The following command sets up:

- A simple test directory for clean start
- A local Kubernetes cluster using [kind](https://kind.sigs.k8s.io/)
- Athenz server deployed into the local Kubernetes cluster using [Athenz Distribution](https://github.com/ctyano/athenz-distribution)
- Clones this project into the test directory & copy necessary certs and keys for Athenz admin user

```sh
brew install kind && kind create cluster

_tmp_dir=$(date +%y%m%d_%H%M%S_k8s_athenz_syncer_the_hard_clean_way)
mkdir -p ~/test_dive/$_tmp_dir && cd ~/test_dive/$_tmp_dir

git clone https://github.com/ctyano/athenz-distribution.git athenz_distribution
make -C ./athenz_distribution clean-kubernetes-athenz deploy-kubernetes-athenz
```

Once the manifests above is done, set up ZMS server:

```sh
kubectl -n athenz port-forward deployment/athenz-ui 4443:4443 &
kubectl -n athenz port-forward deployment/athenz-ui 3000:3000 &
```

Clone this project, with copying necessary certs and keys for Athenz admin user:

```sh
git clone https://github.com/mlajkim/k8s-athenz-syncer-the-hard-clean-way.git k8s_athenz_syncer_the_hard_clean_way

cp ./athenz_distribution/certs/athenz_admin.cert.pem ./k8s_athenz_syncer_the_hard_clean_way/certs/athenz_admin.cert.pem
cp ./athenz_distribution/keys/athenz_admin.private.pem ./k8s_athenz_syncer_the_hard_clean_way/keys/athenz_admin.private.pem
```

Run the following command, and simply hit `Enter` keys with default values:

```sh
make -C ./k8s_athenz_syncer_the_hard_clean_way run
```

### Run locally

> [!TIP]
> If you want to see running without thinking too much, check out: [For those who want to run easy way](#for-those-who-want-to-run-easy-way)

To run this operator locally, do the following:

```sh
git clone https://github.com/mlajkim/k8s-athenz-syncer-the-hard-clean-way.git k8s_athenz_syncer && cd k8s_athenz_syncer
make run
```


<!-- Footnote -->

[^1]: This project's name is inspired by [Kelsey Hightower's Kubernetes The Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way)

<!-- Footnote -->
