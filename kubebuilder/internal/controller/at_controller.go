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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	cnatv1alpha1 "subh.org/learngo/api/v1alpha1"
)

// AtReconciler reconciles a At object
type AtReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cnat.subh.org.kube,resources=ats,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cnat.subh.org.kube,resources=ats/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cnat.subh.org.kube,resources=ats/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the At object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *AtReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := log.Log.WithValues("namespace", req.Namespace, "at", req.Name)
	reqLogger.Info("=== we are inside reconcile loop ===")

	instance := &cnatv1alpha1.At{}
	if err := r.Get(context.TODO(), req.NamespacedName, instance); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Error(err, "this object got deleted", "name", req.NamespacedName)
			return ctrl.Result{}, err
		}
		reqLogger.Error(err, "there is error")
	}

	reqLogger.Info("found object:", "AT", instance)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AtReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cnatv1alpha1.At{}).
		Complete(r)
}
