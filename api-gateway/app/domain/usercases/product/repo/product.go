package repo

import (
	"time"

	"gorm.io/gorm"

	"api-gateway/app/domain/entities"
	"api-gateway/app/domain/usercases/common"
	"api-gateway/app/domain/usercases/product/orm"
	"api-gateway/app/transport/product/request"
	"api-gateway/app/uerror"
)

type ProductRepo struct {
	Id        int
	Name      string
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Active    bool
}
type IProductRepo interface {
	CreateProductTx(newProduct request.ProductRequestBody) (ProductRepo, error)
	GetProductBySectionId(sectionId int, limit int, offset int) ([]ProductRepo, int, error)
}
type productRepo struct{}

func NewProductRepo() *productRepo {
	return &productRepo{}
}

func (productRepo) CreateProductTx(newProduct request.ProductRequestBody) (ProductRepo, error) {
	tx := common.BeginTx()
	defer common.RecoveryTx(tx)

	now := time.Now()
	newProductEntity := entities.Product{
		Name:      newProduct.Name,
		Active:    true,
		Image:     newProduct.Image,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := orm.Product.CreateProductTx(&newProductEntity, tx)
	if err != nil {
		tx.Rollback()
		return ProductRepo{}, uerror.InternalError(err, err.Error())
	}
	return ProductRepo{
		Id:        newProductEntity.Id,
		Active:    newProductEntity.Active,
		Image:     newProductEntity.Image,
		CreatedAt: newProductEntity.CreatedAt,
		UpdatedAt: newProductEntity.UpdatedAt,
	}, nil
}

func (d productRepo) GetProductBySectionId(sectionId int, limit int, offset int) ([]ProductRepo, int, error) {
	productEntity, total, err := orm.Product.GetProductBySectionId(sectionId, limit, offset)
	if err != nil && err == gorm.ErrRecordNotFound {
		return []ProductRepo{}, 0, nil
	}
	if err != nil {
		return []ProductRepo{}, 0, uerror.InternalError(err, err.Error())
	}
	return d.populateProducts(productEntity), total, nil
}

func (d productRepo) populateProducts(products []entities.Product) []ProductRepo {
	if len(products) <= 0 {
		return []ProductRepo{}
	}
	resp := make([]ProductRepo, len(products))
	for i := range products {
		resp[i] = d.populateProduct(products[i])
	}
	return resp
}

func (d productRepo) populateProduct(product entities.Product) ProductRepo {
	return ProductRepo{
		Id:        product.Id,
		Name:      product.Name,
		Image:     product.Image,
		Active:    product.Active,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
