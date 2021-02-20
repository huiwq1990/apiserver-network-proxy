
kubectl delete secret   konnectivity-agentca

kubectl delete secret   konnectivity-agentclient
kubectl delete secret   konnectivity-ca
kubectl delete secret   konnectivity-agentserver
kubectl delete secret   konnectivity-proxyserver

sleep 4

# 1.创建 CA 私钥
openssl genrsa -out ca.key 2048

# 2.生成 CA 的自签名证书
openssl req \
    -subj "/C=CN/ST=mykey/L=mykey/O=mykey/OU=mykey/CN=konnectivity-ca/emailAddress=aa@aa.com'" \
    -new \
    -x509 \
    -days 3650 \
    -key ca.key \
    -out ca.crt

kubectl create secret tls konnectivity-ca \
  --cert=ca.crt \
  --key=ca.key


# agentserver
openssl genrsa -out agentserver.key 2048

#openssl req \
#    -subj "/C=AA/ST=AA/L=AA/O=AA/OU=AA/CN=konnectivity-agentserver/emailAddress=aa@aa.com" \
#    -new \
#    -key agentserver.key \
#    -out agentserver.csr

cat /etc/pki/tls/openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:konnectivity-agentserver,DNS:*.jd-tpaas,DNS:tpaas-openapi.jd-tpaas,IP:172.16.189.196") > agentserver.cnf
openssl req -new \
    -sha256 \
    -key agentserver.key \
    -subj "/C=CN/ST=GD/L=SZ/O=lee/OU=study/CN=konnectivity-agentserver" \
    -reqexts SAN \
    -config agentserver.cnf \
    -out agentserver.csr

#openssl x509 -req -days 3650 -in agentserver.csr -CA ca.crt -CAkey ca.key \
#    -CAcreateserial -CAserial serial \
#    -out agentserver.crt

openssl x509 -req -days 3650 -in agentserver.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out agentserver.crt -extfile agentserver.cnf -extensions req_ext

openssl x509 -text -noout -in agentserver.crt



kubectl create secret tls konnectivity-agentserver \
  --cert=agentserver.crt \
  --key=agentserver.key


#end agentserver

openssl genrsa -out proxyserver.key 2048

openssl req -new \
    -sha256 \
    -key proxyserver.key \
    -subj "/C=CN/ST=GD/L=SZ/O=lee/OU=study/CN=konnectivity-proxyserver" \
    -reqexts req_ext \
    -config <(cat /etc/pki/tls/openssl.cnf \
        <(printf "[req_ext]\nsubjectAltName=DNS:konnectivity-proxyserver,DNS:*.jd-tpaas,DNS:tpaas-openapi.jd-tpaas,IP:172.16.189.196")) \
    -out proxyserver.csr

openssl req -text -in proxyserver.csr


openssl x509 \
    -req \
    -days 3650 \
    -in proxyserver.csr \
    -CA ca.crt \
    -CAkey ca.key \
    -CAcreateserial -CAserial serial \
    -out proxyserver.crt \
     -extensions req_ext

openssl x509 -text -noout -in proxyserver.crt


kubectl create secret tls konnectivity-proxyserver \
  --cert=proxyserver.crt \
  --key=proxyserver.key



openssl genrsa -out agentclient.key 2048
openssl req \
    -subj "/C=AA/ST=AA/L=AA/O=AA/CN=konnectivity-agentclient/emailAddress=aa@aa.com" \
    -new \
    -key agentclient.key \
    -out agentclient.csr
openssl x509 \
    -req \
    -days 3650 \
    -in agentclient.csr \
    -CA ca.crt \
    -CAkey ca.key \
    -CAcreateserial -CAserial serial \
    -out agentclient.crt

kubectl create secret tls konnectivity-agentclient \
  --cert=agentclient.crt \
  --key=agentclient.key


