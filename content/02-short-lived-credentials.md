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

```yaml[|9-26|5-8|1-26|14]
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
- `3607s`, aka `1h7s` looks good (not too long), right?

vvv

## `TokenRequest` API

Notes:
- Short demo creating a token
- Mention that this API can also be called from others (not only kubelet)
- We use this in Gardener as well

vvv

## "Magic" Expiration Time

kubelet silently overwrites the expiration seconds 🥷🏻


[Source](https://github.com/kubernetes/kubernetes/blob/475f9010f5faa7bdd439944a6f5f1ec206297602/pkg/registry/core/serviceaccount/storage/token.go#L177-L180)

```go[|9|3,10]
const (
	WarnOnlyBoundTokenExpirationSeconds = 60 * 60 + 7
	ExpirationExtensionSeconds          = 24 * 365 * 60 * 60
)

// ...

exp := req.Spec.ExpirationSeconds
if r.extendExpiration && pod != nil && req.Spec.ExpirationSeconds == WarnOnlyBoundTokenExpirationSeconds && r.isKubeAudiences(req.Spec.Audiences) {
	exp = ExpirationExtensionSeconds
}
```

Notes:
- This flag is true by default
- Token lifetime is extended to `1y` even though spec says `1h7s`
- Why? Clients must reload the token - prevent unexpected failure in production clusters

vvv

## Conclusion

- Set `--service-account-extend-token-expiration=false` to ensure tokens are indeed only valid for ~`1h`

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

- Such tokens have no expiration date! 😱
- Probably they still exist in your 1.24+ clusters! 👹

Notes:
- You can even now still create them manually if needed
- No auto-cleanup as of today

vvv

## [KEP-2799](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token): Reduction of Secret-based Service Account Tokens

![1](../assets/kep2799-1.png)

vvv

## [KEP-2799](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token): Reduction of Secret-based Service Account Tokens

![1](../assets/kep2799-2.png)

vvv

## [KEP-2799](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token): Reduction of Secret-based Service Account Tokens

![1](../assets/kep2799-3.png)

vvv

## Token Invalidation

![Demo](../assets/demo-time.gif)
<!-- .element: class="r-stretch" -->

vvv

## Conclusion

- If you are on `v1.24` or higher, manually delete still remaining static token secrets

- If you are stuck below `v1.24`, consider invalidating the tokens

Notes:
- Gardener runs business-critical workload
- Brown-field applications cannot update too frequently (agreed MTW with customers), or bugs are blocking updates
- Security standards must also be applied for lower Kubernetes versions
