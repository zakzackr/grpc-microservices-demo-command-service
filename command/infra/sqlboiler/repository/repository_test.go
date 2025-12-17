package repository

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zakzackr/grpc-microservices-demo-command-service/infra/sqlboiler/handler"
)

func TestRepimplPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/repositoryパッケージのテスト")
}

// 全テストが実行される前に1度だけ実行される関数
var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../config/database.toml")
	// database.tomlファイルにパスを環境変数に設定する
	os.Setenv("DATABSE_TOML_PATH", absPath)
	err := handler.DBConnect() // データベースに接続する
	Expect(err).NotTo(HaveOccurred(), "データベース接続が失敗したのでテストを中止します。")
})