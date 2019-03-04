/*
Copyright 2017-2019 Kaloom Inc.

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

package controller

import (
	"context"
	"time"

	"github.com/kaloom/kubernetes-podagent/controller/cni"
	dcri "github.com/kaloom/kubernetes-podagent/controller/docker-runtime"

	"github.com/golang/glog"

	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/kubelet/dockershim/libdocker"
)

// Controller the controller object
type Controller struct {
	kubeClient *kubernetes.Clientset
	runtime    *dcri.DockerRuntime
	cniPlugin  *cni.NetworkPlugin
}

// Run starts a Pod resource controller
func (c *Controller) Run(ctx context.Context, nodeName string) error {
	var where string
	if nodeName == "" {
		where = "all nodes in the cluster"
	} else {
		where = "node: " + nodeName
	}
	glog.Infof("Pod's resource controller watching on %s", where)

	// Watch Pod objects
	_, err := c.watchPods(ctx, nodeName)
	if err != nil {
		glog.Errorf("Failed to register watch for Pod resource: %v", err)
		return err
	}

	<-ctx.Done()
	return ctx.Err()
}

// NewController instantiate a controller object
func NewController(kubeClient *kubernetes.Clientset, dockerEndpoint, cniBinPath, cniConfPath, cniVendor string) (*Controller, error) {
	runtimeRequestTimeout := 2 * time.Minute
	imagePullProgressDeadline := time.Minute
	dockerClient := libdocker.ConnectToDockerOrDie(dockerEndpoint,
		runtimeRequestTimeout,
		imagePullProgressDeadline)
	runTime, err := dcri.NewDockerRuntime(dockerClient)
	if err != nil {
		return nil, err
	}
	cniPlugin, err := cni.NewCNIPlugin(cniBinPath, cniConfPath, cniVendor)
	if err != nil {
		return nil, err
	}
	c := &Controller{
		kubeClient: kubeClient,
		runtime:    runTime,
		cniPlugin:  cniPlugin,
	}
	return c, nil
}