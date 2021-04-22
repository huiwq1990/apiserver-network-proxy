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
	"fmt"
	"net/http"

	"net/url"

	"k8s.io/apimachinery/pkg/util/proxy"

	//
	//"github.com/openyurtio/openyurt/pkg/yurthub/cachemanager"
	//"github.com/openyurtio/openyurt/pkg/yurthub/healthchecker"
	//"github.com/openyurtio/openyurt/pkg/yurthub/transport"
	//"github.com/openyurtio/openyurt/pkg/yurthub/util"
	//
	//apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/klog"
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


func (c *Client)  startDummyServer(o *GrpcProxyClientOptions) {

	m := http.NewServeMux()


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

		dialer, err := c.getDialer(o)
		if err != nil {
			er.Error(rw, req, err)
			return
		}
		transport := &http.Transport{
			DialContext: dialer,
		}


		// 如果本身就是升级请求
		//if httpstream.IsUpgradeRequest(req) {
		//	upgradeProxy := NewUpgradeProxy(&u, transport)
		//	upgradeProxy.ServeHTTP(rw, req)
		//	return
		//}
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
	if err := s.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}
}
