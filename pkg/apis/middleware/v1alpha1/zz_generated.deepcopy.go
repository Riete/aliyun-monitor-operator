// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsAliyunAccount) DeepCopyInto(out *RdsAliyunAccount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsAliyunAccount.
func (in *RdsAliyunAccount) DeepCopy() *RdsAliyunAccount {
	if in == nil {
		return nil
	}
	out := new(RdsAliyunAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsMonitor) DeepCopyInto(out *RdsMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsMonitor.
func (in *RdsMonitor) DeepCopy() *RdsMonitor {
	if in == nil {
		return nil
	}
	out := new(RdsMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RdsMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsMonitorList) DeepCopyInto(out *RdsMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RdsMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsMonitorList.
func (in *RdsMonitorList) DeepCopy() *RdsMonitorList {
	if in == nil {
		return nil
	}
	out := new(RdsMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RdsMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsMonitorSpec) DeepCopyInto(out *RdsMonitorSpec) {
	*out = *in
	if in.RdsAliyunAccount != nil {
		in, out := &in.RdsAliyunAccount, &out.RdsAliyunAccount
		*out = make([]RdsAliyunAccount, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsMonitorSpec.
func (in *RdsMonitorSpec) DeepCopy() *RdsMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(RdsMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsMonitorStatus) DeepCopyInto(out *RdsMonitorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsMonitorStatus.
func (in *RdsMonitorStatus) DeepCopy() *RdsMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(RdsMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisAliyunAccount) DeepCopyInto(out *RedisAliyunAccount) {
	*out = *in
	if in.ExtraMetric != nil {
		in, out := &in.ExtraMetric, &out.ExtraMetric
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisAliyunAccount.
func (in *RedisAliyunAccount) DeepCopy() *RedisAliyunAccount {
	if in == nil {
		return nil
	}
	out := new(RedisAliyunAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisMonitor) DeepCopyInto(out *RedisMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisMonitor.
func (in *RedisMonitor) DeepCopy() *RedisMonitor {
	if in == nil {
		return nil
	}
	out := new(RedisMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisMonitorList) DeepCopyInto(out *RedisMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisMonitorList.
func (in *RedisMonitorList) DeepCopy() *RedisMonitorList {
	if in == nil {
		return nil
	}
	out := new(RedisMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisMonitorSpec) DeepCopyInto(out *RedisMonitorSpec) {
	*out = *in
	if in.RedisAliyunAccount != nil {
		in, out := &in.RedisAliyunAccount, &out.RedisAliyunAccount
		*out = make([]RedisAliyunAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisMonitorSpec.
func (in *RedisMonitorSpec) DeepCopy() *RedisMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(RedisMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisMonitorStatus) DeepCopyInto(out *RedisMonitorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisMonitorStatus.
func (in *RedisMonitorStatus) DeepCopy() *RedisMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(RedisMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlbAliyunAccount) DeepCopyInto(out *SlbAliyunAccount) {
	*out = *in
	if in.InstanceId != nil {
		in, out := &in.InstanceId, &out.InstanceId
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlbAliyunAccount.
func (in *SlbAliyunAccount) DeepCopy() *SlbAliyunAccount {
	if in == nil {
		return nil
	}
	out := new(SlbAliyunAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlbMonitor) DeepCopyInto(out *SlbMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlbMonitor.
func (in *SlbMonitor) DeepCopy() *SlbMonitor {
	if in == nil {
		return nil
	}
	out := new(SlbMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlbMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlbMonitorList) DeepCopyInto(out *SlbMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlbMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlbMonitorList.
func (in *SlbMonitorList) DeepCopy() *SlbMonitorList {
	if in == nil {
		return nil
	}
	out := new(SlbMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlbMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlbMonitorSpec) DeepCopyInto(out *SlbMonitorSpec) {
	*out = *in
	if in.SlbAliyunAccount != nil {
		in, out := &in.SlbAliyunAccount, &out.SlbAliyunAccount
		*out = make([]SlbAliyunAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlbMonitorSpec.
func (in *SlbMonitorSpec) DeepCopy() *SlbMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(SlbMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlbMonitorStatus) DeepCopyInto(out *SlbMonitorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlbMonitorStatus.
func (in *SlbMonitorStatus) DeepCopy() *SlbMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(SlbMonitorStatus)
	in.DeepCopyInto(out)
	return out
}
