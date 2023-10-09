package grpc_bench

import(
    "log"
    "golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SendTest(ctx context.Context, in *TestMessage) (*TestMessage, error){
    log.Printf("Receive message body from client: %s", in.Msg)
    return &TestMessage{Msg: "Hello From the Server!"}, nil
} 
