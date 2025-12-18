package adapter

import (
	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/categories"
	"github.com/zakzackr/grpc-microservices-demo-command-service/errs"
	"github.com/zakzackr/grpc-microservices-demo-pb/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Adapterインターフェースの実装
type CategoryAdapterImpl struct {}

func NewCategoryAdapterImpl() CategoryAdapter {
	return &CategoryAdapterImpl{}
}

// CategoryUpParamをCategoryエンティティに変換
func (ins *CategoryAdapterImpl) ToEntity(param *pb.CategoryUpParam) (*categories.Category, error) {
	switch param.GetCrud() {
	case pb.CRUD_INSERT:
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		category, err := categories.NewCategory(name)
		if err != nil {
			return nil, err
		}
		return category, nil
	case pb.CRUD_UPDATE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, name), nil
	case pb.CRUD_DELETE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, nil), nil
	default:
		return nil, errs.NewDomainError("不明な操作を受け取りました。")
	}
}

// CategoryエンティティをCategoryUpResultに変換
func (ins *CategoryAdapterImpl) ToResult(result any) *pb.CategoryUpResult {
	var up_category *pb.Category
	var up_err *pb.Error

	switch v := result.(type) {
	case *categories.Category:
		if v.Name() == nil {
			up_category = &pb.Category{Id: v.Id().Value(), Name: ""}
		} else {
			up_category = &pb.Category{Id: v.Id().Value(), Name: v.Name().Value()}
		}
	case *errs.DomainError:
		up_err = &pb.Error{Type: "Domain Error", Message: v.Error()}
	case *errs.CRUDError:
		up_err = &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError:
		up_err = &pb.Error{Type: "Internal Error", Message: "サーバ側でエラーが発生しました。"}
	}
	return &pb.CategoryUpResult{
		Category: up_category,
		Error: up_err,
		Timestamp: timestamppb.Now(),
	}
}

