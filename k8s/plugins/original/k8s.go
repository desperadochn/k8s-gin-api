package original

import (
	"context"
	"encoding/base64"
	"fmt"
	"jarvis_server/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

type OriginalK8s struct {
	config    rest.Config
	clientset *kubernetes.Clientset
}

func (a *OriginalK8s) Type() string {
	return "OriginalK8s"
}

func (a *OriginalK8s) Name() string {
	return "k8s"
}

// 初始化
func (a *OriginalK8s) Init(Host, CAData, CertData, KeyData string) {
	a.config.Host = Host
	a.config.CAData, _ = base64.StdEncoding.DecodeString(CAData)
	a.config.CertData, _ = base64.StdEncoding.DecodeString(CertData)
	a.config.KeyData, _ = base64.StdEncoding.DecodeString(KeyData)
}

// 连接测试
func (a *OriginalK8s) TestConnect() error {
	clientset, _ := kubernetes.NewForConfig(&a.config)
	ctx, _ := context.WithCancel(context.Background())
	_, err := clientset.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err == nil {
		a.clientset = clientset
	}
	return err
}

// 获取实例
func (a *OriginalK8s) GetNode() []*k8s.Node {
	var nodes []*k8s.Node
	ctx, _ := context.WithCancel(context.Background())
	nodeList, err := a.clientset.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return []*k8s.Node{}
	}
	for _, node := range nodeList.Items {
		var n k8s.Node
		n.Name = node.Name
		n.Version = node.Status.NodeInfo.KubeletVersion
		for _, value := range node.Status.Conditions {
			if value.Type == "Ready" {
				if value.Status == "True" {
					n.Conditions = "Ready"
				} else {
					n.Conditions = "No Ready"
				}
				t1 := value.LastHeartbeatTime.Time
				sub := t1.Sub(value.LastTransitionTime.Time)
				if hours := sub.Hours(); hours > 0 {
					n.Age = fmt.Sprintf("%.0f", hours/24)
				} else {
					n.Age = "0"
				}
				break
			}
		}
		n.CPU = fmt.Sprintf("%s", node.Status.Capacity.Cpu())
		n.Memory = fmt.Sprintf("%s", node.Status.Capacity.Memory())

		nodes = append(nodes, &n)
	}
	return nodes
}

func (a *OriginalK8s) GetNameSpace() []*k8s.NameSpaces {
	var names []*k8s.NameSpaces
	ctx, _ := context.WithCancel(context.Background())
	nameList, _ := a.clientset.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	for _, name := range nameList.Items {
		var n k8s.NameSpaces
		n.Name = name.Name
		n.Labels = name.Labels
		n.Status = fmt.Sprintf("%s", name.Status.Phase)
		t1 := time.Now()
		sub := t1.Sub(name.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		names = append(names, &n)
	}

	return names
}

func init() {
	k8s.DefaultManager.Register(&OriginalK8s{})
}
