package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getk8sClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, err
}

func printPodStatus(table [][]string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.TabIndent|tabwriter.Debug)
	for _, line := range table {
		fmt.Fprintln(writer, strings.Join(line, "\t")+"\t")
	}
	writer.Flush()
}

func main() {
	k8sClient, err := getk8sClient()
	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := k8sClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Println(err.Error())
		}
		var k8sPod = make([][]string, len(pods.Items))
		for index, pod := range pods.Items {
			k8sPod[index] = []string{pod.Name, string(pod.Status.Phase)}
		}
		printPodStatus(k8sPod)
		log.Printf("There are %d pods in the cluster\n", len(pods.Items))
		log.Println("Sleeping ...")
		time.Sleep(20 * time.Second)
	}
}
