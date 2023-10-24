# k8s playground

Before running you need to give permissions to service account:

```
k apply -f resources/clusterrole.yaml
```

To run inside a pod:

```
k run k8s-playground --image=ghcr.io/nikimoro-qlik/k8s_playground:dev
```

To run inside a deployment:

```
k create deployment k8s-playground --image=ghcr.io/nikimoro-qlik/k8s_playground:dev

or

k apply -f resources/deployment.yaml 
```

How to check deletion cost:

```
k get pod -w -o custom-columns=NAME:.metadata.name,STATUS:.status.phase,"DELETION-COST":".metadata.annotations.controller\.kubernetes\.io/pod-deletion-cost"
```