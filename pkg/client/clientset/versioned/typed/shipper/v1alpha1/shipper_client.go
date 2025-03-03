/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1alpha1 "github.com/bookingcom/shipper/pkg/apis/shipper/v1alpha1"
	"github.com/bookingcom/shipper/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ShipperV1alpha1Interface interface {
	RESTClient() rest.Interface
	ApplicationsGetter
	CapacityTargetsGetter
	ClustersGetter
	InstallationTargetsGetter
	ReleasesGetter
	TrafficTargetsGetter
}

// ShipperV1alpha1Client is used to interact with features provided by the shipper.booking.com group.
type ShipperV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ShipperV1alpha1Client) Applications(namespace string) ApplicationInterface {
	return newApplications(c, namespace)
}

func (c *ShipperV1alpha1Client) CapacityTargets(namespace string) CapacityTargetInterface {
	return newCapacityTargets(c, namespace)
}

func (c *ShipperV1alpha1Client) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *ShipperV1alpha1Client) InstallationTargets(namespace string) InstallationTargetInterface {
	return newInstallationTargets(c, namespace)
}

func (c *ShipperV1alpha1Client) Releases(namespace string) ReleaseInterface {
	return newReleases(c, namespace)
}

func (c *ShipperV1alpha1Client) TrafficTargets(namespace string) TrafficTargetInterface {
	return newTrafficTargets(c, namespace)
}

// NewForConfig creates a new ShipperV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ShipperV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ShipperV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ShipperV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ShipperV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ShipperV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ShipperV1alpha1Client {
	return &ShipperV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ShipperV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
