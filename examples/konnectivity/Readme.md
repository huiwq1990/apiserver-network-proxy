



NS=default
APP=konnectivity-agent
podName=`kubectl get pods -n $NS -l app=${APP} -o jsonpath='{.items[*].metadata.name}'`
kubectl exec  ${podName} -n ${NS} -- tar cf - /agent.pcap | tar xf - -C ~/Desktop/



APP=konnectivity-server
podName=`kubectl get pods -n $NS -l app=${APP} -o jsonpath='{.items[*].metadata.name}'`
kubectl exec  ${podName} -n ${NS} -- tar cf - /server.pcap | tar xf - -C ~/Desktop/



```shell

```


curl -v -p --proxy-key /opt/pki/proxyserver/tls.key --proxy-cert /opt/pki/proxyserver/tls.crt --proxy-cacert /opt/pki/ca/tls.crt --proxy-cert-type PEM -x https://127.0.0.1:8090  http://localhost:8000```

