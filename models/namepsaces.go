package models

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type NameSpaces struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Age    string            `json:"age"`
	Status string            `json:"status"`
}

func QueryNameSpacesAll() *[]NameSpaces {
	var names []NameSpaces
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/tonny/.kube/config")

	clientset, _ := kubernetes.NewForConfig(config)

	ctx, _ := context.WithCancel(context.Background())
	nameList, _ := clientset.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	for _, name := range nameList.Items {
		var n NameSpaces
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
		names = append(names, n)
	}

	return &names
}
