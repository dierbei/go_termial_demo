package kubernetes_client

import (
	"log"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	once sync.Once
	c    *kubernetes.Clientset
)

func GetK8sCli() *kubernetes.Clientset {
	once.Do(func() {
		_config, err := clientcmd.BuildConfigFromFlags("", "config/config")
		_c, err := kubernetes.NewForConfig(_config)
		if err != nil {
			log.Fatalln(err)
		}
		c = _c
	})
	return c
}
