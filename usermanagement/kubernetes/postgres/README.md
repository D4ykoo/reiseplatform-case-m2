#### Install 
```bash
kubectl apply -f pv.yml
```


```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```

```bash
helm install postgres-usermanagement bitnami/postgresql -n usermanagement --values values.yml
```

Notes:
https://github.com/bitnami/charts/tree/main/bitnami/postgresql/#installing-the-chart