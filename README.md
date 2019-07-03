## kubectl-select
This is a simple plugin to add the ability to interactively select kubernetes resources
and print that resource to `stdout`.

### Examples

```sh
# This allows interactive selection of a pod, and then creates a terminal session
kubectl exec `kubectl select pods` -it sh
```

### Building
```
go get github.com/brendandburns/kubectl-select
```

### Installing
```
sudo cp $GOPATH/bin/kubectl-select /usr/local/bin/kubectl-select
```
