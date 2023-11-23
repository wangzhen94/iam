package authorizer

import (
	"github.com/ory/ladon"
	"github.com/wangzhen94/iam/internal/authzserver/authorization"
)

type PolicyGetter interface {
	GetPolicy(key string) ([]*ladon.DefaultPolicy, error)
}

type Authorization struct {
	getter PolicyGetter
}

func NewAuthorization(getter PolicyGetter) authorization.AuthorizationInterface {
	return &Authorization{getter}

}

func (auth *Authorization) LogRejectedAccessRequest(request *ladon.Request, pool ladon.Policies, deciders ladon.Policies) {

	//TODO implement me
	panic("implement me")
}

func (auth *Authorization) LogGrantedAccessRequest(request *ladon.Request, pool ladon.Policies, deciders ladon.Policies) {
	//TODO implement me
	panic("implement me")
}

// Create create a policy.
// Return nil because we use mysql storage to store the policy.
func (auth *Authorization) Create(policy *ladon.DefaultPolicy) error {
	return nil
}

// Update update a policy.
// Return nil because we use mysql storage to store the policy.
func (auth *Authorization) Update(policy *ladon.DefaultPolicy) error {
	return nil
}

// Delete delete a policy by the given identifier.
// Return nil because we use mysql storage to store the policy.
func (auth *Authorization) Delete(id string) error {
	return nil
}

// DeleteCollection batch delete policies by the given identifiers.
// Return nil because we use mysql storage to store the policy.
func (auth *Authorization) DeleteCollection(idList []string) error {
	return nil
}

// Get returns the policy detail by the given identifier.
// Return nil because we use mysql storage to store the policy.
func (auth *Authorization) Get(id string) (*ladon.DefaultPolicy, error) {
	return &ladon.DefaultPolicy{}, nil
}

func (auth *Authorization) List(username string) ([]*ladon.DefaultPolicy, error) {
	return auth.getter.GetPolicy(username)
}
