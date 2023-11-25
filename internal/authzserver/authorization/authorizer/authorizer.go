package authorizer

import (
	"github.com/marmotedu/component-base/pkg/json"
	"github.com/ory/ladon"
	"github.com/wangzhen94/iam/internal/authzserver/authorization"
	"strings"
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
	return
}

func (auth *Authorization) LogGrantedAccessRequest(r *ladon.Request, p ladon.Policies, d ladon.Policies) {
	//conclusion := fmt.Sprintf("policies %s allow access", joinPoliciesNames(d))
	//rstring, pstring, dstring := convertToString(r, p, d)
	//record := analytics.AnalyticsRecord{
	//	TimeStamp:  time.Now().Unix(),
	//	Username:   r.Context["username"].(string),
	//	Effect:     ladon.AllowAccess,
	//	Conclusion: conclusion,
	//	Request:    rstring,
	//	Policies:   pstring,
	//	Deciders:   dstring,
	//}
	//
	//record.SetExpiry(0)
	//_ = analytics.GetAnalytics().RecordHit(&record)
	return
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

func joinPoliciesNames(policies ladon.Policies) string {
	names := []string{}
	for _, policy := range policies {
		names = append(names, policy.GetID())
	}

	return strings.Join(names, ", ")
}

func convertToString(r *ladon.Request, p ladon.Policies, d ladon.Policies) (string, string, string) {
	rbytes, _ := json.Marshal(r)
	pbytes, _ := json.Marshal(p)
	dbytes, _ := json.Marshal(d)

	return string(rbytes), string(pbytes), string(dbytes)
}
