# k8s playground

k8s documentation [here](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/#pod-deletion-cost).

Before running you need to give permissions to service account:

```
kubectl apply -f resources/serviceaccount.yaml
```

To run inside a pod:

```
kubectl run k8s-playground --image=ghcr.io/nikimoro-qlik/k8s_playground:dev
```

To run inside a deployment:

```
kubectl create deployment k8s-playground --image=ghcr.io/nikimoro-qlik/k8s_playground:dev

...or...

kubectl apply -f resources/deployment.yaml 
```

How to check logs:

```
kubectl logs -f -l app=k8s-playground
```

How to check deletion cost:

```
watch -d -t 'kubectl get pod -o custom-columns=NAME:.metadata.name,STATUS:.status.phase,"DELETION-COST":".metadata.annotations.controller\.kubernetes\.io/pod-deletion-cost"'
```