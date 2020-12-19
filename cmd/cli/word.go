package cli

import (
	"github.com/ChaLanZi/gotools/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_UPPER                         = iota + 1 //全部单词转为大写模式
	MODE_LOWER                                    //全部单词转为小写模式
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE            //下画线单词转为大写驼峰模式
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE            //下画线单词转为小写驼峰模式
	MODE_CAMELCASE_TO_UNDERSCORE                  //驼峰模式转为下画线模式
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1:全部单词转为大写模式",
	"2:全部单词转为小写模式",
	"3:下画线单词转为大写驼峰模式",
	"4:下画线单词转为小写驼峰模式",
	"5:驼峰模式转为下画线模式",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该模式转换，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果： %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
