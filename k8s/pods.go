package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPodsByNameSpace(in <-chan string, clientset *kubernetes.Clientset) <-chan *v1.PodList {
	out := make(chan *v1.PodList)
	go func(input <-chan string) {
		for ns := range input {
			pods, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			out <- pods
		}
		close(out)
	}(in)
	return out
}

func GetPodFromPodList(in <-chan *v1.PodList, clientset *kubernetes.Clientset) <-chan v1.Pod {
	out := make(chan v1.Pod)

	go func(inChan <-chan *v1.PodList) {
		for podList := range inChan {
			for _, x := range podList.Items {
				out <- x
			}
		}
		close(out)
	}(in)
	return out

}
