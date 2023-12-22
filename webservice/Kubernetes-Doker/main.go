package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	kubeconfig := filepath.Join(homeDir, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	event := &corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-event-" + time.Now().Format(time.RFC3339),
			Namespace: "default",
		},
		InvolvedObject: corev1.ObjectReference{
			Kind: "Node",
			Name: "minikube",
		},
		Reason:  "TestEvent",
		Message: "This is a test event",
		Source: corev1.EventSource{
			Component: "test-script",
		},
		FirstTimestamp: metav1.NewTime(time.Now()),
		LastTimestamp:  metav1.NewTime(time.Now()),
		Count:          1,
		Type:           "Normal",
	}

	_, err = clientset.CoreV1().Events("default").Create(context.TODO(), event, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Event created successfully")
}
