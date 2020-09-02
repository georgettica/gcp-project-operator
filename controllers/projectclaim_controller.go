/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	gcpmanagedopenshiftiov1alpha1 "github.com/openshift/gcp-project-operator/api/v1alpha1"
	"github.com/openshift/gcp-project-operator/pkg/condition"
	"github.com/openshift/gcp-project-operator/pkg/util"
	gcputil "github.com/openshift/gcp-project-operator/pkg/util"
)

// ProjectClaimReconciler reconciles a ProjectClaim object
type ProjectClaimReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//go:generate mockgen -destination=../../util/mocks/$GOPACKAGE/customeresourceadapter.go -package=$GOPACKAGE github.com/openshift/gcp-project-operator/pkg/controller/projectclaim CustomResourceAdapter
type CustomResourceAdapter interface {
	EnsureProjectClaimDeletionProcessed() (gcputil.OperationResult, error)
	ProjectReferenceExists() (bool, error)
	EnsureProjectClaimInitialized() (gcputil.OperationResult, error)
	EnsureProjectClaimStatePending() (gcputil.OperationResult, error)
	EnsureProjectClaimStatePendingProject() (gcputil.OperationResult, error)
	EnsureRegionSupported() (gcputil.OperationResult, error)
	EnsureProjectReferenceExists() (gcputil.OperationResult, error)
	EnsureProjectReferenceLink() (gcputil.OperationResult, error)
	EnsureFinalizer() (gcputil.OperationResult, error)
	FinalizeProjectClaim() (ObjectState, error)
}

// +kubebuilder:rbac:groups=gcp.managed.openshift.io.my.domain,resources=projectclaims,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gcp.managed.openshift.io.my.domain,resources=projectclaims/status,verbs=get;update;patch
func (r *ProjectClaimReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("projectclaim", req.NamespacedName)

	reqLogger := log.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)

	// Fetch the ProjectClaim instance
	instance := &gcpmanagedopenshiftiov1alpha1.ProjectClaim{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			return r.doNotRequeue()
		}
		return r.requeueOnErr(err)
	}

	conditionManager := condition.NewConditionManager()
	adapter := NewProjectClaimAdapter(instance, reqLogger, r.Client, conditionManager)
	result, err := r.ReconcileHandler(adapter)
	reason := "ReconcileError"
	_ = adapter.SetProjectClaimCondition(reason, err)

	return result, err

}

type ReconcileOperation func() (util.OperationResult, error)

// ReconcileHandler reads that state of the cluster for a ProjectClaim object and makes changes based on the state read
// and what is in the ProjectClaim.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ProjectClaimReconciler) ReconcileHandler(adapter CustomResourceAdapter) (reconcile.Result, error) {
	operations := []ReconcileOperation{
		adapter.EnsureProjectClaimDeletionProcessed,
		adapter.EnsureProjectClaimInitialized,
		adapter.EnsureRegionSupported,
		adapter.EnsureProjectClaimStatePending,
		adapter.EnsureProjectReferenceExists,
		adapter.EnsureProjectReferenceLink,
		adapter.EnsureFinalizer,
		adapter.EnsureProjectClaimStatePendingProject,
	}
	for _, operation := range operations {
		result, err := operation()
		if err != nil || result.RequeueRequest {
			return r.requeueAfter(result.RequeueDelay, err)
		}
		if result.CancelRequest {
			return r.doNotRequeue()
		}
	}
	return r.doNotRequeue()
}

func (r *ProjectClaimReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&gcpmanagedopenshiftiov1alpha1.ProjectClaim{}).
		Complete(r)
}

func (r *ProjectClaimReconciler) doNotRequeue() (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (r *ProjectClaimReconciler) requeueOnErr(err error) (ctrl.Result, error) {
	return ctrl.Result{}, err
}

func (r *ProjectClaimReconciler) requeueAfter(duration time.Duration, err error) (ctrl.Result, error) {
	return ctrl.Result{RequeueAfter: duration}, err
}
