package utils

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	camelCase1 := "CamelCaseToSnakeCase"
	camelCase2 := "camelCaseToSnakeCase"
	snakeCase1 := ToSnakeCase(camelCase1)
	snakeCase2 := ToSnakeCase(camelCase2)
	if snakeCase1 != "camel_case_to_snake_case" {
		t.Fatal("蛇形命名转换失败")
	}
	if snakeCase2 != "camel_case_to_snake_case" {
		t.Fatal("蛇形命名转换失败")
	}
}
