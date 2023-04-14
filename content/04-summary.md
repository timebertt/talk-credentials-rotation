# Key Takeaways ✨

vvv

## Key Takeaways

- use short-lived projected ServiceAccount tokens
- delete remaining static ServiceAccount tokens

notes:

- short-lived wherever possible
- or invalidate them

vvv

## Key Takeaways

- rotate static credentials in two steps
- use bundles in between
- implemented for Kubernetes as workload, concept is applicable to most workloads

notes:

- where short-lived credentials not possible, rotate static one in two steps
- bundles -> disruption-free
- applicable to other workloads
