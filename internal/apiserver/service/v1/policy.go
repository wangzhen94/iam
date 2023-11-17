package v1

import (
	"context"
	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/apiserver/store"
	"github.com/wangzhen94/iam/internal/pkg/code"
)

type PolicySrv interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
	Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}

type policyService struct {
	store store.Factory
}

func newPolices(s *service) *policyService {
	return &policyService{store: s.store}
}

var _ PolicySrv = (*policyService)(nil)

func (p *policyService) Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error {
	if err := p.store.Policies().Create(ctx, policy, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyService) Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error {
	// Save changed fields.
	if err := p.store.Policies().Update(ctx, policy, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyService) Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error {
	if err := p.store.Policies().Delete(ctx, username, name, opts); err != nil {
		return err
	}

	return nil
}

func (p *policyService) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	if err := p.store.Policies().DeleteCollection(ctx, username, names, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p *policyService) Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error) {
	policy, err := p.store.Policies().Get(ctx, username, name, opts)
	if err != nil {
		return nil, err
	}

	return policy, nil
}

func (p *policyService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error) {
	policies, err := p.store.Policies().List(ctx, username, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return policies, nil
}
