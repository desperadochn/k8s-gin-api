package original

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"jarvis_server/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
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
func (a *OriginalK8s) KubeconfigInit() error {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, _ := kubernetes.NewForConfig(config)
	ctx, _ := context.WithCancel(context.Background())
	_, err = clientset.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err == nil {
		a.clientset = clientset
	}
	return err
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

func (a *OriginalK8s) GetNode() []*k8s.Node {
	var nodes []*k8s.Node
	ctx, _ := context.WithCancel(context.Background())
	nodeList, _ := a.clientset.CoreV1().Nodes().List(ctx, v1.ListOptions{})
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
		fmt.Println(nodes)
	}
	return nodes
}

func (a *OriginalK8s)GetStorageclasses() []*k8s.Storageclasses {
	var storageclasses []*k8s.Storageclasses
	ctx, _ := context.WithCancel(context.Background())
	storageclassesList, _ := a.clientset.StorageV1().StorageClasses().List(ctx,v1.ListOptions{})
	for _, storageclass := range storageclassesList.Items {
		var n k8s.Storageclasses
		n.Name = storageclass.Name
		n.NameSpace = storageclass.Namespace
		n.Provisioner = storageclass.Provisioner
		t1 := time.Now()
		sub := t1.Sub(storageclass.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		storageclasses = append(storageclasses,&n)

	}
	return storageclasses
}

func (a *OriginalK8s)GetPvc() []*k8s.Pvc  {
	var Pvc []*k8s.Pvc
	ctx, _ := context.WithCancel(context.Background())
	PvcLister, _ := a.clientset.CoreV1().PersistentVolumeClaims("").List(ctx,v1.ListOptions{})
	for _, Pvcs := range PvcLister.Items{
		var n k8s.Pvc
		n.Name = Pvcs.Name
		n.NameSpace = Pvcs.Namespace
		n.Storageclass = Pvcs.Spec.StorageClassName
		n.Status = Pvcs.Status.Phase
		n.AccessMode = Pvcs.Spec.AccessModes
		n.Capacity = Pvcs.Status.Capacity
		n.Volume = Pvcs.Spec.VolumeName
		t1 := time.Now()
		sub := t1.Sub(Pvcs.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		Pvc = append(Pvc,&n)
	}
	return Pvc
}
func (a *OriginalK8s)GetNamespacePvc(namespace string) []*k8s.Pvc {
	var Pvc []*k8s.Pvc
	ctx, _ := context.WithCancel(context.Background())
	PvcLister, err := a.clientset.CoreV1().PersistentVolumeClaims(namespace).List(ctx,v1.ListOptions{})

	if err != nil {
		fmt.Println(err)
	}
	for _, Pvcs := range PvcLister.Items{
		var n k8s.Pvc
		n.Name = Pvcs.Name
		n.NameSpace = Pvcs.Namespace
		n.Storageclass = Pvcs.Spec.StorageClassName
		n.Status = Pvcs.Status.Phase
		n.AccessMode = Pvcs.Spec.AccessModes
		n.Capacity = Pvcs.Status.Capacity
		n.Volume = Pvcs.Spec.VolumeName
		t1 := time.Now()
		sub := t1.Sub(Pvcs.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		Pvc = append(Pvc,&n)
	}
	println(Pvc)
	return Pvc
}
func (a *OriginalK8s)GetPv() []*k8s.Pv  {

	var PVs []*k8s.Pv
	ctx, _ := context.WithCancel(context.Background())
	PVLister, _ := a.clientset.CoreV1().PersistentVolumes().List(ctx,v1.ListOptions{
		TypeMeta: v1.TypeMeta{Kind: v1.NamespaceAll},
	})
	for _, Pv := range PVLister.Items{
		var n k8s.Pv
		n.Name = Pv.Name
		n.Capacity = Pv.Spec.Capacity
		n.AccessMode = Pv.Spec.AccessModes
		n.ReclaimPolicy = Pv.Spec.PersistentVolumeReclaimPolicy
		n.Status = Pv.Status.Phase
		n.Reason = Pv.Status.Reason
		ClaimRefName :=  Pv.Spec.ClaimRef.Name
		ClaimRefNameSpace := Pv.Spec.ClaimRef.Namespace
		n.Claim = ClaimRefNameSpace + "/" + ClaimRefName
		n.Storageclass = Pv.Spec.StorageClassName
		t1 := time.Now()
		sub := t1.Sub(Pv.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		PVs = append(PVs,&n)

	}
	return PVs
}
func (a *OriginalK8s)GetEvent() []*k8s.Event  {
	var events []*k8s.Event
	ctx, _ := context.WithCancel(context.Background())
	eventsList, _ := a.clientset.CoreV1().Events("").List(ctx,v1.ListOptions{})
	for _, event := range eventsList.Items{
		var n k8s.Event
		n.Name = event.Name
		n.NameSpace = event.Namespace
		n.Meessage = event.Message
		n.Reason = event.Reason
		n.Type = event.Type
		n.Kind = event.Kind
		n.Counts = event.Count
		t1 := time.Now()
		firstSub := t1.Sub(event.FirstTimestamp.Time)
		lastSub := t1.Sub(event.LastTimestamp.Time)
		if firstHours := firstSub.Hours(); firstHours > 0 {
			n.FirstSeen = fmt.Sprintf("%.0f", firstHours/24)
		} else {
			n.FirstSeen = "0"
		}
		if lastHours := lastSub.Hours(); lastHours > 0  {
			n.LastSeen = fmt.Sprintf("%.0f", lastHours/24)
		} else {
			n.LastSeen = "0"
		}
		events = append(events,&n)
	}
	return events
}
func (a *OriginalK8s)GetNamespacedEvent(namespace string) []*k8s.Event  {
	var events []*k8s.Event
	ctx, _ := context.WithCancel(context.Background())
	eventsList, _ := a.clientset.CoreV1().Events(namespace).List(ctx,v1.ListOptions{})
	for _, event := range eventsList.Items{
		var n k8s.Event
		n.Name = event.Name
		n.NameSpace = event.Namespace
		n.Meessage = event.Message
		n.Reason = event.Reason
		n.Type = event.Type
		n.Kind = event.Kind
		n.Counts = event.Count
		t1 := time.Now()
		firstSub := t1.Sub(event.FirstTimestamp.Time)
		lastSub := t1.Sub(event.LastTimestamp.Time)
		if firstHours := firstSub.Hours(); firstHours > 0 {
			n.FirstSeen = fmt.Sprintf("%.0f", firstHours/24)
		} else {
			n.FirstSeen = "0"
		}
		if lastHours := lastSub.Hours(); lastHours > 0  {
			n.LastSeen = fmt.Sprintf("%.0f", lastHours/24)
		} else {
			n.LastSeen = "0"
		}
		events = append(events,&n)
	}
	return events
}

func (a *OriginalK8s) GetEndpoints() []*k8s.Endpoints  {
	var endpoints []*k8s.Endpoints
	ctx, _ := context.WithCancel(context.Background())
	endpointsList, _ := a.clientset.CoreV1().Endpoints("").List(ctx,v1.ListOptions{})
	for _, endpoint := range endpointsList.Items{
		var n k8s.Endpoints
		n.Name = endpoint.Name
		n.NameSpace = endpoint.Namespace
		n.Labels = endpoint.Labels
		n.Endpoint = endpoint.Subsets
		t1 := time.Now()
		sub := t1.Sub(endpoint.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		endpoints = append(endpoints,&n)

	}
	return endpoints
}

func (a *OriginalK8s) GetNamespacedEndpoints(namespcae string) []*k8s.Endpoints {
	var endpoints []*k8s.Endpoints
	ctx, _ := context.WithCancel(context.Background())
	endpointsList, _ := a.clientset.CoreV1().Endpoints(namespcae).List(ctx,v1.ListOptions{})
	for _, endpoint := range endpointsList.Items{
		var n k8s.Endpoints
		n.Name = endpoint.Name
		n.NameSpace = endpoint.Namespace
		n.Labels = endpoint.Labels
		n.Endpoint = endpoint.Subsets
		t1 := time.Now()
		sub := t1.Sub(endpoint.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		endpoints = append(endpoints,&n)

	}
	return endpoints
}
func (a *OriginalK8s) ListNamespacePod (namespace string) []*k8s.NamespcePod  {
	var pods []*k8s.NamespcePod
	ctx, _ := context.WithCancel(context.Background())
	podList, _ := a.clientset.CoreV1().Pods(namespace).List(ctx,v1.ListOptions{})
	for _, pod := range podList.Items{
		var n k8s.NamespcePod
		println(pod.Name)
		n.NameSpaces = pod.Namespace
		n.Status = fmt.Sprintf("%s", pod.Status.Phase)
		for _, value := range pod.OwnerReferences{
			if value.Kind != "" {
				n.Controlled = value.Kind
				break
			}
		}
		n.Qos = fmt.Sprintf("%s", pod.Status.QOSClass)
		t1 := time.Now()
		sub := t1.Sub(pod.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		pods = append(pods, &n)

	}
	return pods

}
func (a *OriginalK8s) GetStatefulsets () []*k8s.Statefulsets {
	var statefulset []*k8s.Statefulsets
	ctx, _ := context.WithCancel(context.Background())
	statefulsetLister, _ := a.clientset.AppsV1().StatefulSets("").List(ctx,v1.ListOptions{})
	for _,  statefulsets := range statefulsetLister.Items{
		var n k8s.Statefulsets
		n.Name = statefulsets.Name
		n.NameSpace = statefulsets.Name
		n.Labels = statefulsets.Labels
		n.Desired = *statefulsets.Spec.Replicas
		n.Current = statefulsets.Status.ReadyReplicas
		n.Uptodate = statefulsets.Status.UpdatedReplicas
		t1 := time.Now()
		sub := t1.Sub(statefulsets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		statefulset = append(statefulset,&n)
	}
	return statefulset

}

func (a *OriginalK8s)GetNamespacedStatefulsets(namespace string) []*k8s.Statefulsets {
	var statefulset []*k8s.Statefulsets
	ctx, _ := context.WithCancel(context.Background())
	statefulsetLister, _ := a.clientset.AppsV1().StatefulSets(namespace).List(ctx,v1.ListOptions{})
	for _,  statefulsets := range statefulsetLister.Items{
		var n k8s.Statefulsets
		n.Name = statefulsets.Name
		n.NameSpace = statefulsets.Name
		n.Labels = statefulsets.Labels
		n.Desired = *statefulsets.Spec.Replicas
		n.Current = statefulsets.Status.ReadyReplicas
		n.Uptodate = statefulsets.Status.UpdatedReplicas
		t1 := time.Now()
		sub := t1.Sub(statefulsets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		statefulset = append(statefulset,&n)
	}
	return statefulset
}

func (a *OriginalK8s)GetSecret() []*k8s.Secret  {
	var sercret []*k8s.Secret
	ctx, _ := context.WithCancel(context.Background())
	sercretLister, _ := a.clientset.CoreV1().Secrets("").List(ctx,v1.ListOptions{})
	for _, sercrets := range sercretLister.Items{
		var n k8s.Secret
		n.Name = sercrets.Name
		n.NameSpace = sercrets.Namespace
		n.Type = string(sercrets.Type)
		n.Data = sercrets.Data
		t1 := time.Now()
		sub := t1.Sub(sercrets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		sercret = append(sercret,&n)
	}
	return sercret
}
func (a *OriginalK8s)GetNamespacedSecret(namespace string) []*k8s.Secret {
	var sercret []*k8s.Secret
	ctx, _ := context.WithCancel(context.Background())
	sercretLister, _ := a.clientset.CoreV1().Secrets(namespace).List(ctx,v1.ListOptions{})
	for _, sercrets := range sercretLister.Items{
		var n k8s.Secret
		n.Name = sercrets.Name
		n.NameSpace = sercrets.Namespace
		n.Type = string(sercrets.Type)
		n.Data = sercrets.Data
		t1 := time.Now()
		sub := t1.Sub(sercrets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		sercret = append(sercret,&n)
	}
	return sercret
}
func (a *OriginalK8s)GetDeployment() []*k8s.Deployment {
	var deploy []*k8s.Deployment
	ctx, _ := context.WithCancel(context.Background())
	deployLister, _ := a.clientset.AppsV1().Deployments("").List(ctx, v1.ListOptions{})
	for _, deployment := range deployLister.Items{
		var n k8s.Deployment
		n.Name = deployment.Name
		n.Labels = deployment.Labels
		n.NameSpace = deployment.Namespace
		n.Available = deployment.Status.AvailableReplicas
		n.Current = deployment.Status.ReadyReplicas
		n.Desired = *deployment.Spec.Replicas
		n.StrategyType = deployment.Spec.Strategy.Type
		t1 := time.Now()
		sub := t1.Sub(deployment.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		deploy =append(deploy,&n)
	}
	return deploy
}
func (a *OriginalK8s) GetConfigMap() []*k8s.Configmap {
	var cm  []*k8s.Configmap
	ctx, _ := context.WithCancel(context.Background())
	configmapLister, _ := a.clientset.CoreV1().ConfigMaps("").List(ctx,v1.ListOptions{})
	for _, configmap := range configmapLister.Items {
		var n k8s.Configmap
		n.Name = configmap.Name
		n.NameSpace = configmap.Namespace
		n.Data = configmap.Data
		t1 := time.Now()
		sub := t1.Sub(configmap.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		cm =append(cm,&n)
	}
	return cm

}
func (a *OriginalK8s) GetNamespacedConfigMap(namespace string) []*k8s.Configmap {
	var cm  []*k8s.Configmap
	ctx, _ := context.WithCancel(context.Background())
	configmapLister, _ := a.clientset.CoreV1().ConfigMaps(namespace).List(ctx,v1.ListOptions{})
	for _, configmap := range configmapLister.Items {
		var n k8s.Configmap
		n.Name = configmap.Name
		n.NameSpace = configmap.Namespace
		n.Data = configmap.Data
		t1 := time.Now()
		sub := t1.Sub(configmap.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		cm =append(cm,&n)
	}
	return cm
}

func (a *OriginalK8s)GetNamespacedCronjob(namespace string) []*k8s.CronJob  {
	var cronjob []*k8s.CronJob
	ctx, _ := context.WithCancel(context.Background())
	cronjobLister,_ := a.clientset.BatchV1beta1().CronJobs(namespace).List(ctx,v1.ListOptions{})
	for _, cronjobs := range cronjobLister.Items {
		var n k8s.CronJob
		n.Name = cronjobs.Name
		n.NameSpace = cronjobs.Namespace
		n.Schedule = cronjobs.Spec.Schedule
		n.Suspend = cronjobs.Spec.Suspend
		n.Active = cronjobs.Status.Active
		n.LastSchedule = cronjobs.Status.LastScheduleTime
		t1 := time.Now()
		sub := t1.Sub(cronjobs.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		cronjob = append(cronjob,&n)
	}
	return cronjob
}
func (a *OriginalK8s)GetCronJob() []*k8s.CronJob  {
	var cronjob []*k8s.CronJob
	ctx, _ := context.WithCancel(context.Background())
	cronjobLister,_ := a.clientset.BatchV1beta1().CronJobs("").List(ctx,v1.ListOptions{})
	for _, cronjobs := range cronjobLister.Items {
		var n k8s.CronJob
		n.Name = cronjobs.Name
		n.NameSpace = cronjobs.Namespace
		n.Schedule = cronjobs.Spec.Schedule
		n.Suspend = cronjobs.Spec.Suspend
		n.Active = cronjobs.Status.Active
		n.LastSchedule = cronjobs.Status.LastScheduleTime
		t1 := time.Now()
		sub := t1.Sub(cronjobs.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		cronjob = append(cronjob,&n)
	}
	return cronjob
}
func (a *OriginalK8s)GetJob() []*k8s.Job {
	var job []*k8s.Job
	ctx, _ := context.WithCancel(context.Background())
	jobLister,_ := a.clientset.BatchV1().Jobs("").List(ctx,v1.ListOptions{})
	for _, jobs := range jobLister.Items{
		var n k8s.Job
		n.Name = jobs.Name
		n.NameSpace = jobs.Namespace
		n.Completions = jobs.Spec.Completions
		t1 := time.Now()
		sub := t1.Sub(jobs.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		job = append(job,&n)
	}
	return job
}
func (a *OriginalK8s)GetRole() []*k8s.Role  {
	var role []*k8s.Role
	ctx, _ := context.WithCancel(context.Background())
	roleLister, _ := a.clientset.RbacV1().Roles("").List(ctx,v1.ListOptions{})
	for _, roles := range roleLister.Items {
		var n k8s.Role
		n.Name = roles.Name
		n.NameSpace = roles.Namespace
		t1 := time.Now()
		sub := t1.Sub(roles.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		role = append(role,&n)
	}
	return role
}
func (a *OriginalK8s)GetNamespaceRole(namespace string) []*k8s.Role  {
	var role []*k8s.Role
	ctx, _ := context.WithCancel(context.Background())
	roleLister, _ := a.clientset.RbacV1().Roles(namespace).List(ctx,v1.ListOptions{})
	for _, roles := range roleLister.Items {
		var n k8s.Role
		n.Name = roles.Name
		n.NameSpace = roles.Namespace
		t1 := time.Now()
		sub := t1.Sub(roles.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		role = append(role,&n)
	}
	return role
}
func (a *OriginalK8s)GetClusterRole() []*k8s.Clusterrole {
	var clusterrole []*k8s.Clusterrole
	ctx, _ := context.WithCancel(context.Background())
	clusterroleLister, _ := a.clientset.RbacV1().ClusterRoles().List(ctx,v1.ListOptions{})
	for _, clusterroles := range clusterroleLister.Items {
		var n k8s.Clusterrole
		n.Name = clusterroles.Name
		t1 := time.Now()
		sub := t1.Sub(clusterroles.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		clusterrole = append(clusterrole,&n)
	}
	return clusterrole

}
func (a *OriginalK8s)GetRolebindings() []*k8s.Rolebindings  {
	var rolebinding []*k8s.Rolebindings
	ctx, _ := context.WithCancel(context.Background())
	rolebindingLister, _ := a.clientset.RbacV1().RoleBindings("").List(ctx,v1.ListOptions{})
	for _, rolebindings := range rolebindingLister.Items {
		var n k8s.Rolebindings
		n.Name = rolebindings.Name
		n.NameSpace = rolebindings.Namespace
		t1 := time.Now()
		sub := t1.Sub(rolebindings.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		rolebinding = append(rolebinding,&n)
	}
	return rolebinding
}
func (a *OriginalK8s)GetNamespaceRolebindings(namespace string) []*k8s.Rolebindings {
	var rolebinding []*k8s.Rolebindings
	ctx, _ := context.WithCancel(context.Background())
	rolebindingLister, _ := a.clientset.RbacV1().RoleBindings(namespace).List(ctx,v1.ListOptions{})
	for _, rolebindings := range rolebindingLister.Items {
		var n k8s.Rolebindings
		n.Name = rolebindings.Name
		n.NameSpace = rolebindings.Namespace
		t1 := time.Now()
		sub := t1.Sub(rolebindings.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		rolebinding = append(rolebinding,&n)
	}
	return rolebinding
}
func (a *OriginalK8s)GetClusterrolebinding() []*k8s.Clusterrolebinding  {
	var clusterrolebinding []*k8s.Clusterrolebinding
	ctx, _ := context.WithCancel(context.Background())
	clusterrolebindingLister, _ := a.clientset.RbacV1().ClusterRoleBindings().List(ctx,v1.ListOptions{})
	for _, clusterrolebindings := range clusterrolebindingLister.Items{
		var n k8s.Clusterrolebinding
		n.Name = clusterrolebindings.Name
		t1 := time.Now()
		sub := t1.Sub(clusterrolebindings.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		clusterrolebinding = append(clusterrolebinding,&n)
	}
	return clusterrolebinding
}
func (a *OriginalK8s)GetNamespacedJob(namespace string) []*k8s.Job {
	var job []*k8s.Job
	ctx, _ := context.WithCancel(context.Background())
	jobLister,_ := a.clientset.BatchV1().Jobs(namespace).List(ctx,v1.ListOptions{})
	for _, jobs := range jobLister.Items{
		var n k8s.Job
		n.Name = jobs.Name
		n.NameSpace = jobs.Namespace
		n.Completions = jobs.Spec.Completions
		t1 := time.Now()
		sub := t1.Sub(jobs.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		job = append(job,&n)
	}
	return job
}
func (a *OriginalK8s) GetDaemonsets() []*k8s.Daemonsets  {
	var daemonset []*k8s.Daemonsets
	ctx, _ := context.WithCancel(context.Background())
	daemonsetsLister, _ := a.clientset.AppsV1().DaemonSets("").List(ctx,v1.ListOptions{})
	for _, daemonsets := range daemonsetsLister.Items{
		var n k8s.Daemonsets
		n.Name = daemonsets.Name
		n.Labels = daemonsets.Labels
		n.NameSpace = daemonsets.Namespace
		n.Desired = daemonsets.Status.DesiredNumberScheduled
		n.Current = daemonsets.Status.CurrentNumberScheduled
		n.Available = daemonsets.Status.NumberAvailable
		n.Uptodate = daemonsets.Status.UpdatedNumberScheduled
		n.Ready = daemonsets.Status.NumberReady
		t1 := time.Now()
		sub := t1.Sub(daemonsets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		daemonset =append(daemonset,&n)

	}
	return daemonset
}

func (a *OriginalK8s) GetNamspacedDaemonsets(namespace string) []*k8s.Daemonsets  {
	var daemonset []*k8s.Daemonsets
	ctx, _ := context.WithCancel(context.Background())
	daemonsetsLister, _ := a.clientset.AppsV1().DaemonSets(namespace).List(ctx,v1.ListOptions{})
	for _, daemonsets := range daemonsetsLister.Items{
		var n k8s.Daemonsets
		n.Name = daemonsets.Name
		n.Labels = daemonsets.Labels
		n.NameSpace = daemonsets.Namespace
		n.Desired = daemonsets.Status.DesiredNumberScheduled
		n.Current = daemonsets.Status.CurrentNumberScheduled
		n.Available = daemonsets.Status.NumberAvailable
		n.Uptodate = daemonsets.Status.UpdatedNumberScheduled
		n.Ready = daemonsets.Status.NumberReady
		t1 := time.Now()
		sub := t1.Sub(daemonsets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		daemonset =append(daemonset,&n)

	}
	return daemonset
}
func (a *OriginalK8s)GetReplicaset() []*k8s.Replicaset{
	var replicaset []*k8s.Replicaset
	ctx, _ := context.WithCancel(context.Background())
	replicasetLister, _ := a.clientset.AppsV1().ReplicaSets("").List(ctx,v1.ListOptions{})
	for _, replicasets := range replicasetLister.Items {
		var n k8s.Replicaset
		n.Name = replicasets.Name
		n.NameSpace = replicasets.Namespace
		n.Labels = replicasets.Labels
		n.Ready = replicasets.Status.ReadyReplicas
		n.Current = replicasets.Status.AvailableReplicas
		n.Desired = replicasets.Spec.Replicas
		n.OwnerReferences = replicasets.OwnerReferences
		n.Desired =replicasets.Spec.Replicas
		t1 := time.Now()
		sub := t1.Sub(replicasets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		replicaset = append(replicaset,&n)
	}
	return replicaset
}
func (a *OriginalK8s)GetNamespaceReplicaset(namespace string) []*k8s.Replicaset {
	var replicaset []*k8s.Replicaset
	ctx, _ := context.WithCancel(context.Background())
	replicasetLister, _ := a.clientset.AppsV1().ReplicaSets(namespace).List(ctx,v1.ListOptions{})
	for _, replicasets := range replicasetLister.Items {
		var n k8s.Replicaset
		n.Name = replicasets.Name
		n.NameSpace = replicasets.Namespace
		n.Labels = replicasets.Labels
		n.Ready = replicasets.Status.ReadyReplicas
		n.Current = replicasets.Status.AvailableReplicas
		n.Desired = replicasets.Spec.Replicas
		n.OwnerReferences = replicasets.OwnerReferences
		n.Desired =replicasets.Spec.Replicas
		t1 := time.Now()
		sub := t1.Sub(replicasets.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		replicaset = append(replicaset,&n)
	}
	return replicaset
}
func (a *OriginalK8s) GetNamespacedDeployment (namespace string) []*k8s.NamespacedDeployment  {
	var deploy []*k8s.NamespacedDeployment
	ctx, _ := context.WithCancel(context.Background())
	deployLister, _ := a.clientset.AppsV1().Deployments(namespace).List(ctx, v1.ListOptions{})
	for _, deployment := range deployLister.Items{
		var n k8s.NamespacedDeployment
		n.Name = deployment.Name
		n.Labels = deployment.Labels
		n.NameSpace = deployment.Namespace
		n.Available = deployment.Status.AvailableReplicas
		n.Current = deployment.Status.ReadyReplicas
		n.Desired = *deployment.Spec.Replicas
		n.StrategyType = deployment.Spec.Strategy.Type
		t1 := time.Now()
		sub := t1.Sub(deployment.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		deploy =append(deploy,&n)
	}
	return deploy
}
func (a *OriginalK8s) Getservice() []*k8s.Service {
	var service []*k8s.Service
	ctx, _ := context.WithCancel(context.Background())
	serviceLister, _ := a.clientset.CoreV1().Services("").List(ctx, v1.ListOptions{})
	for _, services := range serviceLister.Items {
		var n k8s.Service
		n.Name = services.Name
		n.NameSpace = services.Namespace
		n.Type = string(services.Spec.Type)
		n.Labels = services.Labels
		n.ClusterIp = services.Spec.ClusterIP
		n.ExternalIp = services.Spec.ExternalIPs
		n.Selector = services.Spec.Selector
		n.Ports = services.Spec.Ports
		t1 := time.Now()
		sub := t1.Sub(services.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		service = append(service,&n)

	}
	return service
}
func (a *OriginalK8s) GetSA() []*k8s.Serviceaccounts  {
	var serviceaccounts []*k8s.Serviceaccounts
	ctx, _ := context.WithCancel(context.Background())
	serviceaccountsLister, _ := a.clientset.CoreV1().ServiceAccounts("").List(ctx, v1.ListOptions{})
	for _, serviceaccount := range serviceaccountsLister.Items{
		var n k8s.Serviceaccounts
		n.Name = serviceaccount.Name
		n.NameSpace = serviceaccount.Namespace
		n.SecretsName = serviceaccount.Secrets
		t1 := time.Now()
		sub := t1.Sub(serviceaccount.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		serviceaccounts = append(serviceaccounts,&n)
	}
	return serviceaccounts

}
func (a *OriginalK8s) GetNamespacedSA(namespace string) []*k8s.Serviceaccounts  {
	var serviceaccounts []*k8s.Serviceaccounts
	ctx, _ := context.WithCancel(context.Background())
	serviceaccountsLister, _ := a.clientset.CoreV1().ServiceAccounts(namespace).List(ctx, v1.ListOptions{})
	for _, serviceaccount := range serviceaccountsLister.Items{
		var n k8s.Serviceaccounts
		n.Name = serviceaccount.Name
		n.NameSpace = serviceaccount.Namespace
		n.SecretsName = serviceaccount.Secrets
		t1 := time.Now()
		sub := t1.Sub(serviceaccount.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		serviceaccounts = append(serviceaccounts,&n)
	}
	return serviceaccounts
}
func (a *OriginalK8s) GetNamespacedService(namespace string) []*k8s.Service {
	var service []*k8s.Service
	ctx, _ := context.WithCancel(context.Background())
	serviceLister, _ := a.clientset.CoreV1().Services(namespace).List(ctx, v1.ListOptions{})
	for _, services := range serviceLister.Items {
		var n k8s.Service
		n.Name = services.Name
		n.NameSpace = services.Namespace
		n.Type = string(services.Spec.Type)
		n.Labels = services.Labels
		n.ClusterIp = services.Spec.ClusterIP
		n.ExternalIp = services.Spec.ExternalIPs
		n.Selector = services.Spec.Selector
		n.Ports = services.Spec.Ports
		t1 := time.Now()
		sub := t1.Sub(services.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		service = append(service,&n)

	}
	return service
}
func (a *OriginalK8s) GetPod() []*k8s.Pod {
	var pods []*k8s.Pod
	ctx, _ := context.WithCancel(context.Background())
	podList, _ := a.clientset.CoreV1().Pods("").List(ctx, v1.ListOptions{})
	for _, pod := range podList.Items {
		var n k8s.Pod
		n.Name = pod.Name
		n.NameSpaces = pod.Namespace
		n.Status = fmt.Sprintf("%s", pod.Status.Phase)

		for _, value := range pod.OwnerReferences {
			if value.Kind != "" {
				n.Controlled = value.Kind
				break
			}
		}
		for _, value := range pod.Status.ContainerStatuses {
			n.Restarts = value.RestartCount
			break
		}

		n.Qos = fmt.Sprintf("%s", pod.Status.QOSClass)
		t1 := time.Now()
		sub := t1.Sub(pod.CreationTimestamp.Time)
		if hours := sub.Hours(); hours > 0 {
			n.Age = fmt.Sprintf("%.0f", hours/24)
		} else {
			n.Age = "0"
		}
		pods = append(pods, &n)
	}

	return pods
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



