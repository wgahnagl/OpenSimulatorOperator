/*
Copyright 2023.

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

package controllers

import (
	"context"

	routev1 "github.com/openshift/api/route/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	examplecomv1 "github.com/wgahnagl/OpenSimulatorOperator/api/v1"
)

// OpenSimulatorRouteReconciler reconciles a OpenSimulatorRoute object
type OpenSimulatorRouteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=example.com,resources=opensimulatorroutes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=example.com,resources=opensimulatorroutes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=example.com,resources=opensimulatorroutes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OpenSimulatorRoute object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *OpenSimulatorRouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	var OpenSimulatorRoute examplecomv1.OpenSimulatorRoute

	log.Log.Info("reconciling routes")
	route := &routev1.Route{}
	if err := r.Client.Get(ctx, req.NamespacedName, &OpenSimulatorRoute); err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}
	// fix this, create a real route custom resource
	// tell it to create an external ip
	route.Spec.To.Name = "OpenSimulatorRoute"

	if err := r.Client.Update(ctx, route); err != nil {
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *OpenSimulatorReconciler) getOpenShiftRouteExternalIP(ctx context.Context, namespace types.NamespacedName, routeName string) (string, error) {
	var route routev1.Route
	if err := r.Get(ctx, namespace, &route); err != nil {
		return "", err
	}

	existingRoutes := &routev1.Route{}
	err := r.Get(ctx, namespace, existingRoutes)
	if err != nil {
		// Route does not exist, create it
		// newRoute := createRouteFromCR(yourCR)
		// err := r.Create(ctx, newRoute)

		if err != nil {
			// Handle creation error
		}
	} else if err != nil {
		// Handle other errors
	}

	// Assuming the Route has an Ingress with an external IP
	if len(route.Status.Ingress) > 0 {
		// Retrieve the first Ingress IP as the external IP
		externalIP := route.Status.Ingress[0].Host
		return externalIP, nil
	}

	log.Log.Error(nil, "no external IP found")
	return "", nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpenSimulatorRouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplecomv1.OpenSimulatorRoute{}).
		Complete(r)
}
