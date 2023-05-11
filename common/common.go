package common

import (
	"flag"
	"log"
	"os"
)

func GetKubeconfig() string {
	homeDir, err := os.UserHomeDir()
	CheckErrorAndFatal(err)
	kubeconfig := flag.String("kubeconfig", homeDir+"/.kube/config", "local profile kubeconfig default location")
	return *kubeconfig
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
