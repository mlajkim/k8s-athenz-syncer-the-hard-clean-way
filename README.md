# k8s-athenz-syncer-the-hard-clean-way

`k8s-athenz-syncer-the-hard-clean-way` [^1] is a Kubernetes controller that syncs Athenz roles and policies with Kubernetes RBAC, just like [Athenz/k8s-athenz-syncer](https://github.com/AthenZ/k8s-athenz-syncer), but in a more manual and educational way.

## Philosophy

This project tries to demonstrate the sync between Athenz Role and Kubernetes RBAC only using `ZMS API calls`. This project is also designed not to care about the performance or scalability. The main goal is to help users understand how the sync works under the hood.


## Build Locally

To build the manager binary locally you can run:

```sh
make build
```

## Code Structure

### main.go

The entry point for the controller manager. It sets up the manager, registers the controllers, and starts the manager.

### internal/config

Handles required configurations for the controller manager.


### internal/controller

List of controllers that this operator `k8s-athenz-syncer-the-hard-clean-way` can do.

Controllers's code should be neat so that it is easier to grab the flow of reconciliation logic.

Core jobs:

- `NamespaceController`: Use namespaces as SSOT, and syncs:
  - Athenz Sub Domains, if not exist
  - Athenz Default Roles, if not exist
  - Kubernetes necessary RBAC Roles, if not exist
- `AthenzRoleController`: Every minute, check all athenz roles under certain Parent domain, and syncs:
  - Kubernetes RBAC Roles, if not synced


### internal/syncer

List of core syncer logics that controllers use to perform the sync between Athenz and Kubernetes.


### pkg/athenz

> [!TIP]
>`pkg` does not include any business logics, or config imports.

Self-created athenz library to interact with Athenz ZMS server using ZMS APIs. I could have used the official Athenz Go client library, but I wanted to keep this project simple and focused on demonstrating the sync logic.

### pkg/athenzutil

Even lower level utility functions to support `pkg/athenz`.


<!-- Footnote -->

[^1]: This project's name is inspired by [Kelsey Hightower's Kubernetes The Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way)

<!-- Footnote -->
