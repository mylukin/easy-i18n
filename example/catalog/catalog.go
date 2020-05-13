package catalog

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// init
func init() {
	initEn(language.Make("en"))
	initZhHans(language.Make("zh-Hans"))
}

// initEn will init en support.
func initEn(tag language.Tag) {
	message.SetString(tag, "%s has %d cat.", "%s has %d cat.")
	message.SetString(tag, "%s has %d cats.", "%s has %d cats.")
	message.SetString(tag, "%s have %d apples.", "%s have %d apples.")
	message.SetString(tag, "%s have an apple.", "%s have an apple.")
	message.SetString(tag, "%s have two apples.", "%s have two apples.")
	message.SetString(tag, "hello %s!", "hello %s!")
	message.SetString(tag, "hello world!", "hello world!")
}

// initZhHans will init zh-Hans support.
func initZhHans(tag language.Tag) {
	message.SetString(tag, "%s has %d cat.", "%s有%d只猫。")
	message.SetString(tag, "%s has %d cats.", "%s有%d只猫。")
	message.SetString(tag, "%s have %d apples.", "%s有%d个苹果。")
	message.SetString(tag, "%s have an apple.", "%s有一个苹果。")
	message.SetString(tag, "%s have two apples.", "%s有两个苹果。")
	message.SetString(tag, "hello %s!", "你好%s！")
	message.SetString(tag, "hello world!", "你好世界！")
}
