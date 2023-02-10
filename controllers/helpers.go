package controllers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	fullmetalcomv1 "github.com/tapojit047/CRD-Controller-kubebuilder/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func newDeployment(alchemist *fullmetalcomv1.Alchemist) *appsv1.Deployment {
	labels := map[string]string{
		"app":        "main",
		"controller": alchemist.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      alchemist.Spec.DeploymentName,
			Namespace: alchemist.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(alchemist, fullmetalcomv1.GroupVersion.WithKind("Alchemist")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: alchemist.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "mustang",
							Image: alchemist.Spec.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: *alchemist.Spec.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}
}

func newService(alchemist *fullmetalcomv1.Alchemist) *corev1.Service {
	labels := map[string]string{
		"app":        "main",
		"controller": alchemist.Name,
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      alchemist.Spec.DeploymentName + "-service",
			Namespace: alchemist.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(alchemist, fullmetalcomv1.GroupVersion.WithKind("Alchemist")),
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Type:     "NodePort",
			Ports: []corev1.ServicePort{
				{
					NodePort: int32(30047),
					Port:     *alchemist.Spec.ServicePort,
					TargetPort: intstr.IntOrString{
						IntVal: *alchemist.Spec.TargetPort,
					},
				},
			},
		},
	}
}
