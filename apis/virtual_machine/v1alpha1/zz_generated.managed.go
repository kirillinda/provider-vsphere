/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VSphereContentLibrary.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VSphereContentLibrary) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VSphereContentLibrary.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VSphereContentLibrary) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VSphereContentLibrary.
func (mg *VSphereContentLibrary) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VSphereContentLibraryItem.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VSphereContentLibraryItem) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VSphereContentLibraryItem.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VSphereContentLibraryItem) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VSphereContentLibraryItem.
func (mg *VSphereContentLibraryItem) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VSphereVappContainer.
func (mg *VSphereVappContainer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VSphereVappContainer.
func (mg *VSphereVappContainer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VSphereVappContainer.
func (mg *VSphereVappContainer) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VSphereVappContainer.
func (mg *VSphereVappContainer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VSphereVappContainer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VSphereVappContainer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VSphereVappContainer.
func (mg *VSphereVappContainer) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VSphereVappContainer.
func (mg *VSphereVappContainer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VSphereVappContainer.
func (mg *VSphereVappContainer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VSphereVappContainer.
func (mg *VSphereVappContainer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VSphereVappContainer.
func (mg *VSphereVappContainer) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VSphereVappContainer.
func (mg *VSphereVappContainer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VSphereVappContainer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VSphereVappContainer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VSphereVappContainer.
func (mg *VSphereVappContainer) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VSphereVappContainer.
func (mg *VSphereVappContainer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VSphereVappEntity.
func (mg *VSphereVappEntity) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VSphereVappEntity.
func (mg *VSphereVappEntity) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VSphereVappEntity.
func (mg *VSphereVappEntity) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VSphereVappEntity.
func (mg *VSphereVappEntity) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VSphereVappEntity.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VSphereVappEntity) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VSphereVappEntity.
func (mg *VSphereVappEntity) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VSphereVappEntity.
func (mg *VSphereVappEntity) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VSphereVappEntity.
func (mg *VSphereVappEntity) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VSphereVappEntity.
func (mg *VSphereVappEntity) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VSphereVappEntity.
func (mg *VSphereVappEntity) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VSphereVappEntity.
func (mg *VSphereVappEntity) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VSphereVappEntity.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VSphereVappEntity) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VSphereVappEntity.
func (mg *VSphereVappEntity) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VSphereVappEntity.
func (mg *VSphereVappEntity) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VSphereVirtualDisk.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VSphereVirtualDisk) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VSphereVirtualDisk.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VSphereVirtualDisk) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VSphereVirtualMachine.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VSphereVirtualMachine) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VSphereVirtualMachine.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VSphereVirtualMachine) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VSphereVirtualMachineSnapshot.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VSphereVirtualMachineSnapshot) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VSphereVirtualMachineSnapshot.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VSphereVirtualMachineSnapshot) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VSphereVirtualMachineSnapshot.
func (mg *VSphereVirtualMachineSnapshot) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
