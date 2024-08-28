package utils

import (
	"strings"
	"unicode"
)

// ToSnakeCase 驼峰命名转蛇形命名
func ToSnakeCase(str string) string {
	var snake strings.Builder
	for i, r := range str {
		if unicode.IsUpper(r) {
			// 如果是大写字母，且不是第一个字母，则在前面加下划线
			if i > 0 {
				snake.WriteRune('_')
			}
			// 将大写字母转换为小写
			snake.WriteRune(unicode.ToLower(r))
		} else {
			// 非大写字母直接添加
			snake.WriteRune(r)
		}
	}
	return snake.String()
}
