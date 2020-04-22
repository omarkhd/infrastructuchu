This small provisioning project deploys a test cluster of Kubernetes.

It also configures deployment and an admin user for the web ui dashboard.
https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/#command-line-proxy

As the dashboard can only be accesed from the machine where the following command is
executed, it should be run from a `kubectl` installation on the host system.

    kubectl proxy

The `config` file needed to configure `kubectl` is placed in the `__bucket__` folder. The bearer token
needed to access the we ui dashboard is also in the `__bucket__` folder. The dashboard will be available at
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

Note: Remember to delete all *.txt file in __bucket__ directory before recreating the cluster.
