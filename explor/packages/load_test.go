package packages

import (
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestLoad(t *testing.T) {
	// Dir 参数指定执行查询操作时(run the build system's query tool)的目录, 为空表示当前目录
	cfg := &packages.Config{Mode: packages.LoadAllSyntax, Dir: ""}
	// patterns 包模式，指定要加载的包。默认情况下和go list/build/test等命令的包模式相同:
	// ./... 当前目录及子目录下的所有包
	// github.com/user/repo/... 目录及子目录下的所有包
	// std 标准库的所有包
	// all 所有包
	// 包模式是一种指定包列表的抽象概念
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		t.Fatal(err)
	}
	for _, pkg := range pkgs {
		t.Log(pkg.Name, pkg.PkgPath)
	}
}
