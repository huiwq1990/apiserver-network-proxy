module sigs.k8s.io/apiserver-network-proxy

go 1.12

require (
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/prometheus/client_golang v1.7.1
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.5
	github.com/tsenart/vegeta/v12 v12.8.4
	golang.org/x/net v0.0.0-20191004110552-13f9640d40b9 // indirect
	google.golang.org/grpc v1.27.0
	k8s.io/api v0.17.1
	k8s.io/apimachinery v0.17.1
	k8s.io/client-go v0.17.1
	k8s.io/component-base v0.17.1
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.0.0
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.0
)

replace sigs.k8s.io/apiserver-network-proxy/konnectivity-client => ./konnectivity-client
