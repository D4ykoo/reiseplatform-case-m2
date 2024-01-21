while getopts m flag
do
    case "${flag}" in
        m) minikube=true;;
    esac
done

if [ "$minikube" = true ] ; then
    minikube start --cpus 4
    minikube addons enable ingress
    echo "Wait for the “Ingress” startup to complete."
    sleep 10
    kubectl get pods -n ingress-nginx
fi

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
kubectl apply -f checkout/frontend/config.yaml
kubectl apply -f checkout/frontend/nginx.conf.yml
kubectl apply -f checkout/frontend/deployment.yaml
kubectl apply -f checkout/frontend/service.yaml
kubectl apply -f checkout/frontend/ingress.yaml

kubectl apply -f checkout/backend/db/config.yaml
kubectl apply -f checkout/backend/db/secrets.yaml
kubectl apply -f checkout/backend/db/pvc.yaml
kubectl apply -f checkout/backend/db/deployment.yaml
kubectl apply -f checkout/backend/db/service.yaml

kubectl apply -f checkout/backend/app/config.yaml
kubectl apply -f checkout/backend/app/secrets.yaml
kubectl apply -f checkout/backend/app/deployment.yaml
kubectl apply -f checkout/backend/app/service.yaml
kubectl apply -f checkout/backend/app/ingress.yaml

echo "Create Travelmanagement"
sleep 3
kubectl apply -f travelmanagement/frontend/config.yaml
kubectl apply -f travelmanagement/frontend/deployment.yaml
kubectl apply -f travelmanagement/frontend/service.yaml
kubectl apply -f travelmanagement/frontend/ingress.yaml

kubectl apply -f travelmanagement/backend/db/config.yaml
kubectl apply -f travelmanagement/backend/db/secrets.yaml
kubectl apply -f travelmanagement/backend/db/pvc.yaml
kubectl apply -f travelmanagement/backend/db/deployment.yaml
kubectl apply -f travelmanagement/backend/db/service.yaml

kubectl apply -f travelmanagement/backend/app/config.yaml
kubectl apply -f travelmanagement/backend/app/secrets.yaml
kubectl apply -f travelmanagement/backend/app/deployment.yaml
kubectl apply -f travelmanagement/backend/app/service.yaml
kubectl apply -f travelmanagement/backend/app/ingress.yaml

echo "Create usermanagement"
sleep 3
                 
kubectl apply -f usermanagement/frontend/config.yml
kubectl apply -f usermanagement/frontend/nginx.conf.yml
kubectl apply -f usermanagement/frontend/deployment.yml
kubectl apply -f usermanagement/frontend/service.yml
kubectl apply -f usermanagement/frontend/ingress.yml

kubectl apply -f usermanagement/backend/db/config.yaml
kubectl apply -f usermanagement/backend/db/secrets.yaml
kubectl apply -f usermanagement/backend/db/pvc.yaml
kubectl apply -f usermanagement/backend/db/deployment.yaml
kubectl apply -f usermanagement/backend/db/service.yaml

kubectl apply -f usermanagement/backend/app/config.yaml
kubectl apply -f usermanagement/backend/app/secrets.yaml
kubectl apply -f usermanagement/backend/app/deployment.yaml
kubectl apply -f usermanagement/backend/app/service.yaml
kubectl apply -f usermanagement/backend/app/ingress.yaml

echo "Create Monitoring"
sleep 3
kubectl apply -f monitoring/frontend/config.yaml
kubectl apply -f monitoring/frontend/deployment.yml
kubectl apply -f monitoring/frontend/service.yml
kubectl apply -f monitoring/frontend/ingress.yml

kubectl apply -f monitoring/backend/db/config.yaml
kubectl apply -f monitoring/backend/db/secrets.yaml
kubectl apply -f monitoring/backend/db/pvc.yaml
kubectl apply -f monitoring/backend/db/deployment.yaml
kubectl apply -f monitoring/backend/db/service.yaml

kubectl apply -f monitoring/backend/app/config.yaml
kubectl apply -f monitoring/backend/app/secrets.yaml
kubectl apply -f monitoring/backend/app/deployment.yaml
kubectl apply -f monitoring/backend/app/service.yaml
kubectl apply -f monitoring/backend/app/ingress.yaml

if [ "$minikube" = true ] ; then
    echo "To Access the cluster execute minikube tunnel in a separate terminal"
    echo "The DNS name is current mini.local"
    echo "Add an Host in your /etc/hosts"
    echo "eg.: 192.168.49.2 mini.local"
    echo "the travelmanagement is now reachable at mini.local/travma"
else
    echo "the travelmanagement is now reachable at sub.domain/travma"
    echo "Make sure your domain is configured in the web applications configmaps"
fi
