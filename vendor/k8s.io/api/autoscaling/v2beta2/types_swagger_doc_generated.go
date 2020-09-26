/*
Copyright The Kubernetes Authors.

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

package v2beta2

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_CrossVersionObjectReference = map[string]string{
	"":           "CrossVersionObjectReference contains enough information to let you identify the referred resource.",
	"kind":       "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds\"",
	"name":       "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
	"apiVersion": "API version of the referent",
}

func (CrossVersionObjectReference) SwaggerDoc() map[string]string {
	return map_CrossVersionObjectReference
}

var map_ExternalMetricSource = map[string]string{
	"":       "ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
	"metric": "metric identifies the target metric by name and selector",
	"target": "target specifies the target value for the given metric",
}

func (ExternalMetricSource) SwaggerDoc() map[string]string {
	return map_ExternalMetricSource
}

var map_ExternalMetricStatus = map[string]string{
	"":        "ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.",
	"metric":  "metric identifies the target metric by name and selector",
	"current": "current contains the current value for the given metric",
}

func (ExternalMetricStatus) SwaggerDoc() map[string]string {
	return map_ExternalMetricStatus
}

var map_HPAScalingPolicy = map[string]string{
	"":              "HPAScalingPolicy is a single policy which must hold true for a specified past interval.",
	"type":          "Type is used to specify the scaling policy.",
	"value":         "Value contains the amount of change which is permitted by the policy. It must be greater than zero",
	"periodSeconds": "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
}

func (HPAScalingPolicy) SwaggerDoc() map[string]string {
	return map_HPAScalingPolicy
}

var map_HPAScalingRules = map[string]string{
	"":                           "HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead, the safest value from the stabilization window is chosen.",
	"stabilizationWindowSeconds": "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
	"selectPolicy":               "selectPolicy is used to specify which policy should be used. If not set, the default value MaxPolicySelect is used.",
	"policies":                   "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
}

func (HPAScalingRules) SwaggerDoc() map[string]string {
	return map_HPAScalingRules
}

var map_HorizontalPodAutoscaler = map[string]string{
	"":         "HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.",
	"metadata": "metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.",
	"status":   "status is the current information about the autoscaler.",
}

func (HorizontalPodAutoscaler) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscaler
}

var map_HorizontalPodAutoscalerBehavior = map[string]string{
	"":          "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
	"scaleUp":   "scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of:\n  * increase no more than 4 pods per 60 seconds\n  * double the number of pods per 60 seconds\nNo stabilization is used.",
	"scaleDown": "scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).",
}

func (HorizontalPodAutoscalerBehavior) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerBehavior
}

var map_HorizontalPodAutoscalerCondition = map[string]string{
	"":                   "HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.",
	"type":               "type describes the current condition",
	"status":             "status is the status of the condition (True, False, Unknown)",
	"lastTransitionTime": "lastTransitionTime is the last time the condition transitioned from one status to another",
	"reason":             "reason is the reason for the condition's last transition.",
	"message":            "message is a human-readable explanation containing details about the transition",
}

func (HorizontalPodAutoscalerCondition) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerCondition
}

var map_HorizontalPodAutoscalerList = map[string]string{
	"":         "HorizontalPodAutoscalerList is a list of horizontal pod autoscaler objects.",
	"metadata": "metadata is the standard list metadata.",
	"items":    "items is the list of horizontal pod autoscaler objects.",
}

func (HorizontalPodAutoscalerList) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerList
}

var map_HorizontalPodAutoscalerSpec = map[string]string{
	"":               "HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.",
	"scaleTargetRef": "scaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.",
	"minReplicas":    "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
	"maxReplicas":    "maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.",
	"metrics":        "metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization.",
	"behavior":       "behavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively). If not set, the default HPAScalingRules for scale up and scale down are used.",
}

func (HorizontalPodAutoscalerSpec) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerSpec
}

var map_HorizontalPodAutoscalerStatus = map[string]string{
	"":                   "HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.",
	"observedGeneration": "observedGeneration is the most recent generation observed by this autoscaler.",
	"lastScaleTime":      "lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.",
	"currentReplicas":    "currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.",
	"desiredReplicas":    "desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.",
	"currentMetrics":     "currentMetrics is the last read state of the metrics used by this autoscaler.",
	"conditions":         "conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.",
}

func (HorizontalPodAutoscalerStatus) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerStatus
}

var map_MetricIdentifier = map[string]string{
	"":         "MetricIdentifier defines the name and optionally selector for a metric",
	"name":     "name is the name of the given metric",
	"selector": "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
}

func (MetricIdentifier) SwaggerDoc() map[string]string {
	return map_MetricIdentifier
}

var map_MetricSpec = map[string]string{
	"":         "MetricSpec specifies how to scale based on a single metric (only `type` and one other matching field should be set at once).",
	"type":     "type is the type of metric source.  It should be one of \"Object\", \"Pods\", \"Resource\" or \"External\", each mapping to a matching field in the object.",
	"object":   "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
	"pods":     "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
	"resource": "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"external": "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
}

func (MetricSpec) SwaggerDoc() map[string]string {
	return map_MetricSpec
}

var map_MetricStatus = map[string]string{
	"":         "MetricStatus describes the last-read state of a single metric.",
	"type":     "type is the type of metric source.  It will be one of \"Object\", \"Pods\" or \"Resource\", each corresponds to a matching field in the object.",
	"object":   "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
	"pods":     "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
	"resource": "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"external": "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
}

func (MetricStatus) SwaggerDoc() map[string]string {
	return map_MetricStatus
}

var map_MetricTarget = map[string]string{
	"":                   "MetricTarget defines the target value, average value, or average utilization of a specific metric",
	"type":               "type represents whether the metric type is Utilization, Value, or AverageValue",
	"value":              "value is the target value of the metric (as a quantity).",
	"averageValue":       "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
	"averageUtilization": "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type",
}

func (MetricTarget) SwaggerDoc() map[string]string {
	return map_MetricTarget
}

var map_MetricValueStatus = map[string]string{
	"":                   "MetricValueStatus holds the current value for a metric",
	"value":              "value is the current value of the metric (as a quantity).",
	"averageValue":       "averageValue is the current value of the average of the metric across all relevant pods (as a quantity)",
	"averageUtilization": "currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.",
}

func (MetricValueStatus) SwaggerDoc() map[string]string {
	return map_MetricValueStatus
}

var map_ObjectMetricSource = map[string]string{
	"":       "ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).",
	"target": "target specifies the target value for the given metric",
	"metric": "metric identifies the target metric by name and selector",
}

func (ObjectMetricSource) SwaggerDoc() map[string]string {
	return map_ObjectMetricSource
}

var map_ObjectMetricStatus = map[string]string{
	"":        "ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).",
	"metric":  "metric identifies the target metric by name and selector",
	"current": "current contains the current value for the given metric",
}

func (ObjectMetricStatus) SwaggerDoc() map[string]string {
	return map_ObjectMetricStatus
}

var map_PodsMetricSource = map[string]string{
	"":       "PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.",
	"metric": "metric identifies the target metric by name and selector",
	"target": "target specifies the target value for the given metric",
}

func (PodsMetricSource) SwaggerDoc() map[string]string {
	return map_PodsMetricSource
}

var map_PodsMetricStatus = map[string]string{
	"":        "PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).",
	"metric":  "metric identifies the target metric by name and selector",
	"current": "current contains the current value for the given metric",
}

func (PodsMetricStatus) SwaggerDoc() map[string]string {
	return map_PodsMetricStatus
}

var map_ResourceMetricSource = map[string]string{
	"":       "ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.  Only one \"target\" type should be set.",
	"name":   "name is the name of the resource in question.",
	"target": "target specifies the target value for the given metric",
}

func (ResourceMetricSource) SwaggerDoc() map[string]string {
	return map_ResourceMetricSource
}

var map_ResourceMetricStatus = map[string]string{
	"":        "ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"name":    "Name is the name of the resource in question.",
	"current": "current contains the current value for the given metric",
}

func (ResourceMetricStatus) SwaggerDoc() map[string]string {
	return map_ResourceMetricStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
