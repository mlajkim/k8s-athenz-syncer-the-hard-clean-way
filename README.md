# k8s-athenz-syncer-the-hard-clean-way

`k8s-athenz-syncer-the-hard-clean-way` [^1] is a Kubernetes controller that syncs Athenz roles into Kubernetes RBAC, just like [Athenz/k8s-athenz-syncer](https://github.com/AthenZ/k8s-athenz-syncer), but in a more manual and educational way.

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

## To build for the first time

This operator requires the following:

- Running kubernetes cluster
- Running Athenz Sever

### How to build "Running kubernetes cluster"

Simply do the following:

```sh
brew install kind && kind create cluster
```

### How to build "Running Athenz Sever"

[ctyano's athenz-distribution](https://github.com/ctyano/athenz-distribution) is the easiest way to run Athenz server locally.

```sh
git clone https://github.com/ctyano/athenz-distribution.git athenz_dist
cd athenz_dist
make clean-kubernetes-athenz deploy-kubernetes-athenz
kubectl -n athenz port-forward deployment/athenz-ui 4443:4443
kubectl -n athenz port-forward deployment/athenz-ui 3000:3000
```

## To build locally

If you want to see in action, you can build the operator binary locally by running. However, please note that this project depends on the following, which I share it :

- Running kubernetes cluster
- Running Athenz Sever

To build this project locally, do the following:
```sh
make build
```



<!-- Footnote -->

[^1]: This project's name is inspired by [Kelsey Hightower's Kubernetes The Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way)

<!-- Footnote -->
