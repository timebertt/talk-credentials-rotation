# Rotating Static Credentials

---

## Static Credentials 🔐

![Shoot Credentials](../assets/static-credentials.excalidraw.png)
<!-- .element: class="r-stretch" -->

notes:

- every k8s cluster: CAs and server/client certs
- CA: kube-apiserver serving cert
- client CA: kubelet client certs
- CAs: typically valid for long time (gardener: 10y, GKE: 30y TODO)
- other credentials (etcd encryption, SA signing key): no expiration
- frequently rotate static credentials

vvv

## Solution 💡

rotation in two phases

1. issue new credentials, accept both old and new
2. invalidate old credentials

vvv

<!-- https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing -->

## Server Certificates 🗄

|  phase |   trust    | cert signed by |
|-------:|:----------:|----------------|
|      0 |   old CA   | old CA         |
|      1 | old+new CA | old CA         |
|      2 |   new CA   | new CA         |

vvv

## Client Certificates 🧑‍💻

|  phase |   trust    | cert signed by |
|-------:|:----------:|----------------|
|      0 |   old CA   | old CA         |
|      1 | old+new CA | new CA         |
|      2 |   new CA   | new CA         |

notes:

- bundles approach also works for other credentials: SA signing key

vvv

## Key Elements 🔑

- clients need to refresh their credentials after preparation
- clients trigger completion once ready
- automatic rotation for non-user-facing credentials

notes:

- clients: humans (e.g., kubeconfig CA bundle)
- machines (e.g., kubelet client cert)

---

## Secrets Manager 👔

- our implementation in ![go](../assets/gopher.png) <!-- .element: class="img-inline" -->
- manages all types of credentials

vvv

## Requesting a Server Cert

**Live Coding!**

![Live Coding](../assets/live-coding.gif)
<!-- .element: style="height: 300px" -->

notes:

- preparation: `cd code && make kind-up`
- `watch k -n secrets-manager get secret -L name,bundle-for`
- run once w/o server cert
- show generated CA and bundle secret
- add code for server cert
- run code again
- show CA was not regenerated
- show new server cert in YAML
  - plain Kubernetes secrets
  - labels for locating and identifying when to rate
  - immutable secrets: scalability
- rotate CA
- `k -n secrets-manager get secret -l bundle-for=demo-ca -oyaml | yq '.items[].data["bundle.crt"]' | base64 -d`
- wrap-up:
  - plain k8s secrets as source of truth
  - this is how we implement two-phase approach
  - applicable to other applications as well
  - go library available

vvv

## Next Level: Auto-Rotation 🔁

- activated for non-user-facing credentials
- phase 1 triggered when approaching end of validity
- phase 2 triggered 24h later
- fully-automated + disruption-free!

notes:

- more secrets manager features (generate options)
- auto-rotation activated for non-user-facing CAs
- internal webhook CAs: 30d validity
