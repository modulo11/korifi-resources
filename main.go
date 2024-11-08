package main

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getResources(client *kubernetes.Clientset) ([]schema.GroupVersionResource, error) {
	resources := []schema.GroupVersionResource{}

	_, allResources, err := client.Discovery().ServerGroupsAndResources()
	if err != nil {
		return nil, err
	}

	for _, resourceList := range allResources {
		var group, version string
		if resourceList.GroupVersion == "v1" {
			group = ""
			version = "v1"
		} else {
			groupVersion := strings.Split(resourceList.GroupVersion, "/")
			group = groupVersion[0]
			version = groupVersion[1]
		}
		for _, resource := range resourceList.APIResources {
			if resource.Namespaced && slices.Contains(resource.Verbs, "list") {
				kind := resource.Name
				kindRes := schema.GroupVersionResource{Group: group, Version: version, Resource: kind}
				resources = append(resources, kindRes)
			}
		}
	}

	return resources, nil
}

func getNamespaces(ctx context.Context, client *kubernetes.Clientset) ([]string, error) {
	namespaces := []string{}

	allNamespaces, err := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, namespace := range allNamespaces.Items {
		if strings.HasPrefix(namespace.Name, "cf") {
			namespaces = append(namespaces, namespace.Name)
		}
	}

	return namespaces, nil
}

func printResources(ctx context.Context, client *kubernetes.Clientset, dynamicClient dynamic.Interface, resources []schema.GroupVersionResource) {
	// Retrieve all namespaces with cf prefix, includes spaces and orgs
	namespaces, err := getNamespaces(ctx, client)
	if err != nil {
		panic(err)
	}

	for _, namespace := range namespaces {
		fmt.Printf("Namespace: %s\n", namespace)
		for _, resource := range resources {
			// Retrieve all resources within the namespace
			result, err := dynamicClient.Resource(resource).Namespace(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				panic(err)
			}
			if len(result.Items) > 0 {
				group := resource.Group
				if group == "" {
					group = "core"
				}
				fmt.Printf("%s/%s/%s: %d\n", group, resource.Version, resource.Resource, len(result.Items))
			}
		}
		fmt.Println()
	}
}

func main() {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	kubeConfig.WarningHandler = rest.NoWarnings{}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	dynamicClient, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	resources, err := getResources(clientset)
	if err != nil {
		panic(err)
	}

	fmt.Println("Empty cluster ...")
	printResources(ctx, clientset, dynamicClient, resources)

	fmt.Println("Create organization ...")
	orgName := "test"
	spaceName := "test"
	_, err = exec.Command("cf", "create-org", orgName).Output()
	if err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
	printResources(ctx, clientset, dynamicClient, resources)

	fmt.Println("Create space ...")
	_, err = exec.Command("cf", "create-space", "-o", orgName, spaceName).Output()
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	printResources(ctx, clientset, dynamicClient, resources)

	fmt.Println("Push app 1 ...")
	_, err = exec.Command("cf", "target", "-o", orgName, "-s", spaceName).Output()
	if err != nil {
		panic(err)
	}
	_, err = exec.Command("cf", "push", "--path", "./sample", "hello-1").Output()
	if err != nil {
		panic(err)
	}

	printResources(ctx, clientset, dynamicClient, resources)

	fmt.Println("Push app 2 ...")
	_, err = exec.Command("cf", "target", "-o", orgName, "-s", spaceName).Output()
	if err != nil {
		panic(err)
	}
	_, err = exec.Command("cf", "push", "--path", "./sample", "hello-2").Output()
	if err != nil {
		panic(err)
	}

	printResources(ctx, clientset, dynamicClient, resources)

	fmt.Println("Cleanup ...")

	time.Sleep(5 * time.Second)
	_, err = exec.Command("cf", "delete-org", orgName, "-f").Output()
	if err != nil {
		panic(err)
	}
}
