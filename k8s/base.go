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
	Restarts    int32  `json:"restarts"`
	Controlled string `json:"controlled"`
	Qos        string `json:"qos"`
	Age        string `json:"age"`
	Status     string `json:"status"`
}

type NamespcePod struct {
	Name       string `json:"name"`
	NameSpaces string `json:"name_spaces"`
	Restarts    int32  `json:"restarts"`
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

type Deployment struct {
	Name         string            `json:"name"`
	NameSpace    string            `json:"name_space"`
	Replicas     string            `json:"replicas"`
	Labels       map[string]string `json:"labels"`
	Age          string            `json:"age"`
	Desired      int32             `json:"desired"`
	Current      int32             `json:"current"`
	Uptodate     int32             `json:"uptodate"`
	Available    int32             `json:"available"`
	StrategyType interface{}       `json:"strategy_type"`

}

type NamespacedDeployment struct {
	Name         string            `json:"name"`
	NameSpace    string            `json:"name_space"`
	Replicas     string            `json:"replicas"`
	Labels       map[string]string `json:"labels"`
	Age          string            `json:"age"`
	Desired      int32             `json:"desired"`
	Current      int32             `json:"current"`
	Uptodate     int32             `json:"uptodate"`
	Available    int32             `json:"available"`
	StrategyType interface{}       `json:"strategy_type"`
}

type Configmap struct {
	Name      string                `json:"name"`
	NameSpace string                `json:"name_space"`
	Age       string                `json:"age"`
	Data      map[string]string     `json:"data"`

}

type Daemonsets struct {
	Name         string            `json:"name"`
	NameSpace    string            `json:"name_space"`
	Labels       map[string]string `json:"labels"`
	Age          string            `json:"age"`
	Desired      int32             `json:"desired"`
	Current      int32             `json:"current"`
	Uptodate     int32             `json:"uptodate"`
	Available    int32             `json:"available"`
	Ready        int32             `json:"ready"`
}
type Statefulsets struct {
	Name      string               `json:"name"`
	NameSpace string               `json:"name_space"`
	Labels    map[string]string    `json:"labels"`
	Age       string               `json:"age"`
	Desired      int32             `json:"desired"`
	Current      int32             `json:"current"`
	Uptodate     int32             `json:"uptodate"`
	Available    int32             `json:"available"`
}

type Secret struct {
	Name      string               `json:"name"`
	NameSpace string               `json:"name_space"`
	Type      string               `json:"type"`
	Data      map[string][]byte    `json:"data_num"`
	Age string                     `json:"age"`
}
type Service struct {
	Name       string              `json:"name"`
	NameSpace  string              `json:"name_space"`
	Type       string              `json:"type"`
	ClusterIp  string              `json:"cluster_ip"`
	ExternalIp interface{}         `json:"external_ip"`
	Ports      interface{}         `json:"ports"`
	Age        string              `json:"age"`
	Labels     map[string]string   `json:"labels"`
	Selector   map[string]string   `json:"selector"`

}
type Endpoints struct {
	Name      string                  `json:"name"`
	NameSpace string                  `json:"name_space"`
	Labels    map[string]string       `json:"labels"`
	Age       string                  `json:"age"`
	Endpoint  interface{}             `json:"endpoint"`
}
type Serviceaccounts struct {
	Name          string              `json:"name"`
	NameSpace     string              `json:"name_space"`
	Age           string              `json:"age"`
	SecretsName   interface{}         `json:"secrets_name"`

}
type Event struct {
	Name string     `json:"name"`
	NameSpace string `json:"name_space"`
	LastSeen string   `json:"last_seen"`
	FirstSeen string   `json:"first_seen"`
	Kind string        `json:"kind"`
	Type string        `json:"type"`
	Counts int32      `json:"counts"`
	Reason string      `json:"reason"`
	Meessage string    `json:"meessage"`
}

type Job struct {
	Name string  `json:"name"`
	NameSpace string  `json:"name_space"`
	Completions interface{}  `json:"completions"`
	Age  string          `json:"age"`

}
type CronJob struct {
	Name string  `json:"name"`
	NameSpace string  `json:"name_space"`
	Schedule interface{} `json:"schedule"`
	Age  string          `json:"age"`
	Suspend interface{}  `json:"suspend"`
	Active  interface{}   `json:"active"`
	LastSchedule interface{}  `json:"last_schedule"`

}
type Storageclasses struct {
	Name string  `json:"name"`
	NameSpace string `json:"name_space"`
	Age string    `json:"age"`
	Provisioner string  `json:"provisioner"`
}
type Pv struct {
	Name string  `json:"name"`
	Capacity interface{} `json:"capacity"`
	AccessMode interface{} `json:"access_mode"`
	ReclaimPolicy interface{}  `json:"reclaim_policy"`
	Status interface{}   `json:"status"`
	Claim string     `json:"claim"`
	Storageclass string  `json:"storageclass"`
	Reason string    `json:"reason"`
	Age string       `json:"age"`
}
type Pvc struct {
	Name string `json:"name"`
	NameSpace string `json:"name_space"`
	Status interface{}  `json:"status"`
	Volume string  `json:"volume"`
	Capacity interface{} `json:"capacity"`
	AccessMode interface{} `json:"access_mode"`
	Storageclass interface{}  `json:"storageclass"`
	Age string       `json:"age"`
}

type Role struct {
	Name string `json:"name"`
	NameSpace string  `json:"name_space"`
	Age string  `json:"age"`
}
type Clusterrole struct {
	Name string  `json:"name"`
	Age string  `json:"age"`

}

type Rolebindings struct {
	Name string `json:"name"`
	NameSpace string  `json:"name_space"`
	Age string  `json:"age"`
}
type Clusterrolebinding struct {
	Name string  `json:"name"`
	Age string  `json:"age"`
}
type Replicaset struct {
	Name         string            `json:"name"`
	NameSpace    string            `json:"name_space"`
	Replicas     string            `json:"replicas"`
	Labels       map[string]string `json:"labels"`
	Age          string            `json:"age"`
	Desired      interface{}             `json:"desired"`
	Current      int32             `json:"current"`
	Ready        int32             `json:"ready"`
	OwnerReferences interface{}    `json:"owner_references"`

}
type K8s interface {
	Type() string
	Name() string
	Init(string, string, string, string)
	KubeconfigInit() error
	TestConnect() error
	GetNode() []*Node
	GetPod() []*Pod
	GetNameSpace() []*NameSpaces
	ListNamespacePod(namespace string) []*NamespcePod
	GetDeployment() []*Deployment
	GetNamespacedDeployment(namespace string) []*NamespacedDeployment
	GetConfigMap()  []*Configmap
	GetNamespacedConfigMap(namespace string)  []*Configmap
	GetDaemonsets()   []*Daemonsets
	GetNamspacedDaemonsets(namespace string)   []*Daemonsets
	GetStatefulsets()     []*Statefulsets
	GetNamespacedStatefulsets(namespace string)  []*Statefulsets
	GetSecret()   []*Secret
	GetNamespacedSecret(namespace string)  []*Secret
	Getservice()   []*Service
	GetNamespacedService(namespace string)  []*Service
	GetEndpoints()  []*Endpoints
	GetNamespacedEndpoints(namespace string)  []*Endpoints
	GetSA()   []*Serviceaccounts
	GetNamespacedSA(namespace string)  []*Serviceaccounts
	GetEvent()    []*Event
	GetNamespacedEvent(namespace string)  []*Event
	GetJob()   []*Job
	GetNamespacedJob(namespace string)  []*Job
	GetCronJob()  []*CronJob
	GetNamespacedCronjob(namespace string)  []*CronJob
	GetStorageclasses() []*Storageclasses
	GetPv()  []*Pv
	GetPvc() []*Pvc
	GetNamespacePvc(namespace string) []*Pvc
	GetRole()  []*Role
	GetNamespaceRole(namespace string)  []*Role
	GetClusterRole()   []*Clusterrole
	GetRolebindings()  []*Rolebindings
	GetNamespaceRolebindings(namespace string)  []*Rolebindings
	GetClusterrolebinding()  []*Clusterrolebinding
	GetReplicaset()  []*Replicaset
	GetNamespaceReplicaset(namespace string)  []*Replicaset
}