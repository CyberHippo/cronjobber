/*
Copyright 2017 The Kubernetes Authors.
Copyright 2019 The Cronjobber Authors.

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

import (
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TZCronJob represents the configuration of a single cron job.
type TZCronJob struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of a cron job, including the schedule.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec TZCronJobSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Current status of a cron job.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Status TZCronJobStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TZCronJobList is a collection of cron jobs.
type TZCronJobList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of CronJobs.
	Items []TZCronJob `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// TZCronJobSpec describes how the job execution will look like and when it will actually run.
type TZCronJobSpec struct {

	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule string `json:"schedule" protobuf:"bytes,1,opt,name=schedule"`

	// The time zone to be used when scheduling the job given the schedule specified above.
	// Valid options include anything in the IANA Time Zone database.
	// +optional
	TimeZone string `json:"timezone" protobuf:"bytes,8,opt,name=timeZone"`

	// Optional deadline in seconds for starting the job if it misses scheduled
	// time for any reason.  Missed jobs executions will be counted as failed ones.
	// +optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty" protobuf:"varint,2,opt,name=startingDeadlineSeconds"`

	// Specifies how to treat concurrent executions of a Job.
	// Valid values are:
	// - "Allow" (default): allows CronJobs to run concurrently;
	// - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet;
	// - "Replace": cancels currently running job and replaces it with a new one
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty" protobuf:"bytes,3,opt,name=concurrencyPolicy,casttype=ConcurrencyPolicy"`

	// This flag tells the controller to suspend subsequent executions, it does
	// not apply to already started executions.  Defaults to false.
	// +optional
	Suspend *bool `json:"suspend,omitempty" protobuf:"varint,4,opt,name=suspend"`

	// Specifies the job that will be created when executing a TZCronJob.
	JobTemplate batchv1beta1.JobTemplateSpec `json:"jobTemplate" protobuf:"bytes,5,opt,name=jobTemplate"`

	// The number of successful finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 3.
	// +optional
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty" protobuf:"varint,6,opt,name=successfulJobsHistoryLimit"`

	// The number of failed finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 1.
	// +optional
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty" protobuf:"varint,7,opt,name=failedJobsHistoryLimit"`
}

// ConcurrencyPolicy describes how the job will be handled.
// Only one of the following concurrent policies may be specified.
// If none of the following policies is specified, the default one
// is AllowConcurrent.
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent forbids concurrent runs, skipping next run if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// TZCronJobStatus represents the current state of a cron job.
type TZCronJobStatus struct {
	// A list of pointers to currently running jobs.
	// +optional
	Active []v1.ObjectReference `json:"active,omitempty" protobuf:"bytes,1,rep,name=active"`

	// Information when was the last time the job was successfully scheduled.
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty" protobuf:"bytes,4,opt,name=lastScheduleTime"`
}
