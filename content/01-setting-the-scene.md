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

## Project Gardener

![Gardener](../assets/gardener.svg) [github.com/gardener](https://github.com/gardener)

- Started in 2017
- Vanilla, conformant ![Kubernetes](../assets/kubernetes.svg) clusters as a service
- Early adopters of "Kubernetes in Kubernetes" model
- Declarative cluster specification (custom resources)
- 100% open source
- Built for scale

<!--TODO(timebertt): Fix layout-->
![Inception](../assets/inception.gif)
<!-- .element style="width: 20%" -->

Notes:
- Think GKE for any infrastructure/environment
- Homogenous clusters
- Extensible
- Kubernetes-native (declarative cluster specification)

vvv

## Gardener Architecture

![Architecture Kubernetes](../assets/gardener-architecture.png)
<!-- .element: class="r-stretch" -->

Notes:
- Short demo: Show this in a local cluster
- We use Kubernetes everywhere, also for secret management
- Kubernetes is our application/our workload.

vvv

## Rotation At Scale

- **Goals:** Disruption-free, minimal ops, fully automated

- Now think about managing thousands of clusters

- How to orchestrate credentials rotation at such scale?

Notes:
- We talk about credentials for "Kubernetes as application"
- Where are we coming from? -> Mgmt of 1000s of clusters
