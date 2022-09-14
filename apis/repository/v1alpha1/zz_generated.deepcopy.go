//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchesObservation) DeepCopyInto(out *BranchesObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Protected != nil {
		in, out := &in.Protected, &out.Protected
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchesObservation.
func (in *BranchesObservation) DeepCopy() *BranchesObservation {
	if in == nil {
		return nil
	}
	out := new(BranchesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchesParameters) DeepCopyInto(out *BranchesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchesParameters.
func (in *BranchesParameters) DeepCopy() *BranchesParameters {
	if in == nil {
		return nil
	}
	out := new(BranchesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *File) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileList) DeepCopyInto(out *FileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileList.
func (in *FileList) DeepCopy() *FileList {
	if in == nil {
		return nil
	}
	out := new(FileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileObservation) DeepCopyInto(out *FileObservation) {
	*out = *in
	if in.CommitSha != nil {
		in, out := &in.CommitSha, &out.CommitSha
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Sha != nil {
		in, out := &in.Sha, &out.Sha
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileObservation.
func (in *FileObservation) DeepCopy() *FileObservation {
	if in == nil {
		return nil
	}
	out := new(FileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileParameters) DeepCopyInto(out *FileParameters) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.CommitAuthor != nil {
		in, out := &in.CommitAuthor, &out.CommitAuthor
		*out = new(string)
		**out = **in
	}
	if in.CommitEmail != nil {
		in, out := &in.CommitEmail, &out.CommitEmail
		*out = new(string)
		**out = **in
	}
	if in.CommitMessage != nil {
		in, out := &in.CommitMessage, &out.CommitMessage
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.OverwriteOnCreate != nil {
		in, out := &in.OverwriteOnCreate, &out.OverwriteOnCreate
		*out = new(bool)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileParameters.
func (in *FileParameters) DeepCopy() *FileParameters {
	if in == nil {
		return nil
	}
	out := new(FileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpec) DeepCopyInto(out *FileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpec.
func (in *FileSpec) DeepCopy() *FileSpec {
	if in == nil {
		return nil
	}
	out := new(FileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileStatus) DeepCopyInto(out *FileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileStatus.
func (in *FileStatus) DeepCopy() *FileStatus {
	if in == nil {
		return nil
	}
	out := new(FileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagesObservation) DeepCopyInto(out *PagesObservation) {
	*out = *in
	if in.Custom404 != nil {
		in, out := &in.Custom404, &out.Custom404
		*out = new(bool)
		**out = **in
	}
	if in.HTMLURL != nil {
		in, out := &in.HTMLURL, &out.HTMLURL
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagesObservation.
func (in *PagesObservation) DeepCopy() *PagesObservation {
	if in == nil {
		return nil
	}
	out := new(PagesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagesParameters) DeepCopyInto(out *PagesParameters) {
	*out = *in
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]SourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagesParameters.
func (in *PagesParameters) DeepCopy() *PagesParameters {
	if in == nil {
		return nil
	}
	out := new(PagesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryObservation) DeepCopyInto(out *RepositoryObservation) {
	*out = *in
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]BranchesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.FullName != nil {
		in, out := &in.FullName, &out.FullName
		*out = new(string)
		**out = **in
	}
	if in.GitCloneURL != nil {
		in, out := &in.GitCloneURL, &out.GitCloneURL
		*out = new(string)
		**out = **in
	}
	if in.HTMLURL != nil {
		in, out := &in.HTMLURL, &out.HTMLURL
		*out = new(string)
		**out = **in
	}
	if in.HTTPCloneURL != nil {
		in, out := &in.HTTPCloneURL, &out.HTTPCloneURL
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NodeID != nil {
		in, out := &in.NodeID, &out.NodeID
		*out = new(string)
		**out = **in
	}
	if in.RepoID != nil {
		in, out := &in.RepoID, &out.RepoID
		*out = new(float64)
		**out = **in
	}
	if in.SSHCloneURL != nil {
		in, out := &in.SSHCloneURL, &out.SSHCloneURL
		*out = new(string)
		**out = **in
	}
	if in.SvnURL != nil {
		in, out := &in.SvnURL, &out.SvnURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryObservation.
func (in *RepositoryObservation) DeepCopy() *RepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryParameters) DeepCopyInto(out *RepositoryParameters) {
	*out = *in
	if in.AllowAutoMerge != nil {
		in, out := &in.AllowAutoMerge, &out.AllowAutoMerge
		*out = new(bool)
		**out = **in
	}
	if in.AllowMergeCommit != nil {
		in, out := &in.AllowMergeCommit, &out.AllowMergeCommit
		*out = new(bool)
		**out = **in
	}
	if in.AllowRebaseMerge != nil {
		in, out := &in.AllowRebaseMerge, &out.AllowRebaseMerge
		*out = new(bool)
		**out = **in
	}
	if in.AllowSquashMerge != nil {
		in, out := &in.AllowSquashMerge, &out.AllowSquashMerge
		*out = new(bool)
		**out = **in
	}
	if in.ArchiveOnDestroy != nil {
		in, out := &in.ArchiveOnDestroy, &out.ArchiveOnDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Archived != nil {
		in, out := &in.Archived, &out.Archived
		*out = new(bool)
		**out = **in
	}
	if in.AutoInit != nil {
		in, out := &in.AutoInit, &out.AutoInit
		*out = new(bool)
		**out = **in
	}
	if in.DefaultBranch != nil {
		in, out := &in.DefaultBranch, &out.DefaultBranch
		*out = new(string)
		**out = **in
	}
	if in.DeleteBranchOnMerge != nil {
		in, out := &in.DeleteBranchOnMerge, &out.DeleteBranchOnMerge
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GitignoreTemplate != nil {
		in, out := &in.GitignoreTemplate, &out.GitignoreTemplate
		*out = new(string)
		**out = **in
	}
	if in.HasDownloads != nil {
		in, out := &in.HasDownloads, &out.HasDownloads
		*out = new(bool)
		**out = **in
	}
	if in.HasIssues != nil {
		in, out := &in.HasIssues, &out.HasIssues
		*out = new(bool)
		**out = **in
	}
	if in.HasProjects != nil {
		in, out := &in.HasProjects, &out.HasProjects
		*out = new(bool)
		**out = **in
	}
	if in.HasWiki != nil {
		in, out := &in.HasWiki, &out.HasWiki
		*out = new(bool)
		**out = **in
	}
	if in.HomepageURL != nil {
		in, out := &in.HomepageURL, &out.HomepageURL
		*out = new(string)
		**out = **in
	}
	if in.IgnoreVulnerabilityAlertsDuringRead != nil {
		in, out := &in.IgnoreVulnerabilityAlertsDuringRead, &out.IgnoreVulnerabilityAlertsDuringRead
		*out = new(bool)
		**out = **in
	}
	if in.IsTemplate != nil {
		in, out := &in.IsTemplate, &out.IsTemplate
		*out = new(bool)
		**out = **in
	}
	if in.LicenseTemplate != nil {
		in, out := &in.LicenseTemplate, &out.LicenseTemplate
		*out = new(string)
		**out = **in
	}
	if in.MergeCommitMessage != nil {
		in, out := &in.MergeCommitMessage, &out.MergeCommitMessage
		*out = new(string)
		**out = **in
	}
	if in.MergeCommitTitle != nil {
		in, out := &in.MergeCommitTitle, &out.MergeCommitTitle
		*out = new(string)
		**out = **in
	}
	if in.Pages != nil {
		in, out := &in.Pages, &out.Pages
		*out = make([]PagesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(bool)
		**out = **in
	}
	if in.SquashMergeCommitMessage != nil {
		in, out := &in.SquashMergeCommitMessage, &out.SquashMergeCommitMessage
		*out = new(string)
		**out = **in
	}
	if in.SquashMergeCommitTitle != nil {
		in, out := &in.SquashMergeCommitTitle, &out.SquashMergeCommitTitle
		*out = new(string)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = make([]TemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.VulnerabilityAlerts != nil {
		in, out := &in.VulnerabilityAlerts, &out.VulnerabilityAlerts
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryParameters.
func (in *RepositoryParameters) DeepCopy() *RepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceObservation) DeepCopyInto(out *SourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceObservation.
func (in *SourceObservation) DeepCopy() *SourceObservation {
	if in == nil {
		return nil
	}
	out := new(SourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceParameters) DeepCopyInto(out *SourceParameters) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceParameters.
func (in *SourceParameters) DeepCopy() *SourceParameters {
	if in == nil {
		return nil
	}
	out := new(SourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateObservation) DeepCopyInto(out *TemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateObservation.
func (in *TemplateObservation) DeepCopy() *TemplateObservation {
	if in == nil {
		return nil
	}
	out := new(TemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateParameters) DeepCopyInto(out *TemplateParameters) {
	*out = *in
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateParameters.
func (in *TemplateParameters) DeepCopy() *TemplateParameters {
	if in == nil {
		return nil
	}
	out := new(TemplateParameters)
	in.DeepCopyInto(out)
	return out
}