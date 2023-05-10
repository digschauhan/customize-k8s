package lister

import (
	"context"
	"customize-k8s/common"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ListPods(ns string) {
	kubeconfig := common.GetKubeconfig()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	common.CheckErrorAndFatal(err)

	clientSet, err := kubernetes.NewForConfig(config)
	common.CheckErrorAndFatal(err)

	ctx := context.Background()

	fmt.Println("==================== Pods ")
	pods, err := clientSet.CoreV1().Pods(ns).List(ctx, v1.ListOptions{})
	common.CheckErrorAndFatal(err)

	for _, p := range pods.Items {
		fmt.Println(p.Namespace, " | ", p.Name)
	}

}

func ListDeployments(ns string) {
	kubeconfig := common.GetKubeconfig()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	common.CheckErrorAndFatal(err)

	clientSet, err := kubernetes.NewForConfig(config)
	common.CheckErrorAndFatal(err)

	ctx := context.Background()

	fmt.Println("==================== Deployments ")
	deploys, err := clientSet.AppsV1().Deployments(ns).List(ctx, v1.ListOptions{})
	common.CheckErrorAndFatal(err)

	for _, d := range deploys.Items {
		fmt.Println(d.Namespace, " | ", d.Name)
	}
}