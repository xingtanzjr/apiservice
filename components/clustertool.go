package components

import (
	"context"
	"reflect"

	"k8s.io/client-go/kubernetes"
	appsListers "k8s.io/client-go/listers/apps/v1"
	coreLister "k8s.io/client-go/listers/core/v1"
	rbacLister "k8s.io/client-go/listers/rbac/v1"
	asClient "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/clientset/versioned"
	asInformer "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/informers/externalversions/multitenancy/v1"
	asLister "multitenancy.metricsadvisor.ai/apiservice/generated/multitenancy/listers/multitenancy/v1"

	istioClient "istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apis "multitenancy.metricsadvisor.ai/apiservice/apis/multitenancy/v1"
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

func (c *ClusterTool) GetApiService(namespace string, name string) (*apis.ApiService, error) {
	return c.ApiServiceLister.ApiServices(namespace).Get(name)
}

func (c *ClusterTool) CreateApiService(apiService *apis.ApiService) error {
	_, err := c.ApiServiceClient.MultitenancyV1().ApiServices(apiService.Namespace).Create(context.TODO(), c.NewApiServiceForCreate(apiService), metav1.CreateOptions{})
	return err
}

func (c *ClusterTool) UpdateApiService(old, new *apis.ApiService) error {
	old.Spec = *new.Spec.DeepCopy()
	_, err := c.ApiServiceClient.MultitenancyV1().ApiServices(old.Namespace).Update(context.TODO(), old, metav1.UpdateOptions{})
	return err
}

func (c *ClusterTool) NewApiServiceForCreate(apiService *apis.ApiService) *apis.ApiService {
	return &apis.ApiService{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiService.ObjectMeta.Name,
			Namespace: apiService.ObjectMeta.Namespace,
		},
		Spec: *apiService.Spec.DeepCopy(),
	}
}

func (c *ClusterTool) IsApiServiceDifferent(as1, as2 *apis.ApiService) bool {
	if !reflect.DeepEqual(as1.Spec, as2.Spec) {
		return true
	}
	return false
}

func GetResourceUpdateTime(meta metav1.ObjectMeta) *metav1.Time {
	if meta.ManagedFields == nil || len(meta.ManagedFields) == 0 {
		return nil
	}

	var ret = meta.ManagedFields[0].Time
	for _, entity := range meta.ManagedFields {
		if ret == nil || ret.Before(entity.Time) {
			ret = entity.Time
		}
	}
	return ret
}
