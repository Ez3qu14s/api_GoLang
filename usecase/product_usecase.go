package usecase

import "go-api/model"

type productUsecase struct {
	// repository
}

func NewProductUsecase() productUsecase {
	return productUsecase{}
}

func (pu *productUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
