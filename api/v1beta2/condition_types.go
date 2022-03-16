/*
Copyright 2022 The Flux authors

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

package v1beta2

const SourceFinalizer = "finalizers.fluxcd.io"

const (
	// ArtifactOutdatedCondition indicates the current Artifact of the Source
	// is outdated.
	// This is a "negative polarity" or "abnormal-true" type, and is only
	// present on the resource if it is True.
	ArtifactOutdatedCondition string = "ArtifactOutdated"

	// SourceVerifiedCondition indicates the integrity of the Source has been
	// verified. If True, the integrity check succeeded. If False, it failed.
	// The Condition is only present on the resource if the integrity has been
	// verified.
	SourceVerifiedCondition string = "SourceVerified"

	// FetchFailedCondition indicates a transient or persistent fetch failure
	// of an upstream Source.
	// If True, observations on the upstream Source revision may be impossible,
	// and the Artifact available for the Source may be outdated.
	// This is a "negative polarity" or "abnormal-true" type, and is only
	// present on the resource if it is True.
	FetchFailedCondition string = "FetchFailed"

	// BuildFailedCondition indicates a transient or persistent build failure
	// of a Source's Artifact.
	// If True, the Source can be in an ArtifactOutdatedCondition.
	BuildFailedCondition string = "BuildFailed"

	// StorageOperationFailedCondition indicates a transient or persistent
	// failure related to storage. If True, the reconciliation failed while
	// performing some filesystem operation.
	StorageOperationFailedCondition string = "StorageOperationFailed"
)

const (
	// URLInvalidReason signals that a given Source has an invalid URL.
	URLInvalidReason string = "URLInvalid"

	// AuthenticationFailedReason signals that a Secret does not have the
	// required fields, or the provided credentials do not match.
	AuthenticationFailedReason string = "AuthenticationFailed"

	// DirCreationFailedReason signals a failure caused by a directory creation
	// operation.
	DirCreationFailedReason string = "DirectoryCreationFailed"

	// StatOperationFailedReason signals a failure caused by a stat operation on
	// a path.
	StatOperationFailedReason string = "StatOperationFailed"

	// ReadOperationFailedReason signals a failure caused by a read operation.
	ReadOperationFailedReason string = "ReadOperationFailed"

	// AcquireLockFailedReason signals a failure in acquiring lock.
	AcquireLockFailedReason string = "AcquireLockFailed"

	// InvalidPathReason signals a failure caused by an invalid path.
	InvalidPathReason string = "InvalidPath"

	// ArchiveOperationFailedReason signals a failure in archive operation.
	ArchiveOperationFailedReason string = "ArchiveOperationFailed"

	// SymlinkUpdateFailedReason signals a failure in updating a symlink.
	SymlinkUpdateFailedReason string = "SymlinkUpdateFailed"
)