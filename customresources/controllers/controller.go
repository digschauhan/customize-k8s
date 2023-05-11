package controllers

import (
	"customize-k8s/common"
	"fmt"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	appsInformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	appsListers "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	ClientSet         kubernetes.Interface
	DeployLister      appsListers.DeploymentLister
	DeployCacheSynced cache.InformerSynced
	Queue             workqueue.RateLimitingInterface
}

func NewController(cs kubernetes.Interface, inf appsInformers.DeploymentInformer) *Controller {
	c := &Controller{
		ClientSet:         cs,
		DeployLister:      inf.Lister(),
		DeployCacheSynced: inf.Informer().HasSynced,
		Queue:             workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "monitor"),
	}

	inf.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.handleAdd,
		UpdateFunc: c.handleUpdate,
		DeleteFunc: c.handleDelete,
	})
	c.DeployLister = inf.Lister()

	return c
}

func (c *Controller) handleAdd(res interface{}) {
	fmt.Println("New resource created.")
	c.Queue.Add(res)
}

func (c *Controller) handleUpdate(oldRes, newRes interface{}) {
	fmt.Println("Resource Updated.")
}

func (c *Controller) handleDelete(res interface{}) {
	fmt.Println("Resource deleted.")
}

func (c *Controller) Run(ch <-chan struct{}) {
	log.Println("Starting controller ...")
	if !cache.WaitForCacheSync(ch, c.DeployCacheSynced) {
		log.Println("Error while waiting for cache to sync")
	}
	log.Println("Wait done !")
	wait.Until(c.worker, 5*time.Second, ch)
}

func (c *Controller) worker() {
	log.Println("Working on it .....")

	for c.processResource() {

	}
}

func (c *Controller) processResource() bool {
	log.Println("Processing Resource .....")
	result := true
	resouce, shutdown := c.Queue.Get()
	if shutdown {
		return false
	}
	key, err := cache.MetaNamespaceKeyFunc(resouce)
	if common.IsError(err) {
		return false
	}
	log.Println("key : ", key)

	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if common.IsError(err) {
		return false
	}

	log.Println(ns, " | ", name)

	return result
}

func (c *Controller) syncDeploy(ns, name string) error {
	var err error
	log.Println("Sync done and all required resources created.")
	return err
}
