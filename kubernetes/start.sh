minikube start --cpus 4
minikube addons enable ingress
echo "Wait for the “Ingress” startup to complete."
sleep 10
kubectl get pods -n ingress-nginx
sleep 3
echo "Create infrastructure"
kubectl apply -f infrastructure/nginx.yaml
kubectl apply -f infrastructure/zookeeper-deployment.yaml
kubectl apply -f infrastructure/zookeeper-service.yaml
kubectl apply -f infrastructure/kafka-config.yaml
kubectl apply -f infrastructure/kafka-deployment.yaml
kubectl apply -f infrastructure/kafka-service.yaml
sleep 1
kubectl apply -f infrastructure/kafka-topic.yaml

echo "Create Checkout"
sleep 3
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/frontend/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/frontend/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/frontend/service.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/frontend/ingress.yaml

kubectl apply -f travelplatform-case-m2/kubernetes/checkout/db/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/db/secrets.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/db/pvc.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/db/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/db/service.yaml

kubectl apply -f travelplatform-case-m2/kubernetes/checkout/app/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/app/secrets.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/app/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/app/service.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/checkout/app/ingress.yaml

echo "Create Travelmanagement"
sleep 3
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/frontend/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/frontend/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/frontend/service.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/frontend/ingress.yaml

kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/db/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/db/secrets.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/db/pvc.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/db/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/db/service.yaml

kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/app/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/app/secrets.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/app/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/app/service.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/travelmanagement/app/ingress.yaml

echo "Create usermanagement"
sleep 3
# TODO

echo "Create Monitoring"
sleep 3
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/frontend/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/frontend/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/frontend/service.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/frontend/ingress.yaml

kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/db/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/db/secrets.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/db/pvc.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/db/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/db/service.yaml

kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/app/config.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/app/secrets.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/app/deployment.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/app/service.yaml
kubectl apply -f travelplatform-case-m2/kubernetes/monitoring/app/ingress.yaml

echo "To Access the cluster execute minikube tunnel in a separate terminal"
echo "The DNS name is current mini.local (Todo change name)"
echo "Add an Host in your /etc/hosts"
echo "eg.: 192.168.49.2 mini.local"
echo "the travelmanagement is now reachable at mini.local/travma"