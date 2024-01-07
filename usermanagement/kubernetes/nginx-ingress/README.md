#### 1. Add Repo
```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
```

#### 2. Install using helm chart
```bash
helm install ingress-nginx ingress-nginx/ingress-nginx --values values.yml
```