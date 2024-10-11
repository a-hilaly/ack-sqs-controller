// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package queue

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.ContentBasedDeduplication, b.ko.Spec.ContentBasedDeduplication) {
		delta.Add("Spec.ContentBasedDeduplication", a.ko.Spec.ContentBasedDeduplication, b.ko.Spec.ContentBasedDeduplication)
	} else if a.ko.Spec.ContentBasedDeduplication != nil && b.ko.Spec.ContentBasedDeduplication != nil {
		if *a.ko.Spec.ContentBasedDeduplication != *b.ko.Spec.ContentBasedDeduplication {
			delta.Add("Spec.ContentBasedDeduplication", a.ko.Spec.ContentBasedDeduplication, b.ko.Spec.ContentBasedDeduplication)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DelaySeconds, b.ko.Spec.DelaySeconds) {
		delta.Add("Spec.DelaySeconds", a.ko.Spec.DelaySeconds, b.ko.Spec.DelaySeconds)
	} else if a.ko.Spec.DelaySeconds != nil && b.ko.Spec.DelaySeconds != nil {
		if *a.ko.Spec.DelaySeconds != *b.ko.Spec.DelaySeconds {
			delta.Add("Spec.DelaySeconds", a.ko.Spec.DelaySeconds, b.ko.Spec.DelaySeconds)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.FIFOQueue, b.ko.Spec.FIFOQueue) {
		delta.Add("Spec.FIFOQueue", a.ko.Spec.FIFOQueue, b.ko.Spec.FIFOQueue)
	} else if a.ko.Spec.FIFOQueue != nil && b.ko.Spec.FIFOQueue != nil {
		if *a.ko.Spec.FIFOQueue != *b.ko.Spec.FIFOQueue {
			delta.Add("Spec.FIFOQueue", a.ko.Spec.FIFOQueue, b.ko.Spec.FIFOQueue)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSDataKeyReusePeriodSeconds, b.ko.Spec.KMSDataKeyReusePeriodSeconds) {
		delta.Add("Spec.KMSDataKeyReusePeriodSeconds", a.ko.Spec.KMSDataKeyReusePeriodSeconds, b.ko.Spec.KMSDataKeyReusePeriodSeconds)
	} else if a.ko.Spec.KMSDataKeyReusePeriodSeconds != nil && b.ko.Spec.KMSDataKeyReusePeriodSeconds != nil {
		if *a.ko.Spec.KMSDataKeyReusePeriodSeconds != *b.ko.Spec.KMSDataKeyReusePeriodSeconds {
			delta.Add("Spec.KMSDataKeyReusePeriodSeconds", a.ko.Spec.KMSDataKeyReusePeriodSeconds, b.ko.Spec.KMSDataKeyReusePeriodSeconds)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSMasterKeyID, b.ko.Spec.KMSMasterKeyID) {
		delta.Add("Spec.KMSMasterKeyID", a.ko.Spec.KMSMasterKeyID, b.ko.Spec.KMSMasterKeyID)
	} else if a.ko.Spec.KMSMasterKeyID != nil && b.ko.Spec.KMSMasterKeyID != nil {
		if *a.ko.Spec.KMSMasterKeyID != *b.ko.Spec.KMSMasterKeyID {
			delta.Add("Spec.KMSMasterKeyID", a.ko.Spec.KMSMasterKeyID, b.ko.Spec.KMSMasterKeyID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.KMSMasterKeyRef, b.ko.Spec.KMSMasterKeyRef) {
		delta.Add("Spec.KMSMasterKeyRef", a.ko.Spec.KMSMasterKeyRef, b.ko.Spec.KMSMasterKeyRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MaximumMessageSize, b.ko.Spec.MaximumMessageSize) {
		delta.Add("Spec.MaximumMessageSize", a.ko.Spec.MaximumMessageSize, b.ko.Spec.MaximumMessageSize)
	} else if a.ko.Spec.MaximumMessageSize != nil && b.ko.Spec.MaximumMessageSize != nil {
		if *a.ko.Spec.MaximumMessageSize != *b.ko.Spec.MaximumMessageSize {
			delta.Add("Spec.MaximumMessageSize", a.ko.Spec.MaximumMessageSize, b.ko.Spec.MaximumMessageSize)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MessageRetentionPeriod, b.ko.Spec.MessageRetentionPeriod) {
		delta.Add("Spec.MessageRetentionPeriod", a.ko.Spec.MessageRetentionPeriod, b.ko.Spec.MessageRetentionPeriod)
	} else if a.ko.Spec.MessageRetentionPeriod != nil && b.ko.Spec.MessageRetentionPeriod != nil {
		if *a.ko.Spec.MessageRetentionPeriod != *b.ko.Spec.MessageRetentionPeriod {
			delta.Add("Spec.MessageRetentionPeriod", a.ko.Spec.MessageRetentionPeriod, b.ko.Spec.MessageRetentionPeriod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Policy, b.ko.Spec.Policy) {
		delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
	} else if a.ko.Spec.Policy != nil && b.ko.Spec.Policy != nil {
		if *a.ko.Spec.Policy != *b.ko.Spec.Policy {
			delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.PolicyRef, b.ko.Spec.PolicyRef) {
		delta.Add("Spec.PolicyRef", a.ko.Spec.PolicyRef, b.ko.Spec.PolicyRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.QueueName, b.ko.Spec.QueueName) {
		delta.Add("Spec.QueueName", a.ko.Spec.QueueName, b.ko.Spec.QueueName)
	} else if a.ko.Spec.QueueName != nil && b.ko.Spec.QueueName != nil {
		if *a.ko.Spec.QueueName != *b.ko.Spec.QueueName {
			delta.Add("Spec.QueueName", a.ko.Spec.QueueName, b.ko.Spec.QueueName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ReceiveMessageWaitTimeSeconds, b.ko.Spec.ReceiveMessageWaitTimeSeconds) {
		delta.Add("Spec.ReceiveMessageWaitTimeSeconds", a.ko.Spec.ReceiveMessageWaitTimeSeconds, b.ko.Spec.ReceiveMessageWaitTimeSeconds)
	} else if a.ko.Spec.ReceiveMessageWaitTimeSeconds != nil && b.ko.Spec.ReceiveMessageWaitTimeSeconds != nil {
		if *a.ko.Spec.ReceiveMessageWaitTimeSeconds != *b.ko.Spec.ReceiveMessageWaitTimeSeconds {
			delta.Add("Spec.ReceiveMessageWaitTimeSeconds", a.ko.Spec.ReceiveMessageWaitTimeSeconds, b.ko.Spec.ReceiveMessageWaitTimeSeconds)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RedriveAllowPolicy, b.ko.Spec.RedriveAllowPolicy) {
		delta.Add("Spec.RedriveAllowPolicy", a.ko.Spec.RedriveAllowPolicy, b.ko.Spec.RedriveAllowPolicy)
	} else if a.ko.Spec.RedriveAllowPolicy != nil && b.ko.Spec.RedriveAllowPolicy != nil {
		if *a.ko.Spec.RedriveAllowPolicy != *b.ko.Spec.RedriveAllowPolicy {
			delta.Add("Spec.RedriveAllowPolicy", a.ko.Spec.RedriveAllowPolicy, b.ko.Spec.RedriveAllowPolicy)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RedrivePolicy, b.ko.Spec.RedrivePolicy) {
		delta.Add("Spec.RedrivePolicy", a.ko.Spec.RedrivePolicy, b.ko.Spec.RedrivePolicy)
	} else if a.ko.Spec.RedrivePolicy != nil && b.ko.Spec.RedrivePolicy != nil {
		if *a.ko.Spec.RedrivePolicy != *b.ko.Spec.RedrivePolicy {
			delta.Add("Spec.RedrivePolicy", a.ko.Spec.RedrivePolicy, b.ko.Spec.RedrivePolicy)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SQSManagedSSEEnabled, b.ko.Spec.SQSManagedSSEEnabled) {
		delta.Add("Spec.SQSManagedSSEEnabled", a.ko.Spec.SQSManagedSSEEnabled, b.ko.Spec.SQSManagedSSEEnabled)
	} else if a.ko.Spec.SQSManagedSSEEnabled != nil && b.ko.Spec.SQSManagedSSEEnabled != nil {
		if *a.ko.Spec.SQSManagedSSEEnabled != *b.ko.Spec.SQSManagedSSEEnabled {
			delta.Add("Spec.SQSManagedSSEEnabled", a.ko.Spec.SQSManagedSSEEnabled, b.ko.Spec.SQSManagedSSEEnabled)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.VisibilityTimeout, b.ko.Spec.VisibilityTimeout) {
		delta.Add("Spec.VisibilityTimeout", a.ko.Spec.VisibilityTimeout, b.ko.Spec.VisibilityTimeout)
	} else if a.ko.Spec.VisibilityTimeout != nil && b.ko.Spec.VisibilityTimeout != nil {
		if *a.ko.Spec.VisibilityTimeout != *b.ko.Spec.VisibilityTimeout {
			delta.Add("Spec.VisibilityTimeout", a.ko.Spec.VisibilityTimeout, b.ko.Spec.VisibilityTimeout)
		}
	}

	return delta
}
