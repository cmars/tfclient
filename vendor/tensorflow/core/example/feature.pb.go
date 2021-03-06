// Code generated by protoc-gen-go.
// source: tensorflow/core/example/feature.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Containers to hold repeated fundamental values.
type BytesList struct {
	Value [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesList) Reset()                    { *m = BytesList{} }
func (m *BytesList) String() string            { return proto.CompactTextString(m) }
func (*BytesList) ProtoMessage()               {}
func (*BytesList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BytesList) GetValue() [][]byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type FloatList struct {
	Value []float32 `protobuf:"fixed32,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *FloatList) Reset()                    { *m = FloatList{} }
func (m *FloatList) String() string            { return proto.CompactTextString(m) }
func (*FloatList) ProtoMessage()               {}
func (*FloatList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *FloatList) GetValue() []float32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Int64List struct {
	Value []int64 `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *Int64List) Reset()                    { *m = Int64List{} }
func (m *Int64List) String() string            { return proto.CompactTextString(m) }
func (*Int64List) ProtoMessage()               {}
func (*Int64List) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Int64List) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

// Containers for non-sequential data.
type Feature struct {
	// Each feature can be exactly one kind.
	//
	// Types that are valid to be assigned to Kind:
	//	*Feature_BytesList
	//	*Feature_FloatList
	//	*Feature_Int64List
	Kind isFeature_Kind `protobuf_oneof:"kind"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type isFeature_Kind interface {
	isFeature_Kind()
}

type Feature_BytesList struct {
	BytesList *BytesList `protobuf:"bytes,1,opt,name=bytes_list,json=bytesList,oneof"`
}
type Feature_FloatList struct {
	FloatList *FloatList `protobuf:"bytes,2,opt,name=float_list,json=floatList,oneof"`
}
type Feature_Int64List struct {
	Int64List *Int64List `protobuf:"bytes,3,opt,name=int64_list,json=int64List,oneof"`
}

func (*Feature_BytesList) isFeature_Kind() {}
func (*Feature_FloatList) isFeature_Kind() {}
func (*Feature_Int64List) isFeature_Kind() {}

func (m *Feature) GetKind() isFeature_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Feature) GetBytesList() *BytesList {
	if x, ok := m.GetKind().(*Feature_BytesList); ok {
		return x.BytesList
	}
	return nil
}

func (m *Feature) GetFloatList() *FloatList {
	if x, ok := m.GetKind().(*Feature_FloatList); ok {
		return x.FloatList
	}
	return nil
}

func (m *Feature) GetInt64List() *Int64List {
	if x, ok := m.GetKind().(*Feature_Int64List); ok {
		return x.Int64List
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Feature) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Feature_OneofMarshaler, _Feature_OneofUnmarshaler, _Feature_OneofSizer, []interface{}{
		(*Feature_BytesList)(nil),
		(*Feature_FloatList)(nil),
		(*Feature_Int64List)(nil),
	}
}

func _Feature_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Feature)
	// kind
	switch x := m.Kind.(type) {
	case *Feature_BytesList:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BytesList); err != nil {
			return err
		}
	case *Feature_FloatList:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FloatList); err != nil {
			return err
		}
	case *Feature_Int64List:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Int64List); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Feature.Kind has unexpected type %T", x)
	}
	return nil
}

func _Feature_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Feature)
	switch tag {
	case 1: // kind.bytes_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BytesList)
		err := b.DecodeMessage(msg)
		m.Kind = &Feature_BytesList{msg}
		return true, err
	case 2: // kind.float_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FloatList)
		err := b.DecodeMessage(msg)
		m.Kind = &Feature_FloatList{msg}
		return true, err
	case 3: // kind.int64_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Int64List)
		err := b.DecodeMessage(msg)
		m.Kind = &Feature_Int64List{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Feature_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Feature)
	// kind
	switch x := m.Kind.(type) {
	case *Feature_BytesList:
		s := proto.Size(x.BytesList)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Feature_FloatList:
		s := proto.Size(x.FloatList)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Feature_Int64List:
		s := proto.Size(x.Int64List)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Features struct {
	// Map from feature name to feature.
	Feature map[string]*Feature `protobuf:"bytes,1,rep,name=feature" json:"feature,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Features) Reset()                    { *m = Features{} }
func (m *Features) String() string            { return proto.CompactTextString(m) }
func (*Features) ProtoMessage()               {}
func (*Features) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Features) GetFeature() map[string]*Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

// Containers for sequential data.
//
// A FeatureList contains lists of Features.  These may hold zero or more
// Feature values.
//
// FeatureLists are organized into categories by name.  The FeatureLists message
// contains the mapping from name to FeatureList.
//
type FeatureList struct {
	Feature []*Feature `protobuf:"bytes,1,rep,name=feature" json:"feature,omitempty"`
}

func (m *FeatureList) Reset()                    { *m = FeatureList{} }
func (m *FeatureList) String() string            { return proto.CompactTextString(m) }
func (*FeatureList) ProtoMessage()               {}
func (*FeatureList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *FeatureList) GetFeature() []*Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

