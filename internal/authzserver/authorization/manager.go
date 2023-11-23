package authorization

import (
	"github.com/marmotedu/errors"
	"github.com/ory/ladon"
)

func NewPolicyManager(client AuthorizationInterface) ladon.Manager {
	return &PolicyManager{client}
}

type PolicyManager struct {
	client AuthorizationInterface
}

func (m *PolicyManager) Create(policy ladon.Policy) error {
	return nil
}

func (m *PolicyManager) Update(policy ladon.Policy) error {
	return nil
}

func (m *PolicyManager) Get(id string) (ladon.Policy, error) {
	return nil, nil
}

func (m *PolicyManager) GetAll(limit, offset int64) (ladon.Policies, error) {
	return nil, nil
}

// FindRequestCandidates returns candidates that could match the request object. It either returns
// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
// the error.
func (m *PolicyManager) FindRequestCandidates(r *ladon.Request) (ladon.Policies, error) {
	username := ""

	if user, ok := r.Context["username"].(string); ok {
		username = user
	}

	policies, err := m.client.List(username)
	if err != nil {
		return nil, errors.Wrap(err, "list policies failed")
	}

	ret := make([]ladon.Policy, 0, len(policies))
	for _, policy := range policies {
		ret = append(ret, policy)
	}

	return ret, nil
}

func (m *PolicyManager) Delete(id string) error {
	return nil
}

// FindPoliciesForSubject returns policies that could match the subject. It either returns
// a set of policies that applies to the subject, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *PolicyManager) FindPoliciesForSubject(subject string) (ladon.Policies, error) {
	return nil, nil
}

// FindPoliciesForResource returns policies that could match the resource. It either returns
// a set of policies that apply to the resource, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *PolicyManager) FindPoliciesForResource(resource string) (ladon.Policies, error) {
	return nil, nil
}
