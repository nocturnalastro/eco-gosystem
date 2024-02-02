package find

import (
	"fmt"

	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"github.com/openshift-kni/eco-goinfra/pkg/pod"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LinuxPTPDaemonPod returns pod running infrastructure-operator.
func LinuxPTPDaemonPod(apiClient *clients.Settings) (*pod.Builder, error) {
	return getPodBuilder(apiClient, "app=linuxptp-daemon")
}

// getPodBuilder returns a podBuilder of a pod based on provided label.
func getPodBuilder(apiClient *clients.Settings, label string) (*pod.Builder, error) {
	if apiClient == nil {
		return nil, fmt.Errorf("apiClient is nil")
	}

	podList, err := pod.ListInAllNamespaces(apiClient, metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods on cluster: %w", err)
	}

	if len(podList) == 0 {
		return nil, fmt.Errorf("pod with label '%s' not currently running", label)
	}

	if len(podList) > 1 {
		return nil, fmt.Errorf("got unexpected pods when checking for pods with label '%s'", label)
	}

	return podList[0], nil
}
