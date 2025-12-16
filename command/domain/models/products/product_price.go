package products

import (
	"fmt"

	"github.com/zakzackr/grpc-microservices-demo-command-service/errs"
)

// 商品価格を保持する値オブジェクト
type ProductPrice struct {
	value uint32
}

// getter
func (ins *ProductPrice) Value() uint32 {
	return ins.value
}

// コンストラクタ
func NewProductPrice(value uint32) (*ProductPrice, *errs.DomainError) {
	const MIN_VALUE = 50    // 最小単価
	const MAX_VALUE = 10000 // 最大単価
	if value < MIN_VALUE || value > MAX_VALUE {
		return nil, errs.NewDomainError(fmt.Sprintf("単価は%d以上、%d以下です。", MIN_VALUE, MAX_VALUE))
	}
	return &ProductPrice{value: value}, nil
}