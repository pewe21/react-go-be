package product

import (
	"context"
	"errors"
	"github.com/pewe21/newbelajar/dto"
	"github.com/pewe21/newbelajar/model"
)

type Service interface {
	Create(ctx context.Context, product dto.CreateProductDTO) error
	Get(ctx context.Context) (products []dto.ProductDTO, err error)
	GetById(ctx context.Context, id int) (product dto.ProductDTO, err error)
	Update(ctx context.Context, id int, product dto.UpdateProductDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) Create(ctx context.Context, product dto.CreateProductDTO) error {
	data := model.ProductModel{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}
	err := s.repo.Create(ctx, &data)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Get(ctx context.Context) (products []dto.ProductDTO, err error) {
	var datas []dto.ProductDTO

	prod, errs := s.repo.Get(ctx)

	if errs != nil {
		return nil, errs
	}

	for _, p := range prod {
		data := dto.ProductDTO{
			Id:          p.Id,
			Name:        p.Name,
			Stock:       p.Stock,
			Price:       p.Price,
			Description: p.Description,
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (s service) GetById(ctx context.Context, id int) (product dto.ProductDTO, err error) {
	p, errs := s.repo.GetById(ctx, id)
	if errs != nil {
		return product, errs
	}

	if p.Id == 0 {
		return product, errors.New("product not found")
	}

	data := dto.ProductDTO{
		Id:          p.Id,
		Name:        p.Name,
		Stock:       p.Stock,
		Price:       p.Price,
		Description: p.Description,
	}

	return data, nil

}

func (s service) Update(ctx context.Context, id int, product dto.UpdateProductDTO) error {

	p, errs := s.repo.GetById(ctx, id)

	if errs != nil {
		return errs
	}

	if p.Id == 0 {
		return errors.New("product not found")
	}

	prod := model.ProductModel{
		Name:  product.Name,
		Price: product.Price,
	}
	err := s.repo.Update(ctx, id, prod)
	return err
}

func (s service) Delete(ctx context.Context, id int) error {

	p, errs := s.repo.GetById(ctx, id)

	if errs != nil {
		return errs
	}

	if p.Id == 0 {
		return errors.New("product not found")
	}
	errs = s.repo.Delete(ctx, id)
	return errs
}
