# object
minio对象存储相关文档

## 说明
1. 使用uuid作为对象名；对于可以匿名读对象，使用readonly/uuid作为对象名
2. 当内容大于255字节时，使用对象存储；小于等于255字节时，直接存储在数据库，并用一个字段标明使用对象存储还是直接存储