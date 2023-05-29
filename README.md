# Go Deep Proto
This is the generated code for the common APIs for DEEP. This is generated from the project [intergra/deep-proto](https://github.com/intergral/deep-proto).

# Usage
The usage of these packages differs between server and clients. Please see [GRPC docs](https://grpc.io/docs/languages/go/) for more guidance.

## Client
To use this as a client:

```go
import (
    tp "github.com/intergral/go-deep-proto/tracepoint/v1"
    pb "github.com/intergral/go-deep-proto/poll/v1"
)

func sendPoll(){
    // create grpc connection
    var opts []grpc.DialOption
    ...
    conn, err := grpc.Dial(*serverAddr, opts...)
    defer conn.Close()

    // create client
    client := pb.NewPollConfigClient(conn)
	// send poll request
	response, err := client.Poll(context.Background(), &pb.PollRequest{})
}

func sendSnapshot(){
	// create grpc connection
    var opts []grpc.DialOption
    ...
    conn, err := grpc.Dial(*serverAddr, opts...)
    defer conn.Close()

	// create client
    client := tp.NewSnapshotServiceClient(conn)
	// send snapshot
	response, err := client.Send(context.Background(), &pb.PollRequest{})
}
```

## Server
To use this as a server:

```go
import (
    tp "github.com/intergral/go-deep-proto/tracepoint/v1"
    pb "github.com/intergral/go-deep-proto/poll/v1"
)

type TPServer struct {}

func (TPServer) Send(ctx context.Context, in *tp.Snapshot) (*tp.SnapshotResponse, error) {
    // process new snapshot
    return &tp.SnapshotResponse{}, nil
}

type PollServer struct {}

func (PollServer) Poll(ctx context.Context, pollRequest *pb.PollRequest) (*pb.PollResponse, error) {
	// process new poll request
	returng &pb.PollResponse{}, nil
}

func setup() {
	// setup network
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	// setup grpc server
    var opts []grpc.ServerOption
    grpcServer := grpc.NewServer(opts...)
	// register services
    pb.RegisterPollConfigServer(t.Server.GRPC, PollServer())
    tp.RegisterSnapshotServiceServer(t.Server.GRPC, TPServer())
	// start server
	grpcServer.Serve(lis)
}


```

# Documentation
The documentation for this project is available [here](https://intergral.github.io/deep-proto/common/).

# Licensing
This project is licensed as [AGPL-3.0-only](./LICENSE).
