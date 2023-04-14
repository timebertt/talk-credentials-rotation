# Setting the Scene

---

## Problem Statement

Each Kubernetes cluster comes with a lot of credentials:

- 🪪 Certificate authorities
- 📝 Server and client certificates
- 🏷️ ServiceAccount tokens
- 🗝️ Static tokens
- 🔐 ETCD encryption key
- 🔑 SSH key pair

Notes:

- As an application developer/user, you might not know be specifically aware about those.
- A lot of credentials are involved to keep a Kubernetes cluster securely up and running.

vvv

![Rotation](../assets/rotate.gif)

Notes:

- Don't worry if you feel like this now
- Credentials management in Kubernetes is hard

---

## Open-Source Project Gardener

![Gardener](../assets/gardener.svg) <!-- .element: class="img-inline" style="height: 2em;" --> [github.com/gardener](https://github.com/gardener)

- Vanilla, conformant ![Kubernetes](../assets/kubernetes.svg) <!-- .element: class="img-inline" --> clusters as a service
- Early adopters of "Kubernetes in Kubernetes" model
- Native Kubernetes extension
- Universal Kubernetes at scale
- Started in 2017

![Inception](../assets/inception.gif)
<!-- .element style="width: 20%" -->

Notes:

- Think GKE for any infrastructure/environment
- Homogenous clusters
- Extensible
- Kubernetes-native (declarative cluster specification, custom resources)

vvv

## Kubernetes as Workload on Kubernetes

![Architecture Kubernetes](../assets/gardener-architecture.png)
<!-- .element: class="r-stretch" -->

Notes:

- We use Kubernetes everywhere, also for secret management
- Kubernetes is our application/our workload.

vvv

## Rotation At Scale

- Gardener manages thousands of clusters

- How to orchestrate credentials rotation at such scale?

- **Goals:** Disruption-free, minimal ops, fully automated

Notes:

- We talk about credentials for "Kubernetes as application"
- Where are we coming from? -> Mgmt of 1000s of clusters
