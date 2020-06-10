package k8sclient

//go:generate mockgen -destination=../util/mocks/$GOPACKAGE/client.go -package=$GOPACKAGE sigs.k8s.io/controller-runtime/pkg/client Client,StatusWriter
