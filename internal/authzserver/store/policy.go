package store

import pb "github.com/marmotedu/api/proto/apiserver/v1"

type PolicyStore interface {
	List() (map[string]*pb.PolicyInfo, error)
}
