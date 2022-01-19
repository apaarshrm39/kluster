package controller

import (
	"context"
	"fmt"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	klient "github.com/apaarshrm39/Kluster/pkg/client/clientset/versioned"
	kinformer "github.com/apaarshrm39/Kluster/pkg/client/informers/externalversions/apaarshrm.dev/v1alpha1"
	klister "github.com/apaarshrm39/Kluster/pkg/client/listers/apaarshrm.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// The struct for klient controller
type Controller struct {
	// clientset for custom resource Kluster
	clientset klient.Clientset
	// Way to figure out if cache has synced, Informer maintains a local cache
	hasSynced cache.InformerSynced
	// workqueue to store objects
	queue workqueue.RateLimitingInterface
	// we will need a lister
	klusterLister klister.KlusterLister
	// kubernetes clientset
	k8sclient kubernetes.Clientset
}

func New(klient klient.Clientset, klusterInformer kinformer.KlusterInformer, client kubernetes.Clientset) *Controller {
	k := &Controller{
		clientset:     klient,
		hasSynced:     klusterInformer.Informer().HasSynced,
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "kqueue"),
		klusterLister: klusterInformer.Lister(),
		k8sclient:     client,
	}

	// Add Eventhandler
	klusterInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    k.handleAdd,
		DeleteFunc: k.handleDelete,
		UpdateFunc: k.handleUpdate,
	})

	// REturn the controller struct back
	return k
}

func (k Controller) Run(ch <-chan struct{}) {
	fmt.Println("Starting Kluster Controller")

	if !cache.WaitForCacheSync(ch, k.hasSynced) {
		fmt.Println("Cache did not sync")
	}

	go wait.Until(k.worker, 1*time.Second, ch)

	<-ch
}

func (k Controller) worker() {
	for k.process() {
	}
}

func (k Controller) process() bool {
	item, shutdown := k.queue.Get()
	if shutdown {
		return false
	}

	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		fmt.Println(err)
		return false
	}

	ns, n, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		fmt.Println(err)
		return false
	}

	secret, err := k.klusterLister.Klusters(ns).Get(n)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(k.getSecret(secret.Spec.SecretToken))

	return true
}

func (k Controller) handleAdd(obj interface{}) {
	fmt.Println("Add event was executed")
	k.queue.Add(obj)
}

func (k Controller) handleDelete(obj interface{}) {
	fmt.Println("Delete event was executed")
	k.queue.Add(obj)
}

func (k Controller) handleUpdate(oldObj interface{}, newObj interface{}) {
	fmt.Println("Delete event was executed")
	k.queue.Add(newObj)
}

func (k Controller) getSecret(secret string) string {
	fmt.Println("Get SEcret called")
	n := strings.Split(secret, "/")
	fmt.Println("split secret", n)
	sec, err := k.k8sclient.CoreV1().Secrets(n[0]).Get(context.Background(), n[1], metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	//token, err := base64.StdEncoding.DecodeString(sec.Data["token"])
	//if err != nil {
	//	fmt.Println(err)
	//	}

	return string(sec.Data["token"])

}
