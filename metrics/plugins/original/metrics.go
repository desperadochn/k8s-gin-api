package original

import (
	"encoding/base64"
	"k8s.io/client-go/rest"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

type MetricsK8s struct {
	config rest.Config
	clientset  *metricsv.Clientset
}
func (a *MetricsK8s) MetricsType() string {
	return "MetricsK8s"
}

func (a *MetricsK8s) MetricsName() string {
	return "MetricsK8s"
}
func (a *MetricsK8s)MetricsInit(Host, CAData, CertData, KeyData string)  {
	a.config.Host = Host
	a.config.CAData, _ = base64.StdEncoding.DecodeString(CAData)
	a.config.CertData, _ = base64.StdEncoding.DecodeString(CertData)
	a.config.KeyData, _ = base64.StdEncoding.DecodeString(KeyData)
}
func (a *MetricsK8s)MetricsTestConnect() error  {
	clientset, err := metricsv.NewForConfig(&a.config)
	if err == nil {
		a.clientset = clientset
	}
	return err

}
//func (a *MetricsK8s)GetNodeMetrics() []*metrics.NodeMetrics {
//	var nodeMetrics []*metrics.NodeMetrics
//	ctx, _ := context.WithCancel(context.Background())
//	nodeMetricsList, _ := a.clientset.MetricsV1beta1().NodeMetricses().List(ctx,metav1.ListOptions{})
//	for nodeMetricses , _ := range nodeMetricsList.Items {
//		var n metrics.NodeMetrics
//		n.Name = nodeMetricses
//	}
//
//}


