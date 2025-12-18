package adaptor

import (
	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/categories"
	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/products"
	"github.com/zakzackr/grpc-microservices-demo-command-service/errs"
	"github.com/zakzackr/grpc-microservices-demo-pb/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductAdaptorImpl struct {}

func NewProductAdaptorImpl() ProductAdaptor {
	return &ProductAdaptorImpl{}
}
// ProductUpParamをProductエンティティに変換
func (ins *ProductAdaptorImpl) ToEntity(param *pb.ProductUpParam) (*products.Product, error) {
	switch param.GetCrud() {
	case pb.CRUD_INSERT:
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		id, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		product, err := products.NewProduct(name, price, categories.BuildCategory(id, nil))
		if err != nil {
			return nil, err
		}
		return product, nil
	case pb.CRUD_UPDATE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		cid, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, name, price, categories.BuildCategory(cid, nil)), nil
	case pb.CRUD_DELETE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, nil, nil, nil), nil
	default:
		return nil, errs.NewInternalError("不明な操作を受け取りました。")
	}
}

// ProductエンティティをProductUpResultに変換
func (ins *ProductAdaptorImpl) ToResult(result any) *pb.ProductUpResult {
	var up_product *pb.Product
	var up_err *pb.Error
	switch v := result.(type){
	case *products.Product:
		var c *pb.Category
		if v.Category() == nil {
			c = &pb.Category{Id: "", Name: ""}
		} else {
			c = &pb.Category{Id: v.Category().Id().Value(), Name: ""}
		}
		var name string = ""
		if v.Name() != nil {
			name = v.Name().Value()
		}
		var price int32 = 0
		if v.Price() != nil {
			price = int32(v.Price().Value())
		}
		up_product = &pb.Product{Id: v.Id().Value(), Name: name, Price: price, Category: c}

	case *errs.DomainError: 
		up_err = &pb.Error{Type: "Domain Error", Message: v.Error()}
	case *errs.CRUDError: 
		up_err = &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError: 
		up_err = &pb.Error{Type: "Internal Error", Message: "サーバ側でエラーが発生しました。"}
	}

	return &pb.ProductUpResult{
		Product: up_product,
		Error: up_err,
		Timestamp: timestamppb.Now(),
	}
}
