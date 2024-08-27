// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package httpfltr

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tp_debugCallProtocolArgsT struct {
	PidConn    bpf_tp_debugPidConnectionInfoT
	SmallBuf   [24]uint8
	U_buf      uint64
	BytesLen   int32
	Ssl        uint8
	Direction  uint8
	OrigDport  uint16
	PacketType uint8
	_          [7]byte
}

type bpf_tp_debugConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_tp_debugHttp2ConnStreamT struct {
	PidConn  bpf_tp_debugPidConnectionInfoT
	StreamId uint32
}

type bpf_tp_debugHttp2GrpcRequestT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpf_tp_debugConnectionInfoT
	Data            [256]uint8
	RetData         [64]uint8
	Type            uint8
	_               [1]byte
	Len             int32
	_               [4]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Ssl uint8
	_   [3]byte
	Tp  struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_tp_debugHttpConnectionMetadataT struct {
	Pid struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Type uint8
}

type bpf_tp_debugHttpInfoT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpf_tp_debugConnectionInfoT
	_               [2]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [160]uint8
	Len             uint32
	RespLen         uint32
	Status          uint16
	Type            uint8
	Ssl             uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	ExtraId uint64
	TaskTid uint32
	_       [4]byte
}

type bpf_tp_debugPartialConnectionInfoT struct {
	S_addr [16]uint8
	S_port uint16
	D_port uint16
	TcpSeq uint32
}

type bpf_tp_debugPidConnectionInfoT struct {
	Conn bpf_tp_debugConnectionInfoT
	Pid  uint32
}

type bpf_tp_debugPidKeyT struct {
	Pid uint32
	Ns  uint32
}

type bpf_tp_debugRecvArgsT struct {
	SockPtr  uint64
	IovecCtx [40]uint8
}

type bpf_tp_debugSendArgsT struct {
	P_conn bpf_tp_debugPidConnectionInfoT
	Size   uint64
}

type bpf_tp_debugSockArgsT struct {
	Addr       uint64
	AcceptTime uint64
}

type bpf_tp_debugSslArgsT struct {
	Ssl    uint64
	Buf    uint64
	LenPtr uint64
}

type bpf_tp_debugSslPidConnectionInfoT struct {
	P_conn    bpf_tp_debugPidConnectionInfoT
	OrigDport uint16
	_         [2]byte
}

type bpf_tp_debugTcpReqT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpf_tp_debugConnectionInfoT
	_               [2]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	Rbuf            [128]uint8
	Len             uint32
	RespLen         uint32
	Ssl             uint8
	Direction       uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_  [2]byte
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_tp_debugTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid   uint32
	Valid uint8
	_     [3]byte
}

type bpf_tp_debugTraceKeyT struct {
	P_key   bpf_tp_debugPidKeyT
	ExtraId uint64
}

// loadBpf_tp_debug returns the embedded CollectionSpec for bpf_tp_debug.
func loadBpf_tp_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tp_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp_debug: %w", err)
	}

	return spec, err
}

// loadBpf_tp_debugObjects loads bpf_tp_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tp_debugObjects
//	*bpf_tp_debugPrograms
//	*bpf_tp_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tp_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tp_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugSpecs struct {
	bpf_tp_debugProgramSpecs
	bpf_tp_debugMapSpecs
}

// bpf_tp_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugProgramSpecs struct {
	KprobeSysExit           *ebpf.ProgramSpec `ebpf:"kprobe_sys_exit"`
	KprobeTcpCleanupRbuf    *ebpf.ProgramSpec `ebpf:"kprobe_tcp_cleanup_rbuf"`
	KprobeTcpClose          *ebpf.ProgramSpec `ebpf:"kprobe_tcp_close"`
	KprobeTcpConnect        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.ProgramSpec `ebpf:"kprobe_tcp_rcv_established"`
	KprobeTcpRecvmsg        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_recvmsg"`
	KprobeTcpSendmsg        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_sendmsg"`
	KretprobeSockAlloc      *ebpf.ProgramSpec `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysClone       *ebpf.ProgramSpec `ebpf:"kretprobe_sys_clone"`
	KretprobeSysConnect     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_connect"`
	KretprobeTcpRecvmsg     *ebpf.ProgramSpec `ebpf:"kretprobe_tcp_recvmsg"`
	KretprobeTcpSendmsg     *ebpf.ProgramSpec `ebpf:"kretprobe_tcp_sendmsg"`
	ProtocolHttp            *ebpf.ProgramSpec `ebpf:"protocol_http"`
	ProtocolHttp2           *ebpf.ProgramSpec `ebpf:"protocol_http2"`
	ProtocolTcp             *ebpf.ProgramSpec `ebpf:"protocol_tcp"`
	SocketHttpFilter        *ebpf.ProgramSpec `ebpf:"socket__http_filter"`
}

