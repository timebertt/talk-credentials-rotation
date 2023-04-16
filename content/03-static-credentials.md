# Rotating Static Credentials

notes:

- now that we know how to use short-lived credentials
- some credentials cannot be short-lived

---

## Static Credentials 🔐

![Shoot Credentials](../assets/static-credentials.excalidraw.png)
<!-- .element: class="r-stretch" -->

notes:

- static credentials in every k8s cluster
- CA: kube-apiserver serving cert
- client CA: kubelet client certs
- CAs: typically valid for long time (gardener: 10y, GKE: 30y)
- other credentials (etcd encryption, SA signing key): no expiration
- best practice: frequently rotate static credentials
- at scale: highly-automated, disruption-free

vvv

## Solution 💡

Rotation in two phases:

1. issue new credentials, accept both old and new
2. invalidate old credentials

notes:

- phase 1
- phase 2
- clients need to refresh their credentials before triggering phase 2

vvv

<!-- https://github.com/gardener/gardener/blob/master/docs/development/secrets_management.md#certificate-signing -->

## Server Certificates 🗄

<table class="no-borders">
<thead>
<tr>
<th align="right">phase</th>
<th>cert signed by</th>
<th align="center">clients trust</th>
</tr>
</thead>
<tbody>
<tr class="fragment">
<td align="right">0</td>
<td><span class="cred old">old CA</span></td>
<td align="center"><span class="cred old">old CA</span></td>
</tr>
<tr class="fragment">
<td align="right">1</td>
<td><span class="cred old">old CA</span></td>
<td align="center"><span class="cred bundle">old+new CA</span></td>
</tr>
<tr class="fragment">
<td align="right">2</td>
<td><span class="cred new">new CA</span></td>
<td align="center"><span class="cred new">new CA</span></td>
</tr>
</tbody>
</table>

notes:

- before
- phase 1: server certificates signed by old CA, clients add new CA to their CA bundles asynchronously
- phase 2: server certificates signed by new CA, clients drop the old CA from their CA bundles

vvv

## Client Certificates 🧑‍💻

<table class="no-borders">
<thead>
<tr>
<th align="right">phase</th>
<th>cert signed by</th>
<th align="center">servers trust</th>
</tr>
</thead>
<tbody>
<tr class="fragment">
<td align="right">0</td>
<td><span class="cred old">old CA</span></td>
<td align="center"><span class="cred old">old CA</span></td>
</tr>
<tr class="fragment">
<td align="right">1</td>
<td><span class="cred new">new CA</span></td>
<td align="center"><span class="cred bundle">old+new CA</span></td>
</tr>
<tr class="fragment">
<td align="right">2</td>
<td><span class="cred new">new CA</span></td>
<td align="center"><span class="cred new">new CA</span></td>
</tr>
</tbody>
</table>

notes:

- before
- phase 1: servers add new CA to their CA bundles, clients get new certificates asynchronously
- phase 2: servers stops accepting certificates signed by the old CA
- bundles approach also works for other credentials: SA signing key

vvv

## Key Elements 🔑

- clients need to refresh their credentials after preparation
- clients trigger completion once ready
- automatic rotation for non-user-facing credentials

notes:

- clients need to comply -> refresh
  - humans (e.g., kubeconfig CA bundle)
  - cluster components (e.g., kubelet client cert)
- we cannot control all clients
- trigger completion once ready
- if controlling all clients: automatic rotation

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
