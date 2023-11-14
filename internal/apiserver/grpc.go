package apiserver

import (
	"google.golang.org/grpc"
	"net"
)

type grpcAPIServer struct {
	*grpc.Server
	address string
}

func (s *grpcAPIServer) Run() {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		// log fatal
	}

	go func() {
		if err := s.Serve(listen); err != nil {
			// log fatal
		}
	}()

	// log info
}

func (s *grpcAPIServer) Close() {
	s.GracefulStop()
	//log.Infof("GRPC server on %s stopped", s.address)
}
