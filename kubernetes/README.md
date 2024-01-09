# Minikube and Ingress (Docker-Driver)

Here is a brief introduction to configuring Minikube for our application.

1. Start Minikube with `minikube start`.

2. Minikube is not accessible via your network by default. You must create a tunnel with the command `minikube tunnel` in a separate terminal. This tunnel must remain open during your session.

3. Check the connection (e.g. ping). You will get the minikube IP with `minikube ip`.

4. Enbale ingress in minikube with `minikube addons enable ingress`.

5. Configure nginx-ingress-controller. This enables the use of rewriting rules. Apply the ConfigMap with`kubectl apply -f nginx.yaml`.


## Troubleshooting
1. Your Ingress manifest cannot be used in the cluster forbidden by the administrator. Restart the Ingress Controller with `kubectl rollout restart deployment -n ingress-nginx ingress-nginx-controller`.