package main

import (
	"customize-k8s/common"
	"customize-k8s/customresources/controllers"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
)

func main() {
	clientSet, err := common.GetClientSet()
	common.CheckErrorAndFatal(err)

	// Informer
	ch := make(chan struct{})

	informerFactory := informers.NewSharedInformerFactory(clientSet, 10*time.Second)
	deployInformer := informerFactory.Apps().V1().Deployments()
	log.Println(deployInformer.Lister().Deployments("default").Get("default"))
	informerFactory.Start(ch)
	informerFactory.WaitForCacheSync(wait.NeverStop)

	controller := controllers.NewController(clientSet, deployInformer)
	controller.Run(ch)

}
