/*
Copyright 2016 The Kubernetes Authors.

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

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AdmissionHookClientConfig = map[string]string{
	"":         "AdmissionHookClientConfig contains the information to make a TLS connection with the webhook",
	"service":  "Service is a reference to the service for this webhook. If there is only one port open for the service, that port will be used. If there are multiple ports open, port 443 will be used if it is open, otherwise it is an error. Required",
	"urlPath":  "URLPath is an optional field that specifies the URL path to use when posting the AdmissionReview object.",
	"caBundle": "CABundle is a PEM encoded CA bundle which will be used to validate webhook's server certificate. Required",
}

func (AdmissionHookClientConfig) SwaggerDoc() map[string]string {
	return map_AdmissionHookClientConfig
}

var map_ExternalAdmissionHook = map[string]string{
	"":              "ExternalAdmissionHook describes an external admission webhook and the resources and operations it applies to.",
	"name":          "The name of the external admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where \"imagepolicy\" is the name of the webhook, and kubernetes.io is the name of the organization. Required.",
	"clientConfig":  "ClientConfig defines how to communicate with the hook. Required",
	"rules":         "Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule.",
	"failurePolicy": "FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Ignore.",
}

func (ExternalAdmissionHook) SwaggerDoc() map[string]string {
	return map_ExternalAdmissionHook
}

var map_ExternalAdmissionHookConfiguration = map[string]string{
	"":                       "ExternalAdmissionHookConfiguration describes the configuration of initializers.",
	"metadata":               "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
	"externalAdmissionHooks": "ExternalAdmissionHooks is a list of external admission webhooks and the affected resources and operations.",
}

func (ExternalAdmissionHookConfiguration) SwaggerDoc() map[string]string {
	return map_ExternalAdmissionHookConfiguration
}

var map_ExternalAdmissionHookConfigurationList = map[string]string{
	"":         "ExternalAdmissionHookConfigurationList is a list of ExternalAdmissionHookConfiguration.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "List of ExternalAdmissionHookConfiguration.",
}

func (ExternalAdmissionHookConfigurationList) SwaggerDoc() map[string]string {
	return map_ExternalAdmissionHookConfigurationList
}

var map_Initializer = map[string]string{
	"":      "Initializer describes the name and the failure policy of an initializer, and what resources it applies to.",
	"name":  "Name is the identifier of the initializer. It will be added to the object that needs to be initialized. Name should be fully qualified, e.g., alwayspullimages.kubernetes.io, where \"alwayspullimages\" is the name of the webhook, and kubernetes.io is the name of the organization. Required",
	"rules": "Rules describes what resources/subresources the initializer cares about. The initializer cares about an operation if it matches _any_ Rule. Rule.Resources must not include subresources.",
}

func (Initializer) SwaggerDoc() map[string]string {
	return map_Initializer
}

var map_InitializerConfiguration = map[string]string{
	"":             "InitializerConfiguration describes the configuration of initializers.",
	"metadata":     "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
	"initializers": "Initializers is a list of resources and their default initializers Order-sensitive. When merging multiple InitializerConfigurations, we sort the initializers from different InitializerConfigurations by the name of the InitializerConfigurations; the order of the initializers from the same InitializerConfiguration is preserved.",
}

func (InitializerConfiguration) SwaggerDoc() map[string]string {
	return map_InitializerConfiguration
}

var map_InitializerConfigurationList = map[string]string{
	"":         "InitializerConfigurationList is a list of InitializerConfiguration.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "List of InitializerConfiguration.",
}

func (InitializerConfigurationList) SwaggerDoc() map[string]string {
	return map_InitializerConfigurationList
}

var map_Rule = map[string]string{
	"":            "Rule is a tuple of APIGroups, APIVersion, and Resources.It is recommended to make sure that all the tuple expansions are valid.",
	"apiGroups":   "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
	"apiVersions": "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
	"resources":   "Resources is a list of resources this rule applies to.\n\nFor example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources.\n\nIf wildcard is present, the validation rule will ensure resources do not overlap with each other.\n\nDepending on the enclosing object, subresources might not be allowed. Required.",
}

func (Rule) SwaggerDoc() map[string]string {
	return map_Rule
}

var map_RuleWithOperations = map[string]string{
	"":           "RuleWithOperations is a tuple of Operations and Resources. It is recommended to make sure that all the tuple expansions are valid.",
	"operations": "Operations is the operations the admission hook cares about - CREATE, UPDATE, or * for all operations. If '*' is present, the length of the slice must be one. Required.",
}

func (RuleWithOperations) SwaggerDoc() map[string]string {
	return map_RuleWithOperations
}

var map_ServiceReference = map[string]string{
	"":          "ServiceReference holds a reference to Service.legacy.k8s.io",
	"namespace": "Namespace is the namespace of the service Required",
	"name":      "Name is the name of the service Required",
}

func (ServiceReference) SwaggerDoc() map[string]string {
	return map_ServiceReference
}

// AUTO-GENERATED FUNCTIONS END HERE
