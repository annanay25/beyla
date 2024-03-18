// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package grpc

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tpConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_tpGoroutineMetadata struct {
	Parent    uint64
	Timestamp uint64
}

type bpf_tpGrpcClientFuncInvocationT struct {
	StartMonotimeNs uint64
	Cc              uint64
	Method          uint64
	MethodLen       uint64
	Tp              bpf_tpTpInfoT
	Flags           uint64
}

type bpf_tpGrpcFramerFuncInvocationT struct {
	FramerPtr uint64
	Tp        bpf_tpTpInfoT
	Offset    int64
}

type bpf_tpGrpcSrvFuncInvocationT struct {
	StartMonotimeNs uint64
	Stream          uint64
	Tp              bpf_tpTpInfoT
}

type bpf_tpTpInfoPidT struct {
	Tp    bpf_tpTpInfoT
	Pid   uint32
	Valid uint8
	_     [3]byte
}

type bpf_tpTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Ts       uint64
	Flags    uint8
	_        [7]byte
}

// loadBpf_tp returns the embedded CollectionSpec for bpf_tp.
func loadBpf_tp() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tpBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp: %w", err)
	}

	return spec, err
}

// loadBpf_tpObjects loads bpf_tp and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tpObjects
//	*bpf_tpPrograms
//	*bpf_tpMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tpObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tpSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpSpecs struct {
	bpf_tpProgramSpecs
	bpf_tpMapSpecs
}

// bpf_tpSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpProgramSpecs struct {
	UprobeClientConnClose               *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Close"`
	UprobeClientConnInvoke              *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturn        *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke_return"`
	UprobeClientConnNewStream           *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_NewStream"`
	UprobeGrpcFramerWriteHeaders        *ebpf.ProgramSpec `ebpf:"uprobe_grpcFramerWriteHeaders"`
	UprobeGrpcFramerWriteHeadersReturns *ebpf.ProgramSpec `ebpf:"uprobe_grpcFramerWriteHeaders_returns"`
	UprobeServerHandleStream            *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturn      *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream_return"`
	UprobeTransportHttp2ClientNewStream *ebpf.ProgramSpec `ebpf:"uprobe_transport_http2Client_NewStream"`
	UprobeTransportWriteStatus          *ebpf.ProgramSpec `ebpf:"uprobe_transport_writeStatus"`
}

// bpf_tpMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpMapSpecs struct {
	Events                       *ebpf.MapSpec `ebpf:"events"`
	GoTraceMap                   *ebpf.MapSpec `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap    *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	GrpcFramerInvocationMap      *ebpf.MapSpec `ebpf:"grpc_framer_invocation_map"`
	OngoingGoroutines            *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests    *ebpf.MapSpec `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites      *ebpf.MapSpec `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcRequestStatus     *ebpf.MapSpec `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests    *ebpf.MapSpec `ebpf:"ongoing_grpc_server_requests"`
	OngoingHttpServerConnections *ebpf.MapSpec `ebpf:"ongoing_http_server_connections"`
	OngoingStreams               *ebpf.MapSpec `ebpf:"ongoing_streams"`
	TraceMap                     *ebpf.MapSpec `ebpf:"trace_map"`
}

// bpf_tpObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpObjects struct {
	bpf_tpPrograms
	bpf_tpMaps
}

func (o *bpf_tpObjects) Close() error {
	return _Bpf_tpClose(
		&o.bpf_tpPrograms,
		&o.bpf_tpMaps,
	)
}

// bpf_tpMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpMaps struct {
	Events                       *ebpf.Map `ebpf:"events"`
	GoTraceMap                   *ebpf.Map `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap    *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	GrpcFramerInvocationMap      *ebpf.Map `ebpf:"grpc_framer_invocation_map"`
	OngoingGoroutines            *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests    *ebpf.Map `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites      *ebpf.Map `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcRequestStatus     *ebpf.Map `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests    *ebpf.Map `ebpf:"ongoing_grpc_server_requests"`
	OngoingHttpServerConnections *ebpf.Map `ebpf:"ongoing_http_server_connections"`
	OngoingStreams               *ebpf.Map `ebpf:"ongoing_streams"`
	TraceMap                     *ebpf.Map `ebpf:"trace_map"`
}

func (m *bpf_tpMaps) Close() error {
	return _Bpf_tpClose(
		m.Events,
		m.GoTraceMap,
		m.GolangMapbucketStorageMap,
		m.GrpcFramerInvocationMap,
		m.OngoingGoroutines,
		m.OngoingGrpcClientRequests,
		m.OngoingGrpcHeaderWrites,
		m.OngoingGrpcRequestStatus,
		m.OngoingGrpcServerRequests,
		m.OngoingHttpServerConnections,
		m.OngoingStreams,
		m.TraceMap,
	)
}

// bpf_tpPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpPrograms struct {
	UprobeClientConnClose               *ebpf.Program `ebpf:"uprobe_ClientConn_Close"`
	UprobeClientConnInvoke              *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturn        *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke_return"`
	UprobeClientConnNewStream           *ebpf.Program `ebpf:"uprobe_ClientConn_NewStream"`
	UprobeGrpcFramerWriteHeaders        *ebpf.Program `ebpf:"uprobe_grpcFramerWriteHeaders"`
	UprobeGrpcFramerWriteHeadersReturns *ebpf.Program `ebpf:"uprobe_grpcFramerWriteHeaders_returns"`
	UprobeServerHandleStream            *ebpf.Program `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturn      *ebpf.Program `ebpf:"uprobe_server_handleStream_return"`
	UprobeTransportHttp2ClientNewStream *ebpf.Program `ebpf:"uprobe_transport_http2Client_NewStream"`
	UprobeTransportWriteStatus          *ebpf.Program `ebpf:"uprobe_transport_writeStatus"`
}

func (p *bpf_tpPrograms) Close() error {
	return _Bpf_tpClose(
		p.UprobeClientConnClose,
		p.UprobeClientConnInvoke,
		p.UprobeClientConnInvokeReturn,
		p.UprobeClientConnNewStream,
		p.UprobeGrpcFramerWriteHeaders,
		p.UprobeGrpcFramerWriteHeadersReturns,
		p.UprobeServerHandleStream,
		p.UprobeServerHandleStreamReturn,
		p.UprobeTransportHttp2ClientNewStream,
		p.UprobeTransportWriteStatus,
	)
}

func _Bpf_tpClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_bpfel_x86.o
var _Bpf_tpBytes []byte
