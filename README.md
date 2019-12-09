# Introduction

An example on using Go in OpenShift.

# Running locally

Build and run:

```
cd openshift-go-sample
go build
./openshift-go-sample
```

# Importing the base s2i image

First you have to import the base image to your OpenShift installation from
the Red Hat Registry. If you have cluster-admin permissions, import the image
to the `openshift` project:

```
oc import-image ubi8/go-toolset --from=registry.redhat.io/ubi8/go-toolset -n openshift --confirm
```

If you don't have cluster-admin, you'll have to provide an username and
password to registry.redhat.io before importing the image. To this:

```
oc create secret docker-registry redhat-registry \
  --docker-server=registry.redhat.io \
  --docker-username='<username>' \
  --docker-password='<password>' \
  --docker-email='<email>'

oc oc secret link builder redhat-registry

oc import-image ubi8/go-toolset --from=registry.redhat.io/ubi8/go-toolset --confirm
```

# Deploying the app

With the base image imported, do this:

```
oc new-app go-toolset~https://github.com/eitchugo/openshift-go-sample --name=go-sample
oc expose dc/go-sample --port=8080
oc expose svc/go-sample
```

This will build the image using s2i, expose to the router and you can access
the URL.
