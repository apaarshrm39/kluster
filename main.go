package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	klient "github.com/apaarshrm39/Kluster/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "Absoluture Path to the Kubeconfig")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	klientset, err := klient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	kluster, err := klientset.ApaarshrmV1alpha1().Klusters("default").Get(context.TODO(), "kluster-1", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(kluster.Kind)
}