// bpf_tp_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugMapSpecs struct {
	ActiveAcceptArgs        *ebpf.MapSpec `ebpf:"active_accept_args"`
	ActiveConnectArgs       *ebpf.MapSpec `ebpf:"active_connect_args"`
	ActiveNodejsIds         *ebpf.MapSpec `ebpf:"active_nodejs_ids"`
	ActiveRecvArgs          *ebpf.MapSpec `ebpf:"active_recv_args"`
	ActiveSendArgs          *ebpf.MapSpec `ebpf:"active_send_args"`
	ActiveSendSockArgs      *ebpf.MapSpec `ebpf:"active_send_sock_args"`
	ActiveSslConnections    *ebpf.MapSpec `ebpf:"active_ssl_connections"`
	ActiveSslHandshakes     *ebpf.MapSpec `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs       *ebpf.MapSpec `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.MapSpec `ebpf:"active_ssl_write_args"`
	CloneMap                *ebpf.MapSpec `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.MapSpec `ebpf:"connection_meta_mem"`
	DebugEvents             *ebpf.MapSpec `ebpf:"debug_events"`
	Events                  *ebpf.MapSpec `ebpf:"events"`
	Http2InfoMem            *ebpf.MapSpec `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.MapSpec `ebpf:"http_info_mem"`
	IovecMem                *ebpf.MapSpec `ebpf:"iovec_mem"`
	JumpTable               *ebpf.MapSpec `ebpf:"jump_table"`
	NodejsParentMap         *ebpf.MapSpec `ebpf:"nodejs_parent_map"`
	OngoingHttp             *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.MapSpec `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.MapSpec `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.MapSpec `ebpf:"ongoing_tcp_req"`
	PidCache                *ebpf.MapSpec `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.MapSpec `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem         *ebpf.MapSpec `ebpf:"protocol_args_mem"`
	ServerTraces            *ebpf.MapSpec `ebpf:"server_traces"`
	SslToConn               *ebpf.MapSpec `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.MapSpec `ebpf:"ssl_to_pid_tid"`
	TcpConnectionMap        *ebpf.MapSpec `ebpf:"tcp_connection_map"`
	TcpReqMem               *ebpf.MapSpec `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.MapSpec `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.MapSpec `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids               *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_tp_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugObjects struct {
	bpf_tp_debugPrograms
	bpf_tp_debugMaps
}

func (o *bpf_tp_debugObjects) Close() error {
	return _Bpf_tp_debugClose(
		&o.bpf_tp_debugPrograms,
		&o.bpf_tp_debugMaps,
	)
}

