package main

import (
    "context"
    "flag"
    "github.com/golang/glog"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"
    gw "grpc-gateway-server-go/config"
    "net/http"
)

var (
    echoEndpoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of YourService")
)

func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := gw.RegisterConfigStoreHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
    if err != nil {
        return err
    }

    return http.ListenAndServe(":50050", mux)
}

func main() {
    flag.Parse()
    defer glog.Flush()

    if err := run(); err != nil {
        glog.Fatal(err)
    }
}