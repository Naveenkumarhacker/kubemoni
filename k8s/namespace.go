package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getNamespace(clientset *kubernetes.Clientset) <-chan string {

	out := make(chan string)
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	go func(nsList *v1.NamespaceList) {
		for _, namespace := range nsList.Items {
			out <- namespace.Name
		}
		close(out)
	}(namespaceList)
	return out

}
