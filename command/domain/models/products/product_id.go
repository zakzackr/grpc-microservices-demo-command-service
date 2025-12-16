package products

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/zakzackr/grpc-microservices-demo-command-service/errs"
)

// 商品IDを保持する値オブジェクト(UUID)
type ProductId struct {
	value string
}

// valueフィールドのgetter
func (ins *ProductId) Value() string {
	return ins.value
}

// 同一性検証
func (ins *ProductId) Equals(value *ProductId) bool {
	if ins == value {
		return true
	}

	// 値オブジェクトは不変オブジェクトだから値での同一性検証が必要
	return ins.value == value.Value()
}

// コンストラクタ
func NewProductId(value string) (*ProductId, *errs.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("商品IDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("商品IDはUUIDの形式でなければなりません。")
	}
	return &ProductId{value: value}, nil
}
