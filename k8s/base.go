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

type Pod struct {
	Name       string `json:"name"`
	NameSpaces string `json:"name_spaces"`
	Restarts    int32    `json:"restarts"`
	Controlled string `json:"controlled"`
	Qos        string `json:"qos"`
	Age        string `json:"age"`
	Status     string `json:"status"`
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
	GetPod() []*Pod
	GetNameSpace() []*NameSpaces
}
