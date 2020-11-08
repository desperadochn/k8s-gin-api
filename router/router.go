package router

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/controllers"
)

func InitRouter() {
	router := gin.Default()
	k8sWsGr := router.Group("/k8s-ws",CanUserToContainer)
	{
		ParseK8sWebshellRouter(k8sWsGr)
	}

	api := router.Group("api")
	{
		user := api.Group("user")
		{
			user.POST("login", controllers.Login)
			user.GET("info", controllers.Info)
		}
		//prometheus := api.Group("prometheus")
		//{
		//	node:= prometheus.Group("node")
		//	{
		//		node.GET("all")
		//	}
		//	podMetrics := prometheus.Group("pod")
		//	{
		//		podMetrics.Group("all namespace")
		//		{
		//			podMetrics.GET("cpu")
		//			podMetrics.GET("memory")
		//		}
		//	}
		//}
		k8s := api.Group("k8s")
		{
			node := k8s.Group("node")
			{
				node.GET("list", controllers.NodeList)
			}
			cronjob := k8s.Group("cronjob")
			{
				cronjob.GET("list",controllers.CronJobLister)
			}
			role := k8s.Group("role")
			{
				role.GET("list",controllers.GetRoles)
			}
			clusterrole := k8s.Group("clusterrole")
			{
				clusterrole.GET("list",controllers.GetClusterRoles)
			}
			clusterrolebinding := k8s.Group("clusterrolebinding")
			{
				clusterrolebinding.GET("list",controllers.GetClusterrolebinding)
			}
			rolebinding := k8s.Group("rolebinding")
			{
				rolebinding.GET("list",controllers.GetRolebindings)
			}
			pv := k8s.Group("pv")
			{
				pv.GET("list",controllers.PvList)
			}
			pvc := k8s.Group("pvc")
			{
				pvc.GET("list",controllers.PvcList)
			}
			storageclasses := k8s.Group("storageclasses")
			{
				storageclasses.GET("list", controllers.StorageclassesList)
			}
			namespace := k8s.Group("namespace")
			{
				namespace.GET("list", controllers.NameSpacesList)
			}
			pod := k8s.Group("pod")
			{
				pod.GET("list", controllers.PodList)
			}
			replicaset := k8s.Group("replicaset")
			{
				replicaset.GET("list",controllers.ReplicasetLister)
			}
			cm := k8s.Group("configmap")
			{
				cm.GET("list",controllers.ConfigMapList)
			}
			deployment := k8s.Group("deployment")
			{
				deployment.GET("list", controllers.DeploymentLister)
			}
			daemonsets := k8s.Group("daemonsets")
			{
				daemonsets.GET("list",controllers.DaemonsetsLister)
			}
			statefulsets := k8s.Group("statefulsets")
			{
				statefulsets.GET("list",controllers.StatefulsetsList)
			}
			secret := k8s.Group("secret")
			{
				secret.GET("list",controllers.SecretList)
			}
			service := k8s.Group("service")
			{
				service.GET("list",controllers.ServiceList)
			}
			endpoints := k8s.Group("endpoints")
			{
				endpoints.GET("list",controllers.EndpointsLister)
			}
			events := k8s.Group("envents")
			{
				events.GET("list",controllers.EvensLister)
			}
			serviceaccounts := k8s.Group("serviceaccounts")
			{
				serviceaccounts.GET("list",controllers.GetServiceaccounts)
			}
			//serviceaccounts := k8s.Group("serviceaccounts")
			//{
			//	serviceaccounts.GET("list",controllers.Serviceaccounts)
			//}
			container := k8s.Group("container")
			{
				container.GET("deployment/list",controllers.ContainertLister)
			}
			job := k8s.Group("job")
			{
				job.GET("list",controllers.JobLister)
			}

			listNamespacedReouse := k8s.Group("listNamespacedReouse")
			{
				listNamespacedReouse.POST("pods",controllers.NamespacedPodLister)
				listNamespacedReouse.POST("deployments",controllers.NamespacedDeploymentLister)
				listNamespacedReouse.POST("config",controllers.NamespacedConfigmap)
				listNamespacedReouse.POST("daemonsets",controllers.NamespacedDaemonsetsLister)
				listNamespacedReouse.POST("statefulsets",controllers.NamespacedStatefulsetsList)
				listNamespacedReouse.POST("secret",controllers.NamespacedSecret)
				listNamespacedReouse.POST("service",controllers.NamespacedService)
				listNamespacedReouse.POST("endpoints",controllers.NamespacedEndpointsLister)
				listNamespacedReouse.POST("events",controllers.NamespacedEventsLister)
				listNamespacedReouse.POST("serviceaccounts",controllers.GetNamespacedServiceaccounts)
				listNamespacedReouse.POST("job",controllers.NamespacedJobLister)
				listNamespacedReouse.POST("pvc",controllers.NamespacedPvcList)
				listNamespacedReouse.POST("cronjob",controllers.NamespacedCronJobLister)
				listNamespacedReouse.POST("role",controllers.GetNamespacedRoles)
				listNamespacedReouse.POST("rolebinding",controllers.GetNamespacedRolebindings)
				listNamespacedReouse.POST("replicaset",controllers.NamespacedReplicasetLister)
				listNamespacedReouse.POST("deployment/container",controllers.NamespaceLabelContainertLister)
			}

			info := k8s.Group("info")
			{
				info.GET("list",controllers.K8sInfo)
			}

		}
	}
	router.Run(":8088")
}


