package authorization

import "github.com/ory/ladon"

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
