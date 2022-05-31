package main

import (
	"bytes"
	"time"

	"fmt"
	"strings"
	"text/tabwriter"

	k8s "github.com/dunefro/aubserver/k8s"
	slack "github.com/dunefro/aubserver/slack"
	v1 "k8s.io/api/core/v1"
)

func getFormattedPodStatus(table []k8s.K8sPod) string {
	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 0, 8, 4, ' ', tabwriter.TabIndent)
	fmt.Fprintln(writer, strings.Join([]string{"Pod", "Namespace", "Container", "Status", "Reason", "Message"}, "\t")+"\t")
	for _, p := range table {
		for _, c := range p.Containers {
			fmt.Fprintln(writer, strings.Join([]string{p.Name, p.Namespace, c.Name, c.Status, c.Reason, c.Message}, "\t")+"\t")
		}
	}
	writer.Flush()
	return buf.String()
}

func listFailedPods(pods []v1.Pod) []k8s.K8sPod {
	failedPods := make([]k8s.K8sPod, 0, len(pods))
	for _, pod := range pods {
		failedContainers := make([]k8s.PodContainer, 0, len(pod.Status.ContainerStatuses))
		for _, container := range pod.Status.ContainerStatuses {
			if container.State.Running == nil {
				if container.State.Waiting != nil {
					failedContainers = append(failedContainers, k8s.PodContainer{
						Name:    container.Name,
						Status:  "Waiting",
						Reason:  container.State.Waiting.Reason,
						Message: container.State.Waiting.Message,
					})
				}
			}
		}
		// if there is any failed container then append the pod in the failedPod list
		if len(failedContainers) != 0 {
			failedPods = append(failedPods, k8s.K8sPod{
				Name:       pod.Name,
				Namespace:  pod.Namespace,
				Containers: failedContainers,
			})
		}
	}
	return failedPods
}

func main() {
	for {
		pods, err := k8s.GetPods()
		if err != nil {
			panic(err.Error())
		}
		failedPods := listFailedPods(pods)
		// if there is any failed pod, only then send the message on slack
		if len(failedPods) != 0 {
			slackMessge := fmt.Sprintln(getFormattedPodStatus(failedPods))
			slack.SendSlackNotification(slackMessge)
		} else {
			fmt.Println("All well !!")
		}
		fmt.Println("Sleeping ...")
		time.Sleep(60 * time.Second)
	}
}
