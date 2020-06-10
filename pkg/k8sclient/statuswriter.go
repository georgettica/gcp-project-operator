package k8sclient

//go:generate mockgen -destination=../util/mocks/$GOPACKAGE/statuswriter.go -package=$GOPACKAGE sigs.k8s.io/controller-runtime/pkg/client StatusWriter
