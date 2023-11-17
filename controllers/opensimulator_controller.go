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

	route "github.com/openshift/api/route/v1"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

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
	var OpenSimulator examplecomv1.OpenSimulator

	log.Info("reconciling OpenSimulator custom resource")

	if err := r.Get(ctx, req.NamespacedName, &OpenSimulator); err != nil {
		log.Error(err, "unable to fetch pods")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var podList core.PodList

	if err := r.List(ctx, &podList); err != nil {
		log.Error(err, "unable to list pods :()")
	} else {
		for _, item := range podList.Items {
			if item.GetName() == OpenSimulator.Spec.Name {
				log.Info("pod linked to an OpenSimulator custom resource found", "name", item.GetName())
				OpenSimulator.Status.Started = true
				if err := r.Status().Update(ctx, &OpenSimulator); err != nil {
					log.Error(err, "unable to update OpenSimulator's started status")
					return ctrl.Result{}, err
				}
				log.Info("OpenSimulator's status updated", "status")
			}
		}
	}

	if OpenSimulator.Status.Namespace == "" {
		log.Info("Setting up OpenSimlator namespace")

		err := r.Get(ctx, client.ObjectKey{Name: OpenSimulator.Spec.Namespace}, &core.Namespace{})
		if err == nil {
			log.Info("Desired namespace already exists")
		} else {
			namespace := &core.Namespace{
				ObjectMeta: meta.ObjectMeta{
					Name: OpenSimulator.Spec.Namespace,
				},
			}
			if err := r.Create(ctx, namespace); err != nil {
				return reconcile.Result{}, err
			}
		}
		OpenSimulator.Status.Namespace = OpenSimulator.Spec.Namespace
	}

	if !OpenSimulator.Status.NetworkInfo.Configured {
		log.Info("Setting up OpenSimlator route")

		var openSimRoute route.Route
		err := r.Get(ctx, client.ObjectKey{Name: OpenSimulator.Spec.Subdomain, Namespace: OpenSimulator.Spec.Namespace}, &openSimRoute)
		if err == nil {
			log.Info("Desired route already exists")
		} else {
			openSimRoute := &route.Route{
				ObjectMeta: meta.ObjectMeta{
					Name:      OpenSimulator.Spec.Subdomain,
					Namespace: OpenSimulator.Spec.Namespace,
				},
				Spec: route.RouteSpec{
					To: route.RouteTargetReference{
						Kind: "Service",
						Name: OpenSimulator.Spec.Subdomain,
					},
				},
			}

			if err := r.Create(ctx, openSimRoute); err != nil {
				log.Error(err, "failed to create route")
				return reconcile.Result{}, err
			}
		}
		log.Info("Setting host")
		OpenSimulator.Status.NetworkInfo.Host = openSimRoute.Status.Ingress[0].Host

		OpenSimulator.Status.NetworkInfo.Configured = true
	}

	if !OpenSimulator.Status.Configured {
		// Load OpenSimulator config files
		log.Info("loading OpenSimulator config files")
		// this is where you edit the .ini files
		// then we push the files to the opensim instance
		// then we start OpenSim
		OpenSimulator.Status.Configured = true
	}

	log.Info("OpenSimulator custom resource reconciled")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpenSimulatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplecomv1.OpenSimulator{}).
		Watches(
			&core.Pod{},
			handler.EnqueueRequestsFromMapFunc(r.mapPodsReqToOpenSimulatorReq),
		).
		Complete(r)
}

func (r *OpenSimulatorReconciler) mapPodsReqToOpenSimulatorReq(_ context.Context, obj client.Object) []reconcile.Request {
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
