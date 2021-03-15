package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var clientset kubernetes.Interface

func createClientSet() kubernetes.Interface {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		// fmt.Println(home)
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	//fmt.Println(*kubeconfig)
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	//var clientset kubernetes.Interface
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
func getDeployments() {
	fmt.Println("Listing all deployment objects ...")
	//clientset := createClientSet()
	deploymentClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range list.Items {
		fmt.Printf("%s (%d replicas)\n", item.Name, *item.Spec.Replicas)
	}
}
func createDeployment() {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	//clientset := createClientSet()
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bookapi-deployment-go",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "bookapi",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "bookapi",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "bookapi",
							Image: "shohagrana64/bookapi:latest",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}
func deleteDeployment(args []string) {
	//var clientset kubernetes.Interface
	//clientset = CreateClient()

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deletePolicy := metav1.DeletePropagationForeground

	for _, deployName := range args {
		if err := deploymentsClient.Delete(context.TODO(), deployName, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
		fmt.Printf("Deleted deployment : %s\n", deployName)
	}
}

func main() {
	clientset = createClientSet()
	createDeployment()
	getDeployments()
	deleteDeployment([]string{"bookapi-deployment-go"})
	// Update Deployment
	prompt()
	fmt.Println("Updating deployment...")
	//    You have two options to Update() this Deployment:
	//
	//    1. Modify the "deployment" variable and call: Update(deployment).
	//       This works like the "kubectl replace" command and it overwrites/loses changes
	//       made by other clients between you Create() and Update() the object.
	//    2. Modify the "result" returned by Get() and retry Update(result) until
	//       you no longer get a conflict error. This way, you can preserve changes made
	//       by other clients between Create() and Update(). This is implemented below
	//			 using the retry utility package included with client-go. (RECOMMENDED)
	//
	// More Info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency

	//retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
	//	// Retrieve the latest version of Deployment before attempting update
	//	// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
	//	result, getErr := deploymentsClient.Get(context.TODO(), "bookapi-deployment-go", metav1.GetOptions{})
	//	if getErr != nil {
	//		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	//	}
	//
	//	result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
	//	result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
	//	_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
	//	return updateErr
	//})
	//if retryErr != nil {
	//	panic(fmt.Errorf("Update failed: %v", retryErr))
	//}
	//fmt.Println("Updated deployment...")
	//
	//// List Deployments
	//prompt()
	//fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	//list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//for _, d := range list.Items {
	//	fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	//}
	//
	//// Delete Deployment
	//prompt()
	//fmt.Println("Deleting deployment...")
	//deletePolicy := metav1.DeletePropagationForeground
	//if err := deploymentsClient.Delete(context.TODO(), "demo-deployment", metav1.DeleteOptions{
	//	PropagationPolicy: &deletePolicy,
	//}); err != nil {
	//	panic(err)
	//}
	//fmt.Println("Deleted deployment.")
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func int32Ptr(i int32) *int32 { return &i }
