package main

import (
	"context"
	"log"

	k8s "go_terminal/internal/pkg/kubernetes_client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	cli := k8s.GetK8sCli()

	podList, err := cli.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, pod := range podList.Items {
		log.Println(pod.Name)
	}
}
