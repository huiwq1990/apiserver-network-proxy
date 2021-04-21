
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

make deploy-kind
```

### uninstall

```shell
make delete-kind
```



curl -v -p --proxy-key /opt/pki/proxyclient/tls.key --proxy-cert /opt/pki/proxyclient/tls.crt --proxy-cacert /opt/pki/proxyca/tls.crt --proxy-cert-type PEM -x https://127.0.0.1:8090  http://kubia:80