type FeatureLists struct {
	// Map from feature name to feature list.
	FeatureList map[string]*FeatureList `protobuf:"bytes,1,rep,name=feature_list,json=featureList" json:"feature_list,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FeatureLists) Reset()                    { *m = FeatureLists{} }
func (m *FeatureLists) String() string            { return proto.CompactTextString(m) }
func (*FeatureLists) ProtoMessage()               {}
func (*FeatureLists) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *FeatureLists) GetFeatureList() map[string]*FeatureList {
	if m != nil {
		return m.FeatureList
	}
	return nil
}

func init() {
	proto.RegisterType((*BytesList)(nil), "tensorflow.BytesList")
	proto.RegisterType((*FloatList)(nil), "tensorflow.FloatList")
	proto.RegisterType((*Int64List)(nil), "tensorflow.Int64List")
	proto.RegisterType((*Feature)(nil), "tensorflow.Feature")
	proto.RegisterType((*Features)(nil), "tensorflow.Features")
	proto.RegisterType((*FeatureList)(nil), "tensorflow.FeatureList")
	proto.RegisterType((*FeatureLists)(nil), "tensorflow.FeatureLists")
}

func init() { proto.RegisterFile("tensorflow/core/example/feature.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0x49, 0x3e, 0x5b, 0x73, 0x53, 0xa1, 0x8c, 0xff, 0x4a, 0x57, 0x36, 0x50, 0x68,
	0xc1, 0xb6, 0x50, 0xa5, 0x88, 0xba, 0x0a, 0x58, 0x14, 0x0a, 0x96, 0x6c, 0x5c, 0x4a, 0xaa, 0x13,
	0x09, 0x8d, 0x99, 0x92, 0x4c, 0xd5, 0xbe, 0x89, 0x2f, 0xe2, 0xc2, 0x37, 0x73, 0x69, 0x26, 0x99,
	0x4c, 0xa7, 0x4d, 0xdc, 0xe5, 0xce, 0x9c, 0xdf, 0x9d, 0x73, 0x6e, 0x2e, 0xb4, 0x19, 0x09, 0x63,
	0x1a, 0x79, 0x01, 0x7d, 0x1f, 0x3c, 0xd1, 0x88, 0x0c, 0xc8, 0x87, 0xfb, 0xba, 0x08, 0xc8, 0xc0,
	0x23, 0x2e, 0x5b, 0x46, 0xa4, 0xbf, 0x88, 0x28, 0xa3, 0x18, 0xd6, 0x32, 0xab, 0x05, 0x86, 0xbd,
	0x62, 0x24, 0x9e, 0xf8, 0x31, 0xc3, 0x07, 0xb0, 0xf3, 0xe6, 0x06, 0x4b, 0xd2, 0x40, 0x27, 0x7a,
	0xa7, 0xe6, 0x64, 0x85, 0xd5, 0x06, 0x63, 0x1c, 0x50, 0x97, 0xa5, 0x92, 0x86, 0x2a, 0xd1, 0x6c,
	0xad, 0x8e, 0x14, 0xd9, 0x5d, 0xc8, 0x46, 0xe7, 0x45, 0x99, 0xae, 0xca, 0xbe, 0x11, 0x54, 0xc7,
	0x99, 0x1d, 0x3c, 0x02, 0x98, 0xf1, 0xc7, 0x1f, 0x83, 0x84, 0x49, 0xa4, 0xa8, 0x63, 0x0e, 0x0f,
	0xfb, 0x6b, 0x77, 0x7d, 0x69, 0xed, 0xf6, 0x9f, 0x63, 0xcc, 0xa4, 0xcf, 0x84, 0xf3, 0xb8, 0xa3,
	0x8c, 0xd3, 0x8a, 0x9c, 0xf4, 0xcb, 0x39, 0x4f, 0x9a, 0x4f, 0x38, 0x9f, 0x5b, 0xcc, 0x38, 0xbd,
	0xc8, 0xc9, 0x00, 0x9c, 0xf3, 0xf3, 0xc2, 0xae, 0xc0, 0xff, 0xb9, 0x1f, 0x3e, 0x5b, 0x9f, 0x08,
	0x76, 0x85, 0xf7, 0x18, 0x5f, 0x41, 0x55, 0x8c, 0x35, 0x0d, 0x69, 0x0e, 0x5b, 0x1b, 0x0e, 0x84,
	0x2c, 0xff, 0xb8, 0x09, 0x59, 0xb4, 0x72, 0x72, 0xa2, 0x79, 0x0f, 0x35, 0xf5, 0x02, 0xd7, 0x41,
	0x9f, 0x93, 0x55, 0x3a, 0x02, 0xc3, 0xe1, 0x9f, 0xb8, 0x9b, 0x4f, 0x30, 0x8b, 0xb7, 0x5f, 0xd2,
	0x5c, 0x8c, 0xf4, 0x52, 0xbb, 0x40, 0xd6, 0x35, 0x98, 0xe2, 0x34, 0x4d, 0xda, 0xdb, 0x36, 0x57,
	0xca, 0xe7, 0x1a, 0xeb, 0x0b, 0x49, 0x3f, 0x1c, 0x8f, 0xf1, 0x04, 0x6a, 0xe2, 0x2e, 0xff, 0x37,
	0xbc, 0x49, 0xb7, 0xa4, 0x49, 0xaa, 0x57, 0x8b, 0x2c, 0xa9, 0xe9, 0xad, 0x4f, 0x9a, 0x0f, 0x50,
	0xdf, 0x16, 0x94, 0x24, 0xee, 0x6d, 0x26, 0x3e, 0xfe, 0xe3, 0x31, 0x25, 0xb5, 0x7d, 0x0a, 0x47,
	0x34, 0x7a, 0x51, 0x85, 0x62, 0xe3, 0xed, 0x3d, 0x41, 0x4c, 0xf9, 0xc6, 0xc7, 0x53, 0xf4, 0x83,
	0xd0, 0xac, 0x92, 0xae, 0xff, 0xd9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x38, 0x1c, 0x54,
	0x27, 0x03, 0x00, 0x00,
}
