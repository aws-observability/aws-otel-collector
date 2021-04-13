package k8sclient

type jobInfo struct {
	name   string
	owners []*jobOwner
}

type jobOwner struct {
	kind string
	name string
}
