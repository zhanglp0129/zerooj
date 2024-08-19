# key
redis key相关文档，主要包括各种情况的key格式。所有键均采用蛇形命名

## 邮箱验证码
- key格式：**/mail_check_code/服务名称/接口名称/邮箱**
- 表格

|  服务  |  接口  |                 redis键                 |
|:----:|:----:|:--------------------------------------:|
| 用户服务 | 注册接口 | /mail_check_code/user/user_register/邮箱 |

## 缓存
- key格式：**/cache/服务名称/接口名称/相关信息**
- 表格

|  服务  |    接口     |                 redis键                  |
|:----:|:---------:|:---------------------------------------:|
| 用户服务 | 获取用户基本信息  |      /cache/user/get_base_info/id       |
| 用户服务 | 根据用户名搜索用户 | /cache/user/search_by_username/username |

## 访问频率限制
- key格式：**/frequency_limit/服务名称/接口名称/时间/相关信息**
- 表格

|  服务  |   接口    |                        redis键                        |  描述   |
|:----:|:-------:|:----------------------------------------------------:|:-----:|
| 邮箱服务 | 发送邮件验证码 | /frequency_limit/mail/send_mail_check_code/minute/ip | 1分钟1次 |
| 邮箱服务 | 发送邮件验证码 |  /frequency_limit/mail/send_mail_check_code/hour/ip  | 1小时5次 |
| 邮箱服务 | 发送邮件验证码 |  /frequency_limit/mail/send_mail_check_code/day/ip   | 1天10次 |
| 用户服务 | 修改用户名接口 |    /frequency_limit/user/update_username/week/id     | 7天1次  |
