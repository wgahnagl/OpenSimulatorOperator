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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	examplecomv1 "github.com/wgahnagl/OpenSimulatorOperator/api/v1"
)

// OpenSimulatorReconciler reconciles a OpenSimulator object
type OpenSimulatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=example.com,resources=opensimulators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=example.com,resources=opensimulators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=example.com,resources=opensimulators/finalizers,verbs=update

func (r *OpenSimulatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("reconciling OpenSimulator custom resource")

	var OpenSimulator examplecomv1.OpenSimulator
	if err := r.Get(ctx, req.NamespacedName, &OpenSimulator); err != nil {
		log.Error(err, "unable to fetch pods")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var podList corev1.PodList
	var openSimPodStarted bool
	if err := r.List(ctx, &podList); err != nil {
		log.Error(err, "unable to list pods :()")
	} else {
		for _, item := range podList.Items {
			if item.GetName() == OpenSimulator.Spec.Name {
				log.Info("pod linked to an OpenSimulator custom resource found", "name", item.GetName())
				openSimPodStarted = true
			}
		}
	}

	OpenSimulator.Status.Started = openSimPodStarted
	if err := r.Status().Update(ctx, &OpenSimulator); err != nil {
		log.Error(err, "unable to update OpenSimulator's started status", "status", openSimPodStarted)
		return ctrl.Result{}, err
	}
	log.Info("OpenSimulator's status updated", "status", openSimPodStarted)
	log.Info("OpenSimulator custom resource reconciled")

	if !OpenSimulator.Status.Configured {
		// Load OpenSimulator config files
		log.Info("loading OpenSimulator config files")

		externalIP, err := r.getOpenShiftRouteExternalIP(ctx, OpenSimulator.Namespace, "opensimulator")
		if err != nil {
			log.Error(err, "unable to retrieve external IP of the OpenShift Route")
			log.Info("external IP: ", externalIP)
			return ctrl.Result{}, err
		}
		OpenSimulator.Status.Configured = true
	}

	return ctrl.Result{}, nil
}

func (r *OpenSimulatorReconciler) getOpenShiftRouteExternalIP(ctx context.Context, namespace, routeName string) (string, error) {
	var route routev1.Route
	if err := r.Get(ctx, types.NamespacedName{Name: routeName, Namespace: namespace}, &route); err != nil {
		return "", err
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
func (r *OpenSimulatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplecomv1.OpenSimulator{}).
		Watches(
			&source.Kind{Type: &corev1.Pod{}},
			handler.EnqueueRequestsFromMapFunc(r.mapPodsReqToOpenSimulatorReq),
		).
		Complete(r)
}

func (r *OpenSimulatorReconciler) mapPodsReqToOpenSimulatorReq(obj client.Object) []reconcile.Request {
	ctx := context.Background()
	log := log.FromContext(ctx)

	req := []reconcile.Request{}
	var list examplecomv1.OpenSimulatorList
	if err := r.Client.List(context.TODO(), &list); err != nil {
		log.Error(err, "unable to list OpenSimulator custom resources")
	} else {
		for _, item := range list.Items {
			if item.Spec.Name == obj.GetName() {
				req = append(req, reconcile.Request{
					NamespacedName: types.NamespacedName{Name: item.Name, Namespace: item.Namespace},
				})
				log.Info("pod linked to an OpenSimulator custom resource issued an event", "name", obj.GetName())
			}
		}
	}
	return req
}
