package main

import(
    "log"
    "net"
    "google.golang.org/grpc"
    "test/grpc_bench"
)

func main(){
    
    lis, err := net.Listen("tcp", ":9000")

    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc_bench.Server{}

    grpcServer := grpc.NewServer()

    grpc_bench.RegisterTestServiceServer(grpcServer, &s)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %s", err)
    }

    

}
