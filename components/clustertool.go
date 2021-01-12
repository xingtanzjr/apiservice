package components

import (
	"k8s.io/client-go/kubernetes"
	appsListers "k8s.io/client-go/listers/apps/v1"
	coreLister "k8s.io/client-go/listers/core/v1"
	rbacLister "k8s.io/client-go/listers/rbac/v1"
	asClient "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/clientset/versioned"
	asInformer "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/informers/externalversions/multitenancy/v1"
	asLister "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/listers/multitenancy/v1"

	istioClient "istio.io/client-go/pkg/clientset/versioned"
)

type ClusterTool struct {
	ClusterId string

	// Scalfold for ApiService
	ApiServiceLister   asLister.ApiServiceLister
	ApiServiceClient   *asClient.Clientset
	ApiServiceInformer asInformer.ApiServiceInformer

	// Scalfold for basic kubernetes resources
	DeploymentLister         appsListers.DeploymentLister
	ServiceLister            coreLister.ServiceLister
	ServiceAccountLister     coreLister.ServiceAccountLister
	RoleLister               rbacLister.RoleLister
	RoleBindingLister        rbacLister.RoleBindingLister
	ClusterRoleLister        rbacLister.ClusterRoleLister
	ClusterRoleBindingLister rbacLister.ClusterRoleBindingLister
	KubeClient               *kubernetes.Clientset

	// Istio scalfold
	IstioClient *istioClient.Clientset
}
