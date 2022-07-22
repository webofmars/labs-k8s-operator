/*
Copyright 2022.

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
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	pizzav1 "webofmars.com/pizza/api/v1"
)

// PizzaReconciler reconciles a Pizza object
type PizzaReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=pizza.webofmars.com,resources=pizzas,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pizza.webofmars.com,resources=pizzas/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pizza.webofmars.com,resources=pizzas/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Pizza object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *PizzaReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Pizza delivery logic goes here
	var pizza pizzav1.Pizza
	if err := r.Get(ctx, req.NamespacedName, &pizza); err != nil {
		log.Log.Error(err, "unable to fetch Foo")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// fake the delivery
	time.Sleep(60 * time.Second)

	// Update delivery status
	pizza.Status.Delivered = true
	if err := r.Status().Update(ctx, &pizza); err != nil {
		log.Log.Error(err, "unable to update pizza delivered status", "status", true)
		return ctrl.Result{}, err
	}
	log.Log.Info("pizza delivered status updated", "status", true)

	// end of logic

	log.Log.Info("pizza custom resource reconciled")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PizzaReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pizzav1.Pizza{}).
		Complete(r)
}
