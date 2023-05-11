package lister

import (
	"context"
	"customize-k8s/common"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func Listing(ns string) {
	clientSet, err := common.GetClientSet()
	common.CheckErrorAndFatal(err)

	ctx := context.Background()

	fmt.Println("==================== Pods ")
	pods, err := clientSet.CoreV1().Pods(ns).List(ctx, v1.ListOptions{})
	common.CheckErrorAndFatal(err)

	for _, p := range pods.Items {
		fmt.Println(p.Namespace, " | ", p.Name)
	}

	fmt.Println("==================== Deployments ")
	deploys, err := clientSet.AppsV1().Deployments(ns).List(ctx, v1.ListOptions{})
	common.CheckErrorAndFatal(err)

	for _, d := range deploys.Items {
		fmt.Println(d.Namespace, " | ", d.Name)
	}

	// Informer

	informerFactory := informers.NewSharedInformerFactory(clientSet, 30*time.Second)

	podInformer := informerFactory.Core().V1().Pods()

	eventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(newRes interface{}) {
			fmt.Println("New resource created.")
		},
		UpdateFunc: func(oldRes, newRes interface{}) {
			fmt.Println("Resource Updated.")
		},
		DeleteFunc: func(res interface{}) {
			fmt.Println("Resource deleted.")
		},
	}
	podInformer.Informer().AddEventHandler(eventHandler)

	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)
	pod, err := podInformer.Lister().Pods("default").Get("default")

	fmt.Println(pod)

}
