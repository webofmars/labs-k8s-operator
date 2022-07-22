# kubernetes operator

## credits

based on <https://betterprogramming.pub/build-a-kubernetes-operator-in-10-minutes-11eec1492d30>

_Note: this repo is part of webofmars labs project. You can check other labs here <https://github.com/webofmars/labs>.

## requirements

* go version v1.17.9+
* docker version 17.03+
* kubectl version v1.11.3+
* access to a Kubernetes v1.11.3+ cluster (I highly suggest using kind to set up your own local cluster, it’s very easy to use!).

## setup

```console
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH) && \
  chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

or

```console
asdf plugin add kubebuilder
asdf install
```

## init our repo (step-1)

```console
kubebuilder init --domain webofmars.com --repo webofmars.com/pizza
```

## create CRDs (step 2)

```console
kubebuilder create api --group pizza --version v1 --kind Pizza
go mod tidy
make generate
```

## customize CRDs (step 3)

* edit the file `api/v1/pizza_types.go`

## customize the controller (step 4)

* edit the file `controllers/pizza_controller.go` and inseert your own logic for a pizza

## generate the manifests (step 5)

```console
make manifests
```

## install the CRDs (step 6)

NB: you need a working kubectl connected to your dev cluster in order to test it.

```console
make install
kubectl get crds
```

## run the controller (step 7)

Run the controller locally (it can later be packaged through helm or other means and run in a the kubernetes cluster).

```console
make run
```

## create your first pizza (step 8)

* edit file `config/samples/pizza_v1_pizza.yaml` to match a valid pizza spec as defined in `pizza_types.go`

```console
kubectl apply -f config/samples/pizza_v1_pizza.yaml
```

after a while you should see the controller reconcile loop pass over the pizza object and modify it if you instructed to do so in the controller (check step 4 code)

## package the controller (step 9)

```console
make docker-build
make docker-push
```

## going further

You can now add some items to spec and status to reflect a valid pizza from your logic and play with thee controller