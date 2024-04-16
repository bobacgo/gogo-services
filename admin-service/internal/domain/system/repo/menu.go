package repo

import (
	"context"
	"errors"

	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/query"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
	q  *query.Query
}

var _ service.IMenuRepo = (*menuRepo)(nil)

func NewMenuRepo(db *gorm.DB) service.IMenuRepo {
	return &menuRepo{db: db, q: query.Use(db)}
}

func (repo *menuRepo) Find(ctx context.Context, req *v1.MenuListRequest) (result []*model.Menu, count int64, err error) {
	q := repo.q.Menu
	db := q.WithContext(ctx)
	if req.ParentID != nil {
		db = db.Where(q.ParentID.Eq(*req.ParentID))
	}
	return db.FindByPage(req.Offset(), req.Limit())
}

func (repo *menuRepo) FindOne(ctx context.Context, req *v1.MenuRequest) (*model.Menu, error) {
	q := repo.q.Menu
	res, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) { // 错误应该不能依赖于底层错误
		return nil, errs.MenuNotFound
	}
	return res, err
}

func (repo *menuRepo) Create(ctx context.Context, data *model.Menu) error {
	return repo.q.Menu.WithContext(ctx).Create(data)
}

func (repo *menuRepo) Update(ctx context.Context, req *v1.MenuUpdateRequest) error {
	q := repo.q.Menu
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Updates(req)
	return err
}

func (repo *menuRepo) UpdateHidden(ctx context.Context, ID int64, hidden *bool) error {
	q := repo.q.Menu
	_, err := q.WithContext(ctx).Where(q.ID.Eq(ID)).Update(q.Hidden, hidden)
	return err
}

func (repo *menuRepo) Delete(ctx context.Context, req *v1.MenuDeleteRequest) error {
	q := repo.q.Menu
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Delete()
	return err
}
