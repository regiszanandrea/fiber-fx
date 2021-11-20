k3d cluster create dev --port 9443:443@loadbalancer   --port 80:80@loadbalancer  --api-port 6443 --k3s-arg "--disable=traefik@server:0"

istioctl install --set profile=demo -y

kubectl label namespace default istio-injection=enabled

# ArgoCD
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

kubectl apply -f k8s/argocd/fiber-fx-app.yaml

# kubectl apply -f k8s/service.yaml
# kubectl apply -f k8s/deployment.yaml
# kubectl apply -f k8s/istio
kubectl apply -f k8s/istio/ingress.yaml

# to get external ip
kubectl get svc istio-ingressgateway -n istio-system

