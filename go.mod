module github.com/openshift/gcp-project-operator

require (
	git.apache.org/thrift.git v0.12.0 // indirect
	github.com/Azure/go-autorest v12.2.0+incompatible // indirect
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/emicklei/go-restful v2.11.2+incompatible // indirect
	github.com/go-logr/logr v0.1.0
	github.com/go-openapi/spec v0.19.3
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/golang/mock v1.4.3
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/openshift/cluster-api v0.0.0-20191129101638-b09907ac6668
	github.com/operator-framework/operator-sdk v0.18.1
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/rogpeppe/go-internal v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.5.1
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/tools v0.0.0-20200522201501-cb1345f3a375 // indirect
	google.golang.org/api v0.14.0
	google.golang.org/appengine v1.6.5 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.2
	k8s.io/gengo v0.0.0-20200114144118-36b2048a9120
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20200427153329-656914f816f9
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/controller-tools v0.3.0
)

replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.38.1-0.20200424145508-7e176fda06cc
	k8s.io/client-go => k8s.io/client-go v0.18.2
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20181117043124-c2090bec4d9b
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20180711000925-0cf8f7e6ed1d
)

// Pin google api to v0.11.0
replace google.golang.org/api => google.golang.org/api v0.11.0

// Pin hive dep
replace github.com/openshift/cluster-network-operator => github.com/openshift/cluster-network-operator v0.0.0-20190207145423-c226dcab667e

go 1.13
