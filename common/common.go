package common

import (
	"flag"
	"log"
)

func GetKubeconfig() string {
	kubeconfig := flag.String("kubeconfig", "/Users/djaychauhan/.kube/config", "local profile kubeconfig default location")
	return *kubeconfig
}

func CheckErrorAndFatal(err error) {
	if err != nil {
		log.Fatal("error : ", err)
	}
}
