package k8s

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func Run() {
	// k8s configuration path setup
	kubeconfig := flag.String("kubeconfig", "/home/naveenkumar/.kube/config", "absolute path to the kubeconfig file")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the client set
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	nsChan := getNamespace(clientset)
	podListChan := GetPodsByNameSpace(nsChan, clientset)
	podChan := GetPodFromPodList(podListChan, clientset)

	for pod := range podChan {
		cpuRequest := pod.Spec.Containers[0].Resources.Requests.Cpu()
		cpuLimit := pod.Spec.Containers[0].Resources.Limits.Cpu()
		memRequest := pod.Spec.Containers[0].Resources.Requests.Memory()
		memLimit := pod.Spec.Containers[0].Resources.Limits.Memory()
		age := int(time.Since(pod.CreationTimestamp.Time).Hours() / 24)

		fmt.Printf("name : %s \nstatus:%s \nnamespace: %s \nCPU Request: %s \nCPU Limit: %s \nMEM Request: %s \nMEM Limit: %s \nage:%d \n\n", pod.Name, pod.Status.Phase, pod.Namespace, cpuRequest, cpuLimit, memRequest, memLimit, age)
		//fmt.Printf("%+v", pod.Spec.Containers[0].Resources)
	}

}
