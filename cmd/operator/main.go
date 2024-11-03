package main

import (
	"context"
	"flag"
	"kek/generated/clientset"
	informers "kek/generated/informers/externalversions"
	"os"
	"syscall"
	"time"

	"os/signal"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeconfig string
)

func setupSignalHandler() context.Context {
	var c = make(chan os.Signal, 2)
	var ctx, cancel = context.WithCancel(context.Background())
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()
	return ctx
}

func main() {
	klog.InitFlags(nil)
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.Parse()

	var err error
	// set up signals so we handle the shutdown signal gracefully
	var ctx = setupSignalHandler()
	var logger = klog.FromContext(ctx)

	var cfg *rest.Config
	if cfg, err = clientcmd.BuildConfigFromFlags(masterURL, kubeconfig); err != nil {
		logger.Error(err, "Error building kubeconfig")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	var kubeClient *kubernetes.Clientset
	if kubeClient, err = kubernetes.NewForConfig(cfg); err != nil {
		logger.Error(err, "Error building kubernetes clientset")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	var exampleClient *clientset.Clientset
	if exampleClient, err = clientset.NewForConfig(cfg); err != nil {
		logger.Error(err, "Error building kubernetes clientset")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	var kubeInformerFactory = kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	var exampleInformerFactory = informers.NewSharedInformerFactory(exampleClient, time.Second*30)
	var msets = exampleInformerFactory.Sp().V1alpha1().ModulePodSets()
	var pods = kubeInformerFactory.Core().V1().Pods()
	pods.Informer().AddEventHandler(&podController{})
	msets.Informer().AddEventHandler(&podController{})
	kubeInformerFactory.Start(nil)
	exampleInformerFactory.Start(nil)
}

type podController struct {
}

func (p *podController) OnAdd(obj any, isInInitialList bool) {
	var pod = obj.(*corev1.Pod)
}
func (p *podController) OnUpdate(oldObj, newObj any) {
	var opod = oldObj.(*corev1.Pod)
	var npod = newObj.(*corev1.Pod)

}
func (p *podController) OnDelete(obj any) {
	var pod = obj.(*corev1.Pod)
}
