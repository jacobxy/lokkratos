package generate

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	_ "embed"

	"github.com/go-kratos/kratos/cmd/kratos/v2/pkg"
	"github.com/spf13/cobra"
)

// CmdNew represents the new command.
var CmdNew = &cobra.Command{
	Use:   "generate",
	Short: "Create a biz/data template",
	Long:  "Create a layer template  Example: kratos generate -l biz -n User -t internal/biz/user.go",
	Run:   run,
}

var (
	// biz data
	layer  string
	name   string
	target string
)

func init() {
	CmdNew.Flags().StringVarP(&layer, "layer", "l", layer, "layername")
	CmdNew.Flags().StringVarP(&name, "name", "n", name, "objectName")
	CmdNew.Flags().StringVarP(&target, "target", "t", target, "target dir")
}

//go:embed biz.tmpl
var bizTemplate string

//go:embed data.tmpl
var dataTemplate string

type generateFileParams struct {
	Name    string
	ModName string
}

func run(cmd *cobra.Command, args []string) {
	params := generateFileParams{
		Name:    name,
		ModName: pkg.GetModName(),
	}
	var temp *template.Template
	switch layer {
	case "biz":
		temp, _ = template.New("temp").Parse(bizTemplate)
	case "data":
		temp, _ = template.New("temp").Parse(dataTemplate)
	default:
		log.Println("layer not found")
	}

	file, err := CreateFile(target)
	if err != nil {
		log.Println("error happen ", err)
		return
	}
	defer file.Close()

	err = temp.Execute(file, params)
	fmt.Println(err)
}

func CreateFile(filePath string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsExist(err) {
		return nil, err
	}
	os.MkdirAll(path.Dir(filePath), os.ModePerm)

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, err
}
