# Rotating Static Credentials

---

## Requirements

- frequently rotate static credentials
- must be disruption-free
- must be minimal-ops

vvv

## Solution

rotation in two steps

1. issue new credentials, accept both old and new
2. invalidate old credentials

vvv

## Key Elements
 
- clients need to refresh their credentials
- clients trigger completion once ready
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
- show the zoo: `k -n shoot--local--local get secret -l managed-by -L name`

vvv

![Rotation](../assets/rotate.gif)

---

## Bundles

- certificate authorities and related certificates
- service account signing key
- etcd encryption key

notes:
- TODO: move after vertical slides

vvv

<!-- https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing -->

## Server Certificates

- step 1: server certificates signed by old CA, clients add new CA to their CA bundles asynchronously
- step 2: server certificates signed by new CA, clients drop the old CA from their CA bundles

vvv

## Client Certificates

- step 1: servers add new CA to their CA bundles, clients get new certificates asynchronously
- step 2: servers stops accepting certificates signed by the old CA

---

## TODO: trigger CA rotation

- keep previous list of secrets open
- show rotated secrets (bundles)

---

## Secrets Manager

- manages all types of credentials
- TODO: code with example call
- kubernetes primitives: `Secrets`
- TODO: code with example secret yaml
- immutable secrets
  - scalability
  - immutable infrastructure paradigm
- knows when to rotate
  - based on trigger
  - based on validity
- server/clients always use bundle of old and new CA secret during rotation

notes:
- TODO: add link to https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing

vvv

## TODO: auto-rotation for non-user-facing CAs

- internal CAs: 30d validity

---

## TODO: conclusion

- two-step rotation for static credentials
- implemented for kubernetes as workload, concept is applicable to most workloads
