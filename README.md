# OSRM Fetcher
OSRM Fetcher fetchs the [OSRM Builder](https://github.com/mig-elgt/osrm-builder) files from Google Cloud Storage. These files will be store in your container local file system using the path /osrm-data 

# Installation

### Requirements

* Docker Engine
* Docker Image Registry
* Google Cloud Storage
* Google Service Account

Build a Docker Image using this code and publish it in your favorite Docker Image Registry. You will use this Docker Image as the Base Image in [OSRM Kubernetes](https://github.com/mig-elgt/osrm-kubernetes) Dockerfile in order to deploy the OSRM Server for your Kubernetes Cluster.

```
// Replace gcr.io for your Docker Image Registry

$ docker build -t gcr.io/osrm/osrm-fetcher:v1 .
$ docker push gcr.io/osrm/osrm-fetcher:v1

```
