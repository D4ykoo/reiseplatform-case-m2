# How to apply kubernetes manifests
There are two different manifests:
1. allinone.yml -> The whole checkout stack is **one** deployment
1. deployments.yml -> **Every** checkout service as well as the db is a single deployment

Apply the manifest:
```bash
kubectl apply -f <allinone/deployments>.yml
```