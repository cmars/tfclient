// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/function.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function []*FunctionDef `protobuf:"bytes,1,rep,name=function" json:"function,omitempty"`
	Gradient []*GradientDef `protobuf:"bytes,2,rep,name=gradient" json:"gradient,omitempty"`
}

func (m *FunctionDefLibrary) Reset()                    { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string            { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()               {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret map[string]string `protobuf:"bytes,4,rep,name=ret" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FunctionDef) Reset()                    { *m = FunctionDef{} }
func (m *FunctionDef) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()               {}
func (*FunctionDef) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName" json:"function_name,omitempty"`
	GradientFunc string `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc" json:"gradient_func,omitempty"`
}

func (m *GradientDef) Reset()                    { *m = GradientDef{} }
func (m *GradientDef) String() string            { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()               {}
func (*GradientDef) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/function.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0x5f, 0x4f, 0xf2, 0x30,
	0x14, 0xc6, 0x33, 0x06, 0xef, 0xcb, 0x0e, 0x68, 0xb4, 0x6a, 0x5c, 0x76, 0x85, 0x98, 0x18, 0xa2,
	0xc9, 0x96, 0x40, 0x34, 0xc6, 0x3b, 0x89, 0x7f, 0x6e, 0x0c, 0x92, 0x5d, 0xe8, 0xe5, 0x32, 0xa0,
	0x23, 0x0b, 0xd8, 0x92, 0x52, 0x24, 0xdc, 0xf8, 0x5d, 0xfd, 0x16, 0x5e, 0xda, 0x76, 0x2b, 0x6b,
	0xa2, 0xf3, 0xae, 0xe9, 0xf9, 0x3d, 0xcf, 0x73, 0x7a, 0x7a, 0xa0, 0xc3, 0x31, 0x59, 0x52, 0x96,
	0xcc, 0xe9, 0x3a, 0x18, 0x53, 0x86, 0x83, 0x84, 0xc5, 0x6f, 0x78, 0x4d, 0xd9, 0x2c, 0x48, 0x56,
	0x64, 0xcc, 0x53, 0x4a, 0xfc, 0x05, 0xa3, 0x9c, 0x22, 0x28, 0x48, 0xef, 0xbc, 0x5c, 0x15, 0x73,
	0xce, 0xa2, 0xf7, 0x78, 0xbe, 0xc2, 0x99, 0xce, 0xfb, 0x23, 0x81, 0xd0, 0x09, 0x8e, 0x26, 0x38,
	0xc9, 0xc9, 0xb3, 0x72, 0x92, 0x2e, 0x0a, 0xae, 0xfd, 0x01, 0xe8, 0x21, 0xef, 0xed, 0x0e, 0x27,
	0x4f, 0xe9, 0x88, 0xc5, 0x6c, 0x83, 0x7a, 0x50, 0xd7, 0x1d, 0xbb, 0x56, 0xcb, 0xee, 0x34, 0xba,
	0xc7, 0x7e, 0x61, 0xe8, 0x1b, 0x8a, 0x70, 0x0b, 0x4a, 0xd1, 0x94, 0xc5, 0x93, 0x14, 0x13, 0xee,
	0x56, 0x7e, 0x8a, 0x1e, 0xf3, 0x9a, 0x12, 0x69, 0xb0, 0xfd, 0x59, 0x81, 0x86, 0x61, 0x87, 0x02,
	0x70, 0x96, 0xe9, 0x94, 0xc4, 0x7c, 0xc5, 0xb0, 0x88, 0xb6, 0x84, 0xcb, 0xbe, 0xe9, 0xf2, 0xbc,
	0x90, 0xfa, 0x82, 0x41, 0x97, 0x50, 0x95, 0x63, 0x72, 0x6b, 0x2a, 0xf1, 0xa4, 0xa4, 0x4d, 0xff,
	0x56, 0x30, 0xf7, 0x84, 0xb3, 0x4d, 0xa8, 0x70, 0xe4, 0x43, 0x5d, 0x4f, 0xcc, 0xb5, 0x95, 0xf4,
	0xc0, 0x94, 0x0e, 0x44, 0x4d, 0x06, 0xfd, 0x27, 0xd9, 0x01, 0x75, 0xc1, 0x66, 0x98, 0xbb, 0x55,
	0x85, 0xb6, 0xca, 0x52, 0x42, 0xcc, 0xb3, 0x10, 0x09, 0x7b, 0x03, 0x70, 0xb6, 0xb1, 0x68, 0x0f,
	0xec, 0x19, 0xde, 0xa8, 0x27, 0x39, 0xa1, 0x3c, 0xa2, 0x0b, 0xa8, 0xa9, 0xbf, 0x15, 0xc3, 0x92,
	0xcf, 0x3c, 0x32, 0x4d, 0xa5, 0xee, 0x45, 0x16, 0xc3, 0x8c, 0xb9, 0xa9, 0x5c, 0x5b, 0xde, 0x15,
	0xd4, 0x75, 0xc0, 0x2f, 0x76, 0x87, 0xa6, 0x9d, 0x63, 0xe8, 0xda, 0xaf, 0xd0, 0x30, 0x86, 0x8f,
	0x4e, 0x61, 0x47, 0xff, 0x59, 0x44, 0xc4, 0x52, 0xe4, 0x26, 0x4d, 0x7d, 0x39, 0x10, 0x77, 0x12,
	0xd2, 0x7f, 0x14, 0xc9, 0x42, 0xee, 0xda, 0xd4, 0x97, 0xf2, 0xd5, 0xfd, 0x00, 0x5c, 0xca, 0xa6,
	0x66, 0xdf, 0xdb, 0x2d, 0xeb, 0xef, 0xea, 0xb9, 0x0c, 0xe5, 0x9e, 0x2d, 0x87, 0xd6, 0x97, 0x65,
	0x8d, 0xfe, 0xa9, 0xa5, 0xeb, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xb1, 0x76, 0x27, 0x2a,
	0x03, 0x00, 0x00,
}
