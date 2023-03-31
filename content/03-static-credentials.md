# Rotating Static Credentials

---

## Static Credentials

![Shoot Credentials](../assets/01-shoot-credentials-before.excalidraw.png)
<!-- .element: class="r-stretch" -->

notes:
- TODO: redraw diagram, not all service account tokens are static, remove observability credentials
- CAs: valid for 10 years by default
- other credentials: no expiration
- previously: only static token and ssh key pair rotatable

vvv

![Rotation](../assets/rotate.gif)

vvv

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

<!-- https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing -->

## Server Certificates

- step 1: server certificates signed by old CA, clients add new CA to their CA bundles asynchronously
- step 2: server certificates signed by new CA, clients drop the old CA from their CA bundles

vvv

## Client Certificates

- step 1: servers add new CA to their CA bundles, clients get new certificates asynchronously
- step 2: servers stops accepting certificates signed by the old CA

notes:
- bundles approach also works for other credentials: SA signing key

vvv

## Key Elements

- clients need to refresh their credentials after preparation
- clients trigger completion once ready
- automatic rotation for non-user-facing credentials

notes:
- clients: humans (e.g., kubeconfig CA bundle)
- machines (e.g., kubelet client cert)

---

## Secrets Manager

- our implementation in go
- manages all types of credentials

vvv

## Requesting a Server Cert

**Live Coding!**

![Live Coding](../assets/live-coding.gif)
<!-- .element: class="r-stretch" -->

notes:
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
- TODO: rotate CA?

vvv

## Rotation

- secrets manager knows when to rotate
  - based on config change
  - based on trigger
  - based on validity
- secrets manager always returns bundle of CAs
- secrets manager always [signs with correct CA](https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing)

notes:
- bundles: could be one or two CAs

vvv

## Auto-Rotation

- activated for non-user-facing credentials
- rotation is prepared when approaching end of validity
- rotation is completed after 24h
- fully-automated + disruption-free!

notes:
- internal CAs: 30d validity
