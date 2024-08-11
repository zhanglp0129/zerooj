# ZeroOJ

基于go-zero框架开发的OJ在线评测系统

## 目录结构

## 仓库规范

- 分支类型

|  分支类型   |     分支名     |            描述            |
|:-------:|:-----------:|:------------------------:|
|  主干分支   |   master    |   作为稳定主分支，用于部署生产环境的分支    |
|  开发分支   |   develop   |   平时开发用的主分支，永远是功能最全最新    |
|  功能分支   |  feature/*  |      一个事项卡对应一个功能分支       |
|  发布分支   |  release/*  |     一次新版本的发布对应一个发布分支     |
|  热修复分支  |  hotfix/*   |  从主干分支拉出，用于线上版本的 Bug 修复  |

- 合并方向

|   源分支   |  目标分支  |          图示          |
|:-------:|:------:|:--------------------:|
|  功能分支   |  开发分支  | feature/* -> develop |
|  发布分支   |  主干分支  | release/* -> master  |
|  发布分支   |  开发分支  | release/* -> develop |
|  热修复分支  |  主干分支  |  hotfix/* -> master  |
|  热修复分支  |  开发分支  | hotfix/* -> develop  |