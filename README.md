# Vertical Pod Autoscaler

A sample project showcasing the implementation of Vertical Pod Autoscaler.

## Blog Post

> Coming soon!

## Installing

### Autoscaling without Pod Disruption Budget

```shell
kind create cluster --image kindest/node:v1.23.4 --config=kind/cluster.yaml

helm repo add jetstack https://charts.jetstack.io
helm install jetstack/cert-manager --name-template cert-manager --create-namespace -n cert-manager --values kind/cert-manager-values.yaml --version 1.9.1 --wait

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx/ingress-nginx --name-template ingress-nginx --create-namespace -n ingress-nginx --values kind/ingress-nginx-values.yaml --version 4.0.19 --wait

helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server
helm install metrics-server/metrics-server --name-template metrics-server --create-namespace -n metrics-server --values kind/metrics-server-values.yaml --version 3.8.2 --wait

helm install helm-chart --name-template vertical-pod-autoscaler --create-namespace -n vertical-pod-autoscaler --wait

helm install sample-app/helm-chart --name-template sample-app --create-namespace -n sample-app --wait
```

### Autoscaling with Pod Disruption Budget

```shell
kind create cluster --image kindest/node:v1.23.4 --config=kind/cluster.yaml

helm repo add jetstack https://charts.jetstack.io
helm install jetstack/cert-manager --name-template cert-manager --create-namespace -n cert-manager --values kind/cert-manager-values.yaml --version 1.9.1 --wait

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx/ingress-nginx --name-template ingress-nginx --create-namespace -n ingress-nginx --values kind/ingress-nginx-values.yaml --version 4.0.19 --wait

helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server
helm install metrics-server/metrics-server --name-template metrics-server --create-namespace -n metrics-server --values kind/metrics-server-values.yaml --version 3.8.2 --wait

helm install helm-chart --name-template vertical-pod-autoscaler --create-namespace -n vertical-pod-autoscaler --wait

helm install sample-app/helm-chart --name-template sample-app --set podDisruptionBudget.enabled=true --create-namespace -n sample-app --wait
```

## Smoke Testing

```shell
curl localhost/success -H "Host: sample.app" -v
curl localhost/error -H "Host: sample.app" -v
```

## Load Testing

```shell
k6 run k6/script.js
```

## Uninstalling

```shell
kind delete cluster
```
