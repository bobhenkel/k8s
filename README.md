# k8s
K8s Puppet Bolt module to manage a local kind cluster -> https://github.com/kubernetes-sigs/kind. Kind is a program you can run locally to create a mini k8s lab cluster on your workstation or during CI.

This repo was created so I can learn how to create Puppet Bolt tasks that are implemented in go.  The value of this is the only dependency is the go compiled binary, no ruby or python runtime needed on the host your running the task on. So far I'm building the binary and keeping it in git with the module, that approach smells a tad, but for learning I'm ok with that, the binaries are small enough and it keeps things simple. A more robust solution would be to have the module download the needed binary from a proper location the first time the module is executed on a given host.

To experiment with this:

* 1. git clone git@github.com:bobhenkel/k8s.git
* 1.1 for every task you want to work on...
* 2. cd k8s/tasks/src/create_kind_cluster
* 3. env GOOS=linux GOARCH=amd64 go build -o ../../create_kind_cluster-linux-amd64
* 4. bolt task run k8s::create_kind_cluster_linux_amd64 cluster_name=kind-cluster2 image="kindest/node:v1.13.7" --nodes localhost --modulepath .
