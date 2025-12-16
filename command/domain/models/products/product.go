package products

import (
	"github.com/google/uuid"
	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/categories"
	"github.com/zakzackr/grpc-microservices-demo-command-service/errs"
)

// 商品エンティティ
type Product struct {
	id *ProductId
	name *ProductName
	price *ProductPrice
	category *categories.Category 
}

// getter
func (ins *Product) Id() *ProductId {
	return ins.id
}

func (ins *Product) Name() *ProductName {
	return ins.name
}

func (ins *Product) Price() *ProductPrice {
	return ins.price
}

func (ins *Product) Category() *categories.Category {
	return ins.category
}

// 値の変更
func (ins *Product) ChangeProductName(name *ProductName) {
	ins.name = name
}

func (ins *Product) ChangeProductPrice(price *ProductPrice) {
	ins.price = price
}

func (ins *Product) ChangeCategory(category *categories.Category) {
	ins.category = category
}

// 同一性検証
// エンティティはIdで同一性検証を行う
func (ins *Product) Equals(obj *Product) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.id.Equals(obj.Id())
	return result, nil
}

// コンストラクタ
func NewProduct(name *ProductName, price *ProductPrice, category *categories.Category) (*Product, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errs.NewDomainError(err.Error())
	} else {
		// 商品ID用値オブジェクトを生成する
		if id, err := NewProductId(uid.String()); err != nil {
			return nil, err
		} else {
			// 商品エンティティのインスタンスを生成して返す
			return &Product{
				id:       id,
				name:     name,
				price:    price,
				category: category,
			}, nil
		}
	}
}

// 商品エンティティの再構築
func BuildProduct(id *ProductId, name *ProductName, price *ProductPrice, category *categories.Category) *Product {
	product := Product{ // 商品エンティティのインスタンスを生成して返す
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}
	return &product
}