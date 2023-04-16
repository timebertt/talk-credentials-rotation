# Questions? 🙋

Slides Available Online:

<img class="slides-qr-code"></img>
<!-- .element: class="r-stretch" -->

<a class="slides-qr-code"></a>

notes:

- distinguish from cert manager functionality:
  - we manage arbitrary secrets, not only certificates
  - etcd encryption key rotation requires more orchestration, e.g., re-encryption of secrets in cluster
  - cert-manager only renews in one step
    - two-phase rotation is not possible
    - rotation always causes disruption
  - our solution is not generic, specifically tailored to operator case

---

## Thanks 🙌

<table>
<tr>
<td>

Follow us on Twitter:
- [@rafaelfranzke](https://twitter.com/rafaelfranzke)
- [@timebertt](https://twitter.com/timebertt)

</td>
<td>

Check out project Gardener:
- [gardener.cloud](https://gardener.cloud)
- [github.com/gardener/gardener](https://github.com/gardener/gardener)

</td>
</tr>
</table>

![Gardener Logo](../assets/gardener.svg)
