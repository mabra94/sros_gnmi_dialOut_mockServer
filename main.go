package main

import (
	"context"
	"flag"
	"log"
	"time"
	"math/rand"

	"google.golang.org/grpc"
	"github.com/openconfig/gnmi/proto/gnmi"
)

// DialoutTelemetryClient is the client API for DialoutTelemetry service.
type DialoutTelemetryClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (DialoutTelemetry_PublishClient, error)
}

type dialoutTelemetryClient struct {
	cc grpc.ClientConnInterface
}

func NewDialoutTelemetryClient(cc grpc.ClientConnInterface) DialoutTelemetryClient {
	return &dialoutTelemetryClient{cc}
}

func (c *dialoutTelemetryClient) Publish(ctx context.Context, opts ...grpc.CallOption) (DialoutTelemetry_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DialoutTelemetry_serviceDesc.Streams[0], "/Nokia.SROS.DialoutTelemetry/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &dialoutTelemetryPublishClient{stream}
	return x, nil
}

type DialoutTelemetry_PublishClient interface {
	Send(*gnmi.SubscribeResponse) error
	Recv() (*PublishResponse, error)
	grpc.ClientStream
}

type dialoutTelemetryPublishClient struct {
	grpc.ClientStream
}

func (x *dialoutTelemetryPublishClient) Send(m *gnmi.SubscribeResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dialoutTelemetryPublishClient) Recv() (*PublishResponse, error) {
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DialoutTelemetry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Nokia.SROS.DialoutTelemetry",
	HandlerType: nil,
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       nil,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sros_dialout.proto",
}

type PublishResponse struct{}

func main() {
	serverAddr := flag.String("server", "localhost:57400", "The server address in the format of host:port")
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewDialoutTelemetryClient(conn)
	stream, err := client.Publish(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	for {
		resp := createMockSubscribeResponse()
		if err := stream.Send(resp); err != nil {
			log.Printf("Failed to send: %v", err)
			break
		}
		log.Printf("Sent mock data: %v", resp)
		time.Sleep(5 * time.Second)
	}
}

func createMockSubscribeResponse() *gnmi.SubscribeResponse {
	return &gnmi.SubscribeResponse{
		Response: &gnmi.SubscribeResponse_Update{
			Update: &gnmi.Notification{
				Timestamp: time.Now().UnixNano(), // This is now correct as int64
				Prefix: &gnmi.Path{
					Target: "router1",
					Origin: "sros",
				},
				Update: []*gnmi.Update{
					{
						Path: &gnmi.Path{
							Elem: []*gnmi.PathElem{
								{Name: "interface"},
								{Name: "statistics"},
								{Name: "in-octets"},
							},
						},
						Val: &gnmi.TypedValue{
							Value: &gnmi.TypedValue_UintVal{UintVal: uint64(rand.Intn(1000000))},
						},
					},
					{
						Path: &gnmi.Path{
							Elem: []*gnmi.PathElem{
								{Name: "interface"},
								{Name: "statistics"},
								{Name: "out-octets"},
							},
						},
						Val: &gnmi.TypedValue{
							Value: &gnmi.TypedValue_UintVal{UintVal: uint64(rand.Intn(1000000))},
						},
					},
				},
			},
		},
	}
}