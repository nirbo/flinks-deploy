#!/bin/bash

# Deploy the CRUD REST API
kubectl apply -f crud-rest-api-daemonset.yaml -f crud-rest-api-ingress.yaml -f crud-rest-api-service.yaml

