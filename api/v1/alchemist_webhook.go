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

package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	validationutils "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var alchemistlog = logf.Log.WithName("alchemist-resource")

func (r *Alchemist) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-fullmetal-com-my-domain-v1-alchemist,mutating=true,failurePolicy=fail,sideEffects=None,groups=fullmetal.com.my.domain,resources=alchemists,verbs=create;update,versions=v1,name=malchemist.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Alchemist{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Alchemist) Default() {
	alchemistlog.Info("default", "name", r.Name)
	// TODO(user): fill in your defaulting logic.

	if r.Spec.DeploymentName == "" {
		r.Spec.DeploymentName = r.Name + "-deployment"
	}
	if r.Spec.Replicas == nil || *r.Spec.Replicas == 0 {
		r.Spec.Replicas = new(int32)
		*r.Spec.Replicas = 3
	}
	if r.Spec.Image == "" {
		r.Spec.Image = "tapojit047/api-server"
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-fullmetal-com-my-domain-v1-alchemist,mutating=false,failurePolicy=fail,sideEffects=None,groups=fullmetal.com.my.domain,resources=alchemists,verbs=create;update,versions=v1,name=valchemist.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Alchemist{}

// This line is a type assertion that states that the type CronJob implements the webhook.Defaulter interface. The var _ syntax is a blank identifier
// and is used here to assert the implementation of the interface without creating a named variable. The purpose of this line is to ensure that CronJob
// implements the methods required by the webhook.Defaulter interface. This way, if CronJob does not implement the required methods, the code will fail to compile.

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Alchemist) ValidateCreate() error {

	alchemistlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return r.validateAlchemist()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Alchemist) ValidateUpdate(old runtime.Object) error {
	alchemistlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return r.validateAlchemist()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Alchemist) ValidateDelete() error {
	alchemistlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *Alchemist) validateAlchemist() error {
	var allErrs field.ErrorList
	if err := r.validateAlchemistName(); err != nil {

	}
	if err := r.validateAlchemistSpec(); err != nil {

	}
	if len(allErrs) == 0 {
		return nil
	}
	return apierrors.NewInvalid(
		schema.GroupKind{Group: "fullmetal.com", Kind: "Alchemist"},
		r.Name, allErrs)
}

func (r *Alchemist) validateAlchemistName() error {
	if len(r.ObjectMeta.Name) > validationutils.DNS1035LabelMaxLength-11 {
		// The job name length is 63 character like all Kubernetes objects
		// (which must fit in a DNS subdomain). The cronjob controller appends
		// a 11-character suffix to the cronjob (`-$TIMESTAMP`) when creating
		// a job. The job name length limit is 63 characters. Therefore cronjob
		// names must have length <= 63-11=52. If we don't validate this here,
		// then job creation will fail later.
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 52 characters")
	}
	return nil
}

func (r *Alchemist) validateAlchemistSpec() error {
	alchemistlog.Info("Validate Alchemist Spec", "name", r.Name)
	return nil
}
