// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/cilium_l7policy.proto

package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type L7Policy struct {
	// Path to the unix domain socket for the cilium access log.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath" json:"access_log_path,omitempty"`
	// Cilium endpoint security policy to enforce.
	PolicyName string `protobuf:"bytes,2,opt,name=policy_name,json=policyName" json:"policy_name,omitempty"`
	// HTTP response body message for 403 status code.
	// If empty, "Access denied" will be used.
	Denied_403Body string `protobuf:"bytes,3,opt,name=denied_403_body,json=denied403Body" json:"denied_403_body,omitempty"`
}

func (m *L7Policy) Reset()                    { *m = L7Policy{} }
func (m *L7Policy) String() string            { return proto.CompactTextString(m) }
func (*L7Policy) ProtoMessage()               {}
func (*L7Policy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *L7Policy) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func (m *L7Policy) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *L7Policy) GetDenied_403Body() string {
	if m != nil {
		return m.Denied_403Body
	}
	return ""
}

func init() {
	proto.RegisterType((*L7Policy)(nil), "cilium.L7Policy")
}

func init() { proto.RegisterFile("cilium/cilium_l7policy.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xce, 0xcc, 0xc9,
	0x2c, 0xcd, 0xd5, 0x87, 0x50, 0xf1, 0x39, 0xe6, 0x05, 0xf9, 0x39, 0x99, 0xc9, 0x95, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x61, 0xa5, 0x6a, 0x2e, 0x0e, 0x1f, 0xf3, 0x00, 0xb0,
	0x8c, 0x90, 0x1a, 0x17, 0x7f, 0x62, 0x72, 0x72, 0x6a, 0x71, 0x71, 0x7c, 0x4e, 0x7e, 0x7a, 0x7c,
	0x41, 0x62, 0x49, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x2f, 0x44, 0xd8, 0x27, 0x3f,
	0x3d, 0x20, 0xb1, 0x24, 0x43, 0x48, 0x9e, 0x8b, 0x1b, 0x62, 0x56, 0x7c, 0x5e, 0x62, 0x6e, 0xaa,
	0x04, 0x13, 0x58, 0x0d, 0x17, 0x44, 0xc8, 0x2f, 0x31, 0x37, 0x15, 0x64, 0x50, 0x4a, 0x6a, 0x5e,
	0x66, 0x6a, 0x4a, 0xbc, 0x89, 0x81, 0x71, 0x7c, 0x52, 0x7e, 0x4a, 0xa5, 0x04, 0x33, 0xc4, 0x20,
	0x88, 0xb0, 0x89, 0x81, 0xb1, 0x53, 0x7e, 0x4a, 0xa5, 0x13, 0x47, 0x14, 0xd4, 0x19, 0x49, 0x6c,
	0x60, 0x57, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xfc, 0xe6, 0xef, 0xb5, 0x00, 0x00,
	0x00,
}
