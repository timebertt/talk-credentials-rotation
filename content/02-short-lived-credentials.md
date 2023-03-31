# Short-Lived Credentials

---

## `ServiceAccount`

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: robot
```

Pods contacting the API server authenticate as part of a `ServiceAccount`:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  serviceAccountName: robot
```

vvv

## Projected Token Mount

Since Kubernetes `v1.22`, this results in

```yaml[|9-26|5-8]
spec:
  containers:
  - name: nginx
    image: nginx
    volumeMounts:
    - name: kube-api-access-5f65c
      mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      readOnly: true
  volumes:
  - name: kube-api-access-5f65c
    projected:
      sources:
      - serviceAccountToken:
          expirationSeconds: 3607
          path: token
      - configMap:
          items:
          - key: ca.crt
            path: ca.crt
          name: kube-root-ca.crt
      - downwardAPI:
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
            path: namespace
```

Notes:
- kubelet automatically requests and mounts JWT token into pod
- client must regularly reload the token file (supported since long time in all major client libraries like `client-go`)

vvv

## `TokenRequest` API

Notes:
- ```shell
  kubectl create sa robot
  kubectl create token robot -v=9
  ```
- Explain `TokenRequest` body and `serviceaccounts/token` subresource
- Copy token into kubeconfig and use it
- Mention `.spec.boundObjectRef` (invalid after pod/SA deletion) -> delete SA and show it's not working anymore
- Mention that this API can also be called from others (not only kubelet)
- We use this in Gardener as well

vvv

## Projected Token Mount

Since Kubernetes `v1.22`, this results in

```yaml[|14]
spec:
  containers:
  - name: nginx
    image: nginx
    volumeMounts:
    - name: kube-api-access-5f65c
      mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      readOnly: true
  volumes:
  - name: kube-api-access-5f65c
    projected:
      sources:
      - serviceAccountToken:
          expirationSeconds: 3607
          path: token
      - configMap:
          items:
          - key: ca.crt
            path: ca.crt
          name: kube-root-ca.crt
      - downwardAPI:
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
            path: namespace
```

Notes:
- `3607s`, aka `1h7s` looks good (not too long), right?

vvv

## "Magic" Expiration Time

kubelet silently overwrites the expiration seconds 👻


[Source](https://github.com/kubernetes/kubernetes/blob/475f9010f5faa7bdd439944a6f5f1ec206297602/pkg/registry/core/serviceaccount/storage/token.go#L177-L180)

```go[|8-11|10-11,2|3,13]
const (
	WarnOnlyBoundTokenExpirationSeconds = 60*60+7 // 3607
	ExpirationExtensionSeconds          = 24*365*60*60 // 1y
)

exp := req.Spec.ExpirationSeconds

if pod != nil &&
  r.isKubeAudiences(req.Spec.Audiences) &&
  r.extendExpiration &&
  req.Spec.ExpirationSeconds == WarnOnlyBoundTokenExpirationSeconds {

  exp = ExpirationExtensionSeconds
}
```

Notes:
- This flag is true by default
- Token lifetime is extended to `1y` even though spec says `1h7s`
- Why? Clients must reload the token - prevent unexpected failure in production clusters

vvv

## Conclusion

- Set `--service-account-extend-token-expiration=false` to ensure tokens are indeed only valid for `1h`

- If you cannot control the flag (e.g., in managed clusters), overwrite the `expirationSeconds` to ensure short validity

---

## Static Token Secrets

Before Kubernetes `v1.24`, a static token `Secret` was automatically generated for `ServiceAccount`s:

```yaml[|11]
apiVersion: v1
kind: Secret
metadata:
  name: robot
  annotations:
    kubernetes.io/service-account.name: robot
    kubernetes.io/service-account.uid: da68f9c6-9d26-11e7-b84e-002dc52800da
data:
  ca.crt: <cluster-ca>
  namespace: default
  token: <some-static-token>
```

Such tokens have no expiration date! 😱

Notes:
- You can even now still create them manually if needed

vvv

## [KEP-2799](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token): Reduction of Secret-based Service Account Tokens

| Feature | Alpha | Beta | GA |
| -------- | ---- | ---- | -- |
| No auto-generation ✅ | - | 1.24 | 1.26 |
| Tracking metrics ✅ | 1.26 | 1.27 | 1.28 |
| Auto-cleanup ⚠️ | 1.27 | 1.28 | 1.29 |

Most probably static tokens still exist in your clusters! 👹

Notes:
- No auto-cleanup as of today

vvv

## Token Invalidation

![Demo](../assets/demo-time.gif)
<!-- .element: class="r-stretch" -->

Notes:
- Prepare: `kubectl create sa invalidation && kubectl patch sa invalidation --type=merge --patch='{"automountServiceAccountToken":false}'`
- ```shell
  kubectl apply -f - <<EOF
  apiVersion: v1
  kind: Secret
  metadata:
    name: invalidation
    annotations:
      kubernetes.io/service-account.name: invalidation
  type: kubernetes.io/service-account-token
  EOF
  ```
- Show invalidated token
- Gardener runs business-critical workload
- Brown-field applications cannot update too frequently (agreed MTW with customers), or bugs are blocking updates
- Security standards must also be applied for lower Kubernetes versions

vvv

## Conclusion

- If you are on `v1.24` or higher, manually delete still remaining static token secrets

- If you are stuck below `v1.24`, consider invalidating the tokens
