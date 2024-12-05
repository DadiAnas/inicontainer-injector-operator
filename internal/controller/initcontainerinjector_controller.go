/*
Copyright 2024.

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
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// InitContainerInjectorReconciler reconciles a InitContainerInjector object
type InitContainerInjectorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;update;patch

func (r *InitContainerInjectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Get the Deployment
	var deployment appsv1.Deployment
	if err := r.Get(ctx, req.NamespacedName, &deployment); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Check if deployment has required annotation
	annotations := deployment.GetAnnotations()
	if args, exists := annotations["initcontainer_injector_args"]; exists {
		modified := false

		// Get values from annotations with defaults
		registry := getDefaultIfEmpty(annotations["initcontainer_injector_registry"], "docker.io")
		image := getDefaultIfEmpty(annotations["initcontainer_injector_image"], "default")
		commandStr := getDefaultIfEmpty(annotations["initcontainer_injector_command"], "/bin/sh,-c,echo")

		// Split command and args
		command := strings.Split(commandStr, ",")
		argsList := strings.Split(args, ",")

		// Create init container
		initContainer := corev1.Container{
			Name:    "injected-init",
			Image:   fmt.Sprintf("%s/%s:latest", registry, image),
			Command: command,
			Args:    argsList,
		}

		// Check if init container already exists
		exists := false
		for _, container := range deployment.Spec.Template.Spec.InitContainers {
			if container.Name == initContainer.Name {
				exists = true
				break
			}
		}

		if !exists {
			deployment.Spec.Template.Spec.InitContainers = append(
				deployment.Spec.Template.Spec.InitContainers,
				initContainer,
			)
			modified = true
		}

		if modified {
			// Create a deep copy of the deployment to modify
			deploymentCopy := deployment.DeepCopy()

			// Update only the spec
			if err := r.Update(ctx, deploymentCopy); err != nil {
				log.Error(err, "Failed to update deployment", "deployment", deployment.Name)
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

func getDefaultIfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

// SetupWithManager sets up the controller with the Manager.
func (r *InitContainerInjectorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Deployment{}).
		Complete(r)
}
