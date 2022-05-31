package main

import (
	"bytes"

	"fmt"
	"strings"
	"text/tabwriter"

	k8s "github.com/dunefro/aubserver/k8s"
	slack "github.com/dunefro/aubserver/slack"
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

func main() {
	Pods, err := k8s.GetFailedPods()
	if err != nil {
		panic(err.Error())
	}
	failedPods := make([]k8s.K8sPod, 0, len(Pods))
	for _, pod := range Pods {
		containerStatus := make([]k8s.PodContainer, 0, len(pod.Status.ContainerStatuses))
		for _, container := range pod.Status.ContainerStatuses {
			if container.State.Running == nil {
				// fmt.Println(index, container.Name, container.State.Running, container.State.Terminated, container.State.Waiting)
				if container.State.Waiting != nil {
					containerStatus = append(containerStatus, k8s.PodContainer{
						Name:    container.Name,
						Status:  "Waiting",
						Reason:  container.State.Waiting.Reason,
						Message: container.State.Waiting.Message,
					})
				}
			}
		}
		failedPods = append(failedPods, k8s.K8sPod{
			Name:       pod.Name,
			Namespace:  pod.Namespace,
			Containers: containerStatus,
		})
	}
	slackMessge := fmt.Sprintln(getFormattedPodStatus(failedPods))
	slack.SendSlackNotification(slackMessge)
}
