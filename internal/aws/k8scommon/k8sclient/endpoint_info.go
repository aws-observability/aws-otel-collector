package k8sclient

type endpointInfo struct {
	name       string //service name
	namespace  string //namespace name
	podKeyList []string
}
