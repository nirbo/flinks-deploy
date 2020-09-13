#!/bin/bash

# Deploy the CRUD REST API
kubectl apply -f crud-rest-api-deployment.yaml -f crud-rest-api-ingress.yaml -f crud-rest-api-service.yaml

