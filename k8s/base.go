package k8s

type Node struct {
	Name       string `json:"name"`
	CPU        string `json:"cpu"`
	Memory     string `json:"memory"`
	Disk       string `json:"disk"`
	Taints     string `json:"taints"`
	Roles      string `json:"roles"`
	Version    string `json:"version"`
	Age        string `json:"age"`
	Conditions string `json:"conditions"`
}

type NameSpaces struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Age    string            `json:"age"`
	Status string            `json:"status"`
}

type K8s interface {
	Type() string
	Name() string
	Init(string, string, string, string)
	TestConnect() error
	GetNode() []*Node
	GetNameSpace() []*NameSpaces
}
