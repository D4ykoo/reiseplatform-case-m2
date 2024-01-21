# How to run
The kubernetes manifests are in a separate folder for each microservice. 
The infrastructure folder contains service-wide applications and settings (e.g. Kafka, nginx-ingress-controller).
A script is provided that creates and configures all necessary components of the application. 
Each Kubernetes manifest can also be started manually (kubectl apply -f 'path to manifest'). 

## Run per script
The script has currently only been tested with Minikube on Linux hosts.
Make sure your system provide a persistent volume “standard”.

Run with the start script:
```bash
chmod +x start.sh && ./start.sh
```

When using Minikube:
```bash
./start.sh -m
```
It will automatically create a new or start an existing k8s cluster as well as activating the ingress addon.

## Run manually
All Kubernetes Manifest can be applied manually. The persistent volume “standard” is used. 
To use the the ingress-rules apply the "infrastructure/nginx.yaml" at first.
The following section describes how to start Minikube with ingress.

### Minikube and Ingress (Docker-Driver)
Here is a brief introduction to configuring Minikube for our application. 

1. Start Minikube with `minikube start`.

2. Minikube is not accessible via your network by default. You must create a tunnel with the command `minikube tunnel` in a separate terminal. This tunnel must remain open during your session.

3. Check the connection (e.g. ping). You will get the minikube IP with `minikube ip`.

4. Enbale ingress in minikube with `minikube addons enable ingress`.

5. Configure nginx-ingress-controller. This enables the use of rewriting rules. Apply the ConfigMap with`kubectl apply -f infrastructure/nginx.yaml`.

### Troubleshooting
1. Your Ingress manifest cannot be used in the cluster forbidden by the administrator. Restart the Ingress Controller with `kubectl rollout restart deployment -n ingress-nginx ingress-nginx-controller`.

## Access the minikube network
Start in another terminal `minikube tunnel`. It shows the status and a route. Paste the ip after the `->` inside `/etc/hosts` e.g. 
```bash
192.168.39.8    mini.local
```
The DNS record is required for the web applications to work successfully with the backend. It is also possible to use a different subdomain, but don't forget to configure the configmaps.(e.g. [monitoring-configmap](monitoring/frontend/config.yaml))

