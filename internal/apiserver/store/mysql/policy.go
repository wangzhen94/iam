package mysql

import (
	"context"
	v1 "github.com/marmotedu/api/apiserver/v1"
	"github.com/marmotedu/component-base/pkg/fields"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/wangzhen94/iam/internal/pkg/code"
	"github.com/wangzhen94/iam/internal/pkg/util/gormutil"
	"gorm.io/gorm"
)

type policies struct {
	db *gorm.DB
}

func newPolicies(ds *datastore) *policies {
	return &policies{ds.db}
}

func (p *policies) Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error {
	return p.db.Create(&policy).Error
}

func (p *policies) Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error {
	return p.db.Save(policy).Error
}

func (p *policies) Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	err := p.db.Where("username = ? and name = ?", username, name).Delete(&v1.Policy{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (p *policies) DeleteByUser(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username = ?", username).Delete(&v1.Policy{}).Error
}

func (p *policies) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username = ? and name in (?)", username, names).Delete(&v1.Policy{}).Error
}

// DeleteCollectionByUser batch deletes policies usernames.
func (p *policies) DeleteCollectionByUser(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username in (?)", usernames).Delete(&v1.Policy{}).Error
}

func (p *policies) Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error) {
	policy := &v1.Policy{}
	err := p.db.Where("username = ? and name = ?", username, name).First(&policy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrPolicyNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return policy, nil
}

func (p *policies) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error) {
	ret := &v1.PolicyList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	if username != "" {
		p.db = p.db.Where("username = ?", username)
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	name, _ := selector.RequiresExactMatch("name")

	d := p.db.Where("name like ?", "%"+name+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
