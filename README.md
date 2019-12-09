# Introduction

An example on using Go in OpenShift.

# Running locally

Build and run:

```
cd openshift-go-sample
go build
./openshift-go-sample
```

# Running on OpenShift

Just do this:

```
oc new-app https://github.com/eitchugo/openshift-go-sample --name=go-sample
oc expose dc/go-sample --port=8080
oc expose svc/go-sample
```

This will build the image using s2i, expose to the router and you can access
the URL.
