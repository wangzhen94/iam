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

// secret store
type Secrets struct {
	db *gorm.DB
}

func newSecrets(db *datastore) *Secrets {
	return &Secrets{db.db}
}

func (s *Secrets) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error {
	return s.db.Create(&secret).Error
}
func (s *Secrets) Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) error {
	return s.db.Updates(&secret).Error
}
func (s *Secrets) Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		s.db.Unscoped()
	}
	err := s.db.Where("username = ? and name = ?", username, name).Delete(&v1.Secret{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}
func (s *Secrets) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		s.db.Unscoped()
	}

	return s.db.Where("username = ? and name in (?)", username, names).Delete(&v1.Secret{}).Error
}
func (s *Secrets) Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Secret, error) {
	secret := &v1.Secret{}
	err := s.db.Where("username = ? and name = ?", username, name).First(&secret).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrSecretNotFound, err.Error())
		}
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return secret, nil
}
func (s *Secrets) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.SecretList, error) {
	ret := &v1.SecretList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	if username != "" {
		s.db = s.db.Where("username = ?", username)
	}
	selector, _ := fields.ParseSelector(opts.FieldSelector)
	name, _ := selector.RequiresExactMatch("name")

	d := s.db.Where(" name like ?", "%"+name+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
