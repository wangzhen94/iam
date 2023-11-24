package authorization

import (
	authzv1 "github.com/marmotedu/api/authz/v1"
	"github.com/ory/ladon"
	"github.com/wangzhen94/iam/pkg/log"
)

type Authorizer struct {
	warden ladon.Warden
}

func NewAuthorizer(authorizationClient AuthorizationInterface) *Authorizer {
	return &Authorizer{
		warden: &ladon.Ladon{
			Manager:     NewPolicyManager(authorizationClient),
			AuditLogger: NewAuditLogger(authorizationClient),
		},
	}
}

func (a *Authorizer) Authorize(request *ladon.Request) *authzv1.Response {
	log.Info("authorize request", log.Any("request", request))

	if err := a.warden.IsAllowed(request); err != nil {
		log.Error("authorize failed.")
		return &authzv1.Response{
			Allowed: false,
			Error:   err.Error(),
		}
	}

	log.Info("authorize success.")
	return &authzv1.Response{
		Allowed: true,
	}
}
