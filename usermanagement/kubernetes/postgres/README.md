#### Install 
```bash
kubectl apply -f pv.yml
```

```bash
helm install postgres-usermanagement oci://registry-1.docker.io/bitnamicharts/postgresql -n usermangement --values values.yml
```

Notes:
https://github.com/bitnami/charts/tree/main/bitnami/postgresql/#installing-the-chart