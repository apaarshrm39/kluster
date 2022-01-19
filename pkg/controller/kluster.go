package controller

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	klient "github.com/apaarshrm39/Kluster/pkg/client/clientset/versioned"
	kinformer "github.com/apaarshrm39/Kluster/pkg/client/informers/externalversions/apaarshrm.dev/v1alpha1"
	klister "github.com/apaarshrm39/Kluster/pkg/client/listers/apaarshrm.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/util/wait"
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
}

func New(klient klient.Clientset, klusterInformer kinformer.KlusterInformer) *Controller {
	k := &Controller{
		clientset:     klient,
		hasSynced:     klusterInformer.Informer().HasSynced,
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "kqueue"),
		klusterLister: klusterInformer.Lister(),
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

	kluster, err := k.clientset.ApaarshrmV1alpha1().Klusters(ns).Get(context.Background(), n, metav1.GetOptions{})

	spec := kluster.Spec

	fmt.Println(spec)

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
