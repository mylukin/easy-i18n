package catalog

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// init
func init() {
	initEn(language.Make("en"))
	initZhHans(language.Make("zh-Hans"))
	initZhHant(language.Make("zh-Hant"))
}

// initEn will init en support.
func initEn(tag language.Tag) {
	message.SetString(tag, "%s extract [path] [outfile]", "%s extract [path] [outfile]")
	message.SetString(tag, "%s generate [path] [outfile]", "%s generate [path] [outfile]")
	message.SetString(tag, "%s update srcfile destfile", "%s update srcfile destfile")
	message.SetString(tag, "a tool for managing message translations.", "a tool for managing message translations.")
	message.SetString(tag, "destfile cannot be empty", "destfile cannot be empty")
	message.SetString(tag, "extracts strings to be translated from code", "extracts strings to be translated from code")
	message.SetString(tag, "generated go file package name", "generated go file package name")
	message.SetString(tag, "generates code to insert translated messages", "generates code to insert translated messages")
	message.SetString(tag, "merge translations and generate catalog", "merge translations and generate catalog")
	message.SetString(tag, "srcfile cannot be empty", "srcfile cannot be empty")
}

// initZhHans will init zh-Hans support.
func initZhHans(tag language.Tag) {
	message.SetString(tag, "%s extract [path] [outfile]", "%s 提取 [路径] [输出文件]")
	message.SetString(tag, "%s generate [path] [outfile]", "%s 生成 [路径] [输出文件]")
	message.SetString(tag, "%s update srcfile destfile", "%s 更新 源文件 输出文件")
	message.SetString(tag, "a tool for managing message translations.", "用于管理消息翻译的工具。")
	message.SetString(tag, "destfile cannot be empty", "输出文件不能为空")
	message.SetString(tag, "extracts strings to be translated from code", "从代码中提取要翻译的字符串")
	message.SetString(tag, "generated go file package name", "生成的go文件包名称")
	message.SetString(tag, "generates code to insert translated messages", "生成代码以插入翻译后的消息")
	message.SetString(tag, "merge translations and generate catalog", "合并翻译并生成目录")
	message.SetString(tag, "srcfile cannot be empty", "源文件不能为空")
}

// initZhHant will init zh-Hant support.
func initZhHant(tag language.Tag) {
	message.SetString(tag, "%s extract [path] [outfile]", "%s 提取 [路徑] [輸出文件]")
	message.SetString(tag, "%s generate [path] [outfile]", "%s 生成 [路徑] [輸出文件]")
	message.SetString(tag, "%s update srcfile destfile", "%s 更新 源文件 輸出文件")
	message.SetString(tag, "a tool for managing message translations.", "用於管理消息翻譯的工具。")
	message.SetString(tag, "destfile cannot be empty", "輸出文件不能為空")
	message.SetString(tag, "extracts strings to be translated from code", "從代碼中提取要翻譯的字符串")
	message.SetString(tag, "generated go file package name", "生成的go文件包名稱")
	message.SetString(tag, "generates code to insert translated messages", "生成代碼以插入翻譯後的消息")
	message.SetString(tag, "merge translations and generate catalog", "合併翻譯並生成目錄")
	message.SetString(tag, "srcfile cannot be empty", "源文件不能為空")
}
