package k8sclient

type replicaSetInfo struct {
	name   string
	owners []*replicaSetOwner
}

type replicaSetOwner struct {
	kind string
	name string
}
