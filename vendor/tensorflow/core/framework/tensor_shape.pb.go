// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/tensor_shape.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Dimensions of a tensor.
type TensorShapeProto struct {
	// Dimensions of the tensor, such as {"input", 30}, {"output", 40}
	// for a 30 x 40 2D tensor.  If an entry has size -1, this
	// corresponds to a dimension of unknown size. The names are
	// optional.
	//
	// The order of entries in "dim" matters: It indicates the layout of the
	// values in the tensor in-memory representation.
	//
	// The first entry in "dim" is the outermost dimension used to layout the
	// values, the last entry is the innermost dimension.  This matches the
	// in-memory layout of RowMajor Eigen tensors.
	//
	// If "dim.size()" > 0, "unknown_rank" must be false.
	Dim []*TensorShapeProto_Dim `protobuf:"bytes,2,rep,name=dim" json:"dim,omitempty"`
	// If true, the number of dimensions in the shape is unknown.
	//
	// If true, "dim.size()" must be 0.
	UnknownRank bool `protobuf:"varint,3,opt,name=unknown_rank,json=unknownRank" json:"unknown_rank,omitempty"`
}

func (m *TensorShapeProto) Reset()                    { *m = TensorShapeProto{} }
func (m *TensorShapeProto) String() string            { return proto.CompactTextString(m) }
func (*TensorShapeProto) ProtoMessage()               {}
func (*TensorShapeProto) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *TensorShapeProto) GetDim() []*TensorShapeProto_Dim {
	if m != nil {
		return m.Dim
	}
	return nil
}

// One dimension of the tensor.
type TensorShapeProto_Dim struct {
	// Size of the tensor in that dimension.
	// This value must be >= -1, but values of -1 are reserved for "unknown"
	// shapes (values of -1 mean "unknown" dimension).  Certain wrappers
	// that work with TensorShapeProto may fail at runtime when deserializing
	// a TensorShapeProto containing a dim value of -1.
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Optional name of the tensor dimension.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *TensorShapeProto_Dim) Reset()                    { *m = TensorShapeProto_Dim{} }
func (m *TensorShapeProto_Dim) String() string            { return proto.CompactTextString(m) }
func (*TensorShapeProto_Dim) ProtoMessage()               {}
func (*TensorShapeProto_Dim) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0, 0} }

func init() {
	proto.RegisterType((*TensorShapeProto)(nil), "tensorflow.TensorShapeProto")
	proto.RegisterType((*TensorShapeProto_Dim)(nil), "tensorflow.TensorShapeProto.Dim")
}

func init() { proto.RegisterFile("tensorflow/core/framework/tensor_shape.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2b, 0x4a, 0xcc,
	0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0xd6, 0x87, 0xc8, 0xc4, 0x17, 0x67, 0x24, 0x16, 0xa4, 0xea, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0x54, 0x2b, 0xcd, 0x60, 0xe4, 0x12, 0x08, 0x01, 0x73,
	0x83, 0x41, 0x2a, 0x02, 0xc0, 0x0a, 0x8c, 0xb8, 0x98, 0x53, 0x32, 0x73, 0x25, 0x98, 0x14, 0x98,
	0x35, 0xb8, 0x8d, 0x14, 0xf4, 0x10, 0xca, 0xf5, 0xd0, 0x95, 0xea, 0xb9, 0x64, 0xe6, 0x06, 0x81,
	0x14, 0x0b, 0x29, 0x72, 0xf1, 0x94, 0xe6, 0x65, 0xe7, 0xe5, 0x97, 0xe7, 0xc5, 0x17, 0x25, 0xe6,
	0x65, 0x4b, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x04, 0x71, 0x43, 0xc5, 0x82, 0x80, 0x42, 0x52, 0xba,
	0x5c, 0xcc, 0x40, 0xe5, 0x42, 0x42, 0x5c, 0x2c, 0xc5, 0x99, 0x55, 0xa9, 0x12, 0x8c, 0x40, 0x15,
	0xcc, 0x41, 0x60, 0x36, 0x48, 0x2c, 0x0f, 0xe8, 0x62, 0xa0, 0x95, 0x8c, 0x1a, 0x9c, 0x41, 0x60,
	0xb6, 0x93, 0x11, 0x97, 0x44, 0x7e, 0x51, 0x3a, 0xb2, 0xed, 0x70, 0x5f, 0x39, 0x09, 0xa2, 0x3b,
	0xa4, 0x38, 0x80, 0xf1, 0x07, 0x23, 0x63, 0x12, 0x1b, 0xd8, 0x87, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x34, 0xe4, 0x72, 0x62, 0x11, 0x01, 0x00, 0x00,
}
