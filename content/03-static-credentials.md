# Rotating Static Credentials

---

## Static Credentials

![Shoot Credentials](../assets/01-shoot-credentials-before.excalidraw.png)
<!-- .element: class="r-stretch" -->

notes:
- TODO: redraw diagram, not all service account tokens are static, exclude observability credentials
- CAs: valid for 10 years by default
- other credentials: no expiration
- previously: only static token and ssh key pair rotatable
- show the zoo: `k -n shoot--local--local get secret -l managed-by -L name`

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

## Key Elements

- clients need to refresh their credentials after preparation
- clients trigger completion once ready
- automatic rotation for non-user-facing credentials

notes:
- clients: humans (e.g., kubeconfig CA bundle)
- machines (e.g., kubelet client cert)

---

<!-- https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing -->

## Server Certificates

- step 1: server certificates signed by old CA, clients add new CA to their CA bundles asynchronously
- step 2: server certificates signed by new CA, clients drop the old CA from their CA bundles

vvv

## Client Certificates

- step 1: servers add new CA to their CA bundles, clients get new certificates asynchronously
- step 2: servers stops accepting certificates signed by the old CA

vvv

## Bundles

- certificate authorities and related certificates
- service account signing key
- etcd encryption key

vvv

## Credentials Rotation

![Demo](../assets/show-me-a-demo.jpg)
<!-- .element: class="r-stretch" -->

notes:
- keep previous list of secrets open
- trigger rotation of all shoot credentials
- show rotated secrets (bundles)

---

## Secrets Manager

- our implementation in go
- manages all types of credentials

vvv

## Requesting a Server Cert

```go[|1|3-8|9|10]
secret, err := secretsManager.Generate(
  ctx,
  &secrets.CertificateSecretConfig{
    CertType:   secrets.ServerCert,
    Name:       "kube-apiserver",
    CommonName: "kube-apiserver",
    DNSNames:   []string{"kube-apiserver", "kube-apiserver.shoot--local--local.svc", /* ... */},
  },
  secretsmanager.SignedByCA("ca"),
  secretsmanager.Rotate(secretsmanager.InPlace),
)
```

notes:
- generates a new server certificate for kube-apiserver
- or retrieves an existing certificate from the cluster
- ensures the secret has the given config
- signs with the correct version of the cluster CA
- drops the old certificate when rotating

vvv

## Result

```yaml[1-5|18-19|6-9|10-11|12-14|16]
apiVersion: v1
kind: Secret
metadata:
  name: kube-apiserver-b937fe88
  namespace: shoot--local--local
  labels:
    managed-by: secrets-manager
    manager-identity: gardenlet
    name: kube-apiserver
    checksum-of-config: "11635920296491681218"
    checksum-of-signing-ca: 2ade89ab87c3b21c3eef0e015d0ee98507374e030509ea4cbe2503feab43fbd
    issued-at-time: "1679647057"
    valid-until-time: "1995266257"
    last-rotation-initiation-time: ""
type: kubernetes.io/tls
immutable: true
data:
  tls.crt: LS0t...
  tls.key: LS0t...
```

vvv

## Kubernetes Primitives

- plain Kubernetes `Secrets`
- immutable secrets
  - scalability
  - immutable infrastructure paradigm

vvv

## Rotation

- secrets manager knows when to rotate
  - based on config change
  - based on trigger
  - based on validity
- clients use bundle of old/new server CA
- servers use bundle of old/new client CA

notes:
- TODO: add link to https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing

vvv

## Auto-Rotation

- activated for non-user-facing CAs
  - internal CAs: 30d validity
- rotation is prepared when approaching end of validity
  - use bundles everywhere
- rotation is completed after 24h
- fully-automated + disruption-free!

---

## TODO: conclusion

- two-step rotation for static credentials
- implemented for kubernetes as workload, concept is applicable to most workloads

notes:
- move to general summary?
