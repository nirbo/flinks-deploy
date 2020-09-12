#!/bin/bash

# Deploy nginx ingress controller (values.yaml kind if set to DaemonSet)
helm install ingress-nginx ./ingress-nginx/ --namespace ingress-controller

# Deploy the environment
kubectl apply -f nginx-web-server-daemonset.yaml -f nginx-web-server-ingress.yaml -f nginx-web-server-service.yaml

