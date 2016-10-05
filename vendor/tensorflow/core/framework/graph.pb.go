// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/graph.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents the graph of operations
type GraphDef struct {
	Node []*NodeDef `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
	// Compatibility versions of the graph.  See core/public/version.h for version
	// history.  The GraphDef version is distinct from the TensorFlow version, and
	// each release of TensorFlow will support a range of GraphDef versions.
	Versions *VersionDef `protobuf:"bytes,4,opt,name=versions" json:"versions,omitempty"`
	// Deprecated single version field; use versions above instead.  Since all
	// GraphDef changes before "versions" was introduced were forward
	// compatible, this field is entirely ignored.
	Version int32 `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	// EXPERIMENTAL. DO NOT USE OR DEPEND ON THIS YET.
	//
	// "library" provides user-defined functions.
	//
	// Naming:
	//   * library.function.name are in a flat namespace.
	//     NOTE: We may need to change it to be hierarchical to support
	//     different orgs. E.g.,
	//     { "/google/nn", { ... }},
	//     { "/google/vision", { ... }}
	//     { "/org_foo/module_bar", { ... }}
	//     map<string, FunctionDefLib> named_lib;
	//   * If node[i].op is the name of one function in "library",
	//     node[i] is deemed as a function call. Otherwise, node[i].op
	//     must be a primitive operation supported by the runtime.
	//
	//
	// Function call semantics:
	//
	//   * The callee may start execution as soon as some of its inputs
	//     are ready. The caller may want to use Tuple() mechanism to
	//     ensure all inputs are ready in the same time.
	//
	//   * The consumer of return values may start executing as soon as
	//     the return values the consumer depends on are ready.  The
	//     consumer may want to use Tuple() mechanism to ensure the
	//     consumer does not start until all return values of the callee
	//     function are ready.
	Library *FunctionDefLibrary `protobuf:"bytes,2,opt,name=library" json:"library,omitempty"`
}

func (m *GraphDef) Reset()                    { *m = GraphDef{} }
func (m *GraphDef) String() string            { return proto.CompactTextString(m) }
func (*GraphDef) ProtoMessage()               {}
func (*GraphDef) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *GraphDef) GetNode() []*NodeDef {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *GraphDef) GetVersions() *VersionDef {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *GraphDef) GetLibrary() *FunctionDefLibrary {
	if m != nil {
		return m.Library
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphDef)(nil), "tensorflow.GraphDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/graph.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2b, 0x4a, 0xcc,
	0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0xd6, 0x4f, 0x2f, 0x4a, 0x2c, 0xc8, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x42, 0x28, 0x93, 0xd2, 0xc0, 0xad, 0x25, 0x2f, 0x3f, 0x25, 0x35, 0x3e, 0x25,
	0x35, 0x0d, 0xa2, 0x0b, 0x9f, 0xca, 0xb4, 0xd2, 0xbc, 0xe4, 0x92, 0xcc, 0xfc, 0x3c, 0xc2, 0x2a,
	0xcb, 0x52, 0x8b, 0x8a, 0x81, 0x0a, 0x8b, 0x21, 0x2a, 0x95, 0xf6, 0x33, 0x72, 0x71, 0xb8, 0x83,
	0x5c, 0xe6, 0x92, 0x9a, 0x26, 0xa4, 0xce, 0xc5, 0x02, 0xb2, 0x52, 0x82, 0x51, 0x81, 0x59, 0x83,
	0xdb, 0x48, 0x58, 0x0f, 0x61, 0x8a, 0x9e, 0x1f, 0x50, 0x1c, 0xa8, 0x24, 0x08, 0xac, 0x40, 0xc8,
	0x88, 0x8b, 0x03, 0x66, 0x8e, 0x04, 0x8b, 0x02, 0x23, 0x50, 0xb1, 0x18, 0xb2, 0xe2, 0x30, 0x88,
	0x1c, 0x48, 0x3d, 0x5c, 0x9d, 0x90, 0x0c, 0x17, 0x3b, 0x94, 0x2d, 0xc1, 0x0c, 0xd4, 0xc2, 0xea,
	0xc4, 0x24, 0xc1, 0x18, 0x04, 0x13, 0x12, 0xb2, 0xe0, 0x62, 0xcf, 0xc9, 0x4c, 0x2a, 0x4a, 0x2c,
	0xaa, 0x94, 0x60, 0x02, 0x1b, 0x28, 0x87, 0x6c, 0xa0, 0x1b, 0xd4, 0x7b, 0x40, 0x13, 0x7d, 0x20,
	0xaa, 0x82, 0x60, 0xca, 0x9d, 0x74, 0xb8, 0x24, 0xf2, 0x8b, 0xd2, 0x91, 0x55, 0xc3, 0x3d, 0xeb,
	0xc4, 0x0d, 0xf6, 0x5a, 0x00, 0xc8, 0xa7, 0xc5, 0x01, 0x8c, 0x3f, 0x18, 0x19, 0x93, 0xd8, 0xc0,
	0xde, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x84, 0xf3, 0xe9, 0xda, 0xa9, 0x01, 0x00, 0x00,
}
