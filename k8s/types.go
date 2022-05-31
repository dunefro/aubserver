package k8s

type PodContainer struct {
	Name    string
	Status  string
	Reason  string
	Message string
}
type K8sPod struct {
	Name       string
	Namespace  string
	Containers []PodContainer
}
