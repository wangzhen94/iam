package authorization

import "github.com/ory/ladon"

//go:generate mockgen -self_package=github.com/wangzhen94/iam/internal/authzserver/authorization -destination mock_authorization.go -package authorization github.com/wangzhen94/iam/internal/authzserver/authorization AuthorizationInterface

type AuthorizationInterface interface {
	Create(*ladon.DefaultPolicy) error
	Update(*ladon.DefaultPolicy) error
	Delete(id string) error
	DeleteCollection(idList []string) error
	Get(id string) (*ladon.DefaultPolicy, error)
	List(username string) ([]*ladon.DefaultPolicy, error)

	// The following two functions tracks denied and granted authorizations.
	LogRejectedAccessRequest(request *ladon.Request, pool ladon.Policies, deciders ladon.Policies)
	LogGrantedAccessRequest(request *ladon.Request, pool ladon.Policies, deciders ladon.Policies)
}
