package models

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

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

func QueryNodeAll() *[]Node {
	var nodes []Node
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/tonny/.kube/config")

	clientset, _ := kubernetes.NewForConfig(config)

	ctx, _ := context.WithCancel(context.Background())
	nodeList, _ := clientset.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	for _, node := range nodeList.Items {
		var n Node
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

		nodes = append(nodes, n)
	}
	return &nodes
}
