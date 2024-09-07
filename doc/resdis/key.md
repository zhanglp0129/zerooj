# key
redis key相关文档，主要包括各种情况的key格式。所有键均采用蛇形命名

## 邮箱验证码
- key格式：**mail_check_code:验证码用途:邮箱**
- 表格

|   用途    |               redis键                |
|:-------:|:-----------------------------------:|
|   注册    |     mail_check_code:register:邮箱     |
|  忘记密码   | mail_check_code:forget_password:邮箱  |
| 修改邮箱(旧) | mail_check_code:update_email_old:邮箱 |
| 修改邮箱(新) | mail_check_code:update_email_new:邮箱 |

## 缓存
- key格式：**cache:缓存描述:相关标识**
- 表格

|   缓存描述    |              redis键               |
|:---------:|:---------------------------------:|
|   用户权限    |     cache:user_permission:id      |
|  用户关注数量   |  cache:user_followings_count:id   |
|  用户粉丝数量   |     cache:user_fans_count:id      |
