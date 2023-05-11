package common

import (
	"flag"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeconfig() string {
	homeDir, err := os.UserHomeDir()
	CheckErrorAndFatal(err)
	kubeconfig := flag.String("kubeconfig", homeDir+"/.kube/config", "local profile kubeconfig default location")
	return *kubeconfig
}

func GetClientSet() (*kubernetes.Clientset, error) {
	kubeconfig := GetKubeconfig()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if IsError(err) {
		config, err = rest.InClusterConfig()
		if IsError(err) {
			return nil, err
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)

	return clientSet, err
}

func CheckErrorAndFatal(err error) {
	if err != nil {
		log.Fatal("error : ", err)
	}
}
func IsError(err error) bool {
	result := false
	if err != nil {
		log.Println("error : ", err)
		result = true
	}
	return result
}
