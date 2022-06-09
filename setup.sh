#!/bin/bash

#Cluster setup
sysctl -w fs.inotify.max_queued_events=1048576
sysctl -w fs.inotify.max_user_watches=1048576
sysctl -w fs.inotify.max_user_instances=1048576
kind create cluster --config dir/cluster-config/kind-config.yaml
kubectl apply -f dir/cluster-config/metricserver.yaml

#Install Kyverno & the policies
helm install kyverno kyverno/kyverno --namespace kyverno --create-namespace 
helm install kyverno-policies kyverno/kyverno-policies --namespace kyverno

#Add extra policies
#bash setup.sh https://github.com/kyverno/policies/blob/main/best-practices/require_probes/require_probes.yaml
if [ ! -z $1 ]
then
    kubectl apply -f "$1"
fi