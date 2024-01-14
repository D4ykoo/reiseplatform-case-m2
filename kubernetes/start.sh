minikube start --cpus 4
minikube addons enable ingress
echo "Wait for the “Ingress” startup to complete."
sleep 10
kubectl get pods -n ingress-nginx
sleep 3
echo "Setup Infrastructure"
kubectl apply -f infrastructure/nginx.yaml
kubectl apply -f infrastructure/zookeeper-deployment.yaml
kubectl apply -f infrastructure/zookeeper-service.yaml
kubectl apply -f infrastructure/kafka-config.yaml
kubectl apply -f infrastructure/kafka-deployment.yaml
kubectl apply -f infrastructure/kafka-service.yaml
sleep 1
kubectl apply -f infrastructure/kafka-topic.yaml

echo "Make sure your persistence volume 'standard' exist"
sleep 3
echo "Setup Travelmanagement"
echo "Travelmanagement - Frontend"
kubectl apply -f travelmanagement/frontend
sleep 3
echo "Travelmanagement - DB"
kubectl apply -f travelmanagement/backend/db
sleep 3
echo "Travelmanagement - App"
kubectl apply -f travelmanagement/backend/app
sleep 3

echo "usermanagement todo"

echo "Setup Checkout"
echo "Checkout - Frontend"
kubectl apply -f checkout/frontend
sleep 3
echo "Checkout - DB"
kubectl apply -f checkout/backend/db
sleep 3
echo "Checkout - App"
kubectl apply -f checkout/backend/app
sleep 3

echo "Setup Monitoring"
echo "Monitoring - Frontend"
kubectl apply -f monitoring/frontend
sleep 3
echo "Monitoring - DB"
kubectl apply -f monitoring/backend/db
sleep 3
echo "Monitoring - App"
kubectl apply -f monitoring/backend/app
sleep 3

echo "To Access the cluster execute minikube tunnel in a separate terminal"
echo "The DNS name is current mini.local (Todo change name)"
echo "Add an Host in your /etc/hosts"
echo "eg.: 192.168.49.2 mini.local"
echo "the travelmanagement is now reachable at mini.local/travma"