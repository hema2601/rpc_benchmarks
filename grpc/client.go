package main

import(
    "log"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "test/grpc_bench"
    //"os"
    "fmt"
    "github.com/pborman/getopt/v2"
)


var helpFlag = getopt.Bool('h', "display help")
var ip = getopt.String('i', "127.0.0.1", "set ip")
var port = getopt.Int16('p', 9000, "set port")
var size = getopt.Int16('s', 64, "set size")
//var protocol = getopt.Enum('t', []string{"UDP", "TCP"}, "UDP", "Set protocol")


func main() {
    
    getopt.Parse()


    fmt.Println(*ip)
    fmt.Println(*port)
    fmt.Println(*size)
    //fmt.Println(*protocol)
    
    bytes := make([]byte, *size)

    for i := *size - 1; i >= 0; i-- {
        bytes[i] = 0x61
    }

    payload := string(bytes[:])

    var conn *grpc.ClientConn

    conn, err := grpc.Dial( fmt.Sprintf("%s:%d", *ip, *port), grpc.WithInsecure())

    if err != nil {
        log.Fatalf("did not connect: %s", err)
    }

    defer conn.Close()

    c := grpc_bench.NewTestServiceClient(conn)

    response, err := c.SendTest(context.Background(), &grpc_bench.TestMessage{Msg: payload})

    if err != nil {
        log.Fatalf("error when calling SayHello: %s", err)
    }

    log.Printf("Response from server: %s", response.Msg)

}
