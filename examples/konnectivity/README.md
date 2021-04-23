
A helm chart for `apiserver-network-proxy`, make it easy deploy and test.

## User Guide


### build image

```
export REGISTRY=gcr.io/apiserver-network-proxy
make docker-build
```

### download binaries

```
./scripts/download-binaries.sh
```

### create kind cluster

```shell
export PATH=$(pwd)/bin:${PATH}

kind create cluster
#export TAG="v0.0.1"
make deploy-kind
```

### uninstall

```shell
make delete-kind
```



curl -v -p --proxy-key /opt/pki/proxyclient/tls.key --proxy-cert /opt/pki/proxyclient/tls.crt --proxy-cacert /opt/pki/proxyca/tls.crt --proxy-cert-type PEM -x https://127.0.0.1:8090  http://kubia:80





curl  -v -X CONNECT http://konnectivity-proxyserver:8090/




git pull
export TAG="v0.0.1"
make docker-build/proxy-test-client

kind load docker-image "gcr.io/apiserver-network-proxy/proxy-test-client-amd64:v0.0.1" --name="${KIND_CLUSTER_NAME:-kind}"




NS=default
APP=konnectivity-test-client

podName=`kubectl get pods -n $NS -l app=${APP} -o jsonpath='{.items[*].metadata.name}'`

kubectl delete pod -n $NS $podName




curl http://konnectivity-test-client/ok
curl http://konnectivity-test-client/k8s/clusters/xx/


curl http://konnectivity-test-client/k8s/clusters/xx/api/v1/namespaces/default/pods

curl http://konnectivity-test-client/k8s/clusters/xx/api/v1/namespaces/default/pods?auth=false&cacert=false



kind load docker-image "nicolaka/netshoot:latest" --name="${KIND_CLUSTER_NAME:-kind}"

cat<<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: test-press
spec:
  containers:
  - name: bb
    image: nicolaka/netshoot
    command: ["/bin/sh"]
    args: ["-c", "while true; do echo hello; sleep 10;done"]
    resources:
      requests:
        cpu: 200m
        memory: 200Mi
      limits:
       cpu: 200m
       memory: 200Mi
EOF



ab -n 10000 -c 50 http://konnectivity-test-client/k8s/clusters/xx/api/v1/namespaces/default/pods
