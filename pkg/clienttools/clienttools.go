package clienttools

import (
	osappsv1 "github.com/openshift/api/apps/v1"
	projectv1 "github.com/openshift/api/project/v1"
	olmapiv1 "github.com/operator-framework/api/pkg/operators/v1"
	olmapiv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	olmapiv1alpha2 "github.com/operator-framework/api/pkg/operators/v1alpha2"
	olmapiv2 "github.com/operator-framework/api/pkg/operators/v2"
	olmcli "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	olmpkgsvr "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators/v1"
	configclient "github.com/openshift/client-go/config/clientset/versioned"
	"github.com/syndesisio/syndesis/install/operator/pkg/util"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	kclient "k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	oauthclient "github.com/openshift/client-go/oauth/clientset/versioned"
)

type ClientTools struct {
	runtimeClient *client.Client
	coreV1Client  corev1client.CoreV1Interface
	dynamicClient dynamic.Interface
	configClient  *configclient.configClient
	apiClient     kubernetes.Interface
	kClient		  kubernetes.Interface
	oauthClient   kubernetes.Interface
	restConfig    *rest.Config
	scheme        *runtime.Scheme
//	olmClient     olmcli.Interface
}

func (ck *ClientTools) RestConfig() (c *rest.Config) {
	if ck.restConfig == nil {
		config, err := config.GetConfig()
		util.ExitOnError(err)
		ck.restConfig = config
	}

	return ck.restConfig
}

func (ck *ClientTools) GetScheme() *runtime.Scheme {
	if ck.scheme == nil {
		ck.scheme = scheme.Scheme
		osappsv1.AddToScheme(ck.scheme)
		olmapiv1alpha2.SchemeBuilder.AddToScheme(ck.scheme)
		olmapiv1alpha1.SchemeBuilder.AddToScheme(ck.scheme)
		olmapiv1.SchemeBuilder.AddToScheme(ck.scheme)
		olmapiv2.AddToScheme(ck.scheme)
		olmpkgsvr.SchemeBuilder.AddToScheme(ck.scheme)
		projectv1.AddToScheme(ck.scheme)
	}

	return ck.scheme
}

func (ck *ClientTools) RuntimeClient() (c client.Client, err error) {
	if ck.runtimeClient == nil {

		s := ck.GetScheme()

		// Register
		options := client.Options{
			Scheme: s,
		}

		cl, err := client.New(ck.RestConfig(), options)
		if err != nil {
			return nil, err
		}
		ck.runtimeClient = &cl
	}

	return *ck.runtimeClient, nil
}

func (ck *ClientTools) SetRuntimeClient(c client.Client) {
	ck.runtimeClient = &c
}

func (ck *ClientTools) DynamicClient() (c dynamic.Interface, err error) {
	if ck.dynamicClient == nil {
		dyncl, err := dynamic.NewForConfig(ck.RestConfig())
		if err != nil {
			return nil, err
		}
		ck.dynamicClient = dyncl
	}
	return ck.dynamicClient, nil
}

func (ck *ClientTools) SetDynamicClient(d dynamic.Interface) {
	ck.dynamicClient = d
}

func (ck *ClientTools) ApiClient() (kubernetes.Interface, error) {
	if ck.apiClient == nil {
		apicl, err := kubernetes.NewForConfig(ck.RestConfig())
		if err != nil {
			return nil, err
		}
		ck.apiClient = apicl
	}
	return ck.apiClient, nil
}

func (ck *ClientTools) SetApiClient(a kubernetes.Interface) {
	ck.apiClient = a
}

func (ck *ClientTools) KClient() (kubernetes.Interface, error) {
	if ck.kClient == nil {
		client, err := kClient.NewForConfig(ck.RestConfig())
		if err != nil {
			return nil, err
		}
		ck.kClient = client
	}

	return ck.kClient, nil
}

func (ck *ClientTools) SetKClient(c kubernetes.Interface) {
	ck.kclient = c
}

func (ck *ClientTools) OAuthClient() (kubernetes.Interface, error) {
	if ck.oauthClient == nil {
		client, err := oauthclient.NewForConfig(ck.RestConfig())
		if err != nil {
			return nil, err
		}
		ck.oauthClient = client
	}

	return ck.oauthClient, nil
}

func (ck *ClientTools) SetOAuthClient(c kubernetes.Interface) {
	ck.oauthClient = c
}

func (ck *ClientTools) CoreV1Client() (corev1.CoreV1Interface, error) {
	if ck.coreV1Client == nil {
		client, err := corev1.NewForConfig(ck.RestConfig())
		if err != nil {
			return nil, err
		}
		ck.coreV1Client = client
	}

	return ck.coreV1Client, nil
}

func (ck *ClientTools) SetCoreV1Client(c corev1.CoreV1Interface) {
	ck.coreV1Client = c
}

/*func (ck *ClientTools) OlmClient() (olmcli.Interface, error) {
	if ck.olmClient == nil {
		client, err := olmcli.NewForConfig(ck.RestConfig())
		if err != nil {
			return nil, err
		}
		ck.olmClient = client
	}

	return ck.olmClient, nil
}

func (ck *ClientTools) SetOlmClient(c olmcli.Interface) {
	ck.olmClient = c
}*/
