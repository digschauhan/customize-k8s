package main

import (
	"customize-k8s/common"
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {

	var res string

	flag.StringVar(&res, "res", "Pods", "Provide resource name i.e. Pods")
	flag.Parse()

	configFlag := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	matchVersionFlags := cmdutil.NewMatchVersionFlags(configFlag)

	mapper, err := cmdutil.NewFactory(matchVersionFlags).ToRESTMapper()
	common.CheckErrorAndFatal(err)

	gvr, err := mapper.ResourceFor(schema.GroupVersionResource{Resource: res})
	common.CheckErrorAndFatal(err)

	fmt.Println(gvr)
}