// bpf_tp_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugMaps struct {
	ActiveAcceptArgs        *ebpf.Map `ebpf:"active_accept_args"`
	ActiveConnectArgs       *ebpf.Map `ebpf:"active_connect_args"`
	ActiveNodejsIds         *ebpf.Map `ebpf:"active_nodejs_ids"`
	ActiveRecvArgs          *ebpf.Map `ebpf:"active_recv_args"`
	ActiveSendArgs          *ebpf.Map `ebpf:"active_send_args"`
	ActiveSendSockArgs      *ebpf.Map `ebpf:"active_send_sock_args"`
	ActiveSslConnections    *ebpf.Map `ebpf:"active_ssl_connections"`
	ActiveSslHandshakes     *ebpf.Map `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs       *ebpf.Map `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.Map `ebpf:"active_ssl_write_args"`
	CloneMap                *ebpf.Map `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.Map `ebpf:"connection_meta_mem"`
	DebugEvents             *ebpf.Map `ebpf:"debug_events"`
	Events                  *ebpf.Map `ebpf:"events"`
	Http2InfoMem            *ebpf.Map `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.Map `ebpf:"http_info_mem"`
	IovecMem                *ebpf.Map `ebpf:"iovec_mem"`
	JumpTable               *ebpf.Map `ebpf:"jump_table"`
	NodejsParentMap         *ebpf.Map `ebpf:"nodejs_parent_map"`
	OngoingHttp             *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.Map `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.Map `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.Map `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.Map `ebpf:"ongoing_tcp_req"`
	PidCache                *ebpf.Map `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.Map `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem         *ebpf.Map `ebpf:"protocol_args_mem"`
	ServerTraces            *ebpf.Map `ebpf:"server_traces"`
	SslToConn               *ebpf.Map `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.Map `ebpf:"ssl_to_pid_tid"`
	TcpConnectionMap        *ebpf.Map `ebpf:"tcp_connection_map"`
	TcpReqMem               *ebpf.Map `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.Map `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.Map `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.Map `ebpf:"trace_map"`
	ValidPids               *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_tp_debugMaps) Close() error {
	return _Bpf_tp_debugClose(
		m.ActiveAcceptArgs,
		m.ActiveConnectArgs,
		m.ActiveNodejsIds,
		m.ActiveRecvArgs,
		m.ActiveSendArgs,
		m.ActiveSendSockArgs,
		m.ActiveSslConnections,
		m.ActiveSslHandshakes,
		m.ActiveSslReadArgs,
		m.ActiveSslWriteArgs,
		m.CloneMap,
		m.ConnectionMetaMem,
		m.DebugEvents,
		m.Events,
		m.Http2InfoMem,
		m.HttpInfoMem,
		m.IovecMem,
		m.JumpTable,
		m.NodejsParentMap,
		m.OngoingHttp,
		m.OngoingHttp2Connections,
		m.OngoingHttp2Grpc,
		m.OngoingHttpFallback,
		m.OngoingTcpReq,
		m.PidCache,
		m.PidTidToConn,
		m.ProtocolArgsMem,
		m.ServerTraces,
		m.SslToConn,
		m.SslToPidTid,
		m.TcpConnectionMap,
		m.TcpReqMem,
		m.TpCharBufMem,
		m.TpInfoMem,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpf_tp_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugPrograms struct {
	KprobeSysExit           *ebpf.Program `ebpf:"kprobe_sys_exit"`
	KprobeTcpCleanupRbuf    *ebpf.Program `ebpf:"kprobe_tcp_cleanup_rbuf"`
	KprobeTcpClose          *ebpf.Program `ebpf:"kprobe_tcp_close"`
	KprobeTcpConnect        *ebpf.Program `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.Program `ebpf:"kprobe_tcp_rcv_established"`
	KprobeTcpRecvmsg        *ebpf.Program `ebpf:"kprobe_tcp_recvmsg"`
	KprobeTcpSendmsg        *ebpf.Program `ebpf:"kprobe_tcp_sendmsg"`
	KretprobeSockAlloc      *ebpf.Program `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.Program `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysClone       *ebpf.Program `ebpf:"kretprobe_sys_clone"`
	KretprobeSysConnect     *ebpf.Program `ebpf:"kretprobe_sys_connect"`
	KretprobeTcpRecvmsg     *ebpf.Program `ebpf:"kretprobe_tcp_recvmsg"`
	KretprobeTcpSendmsg     *ebpf.Program `ebpf:"kretprobe_tcp_sendmsg"`
	ProtocolHttp            *ebpf.Program `ebpf:"protocol_http"`
	ProtocolHttp2           *ebpf.Program `ebpf:"protocol_http2"`
	ProtocolTcp             *ebpf.Program `ebpf:"protocol_tcp"`
	SocketHttpFilter        *ebpf.Program `ebpf:"socket__http_filter"`
}

func (p *bpf_tp_debugPrograms) Close() error {
	return _Bpf_tp_debugClose(
		p.KprobeSysExit,
		p.KprobeTcpCleanupRbuf,
		p.KprobeTcpClose,
		p.KprobeTcpConnect,
		p.KprobeTcpRcvEstablished,
		p.KprobeTcpRecvmsg,
		p.KprobeTcpSendmsg,
		p.KretprobeSockAlloc,
		p.KretprobeSysAccept4,
		p.KretprobeSysClone,
		p.KretprobeSysConnect,
		p.KretprobeTcpRecvmsg,
		p.KretprobeTcpSendmsg,
		p.ProtocolHttp,
		p.ProtocolHttp2,
		p.ProtocolTcp,
		p.SocketHttpFilter,
	)
}

func _Bpf_tp_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_debug_arm64_bpfel.o
var _Bpf_tp_debugBytes []byte
