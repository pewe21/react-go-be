package product

import (
	"context"
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/pewe21/newbelajar/model"
)

type Repository interface {
	Create(ctx context.Context, product *model.ProductModel) error
	Get(ctx context.Context) (products []model.ProductModel, err error)
	GetById(ctx context.Context, id int) (product model.ProductModel, err error)
	Update(ctx context.Context, id int, product model.ProductModel) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *goqu.Database
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: goqu.New("default", db)}
}

func (r repository) Create(ctx context.Context, product *model.ProductModel) error {
	log.Println(product)
	executor := r.db.Insert("products").Rows(product).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) Get(ctx context.Context) (products []model.ProductModel, err error) {
	dataset := r.db.From("products").Executor()
	err = dataset.ScanStructsContext(ctx, &products)
	return
}

func (r repository) GetById(ctx context.Context, id int) (product model.ProductModel, err error) {
	dataset := r.db.From("products").Where(goqu.C("id").Eq(id)).Executor()
	_, err = dataset.ScanStructContext(ctx, &product)
	return
}

func (r repository) Update(ctx context.Context, id int, product model.ProductModel) error {
	executor := r.db.Update("products").Set(goqu.Ex{
		"name":        product.Name,
		"price":       product.Price,
		"description": product.Description,
		"updated_at":  goqu.L("NOW()"),
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) Delete(ctx context.Context, id int) error {
	executor := r.db.Delete("products").Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
