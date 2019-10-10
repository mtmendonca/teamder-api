package grpc

import "google.golang.org/grpc"

// GetConn returns a gRPC connection for given upstream endpoint
func GetConn(upstream string) (*grpc.ClientConn, error) {
	return grpc.Dial(upstream, grpc.WithInsecure())
}
