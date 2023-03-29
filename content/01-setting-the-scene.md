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

vvv

## Rotation At Scale

- **Goals:** Disruption-free, minimal ops, fully automated

- Now think about managing thousands of clusters

- How to orchestrate credentials rotation at such scale?

Notes:
- We talk about credentials for "Kubernetes as application"
- Where are we coming from? -> Mgmt of 1000s of clusters

---

## Project Gardener

![Gardener](../assets/gardener.svg) [github.com/gardener](https://github.com/gardener)

- Started in 2017
- Vanilla, conformant ![Kubernetes](../assets/kubernetes.svg) clusters as a service
- Early adopters of "Kubernetes in Kubernetes" model
- 100% open source
- Built for scale

![Inception](../assets/inception.gif)
<!-- .element style="width: 20%" -->

Notes:
- Think GKE for any infrastructure/environment
- Homogenous clusters
- Extensible
- Kubernetes-native (declarative cluster specification)

vvv

## Project Gardener

Declarative cluster specification:

```yaml
apiVersion: core.gardener.cloud/v1beta1
kind: Shoot
metadata:
  name: my-cluster
  namespace: garden-dev
spec:
  kubernetes:
    version: 1.26.1
  region: europe-central-1
  secretBindingName: my-cloud-provider-account
  provider:
    type: gcp
    workers:
    - name: pool-1
      minimum: 3
      maximum: 5
      machine:
        type: m5.large
        image:
          name: ubuntu
          version: 22.04
  networking:
    type: calico
    pods: 100.96.0.0/11
    nodes: 10.250.0.0/16
    services: 100.64.0.0/13
```

vvv

# TODO: Drop the following two slides. Replace them with a more "lightweight" (as far as this is possible in this project) architecture diagram.
# Goal/Explanation: We use Kubernetes everywhere, also for secrets management. Kubernetes components are our "workload/application".

## Kubernetes Architecture

![Architecture Kubernetes](../assets/architecture-kubernetes.png)
<!-- .element: class="r-stretch" -->

vvv

## Gardener Architecture

![Architecture Gardener](../assets/architecture-gardener.png)
<!-- .element: class="r-stretch" -->

Notes:
- Go back to previous slide to explain what exactly a "Shoot cluster" is
- Main difference: No master VMs, hence better resource utilization and less TCO
- Kubernetes API everywhere available for management
