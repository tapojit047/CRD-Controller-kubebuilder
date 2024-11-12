# CRD-Controller-Kubebuilder: Alchemist Custom Resource Controller

This guide provides an overview of the `CRD-Controller-Kubebuilder`, a controller that manages a custom Kubernetes resource called `Alchemist`. The controller is built with Kubebuilder. It creates and manages deployments and services based on the `Alchemist` resource specifications.

## Custom Resource Definition: Alchemist
The Alchemist resource allows users to specify deployment and service details through its spec fields:

- `deploymentName`: Name of the deployment to be created.
- `replicas`: Number of replicas for the deployment.
- `image`: Container image to use in the deployment.
- `containerPort` (optional): Port on the container.
- `servicePort` (optional): Service port for external access.
- `targetPort` (optional): Target port within the container.

The `Alchemist` has the following fields:
- `deploymentName`
- `replicas`
- `image`
- `containerPort` (optional)
- `servicePort` (optional)
- `targetPort` (optional)

Complete definition details are available in the [Alchemist-API](https://github.com/tapojit047/CRD-Controller-kubebuilder/blob/master/api/v1/alchemist_types.go#L27)

**Example `Alchemist` Manifest**:
```yaml
apiVersion: fullmetal.com.my.domain/v1
kind: Alchemist
metadata:
  name: elric
  namespace: demo
spec:
  deploymentName: alchemist
  replicas: 3
  image: tapojit047/api-server
  containerPort: 8000
  servicePort: 8000
  targetPort: 8000
```

When deployed, this manifest will create a deployment named `alchemist` with 3 replicas, using the image `tapojit047/api-server`, and a service exposing it on port 8000.

## Controller:

### Prerequisites:
- Make installed on your local machine.
- You need to have a Kubernetes cluster, and the `kubectl` command-line tool must be configured to communicate with your cluster. If you do not already have a cluster, you can create one by using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/).

### Install Custom Resource Definitions
```bash
make install
```

### Run the Controller

There are two ways to deploy the controller:

#### Option 1: Use a Pre-built Image
Deploy the controller using the pre-built image `tapojit047/crd-controller`:

```bash
make deploy IMG=tapojit047:crd-controller
```

#### Option 2: Build and Deploy Locally
To build the controller locally, you’ll need `Go` installed. Follow these steps to set up Go and the controller:

- **Install Go:**
```bash
$ sudo rm -rf /usr/local/go
$ go_version=1.23.3
$ cd ~/Downloads
$ sudo apt-get update
$ sudo apt-get install -y build-essential git curl wget
$ wget https://go.dev/dl/go${go_version}.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go${go_version}.linux-amd64.tar.gz
$ sudo chown -R $(id -u):$(id -g) /usr/local/go
$ rm go${go_version}.linux-amd64.tar.gz
```

- **Add go to your `$PATH` variable:**
```bash
$ mkdir $HOME/go
$ nano ~/.bashrc
$ export GOPATH=$HOME/go
$ export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
$ source ~/.bashrc
$ go version
```
> Note: Make sure that this project is inside the `GOPATH`, which is `$HOME/go`

- Build and Push Docker Image:
```bash
$ make docker-build docker-push IMG=<some-registry>/<project-name>:tag
```

- Deploy the controller to the cluster with image specified by IMG:
```bash
make deploy IMG=<some-registry>/<project-name>:tag
```

- Verify the Controller Deployment:
```bash
$ kubectl get pods -n crd-controller-kubebuilder-system
NAME                                                             READY   STATUS    RESTARTS   AGE
crd-controller-kubebuilder-controller-manager-5b594f97bd-6gfpb   2/2     Running   0          170m
```
You should see the `controller-manager` pod in the `Running` state.

### Deploying an `Alchemist` Resource
Once the controller is deployed, create a sample Alchemist resource to test it:

```bash
$ kubectl create ns demo
namespace/demo created

$ kubectl apply -f config/samples/fullmetal.com_v1_alchemist.yaml
alchemist.fullmetal.com.my.domain/elric created
```

The controller will create a deployment, pods, and a service based on the `Alchemist` specification. Use the following command to monitor the resources:

```bash
$ watch kubectl get alchemist,deploy,pod,svc -n demo
Every 2.0s: kubectl get alchemist,deploy,pod,svc -n demo                                                                                                                  

NAME                                      AGE
alchemist.fullmetal.com.my.domain/elric   49s

NAME                        READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/alchemist   3/3     3            3           49s

NAME                            READY   STATUS    RESTARTS         AGE
pod/alchemist-f66994ffb-2qrw8   1/1     Running   0                49s
pod/alchemist-f66994ffb-8qt92   1/1     Running   0                49s
pod/alchemist-f66994ffb-z5tq9   1/1     Running   0                49s

NAME                           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
service/alchemist-service      NodePort    10.96.236.228   <none>        8000:30047/TCP               49s
```

This output shows that the `Alchemist` deployment, pods, and a NodePort service have been successfully created.

## Clean Up
```bash
$ make undeploy
namespace "crd-controller-kubebuilder-system" deleted
customresourcedefinition.apiextensions.k8s.io "alchemists.fullmetal.com.my.domain" deleted
serviceaccount "crd-controller-kubebuilder-controller-manager" deleted
role.rbac.authorization.k8s.io "crd-controller-kubebuilder-leader-election-role" deleted
clusterrole.rbac.authorization.k8s.io "crd-controller-kubebuilder-manager-role" deleted
clusterrole.rbac.authorization.k8s.io "crd-controller-kubebuilder-metrics-reader" deleted
clusterrole.rbac.authorization.k8s.io "crd-controller-kubebuilder-proxy-role" deleted
rolebinding.rbac.authorization.k8s.io "crd-controller-kubebuilder-leader-election-rolebinding" deleted
clusterrolebinding.rbac.authorization.k8s.io "crd-controller-kubebuilder-manager-rolebinding" deleted
clusterrolebinding.rbac.authorization.k8s.io "crd-controller-kubebuilder-proxy-rolebinding" deleted
service "crd-controller-kubebuilder-controller-manager-metrics-service" deleted
deployment.apps "crd-controller-kubebuilder-controller-manager" deleted

$ kubectl delete ns demo
namespace "demo" deleted
```

## References
* #### [Kubebuilder Quickstare](https://book.kubebuilder.io/quick-start)
* #### [Kubebuilder Documentation][1]
* #### [Kubebuilder GitHub][2]

[1]:https://book.kubebuilder.io/ "Kubebuilder Book"
[2]:https://github.com/kubernetes-sigs/kubebuilder "Kubebuilder"