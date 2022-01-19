package main

import (
	"flag"
	"path/filepath"
	"time"

	klient "github.com/apaarshrm39/Kluster/pkg/client/clientset/versioned"
	kinformer "github.com/apaarshrm39/Kluster/pkg/client/informers/externalversions"
	kontroller "github.com/apaarshrm39/Kluster/pkg/controller"
	"k8s.io/client-go/kubernetes"
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
	// create kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// create kluster clientset
	klientset, err := klient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	infofac := kinformer.NewSharedInformerFactory(klientset, 10*time.Minute)

	// CREATE kluster informer
	klusterInformer := infofac.Apaarshrm().V1alpha1().Klusters()

	k := kontroller.New(*klientset, klusterInformer, *clientset)
	ch := make(chan struct{})

	//STart infofac
	infofac.Start(ch)
	k.Run(ch)
}
