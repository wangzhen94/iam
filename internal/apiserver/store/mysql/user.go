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

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

func (u *users) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	return u.db.Create(&user).Error
}

func (u *users) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	return u.db.Updates(&user).Error
}

func (u *users) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	err := u.db.Where("name = ?", username).Delete(&v1.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (u *users) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	return u.db.Where("name in (?)", usernames).Delete(&v1.User{}).Error
}

func (u *users) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("name = ? and status = 1", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return user, nil
}

func (u *users) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, _ := selector.RequiresExactMatch("name")
	d := u.db.Where("name like ? and status = 1", "%"+username+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}

func (u *users) ListOptional(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	where := v1.User{}
	whereNot := v1.User{
		IsAdmin: 0,
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, found := selector.RequiresExactMatch("name")

	if found {
		where.Name = username
	}

	d := u.db.Where(where).
		Not(whereNot).
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
