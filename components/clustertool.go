package components

import (
	"k8s.io/client-go/kubernetes"
	appsListers "k8s.io/client-go/listers/apps/v1"
	coreLister "k8s.io/client-go/listers/core/v1"
	rbacLister "k8s.io/client-go/listers/rbac/v1"
	asClient "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/clientset/versioned"
	asInformer "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/informers/externalversions/multitenancy/v1"
	asLister "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/listers/multitenancy/v1"
)

type ClusterTool struct {
	clusterId string

	// Scalfold for ApiService
	apiServiceLister   asLister.ApiServiceLister
	apiServiceClient   asClient.Clientset
	apiServiceInformer asInformer.ApiServiceInformer

	// Scalfold for basic kubernetes resources
	deploymentLister         appsListers.DeploymentLister
	serviceLister            coreLister.ServiceLister
	serviceAccountLister     coreLister.ServiceAccountLister
	roleLister               rbacLister.RoleLister
	roleBindingLister        rbacLister.RoleBindingLister
	clusterRoleLister        rbacLister.ClusterRoleLister
	clusterRoleBindingLister rbacLister.ClusterRoleBindingLister
	kubeClient               kubernetes.Clientset
}
