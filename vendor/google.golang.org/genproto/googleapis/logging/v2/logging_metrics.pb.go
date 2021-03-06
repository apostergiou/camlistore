// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/logging/v2/logging_metrics.proto
// DO NOT EDIT!

package v2 // import "google.golang.org/genproto/googleapis/logging/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf5 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Stackdriver Logging API version.
type LogMetric_ApiVersion int32

const (
	// Stackdriver Logging API v2.
	LogMetric_V2 LogMetric_ApiVersion = 0
	// Stackdriver Logging API v1.
	LogMetric_V1 LogMetric_ApiVersion = 1
)

var LogMetric_ApiVersion_name = map[int32]string{
	0: "V2",
	1: "V1",
}
var LogMetric_ApiVersion_value = map[string]int32{
	"V2": 0,
	"V1": 1,
}

func (x LogMetric_ApiVersion) String() string {
	return proto.EnumName(LogMetric_ApiVersion_name, int32(x))
}
func (LogMetric_ApiVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

// Describes a logs-based metric.  The value of the metric is the
// number of log entries that match a logs filter.
type LogMetric struct {
	// Required. The client-assigned metric identifier. Example:
	// `"severe_errors"`.  Metric identifiers are limited to 100
	// characters and can include only the following characters: `A-Z`,
	// `a-z`, `0-9`, and the special characters `_-.,+!*',()%/`.  The
	// forward-slash character (`/`) denotes a hierarchy of name pieces,
	// and it cannot be the first character of the name.  The '%' character
	// is used to URL encode unsafe and reserved characters and must be
	// followed by two hexadecimal digits according to RFC 1738.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optional. A description of this metric, which is used in documentation.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Required. An [advanced logs filter](/logging/docs/view/advanced_filters).
	// Example: `"resource.type=gae_app AND severity>=ERROR"`.
	Filter string `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
	// Output only. The API version that created or updated this metric.
	// The version also dictates the syntax of the filter expression. When a value
	// for this field is missing, the default value of V2 should be assumed.
	Version LogMetric_ApiVersion `protobuf:"varint,4,opt,name=version,enum=google.logging.v2.LogMetric_ApiVersion" json:"version,omitempty"`
}

func (m *LogMetric) Reset()                    { *m = LogMetric{} }
func (m *LogMetric) String() string            { return proto.CompactTextString(m) }
func (*LogMetric) ProtoMessage()               {}
func (*LogMetric) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// The parameters to ListLogMetrics.
type ListLogMetricsRequest struct {
	// Required. The resource name containing the metrics.
	// Example: `"projects/my-project-id"`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListLogMetricsRequest) Reset()                    { *m = ListLogMetricsRequest{} }
func (m *ListLogMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsRequest) ProtoMessage()               {}
func (*ListLogMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// Result returned from ListLogMetrics.
type ListLogMetricsResponse struct {
	// A list of logs-based metrics.
	Metrics []*LogMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// If there might be more results than appear in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogMetricsResponse) Reset()                    { *m = ListLogMetricsResponse{} }
func (m *ListLogMetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsResponse) ProtoMessage()               {}
func (*ListLogMetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ListLogMetricsResponse) GetMetrics() []*LogMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// The parameters to GetLogMetric.
type GetLogMetricRequest struct {
	// The resource name of the desired metric.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *GetLogMetricRequest) Reset()                    { *m = GetLogMetricRequest{} }
func (m *GetLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogMetricRequest) ProtoMessage()               {}
func (*GetLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// The parameters to CreateLogMetric.
type CreateLogMetricRequest struct {
	// The resource name of the project in which to create the metric.
	// Example: `"projects/my-project-id"`.
	//
	// The new metric must be provided in the request.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The new logs-based metric, which must not have an identifier that
	// already exists.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *CreateLogMetricRequest) Reset()                    { *m = CreateLogMetricRequest{} }
func (m *CreateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateLogMetricRequest) ProtoMessage()               {}
func (*CreateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CreateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to UpdateLogMetric.
type UpdateLogMetricRequest struct {
	// The resource name of the metric to update.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	//
	// The updated metric must be provided in the request and have the
	// same identifier that is specified in `metricName`.
	// If the metric does not exist, it is created.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
	// The updated metric, whose name must be the same as the
	// metric identifier in `metricName`. If `metricName` does not
	// exist, then a new metric is created.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *UpdateLogMetricRequest) Reset()                    { *m = UpdateLogMetricRequest{} }
func (m *UpdateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLogMetricRequest) ProtoMessage()               {}
func (*UpdateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *UpdateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to DeleteLogMetric.
type DeleteLogMetricRequest struct {
	// The resource name of the metric to delete.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *DeleteLogMetricRequest) Reset()                    { *m = DeleteLogMetricRequest{} }
func (m *DeleteLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogMetricRequest) ProtoMessage()               {}
func (*DeleteLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func init() {
	proto.RegisterType((*LogMetric)(nil), "google.logging.v2.LogMetric")
	proto.RegisterType((*ListLogMetricsRequest)(nil), "google.logging.v2.ListLogMetricsRequest")
	proto.RegisterType((*ListLogMetricsResponse)(nil), "google.logging.v2.ListLogMetricsResponse")
	proto.RegisterType((*GetLogMetricRequest)(nil), "google.logging.v2.GetLogMetricRequest")
	proto.RegisterType((*CreateLogMetricRequest)(nil), "google.logging.v2.CreateLogMetricRequest")
	proto.RegisterType((*UpdateLogMetricRequest)(nil), "google.logging.v2.UpdateLogMetricRequest")
	proto.RegisterType((*DeleteLogMetricRequest)(nil), "google.logging.v2.DeleteLogMetricRequest")
	proto.RegisterEnum("google.logging.v2.LogMetric_ApiVersion", LogMetric_ApiVersion_name, LogMetric_ApiVersion_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MetricsServiceV2 service

type MetricsServiceV2Client interface {
	// Lists logs-based metrics.
	ListLogMetrics(ctx context.Context, in *ListLogMetricsRequest, opts ...grpc.CallOption) (*ListLogMetricsResponse, error)
	// Gets a logs-based metric.
	GetLogMetric(ctx context.Context, in *GetLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Creates a logs-based metric.
	CreateLogMetric(ctx context.Context, in *CreateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Creates or updates a logs-based metric.
	UpdateLogMetric(ctx context.Context, in *UpdateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Deletes a logs-based metric.
	DeleteLogMetric(ctx context.Context, in *DeleteLogMetricRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error)
}

type metricsServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceV2Client(cc *grpc.ClientConn) MetricsServiceV2Client {
	return &metricsServiceV2Client{cc}
}

func (c *metricsServiceV2Client) ListLogMetrics(ctx context.Context, in *ListLogMetricsRequest, opts ...grpc.CallOption) (*ListLogMetricsResponse, error) {
	out := new(ListLogMetricsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/ListLogMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) GetLogMetric(ctx context.Context, in *GetLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) {
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/GetLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) CreateLogMetric(ctx context.Context, in *CreateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) {
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/CreateLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) UpdateLogMetric(ctx context.Context, in *UpdateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) {
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/UpdateLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceV2Client) DeleteLogMetric(ctx context.Context, in *DeleteLogMetricRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error) {
	out := new(google_protobuf5.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/DeleteLogMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetricsServiceV2 service

type MetricsServiceV2Server interface {
	// Lists logs-based metrics.
	ListLogMetrics(context.Context, *ListLogMetricsRequest) (*ListLogMetricsResponse, error)
	// Gets a logs-based metric.
	GetLogMetric(context.Context, *GetLogMetricRequest) (*LogMetric, error)
	// Creates a logs-based metric.
	CreateLogMetric(context.Context, *CreateLogMetricRequest) (*LogMetric, error)
	// Creates or updates a logs-based metric.
	UpdateLogMetric(context.Context, *UpdateLogMetricRequest) (*LogMetric, error)
	// Deletes a logs-based metric.
	DeleteLogMetric(context.Context, *DeleteLogMetricRequest) (*google_protobuf5.Empty, error)
}

func RegisterMetricsServiceV2Server(s *grpc.Server, srv MetricsServiceV2Server) {
	s.RegisterService(&_MetricsServiceV2_serviceDesc, srv)
}

func _MetricsServiceV2_ListLogMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).ListLogMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/ListLogMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).ListLogMetrics(ctx, req.(*ListLogMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_GetLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).GetLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/GetLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).GetLogMetric(ctx, req.(*GetLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_CreateLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).CreateLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/CreateLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).CreateLogMetric(ctx, req.(*CreateLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_UpdateLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).UpdateLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/UpdateLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).UpdateLogMetric(ctx, req.(*UpdateLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsServiceV2_DeleteLogMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceV2Server).DeleteLogMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/DeleteLogMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceV2Server).DeleteLogMetric(ctx, req.(*DeleteLogMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.MetricsServiceV2",
	HandlerType: (*MetricsServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLogMetrics",
			Handler:    _MetricsServiceV2_ListLogMetrics_Handler,
		},
		{
			MethodName: "GetLogMetric",
			Handler:    _MetricsServiceV2_GetLogMetric_Handler,
		},
		{
			MethodName: "CreateLogMetric",
			Handler:    _MetricsServiceV2_CreateLogMetric_Handler,
		},
		{
			MethodName: "UpdateLogMetric",
			Handler:    _MetricsServiceV2_UpdateLogMetric_Handler,
		},
		{
			MethodName: "DeleteLogMetric",
			Handler:    _MetricsServiceV2_DeleteLogMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google.golang.org/genproto/googleapis/logging/v2/logging_metrics.proto",
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/logging/v2/logging_metrics.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0x7e, 0x97, 0x8f, 0xf2, 0x72, 0x50, 0x8a, 0x63, 0x68, 0x9a, 0x8a, 0xa1, 0xee, 0x05, 0x14,
	0x2e, 0x76, 0x70, 0x31, 0x24, 0x9a, 0x78, 0x01, 0x7e, 0xc5, 0x04, 0x0d, 0x29, 0xca, 0x85, 0x5e,
	0x34, 0xcb, 0x72, 0x3a, 0x8e, 0xb4, 0x3b, 0xe3, 0xce, 0xd0, 0x00, 0x86, 0x1b, 0xc3, 0x9d, 0x89,
	0x17, 0xfa, 0x6f, 0xfc, 0x1b, 0xfa, 0x13, 0xfc, 0x21, 0x66, 0x67, 0x76, 0x4b, 0x6d, 0x57, 0x5a,
	0xb9, 0x69, 0x67, 0xce, 0xd7, 0xf3, 0x9c, 0x73, 0x9e, 0x4e, 0xe1, 0x29, 0x13, 0x82, 0xb5, 0xd0,
	0x63, 0xa2, 0x15, 0x44, 0xcc, 0x13, 0x31, 0xa3, 0x0c, 0x23, 0x19, 0x0b, 0x2d, 0xa8, 0x75, 0x05,
	0x92, 0x2b, 0xda, 0x12, 0x8c, 0xf1, 0x88, 0xd1, 0x8e, 0x9f, 0x1d, 0x1b, 0x6d, 0xd4, 0x31, 0x0f,
	0x95, 0x67, 0x62, 0xc9, 0x8d, 0xb4, 0x4e, 0xea, 0xf5, 0x3a, 0x7e, 0xe5, 0xf9, 0x68, 0xa5, 0x03,
	0xc9, 0xa9, 0xc2, 0xb8, 0xc3, 0x43, 0x0c, 0x45, 0xd4, 0xe4, 0x8c, 0x06, 0x51, 0x24, 0x74, 0xa0,
	0xb9, 0x88, 0xd2, 0xea, 0x95, 0x75, 0xc6, 0xf5, 0xbb, 0xa3, 0x7d, 0x2f, 0x14, 0x6d, 0x6a, 0xcb,
	0x51, 0xe3, 0xd8, 0x3f, 0x6a, 0x52, 0xa9, 0x4f, 0x24, 0x2a, 0x8a, 0x6d, 0xa9, 0x4f, 0xec, 0xa7,
	0x4d, 0x72, 0xbf, 0x3b, 0x30, 0xbd, 0x2d, 0xd8, 0x0b, 0xc3, 0x93, 0x10, 0x98, 0x88, 0x82, 0x36,
	0x96, 0x9d, 0xaa, 0x53, 0x9b, 0xae, 0x9b, 0x33, 0xa9, 0xc2, 0xcc, 0x01, 0xaa, 0x30, 0xe6, 0x32,
	0x01, 0x2b, 0x8f, 0x19, 0x57, 0xaf, 0x89, 0x94, 0xa0, 0xd0, 0xe4, 0x2d, 0x8d, 0x71, 0x79, 0xdc,
	0x38, 0xd3, 0x1b, 0xd9, 0x84, 0xa9, 0x0e, 0xc6, 0x2a, 0xc9, 0x9a, 0xa8, 0x3a, 0xb5, 0x59, 0x7f,
	0xd9, 0x1b, 0x18, 0x80, 0xd7, 0x05, 0xf7, 0x36, 0x25, 0xdf, 0xb3, 0xe1, 0xf5, 0x2c, 0xcf, 0x5d,
	0x00, 0xb8, 0x30, 0x93, 0x02, 0x8c, 0xed, 0xf9, 0x73, 0xff, 0x99, 0xef, 0xbb, 0x73, 0x8e, 0x7b,
	0x08, 0xf3, 0xdb, 0x5c, 0xe9, 0x6e, 0x09, 0x55, 0xc7, 0x0f, 0x47, 0xa8, 0x74, 0xc2, 0x48, 0x06,
	0x31, 0x46, 0x3a, 0xed, 0x24, 0xbd, 0x91, 0xdb, 0x00, 0x32, 0x60, 0xd8, 0xd0, 0xe2, 0x10, 0xb3,
	0x56, 0xa6, 0x13, 0xcb, 0xab, 0xc4, 0x40, 0x6e, 0x81, 0xb9, 0x34, 0x14, 0x3f, 0x45, 0xd3, 0xcb,
	0x64, 0xfd, 0xff, 0xc4, 0xb0, 0xcb, 0x4f, 0xd1, 0x3d, 0x86, 0x52, 0x3f, 0x98, 0x92, 0x22, 0x52,
	0x48, 0x36, 0x60, 0x2a, 0xdd, 0x73, 0xd9, 0xa9, 0x8e, 0xd7, 0x66, 0xfc, 0x85, 0xcb, 0xfa, 0xac,
	0x67, 0xc1, 0x64, 0x09, 0x8a, 0x11, 0x1e, 0xeb, 0xc6, 0x00, 0xa5, 0xeb, 0x89, 0x79, 0x27, 0xa3,
	0xe5, 0x6e, 0xc0, 0xcd, 0x67, 0x78, 0x01, 0x9c, 0x35, 0xb9, 0x08, 0x33, 0xb6, 0x52, 0xa3, 0x67,
	0x67, 0x60, 0x4d, 0x2f, 0x83, 0x36, 0xba, 0x4d, 0x28, 0x3d, 0x8a, 0x31, 0xd0, 0x38, 0x90, 0xfa,
	0xb7, 0xf9, 0xdc, 0x83, 0x82, 0xcd, 0x37, 0x44, 0x86, 0x35, 0x92, 0xc6, 0xba, 0x02, 0x4a, 0xaf,
	0xe5, 0x41, 0x1e, 0xce, 0x30, 0x8a, 0x57, 0x04, 0xbc, 0x0f, 0xa5, 0xc7, 0xd8, 0xc2, 0x2b, 0x00,
	0xfa, 0x3f, 0x27, 0x61, 0x2e, 0xdd, 0xdf, 0xae, 0xfd, 0x3d, 0xed, 0xf9, 0xe4, 0x8b, 0x03, 0xb3,
	0x7f, 0xee, 0x96, 0xd4, 0xf2, 0x88, 0xe4, 0x69, 0xad, 0xb2, 0x32, 0x42, 0xa4, 0x15, 0x8a, 0xbb,
	0xfc, 0xe9, 0xc7, 0xaf, 0x6f, 0x63, 0x77, 0xc8, 0x62, 0xf2, 0x44, 0x7c, 0xb4, 0x33, 0x7f, 0x28,
	0x63, 0xf1, 0x1e, 0x43, 0xad, 0xe8, 0xea, 0x19, 0xcd, 0x94, 0x71, 0xee, 0xc0, 0xb5, 0xde, 0x95,
	0x93, 0xa5, 0x1c, 0x90, 0x1c, 0x4d, 0x54, 0x2e, 0x9d, 0x9f, 0xeb, 0x19, 0xfc, 0x1a, 0x59, 0x32,
	0xf8, 0x3d, 0x83, 0xea, 0x21, 0x91, 0x71, 0xa0, 0xab, 0x67, 0xe4, 0xb3, 0x03, 0xc5, 0x3e, 0x05,
	0x91, 0xbc, 0x76, 0xf3, 0x55, 0x36, 0x84, 0x0c, 0x35, 0x64, 0x56, 0xdc, 0x61, 0xc3, 0x78, 0x90,
	0x6e, 0x9d, 0x7c, 0x75, 0xa0, 0xd8, 0xa7, 0xb3, 0x5c, 0x36, 0xf9, 0x5a, 0x1c, 0xc2, 0x66, 0xc3,
	0xb0, 0x59, 0xab, 0x8c, 0x38, 0x9a, 0x2e, 0xa9, 0x73, 0x07, 0x8a, 0x7d, 0x5a, 0xcc, 0x25, 0x95,
	0xaf, 0xd7, 0x4a, 0x29, 0x0b, 0xcd, 0x5e, 0x6a, 0xef, 0x49, 0xf2, 0x38, 0x67, 0x9b, 0x5a, 0x1d,
	0x91, 0xce, 0xd6, 0x5b, 0x98, 0x0f, 0x45, 0x7b, 0x10, 0x77, 0x6b, 0x76, 0xdb, 0x9e, 0x53, 0x29,
	0xee, 0x38, 0x6f, 0xd6, 0xfe, 0xf5, 0xcf, 0x6c, 0xbf, 0x60, 0x9c, 0xeb, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x46, 0x24, 0x21, 0x19, 0x07, 0x07, 0x00, 0x00,
}
