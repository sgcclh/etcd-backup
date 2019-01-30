// +build !ignore_autogenerated

/*
Copyright 2019 Giant Swarm GmbH.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *App) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalog) DeepCopyInto(out *AppCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalog.
func (in *AppCatalog) DeepCopy() *AppCatalog {
	if in == nil {
		return nil
	}
	out := new(AppCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogList) DeepCopyInto(out *AppCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogList.
func (in *AppCatalogList) DeepCopy() *AppCatalogList {
	if in == nil {
		return nil
	}
	out := new(AppCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpec) DeepCopyInto(out *AppCatalogSpec) {
	*out = *in
	out.CatalogStorage = in.CatalogStorage
	out.Config = in.Config
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpec.
func (in *AppCatalogSpec) DeepCopy() *AppCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecCatalogStorage) DeepCopyInto(out *AppCatalogSpecCatalogStorage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecCatalogStorage.
func (in *AppCatalogSpecCatalogStorage) DeepCopy() *AppCatalogSpecCatalogStorage {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecCatalogStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecConfig) DeepCopyInto(out *AppCatalogSpecConfig) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecConfig.
func (in *AppCatalogSpecConfig) DeepCopy() *AppCatalogSpecConfig {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecConfigConfigMap) DeepCopyInto(out *AppCatalogSpecConfigConfigMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecConfigConfigMap.
func (in *AppCatalogSpecConfigConfigMap) DeepCopy() *AppCatalogSpecConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCatalogSpecConfigSecret) DeepCopyInto(out *AppCatalogSpecConfigSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCatalogSpecConfigSecret.
func (in *AppCatalogSpecConfigSecret) DeepCopy() *AppCatalogSpecConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppCatalogSpecConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppList) DeepCopyInto(out *AppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]App, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppList.
func (in *AppList) DeepCopy() *AppList {
	if in == nil {
		return nil
	}
	out := new(AppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	out.Config = in.Config
	out.KubeConfig = in.KubeConfig
	out.UserConfig = in.UserConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec.
func (in *AppSpec) DeepCopy() *AppSpec {
	if in == nil {
		return nil
	}
	out := new(AppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecConfig) DeepCopyInto(out *AppSpecConfig) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecConfig.
func (in *AppSpecConfig) DeepCopy() *AppSpecConfig {
	if in == nil {
		return nil
	}
	out := new(AppSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecConfigConfigMap) DeepCopyInto(out *AppSpecConfigConfigMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecConfigConfigMap.
func (in *AppSpecConfigConfigMap) DeepCopy() *AppSpecConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(AppSpecConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecConfigSecret) DeepCopyInto(out *AppSpecConfigSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecConfigSecret.
func (in *AppSpecConfigSecret) DeepCopy() *AppSpecConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppSpecConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecKubeConfig) DeepCopyInto(out *AppSpecKubeConfig) {
	*out = *in
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecKubeConfig.
func (in *AppSpecKubeConfig) DeepCopy() *AppSpecKubeConfig {
	if in == nil {
		return nil
	}
	out := new(AppSpecKubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecKubeConfigSecret) DeepCopyInto(out *AppSpecKubeConfigSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecKubeConfigSecret.
func (in *AppSpecKubeConfigSecret) DeepCopy() *AppSpecKubeConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppSpecKubeConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecUserConfig) DeepCopyInto(out *AppSpecUserConfig) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecUserConfig.
func (in *AppSpecUserConfig) DeepCopy() *AppSpecUserConfig {
	if in == nil {
		return nil
	}
	out := new(AppSpecUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecUserConfigConfigMap) DeepCopyInto(out *AppSpecUserConfigConfigMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecUserConfigConfigMap.
func (in *AppSpecUserConfigConfigMap) DeepCopy() *AppSpecUserConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(AppSpecUserConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpecUserConfigSecret) DeepCopyInto(out *AppSpecUserConfigSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpecUserConfigSecret.
func (in *AppSpecUserConfigSecret) DeepCopy() *AppSpecUserConfigSecret {
	if in == nil {
		return nil
	}
	out := new(AppSpecUserConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	in.LastDeployed.DeepCopyInto(&out.LastDeployed)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Chart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartList) DeepCopyInto(out *ChartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Chart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartList.
func (in *ChartList) DeepCopy() *ChartList {
	if in == nil {
		return nil
	}
	out := new(ChartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpec) DeepCopyInto(out *ChartSpec) {
	*out = *in
	out.Config = in.Config
	out.KubeConfig = in.KubeConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpec.
func (in *ChartSpec) DeepCopy() *ChartSpec {
	if in == nil {
		return nil
	}
	out := new(ChartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecConfig) DeepCopyInto(out *ChartSpecConfig) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecConfig.
func (in *ChartSpecConfig) DeepCopy() *ChartSpecConfig {
	if in == nil {
		return nil
	}
	out := new(ChartSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecConfigConfigMap) DeepCopyInto(out *ChartSpecConfigConfigMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecConfigConfigMap.
func (in *ChartSpecConfigConfigMap) DeepCopy() *ChartSpecConfigConfigMap {
	if in == nil {
		return nil
	}
	out := new(ChartSpecConfigConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecConfigSecret) DeepCopyInto(out *ChartSpecConfigSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecConfigSecret.
func (in *ChartSpecConfigSecret) DeepCopy() *ChartSpecConfigSecret {
	if in == nil {
		return nil
	}
	out := new(ChartSpecConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecKubeConfig) DeepCopyInto(out *ChartSpecKubeConfig) {
	*out = *in
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecKubeConfig.
func (in *ChartSpecKubeConfig) DeepCopy() *ChartSpecKubeConfig {
	if in == nil {
		return nil
	}
	out := new(ChartSpecKubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpecKubeConfigSecret) DeepCopyInto(out *ChartSpecKubeConfigSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpecKubeConfigSecret.
func (in *ChartSpecKubeConfigSecret) DeepCopy() *ChartSpecKubeConfigSecret {
	if in == nil {
		return nil
	}
	out := new(ChartSpecKubeConfigSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartStatus) DeepCopyInto(out *ChartStatus) {
	*out = *in
	in.LastDeployed.DeepCopyInto(&out.LastDeployed)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartStatus.
func (in *ChartStatus) DeepCopy() *ChartStatus {
	if in == nil {
		return nil
	}
	out := new(ChartStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeepCopyTime.
func (in *DeepCopyTime) DeepCopy() *DeepCopyTime {
	if in == nil {
		return nil
	}
	out := new(DeepCopyTime)
	in.DeepCopyInto(out)
	return out
}