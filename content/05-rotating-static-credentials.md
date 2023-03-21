# Rotating Static Credentials

---

## Requirements

- frequently rotate static credentials
- must be disruption-free
- must be minimal-ops

---

## Solution

rotation in two phases

1. issue new credentials, accept both old and new
2. invalidate old credentials

vvv

## Key Elements
 
- clients need to refresh their credentials
- clients trigger completion once ready
- bundle secrets
- automatic rotation for non-user-facing credentials

notes:
- clients: humans (e.g., kubeconfig CA bundle), machines (e.g., kubelet client cert)

---

## Involved Credentials

![Kubernetes logo](../assets/01-shoot-credentials-before.excalidraw.png)
<!-- .element: class="r-stretch" -->

<!--
- certificate authorities and related certificates
- service account signing key
- etcd encryption key
- ssh key pair for worker nodes
- observability credentials
- (static token kubeconfig – if enabled)
-->

notes:
- CAs: valid for 10 years by default
- other credentials: no expiration
- previously: only static token and ssh key pair rotatable
- `k -n shoot--local--local get secret -l managed-by -L name`

---

## Bundles

- certificate authorities and related certificates
- service account signing key
- etcd encryption key

vvv

<!-- https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing -->

## Server Certificates

- phase 1: servers use old certificates, clients add new CA to their CA bundles asynchronously
- phase 2: servers get new certificates, clients drop the old CA from their CA bundles

vvv

## Client Certificates

- phase 1: servers add new CA to their CA bundles, clients get new certificates asynchronously
- phase 2: servers stops accepting certificates signed by the old CA

---

## Secrets Manager

- manages all types of credentials
- TODO: code with example call
- kubernetes primitives: `Secrets`
- TODO: code with example secret
- immutable secrets
  - scalability
  - immutable infrastructure paradigm
- knows when to rotate
  - based on trigger
  - based on validity
- clients always use bundle
