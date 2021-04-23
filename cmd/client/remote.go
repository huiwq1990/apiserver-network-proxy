/*
Copyright 2020 The OpenYurt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"k8s.io/apimachinery/pkg/util/httpstream"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"k8s.io/apimachinery/pkg/util/proxy"

	//
	//"github.com/openyurtio/openyurt/pkg/yurthub/cachemanager"
	//"github.com/openyurtio/openyurt/pkg/yurthub/healthchecker"
	//"github.com/openyurtio/openyurt/pkg/yurthub/transport"
	//"github.com/openyurtio/openyurt/pkg/yurthub/util"
	//
	//apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/klog/v2"
)

var (
	er = &errorResponder{}
)
type errorResponder struct {
}

func (e *errorResponder) Error(w http.ResponseWriter, req *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

var token = "eyJhbGciOiJSUzI1NiIsImtpZCI6InZpTTVMQlhQT3V2SUctWjNTbnJaUVJ0WWc1WnRkUDk2b1JhSmdBVnQwN1kifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InRlc3Qta29ubmVjdGl2aXR5LXRva2VuLXJuemZ2Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InRlc3Qta29ubmVjdGl2aXR5Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiMDZhYmEyY2MtZWIyNy00YzhkLWEzYzUtMGVlZThhNDRlNWI1Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6dGVzdC1rb25uZWN0aXZpdHkifQ.g-l4DzsJEP5hRtaz5tCjSOOd9ZxPs0k4y37vB0z3j5896a6inE9-q9eACOCLz-Cstel_y5Ft5PiBRWlVggdSkqZeqBT_wisOJaDeeRDV_LSgysSk08nLQWtVhC7ivS6sVdabrziRWIPhtKk8YtD-YdVrBrEOyB8A_JSc9AUvrE_dpTJOWFOaxF6YWXRfbQKlit1TrExiPz5lSj1kkWV8s-NfYM9oY3QPfQYbpcriHbbGiHzz55vL6vCtTyTnC840bBZTjIIWYCpeVrb57XgXYYMAeEpfwcvmLusPd2Y66eZApRPZRwxuQu1noGUfryrAmfNaoBfQ_VaqkZ5m_hgfug"


var cacert = `
-----BEGIN CERTIFICATE-----
MIIC5zCCAc+gAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTIxMDMwMjA5NTgxMVoXDTMxMDIyODA5NTgxMVowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJhM
HwfzDcYoWhOAZcOlNh6rg1Iini1KU7rPdEKP6FPeMjzIBI8rU4NzgYj95hOjGMCo
7auuTHgzdXHPWPV9J/ZrKNW+sx2CW8qbH0rWCe33a8aXLMd/3B3FCf6Hk3lj3pz6
5XyztAzRD79yVu84uv6sfPEKWqIe55+5zmuR5LFF8A13O8PDb+GvZPpWjdYWeRNp
OawPryty7kssZGRLRM/bV1KoWhVXjUl7d+SzLRTLG7Duhf4U+eCg8QkYGOrafmJE
DF7c7l01OoZ0OS64o0D/DbySgCZqACT4Y4x9UfURDo27BNPYFuqnS4KHBVYkrp3+
46XXq2OqiGHIbH54Uk0CAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFM/Ziuqiip55g5AQEZyY314FOGmfMA0GCSqGSIb3
DQEBCwUAA4IBAQBgXPrkcC1u2GJo5CFtmhDoMdxmiRFet+6cq6KEI3AxqnZNcVy4
jmL38W4moYrNUIxhw/fvZ5sCAGh/WnngUvP+EqOzjQnrQ5WqRh5a8gpB2CHAu/ps
lvlPjXzXm4SuIwvIuA1ePfr3wVGAF2E5zNRBAf/ftYlqmUXBSyyNY079CuASGH2Y
ZxIILSJVIaUPaRKtxQWpM1DlgeSZNs7Zgm+notvb//BN46aqFtV1u+vdZC8FHT/1
uLqdph1nhdd1qw6BthAKrcKxBrRQ7+d39aUGTMI0aCa+FGhYC9GrzgqfSh6GIVqV
hvEG24B/Z3RQyVaYQnH2CcKeavnUdlnQ7AAL
-----END CERTIFICATE-----
`
type ClusterInfo struct {
	Token string
	Server string
	CaCert []byte
}




func (c *Client)  startDummyServer(o *GrpcProxyClientOptions) {

	m := http.NewServeMux()
	m.HandleFunc("/ok", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	demoK8s := ClusterInfo{
		Token: token,
		Server:  "https://kubernetes.default",
		CaCert: []byte(cacert),
	}

	m.HandleFunc("/k8s", func(rw http.ResponseWriter, req *http.Request) {

		enableCacert := true
		enableCacertStr,ok := req.URL.Query()["cacert"]
		if ok && len(enableCacertStr)>0 && enableCacertStr[0]=="false" {
			enableCacert = false
		}

		enableAuth := true
		enableAuthStr,ok := req.URL.Query()["auth"]
		if ok && len(enableAuthStr)>0 && enableAuthStr[0]=="false" {
			enableAuth = false
		}

		// apiserver的地址
		u, err := url.Parse(demoK8s.Server)
		if err != nil {
			er.Error(rw, req, err)
			return
		}

		cid := GetHttpsClusterID(req)

		// 去除cluster前缀
		clusterPrefix := fmt.Sprintf("/k8s/clusters/%s", cid)
		u.Path = strings.TrimPrefix(req.URL.Path, clusterPrefix)
		u.RawQuery = req.URL.RawQuery

		klog.InfoS("request","origin path",req.URL.Path,"dest path",u.Path)

		u.RawQuery = req.URL.RawQuery

		req.URL.Scheme = "https"

		req.URL.Host = req.Host

		o.requestHost = "10.96.0.1"
		o.requestPort = 443

		dialer, err := c.getDialer(o)
		if err != nil {
			er.Error(rw, req, err)
			return
		}
		transport := &http.Transport{
			DialContext: dialer,
		}

		if enableCacert {
			certs := x509.NewCertPool()
			certs.AppendCertsFromPEM(demoK8s.CaCert)
			transport.TLSClientConfig = &tls.Config{
				RootCAs: certs,
			}
		}

		if enableAuth {
			req.Header.Add("Authorization",fmt.Sprintf("Bearer %s",token))
		}

		// 如果本身就是升级请求
		if httpstream.IsUpgradeRequest(req) {
			upgradeProxy := NewUpgradeProxy(u, transport)
			upgradeProxy.ServeHTTP(rw, req)
			return
		}
		klog.InfoS("request","url",u)
		// 支持升级的代理
		httpProxy := proxy.NewUpgradeAwareHandler(u, transport, true, false, er)
		httpProxy.ServeHTTP(rw, req)

	})

	m.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {

		u, err := url.Parse("http://kubia:80/")
		if err != nil {
			er.Error(rw, req, err)
			return
		}
		// 去除cluster前缀
		//u.Path = strings.TrimPrefix(req.URL.Path, prefix(r.cluster))
		u.RawQuery = req.URL.RawQuery

		req.URL.Scheme = "https"

		req.URL.Host = req.Host

		o.requestHost = "kubia"
		o.requestPort= 80

		dialer, err := c.getDialer(o)
		if err != nil {
			er.Error(rw, req, err)
			return
		}
		transport := &http.Transport{
			DialContext: dialer,
		}


		// 如果本身就是升级请求
		if httpstream.IsUpgradeRequest(req) {
			upgradeProxy := NewUpgradeProxy(u, transport)
			upgradeProxy.ServeHTTP(rw, req)
			return
		}
		// 支持升级的代理
		httpProxy := proxy.NewUpgradeAwareHandler(u, transport, true, false, er)
		httpProxy.ServeHTTP(rw, req)

		//w.WriteHeader(http.StatusOK)
		//n, err := w.Write([]byte(DummyServerResponse))
		//if err != nil {
		//	t.Fatalf("fail to write response: %v", err)
		//}
		//if n != len([]byte(DummyServerResponse)) {
		//	t.Fatalf("fail to write response: write %d of the %d bytes",
		//		n, len([]byte(DummyServerResponse)))
		//}
	})

	s := &http.Server{
		//TLSConfig: &tls.Config{
		//	Certificates: []tls.Certificate{
		//		genCert(t, GeneircCertFile, GenericKeyFile),
		//	},
		//	ClientCAs: genCAPool(t, RootCAFile),
		//},
		Addr:    fmt.Sprintf(":%d", 8088),
		Handler: m,
	}

	klog.Infof("[TEST] dummy-server is listening on :%d", 8088)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}





// http升级请求的代理
type UpgradeProxy struct {
	Location  *url.URL
	Transport http.RoundTripper
}

func NewUpgradeProxy(location *url.URL, transport http.RoundTripper) *UpgradeProxy {
	return &UpgradeProxy{
		Location:  location,
		Transport: transport,
	}
}

func (p *UpgradeProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	loc := *p.Location
	loc.RawQuery = req.URL.RawQuery

	newReq := req.WithContext(req.Context())
	newReq.Header = CloneHeader(req.Header)
	newReq.URL = &loc

	httpProxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: p.Location.Scheme, Host: p.Location.Host})
	httpProxy.Transport = p.Transport
	httpProxy.ServeHTTP(rw, newReq)
}

func CloneHeader(in http.Header) http.Header {
	out := make(http.Header, len(in))
	for key, values := range in {
		newValues := make([]string, len(values))
		copy(newValues, values)
		out[key] = newValues
	}
	return out
}


func GetHttpsClusterID(req *http.Request) string {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) > 3 && parts[0] == "" &&
		parts[1] == "k8s" && parts[2] == "clusters" {
		return parts[3]
	}
	return ""
}
