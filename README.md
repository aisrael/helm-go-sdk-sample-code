

## Starting Minikube

On Mac OS Big Sur, you may encounter problems using VirtualBox, so use hyperkit instead:

```
minikube start --driver=hyperkit --memory=4g --kubernetes-version=1.15.9
```
