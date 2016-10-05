// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/tensor_description.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TensorDescription struct {
	// Data type of tensor elements
	Dtype DataType `protobuf:"varint,1,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Shape of the tensor.
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
	// Information about the size and allocator used for the data
	AllocationDescription *AllocationDescription `protobuf:"bytes,4,opt,name=allocation_description,json=allocationDescription" json:"allocation_description,omitempty"`
}

func (m *TensorDescription) Reset()                    { *m = TensorDescription{} }
func (m *TensorDescription) String() string            { return proto.CompactTextString(m) }
func (*TensorDescription) ProtoMessage()               {}
func (*TensorDescription) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *TensorDescription) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *TensorDescription) GetAllocationDescription() *AllocationDescription {
	if m != nil {
		return m.AllocationDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorDescription)(nil), "tensorflow.TensorDescription")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/tensor_description.proto", fileDescriptor12)
}

var fileDescriptor12 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x2a, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2b, 0x4a, 0xcc,
	0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0xd6, 0x87, 0xc8, 0xc4, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16,
	0x94, 0x64, 0xe6, 0xe7, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0xf4, 0x48, 0xa9,
	0xe2, 0xd1, 0x5f, 0x59, 0x90, 0x5a, 0x0c, 0xd1, 0x22, 0xa5, 0x43, 0xd0, 0x9a, 0xe2, 0x8c, 0xc4,
	0x82, 0x54, 0xa8, 0x6a, 0x33, 0xdc, 0xaa, 0x13, 0x73, 0x72, 0xf2, 0x93, 0x13, 0x41, 0x8e, 0xc1,
	0x74, 0x98, 0xd2, 0x59, 0x46, 0x2e, 0xc1, 0x10, 0xb0, 0x56, 0x17, 0x84, 0x9c, 0x90, 0x16, 0x17,
	0x6b, 0x0a, 0xc8, 0x2d, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x22, 0x7a, 0x08, 0xd3, 0xf5,
	0x5c, 0x12, 0x4b, 0x12, 0x43, 0x80, 0x72, 0x41, 0x10, 0x25, 0x42, 0x46, 0x5c, 0xac, 0x60, 0x87,
	0x48, 0x30, 0x01, 0xd5, 0x72, 0x1b, 0xc9, 0x20, 0xab, 0x85, 0x98, 0x1c, 0x0c, 0x92, 0x0e, 0x00,
	0x59, 0x17, 0x04, 0x51, 0x2a, 0x14, 0xc1, 0x25, 0x86, 0xdd, 0x55, 0x12, 0x2c, 0x60, 0x43, 0x14,
	0x91, 0x0d, 0x71, 0x84, 0xab, 0x44, 0x72, 0x62, 0x90, 0x68, 0x22, 0x36, 0x61, 0x27, 0x0b, 0x2e,
	0x89, 0xfc, 0xa2, 0x74, 0x64, 0xed, 0xf0, 0x80, 0x70, 0x12, 0xc7, 0xf0, 0x28, 0xd8, 0x51, 0xc5,
	0x01, 0x8c, 0x3f, 0x18, 0x19, 0x93, 0xd8, 0xc0, 0x01, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x26, 0xb4, 0x9a, 0xb7, 0xdf, 0x01, 0x00, 0x00,
}
