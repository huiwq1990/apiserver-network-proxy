
kubectl delete secret konnectivity-agentca
kubectl delete secret konnectivity-agentclient
kubectl delete secret konnectivity-agentserver
kubectl delete secret konnectivity-proxyca
kubectl delete secret konnectivity-proxyserver

sleep 2

rm -rf easy-rsa-master/k8smaster
rm -rf easy-rsa-master/k8sagent

rm -rf certs/*.crt
rm -rf certs/*.key

	# set up easy-rsa
	cp -rf easy-rsa-master/easyrsa3 easy-rsa-master/k8smaster
	cp -rf easy-rsa-master/easyrsa3 easy-rsa-master/k8sagent
	# create the client <-> server-proxy connection certs
	cd easy-rsa-master/k8smaster; \
	./easyrsa init-pki; \
	./easyrsa --batch "--req-cn=127.0.0.1@$(date +%s)" build-ca nopass; \
	./easyrsa --subject-alt-name="DNS:konnectivity-proxyserver,DNS:localhost,IP:127.0.0.1" build-server-full "proxyserver" nopass; \
	./easyrsa build-client-full "proxyclient" nopass; \
	echo '{"signing":{"default":{"expiry":"43800h","usages":["signing","key encipherment","client auth"]}}}' > "ca-config.json"; \
	echo '{"CN":"proxy","names":[{"O":"system:nodes"}],"hosts":[""],"key":{"algo":"rsa","size":2048}}' | cfssl gencert -ca=pki/ca.crt -ca-key=pki/private/ca.key -config=ca-config.json - | cfssljson -bare proxy


cd ../../
cp -r easy-rsa-master/k8smaster/pki/private/proxyserver.key certs/
cp -r easy-rsa-master/k8smaster/pki/private/proxyclient.key certs/
cp easy-rsa-master/k8smaster/pki/private/ca.key certs/proxyca.key
cp -r easy-rsa-master/k8smaster/pki/issued/* certs/
cp easy-rsa-master/k8smaster/pki/ca.crt certs/proxyca.crt


	# create the agent <-> server-proxy connection certs
	cd easy-rsa-master/k8sagent; \
	./easyrsa init-pki; \
	./easyrsa --batch "--req-cn=127.0.0.1@$(date +%s)" build-ca nopass; \
	./easyrsa --subject-alt-name="DNS:konnectivity-agentserver,DNS:kubernetes,DNS:localhost,IP:127.0.0.1" build-server-full "agentserver" nopass; \
	./easyrsa build-client-full "agentclient" nopass; \
	echo '{"signing":{"default":{"expiry":"43800h","usages":["signing","key encipherment","agent auth"]}}}' > "ca-config.json"; \
	echo '{"CN":"proxy","names":[{"O":"system:nodes"}],"hosts":[""],"key":{"algo":"rsa","size":2048}}' | cfssl gencert -ca=pki/ca.crt -ca-key=pki/private/ca.key -config=ca-config.json - | cfssljson -bare proxy
	
cd ../../
cp -r easy-rsa-master/k8sagent/pki/private/agentserver.key certs/
cp -r easy-rsa-master/k8sagent/pki/private/agentclient.key certs/
cp easy-rsa-master/k8sagent/pki/private/ca.key certs/agentca.key
cp -r easy-rsa-master/k8sagent/pki/issued/* certs/
cp easy-rsa-master/k8sagent/pki/ca.crt certs/agentca.crt


cd certs

kubectl create secret tls konnectivity-proxyca --cert=proxyca.crt --key=proxyca.key
kubectl create secret tls konnectivity-proxyserver --cert=proxyserver.crt --key=proxyserver.key
kubectl create secret tls konnectivity-proxyclient --cert=proxyclient.crt --key=proxyclient.key

kubectl create secret tls konnectivity-agentca --cert=agentca.crt --key=agentca.key
kubectl create secret tls konnectivity-agentserver --cert=agentserver.crt --key=agentserver.key
kubectl create secret tls konnectivity-agentclient --cert=agentclient.crt --key=agentclient.key



