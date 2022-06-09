package objects

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateNamespace(clientset kubernetes.Clientset, name string) {
	nsSpec := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "namespace-" + name,
		},
	}
	namespace, err := clientset.CoreV1().Namespaces().Create(context.Background(), nsSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Namespace successfully created:", namespace.GetName())
}

func DeleteNamespace(clientset kubernetes.Clientset, name string) {
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.CoreV1().Namespaces().Delete(context.TODO(), "namespace-"+name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Namespace deleted:", "namespace-"+name)
}

func CreateConfigmap(clientset kubernetes.Clientset, name string, namespace string) {
	configmapSpec := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cm-" + name,
			Namespace: namespace,
		},
		Data: map[string]string{"color.good": "purple", "color.bad": "yellow"},
	}
	configMap, err := clientset.CoreV1().ConfigMaps("default").Create(context.Background(), configmapSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ConfigMap created:", configMap.GetName())
}

func DeleteConfigmap(clientset kubernetes.Clientset, name string, namespace string) {
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), "cm-"+name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ConfigMap deleted:", "cm-"+name)
}

func CreateSecret(clientset kubernetes.Clientset, name string, namespace string) {
	secretSpec := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "secret-" + name,
			Namespace: namespace,
		},
		Data: map[string][]byte{"test": []byte("test")},
	}

	secret, err := clientset.CoreV1().Secrets("default").Create(context.Background(), secretSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Secret created", secret.GetName())
}

func DeleteSecret(clientset kubernetes.Clientset, name string, namespace string) {
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.CoreV1().Secrets(namespace).Delete(context.TODO(), "secret-"+name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Secret deleted:", "secret-"+name)
}

func CreatePod(clientset kubernetes.Clientset, name string, namespace string, image string) {
	podSpec := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod-" + name,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{Name: "busybox", Image: "busybox:latest", ImagePullPolicy: "IfNotPresent", Command: []string{"sleep", "100000"}},
			},
		},
	}
	pod, err := clientset.CoreV1().Pods("default").Create(context.Background(), podSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pod created:", pod.GetName())
}

func DeletePod(clientset kubernetes.Clientset, name string, namespace string) {
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), "pod-"+name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pod deleted:", "pod-"+name)
}

func CreateCronjob(clientset kubernetes.Clientset, name string, namespace string, schedule string) {
	cronjobSpec := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cronjob-" + name,
			Namespace: namespace,
		},
		Spec: batchv1.CronJobSpec{
			Schedule: schedule,
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{Name: "busybox", Image: "busybox", ImagePullPolicy: "IfNotPresent", Command: []string{"sleep", "60"}},
							},
							RestartPolicy: "OnFailure",
						},
					},
				},
			},
		},
	}
	cronjob, err := clientset.BatchV1().CronJobs("default").Create(context.Background(), cronjobSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Cronjob created :", cronjob.GetName())
}

func DeleteCronjob(clientset kubernetes.Clientset, name string, namespace string) {
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.BatchV1().CronJobs(namespace).Delete(context.TODO(), "cronjob-"+name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Cronjob deleted:", "cronjob-"+name)
}

func CreateDeployment(clientset kubernetes.Clientset, name string, namespace string, image string, label map[string]string) {
	deploymentSpec := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "deployment-" + name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: label,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: label,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "web",
							Image:           image,
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	deployment, err := clientset.AppsV1().Deployments("default").Create(context.Background(), deploymentSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("deployment sucessfully created:", deployment.GetName())
}

func DeleteDeployment(clientset kubernetes.Clientset, name string, namespace string) {
	deletePolicy := metav1.DeletePropagationForeground
	if err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), "deployment-"+name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Deployment deleted:", "deployment-"+name)
}

func int32Ptr(i int32) *int32 { return &i }
