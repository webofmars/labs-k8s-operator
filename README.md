# kubernetes operator

## credits

based on <https://betterprogramming.pub/build-a-kubernetes-operator-in-10-minutes-11eec1492d30>

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
