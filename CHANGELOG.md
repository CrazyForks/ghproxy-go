# 更新日志

v0.2.0
---
底层核心代码重写,彻底代表着项目进入自主可控的开发阶段,彻底脱离原有实现
- ADD: 增加多处日志记录,便于审计与排障
- CHANGE: 优化代码结构,进一步模块化,同时提升性能
- ADD： 使用req库重构代码,提升请求伪装能力,尽可能bypass反爬机制

24w05b
---
- CHANGE: 重命名proxychrome函数
- ADD: 增加多处日志记录,便于审计与排障

24w05a
---
- FIX： 修正上一版本的req请求未继承请求方式的问题
- CHANGE: 优化代码结构,进一步模块化,同时提升性能

24w04b
---
- CHANGE: 更换Docker基础镜像为daily版本
- ADD： 新增使用req库实现代理请求,使用chrome TLS指纹发起请求

24w04a
---
- CHANGE: 调整程序结构,使用init函数初始化配置,并优化代码结构

v0.1.7
---
- CHANGE: 合入上游(wjqserver/caddy:latest)安全更新, 增强镜像安全性

24w03b
---
- CHANGE: 合入上游(wjqserver/caddy:latest)安全更新, 增强镜像安全性

v0.1.6
---
- ADD: 新增跨域配置选项
- CHANGE: 更新UA标识

24w03a
---
- CHANGE: 改进Docker代理相关Caddy配置
- ADD: 新增跨域配置选项

v0.1.5
---
- CHANGE: 更新Go版本至v1.23.1
- CHANGE: 优化代码行为

24w02b
---
- ADD: 新增Docker代理 (未并入正式版)

24w02a
---
- CHANGE: 更新Go版本至v1.23.1
- CHANGE: 优化代码行为

v0.1.4
---
正式版24w01b内容更新
- ADD: 新增外部文件配置功能
- ADD: 新增日志功能
- CHANGE: 优化代码结构,提升性能
- CHANGE: 改进前端界面,加入页脚

24w01b
---
标志着项目正式进入自主开发阶段
- ADD: 新增外部文件配置功能
- ADD: 新增日志功能
- CHANGE: 优化代码结构,提升性能
- CHANGE: 改进前端界面,加入页脚

v0.1.3
---
开始自行维护项目,脱离上游更新
- CHANGE: 改进已有实现,增强程序稳定性

24w01a
---
首个DEV版本
- CHANGE: 同步更新

v0.1.2
---
- ADD: 新增项目介绍
- CHANGE: 限制默认文件大小限制到256M

v0.1.1
---
- ADD: Apache License Version 2.0
- FIX: 改进部分代码逻辑
- CHANGE: 将Go升级至v1.23.0

v0.1.0
---
项目的第一个版本
- ADD: 实现速率限制
- ADD: 实现符合[RFC 7234](https://httpwg.org/specs/rfc7234.html)的HTTP缓存机制
- ADD: 实现action编译
- ADD: 实现Docker部署
- INFO: 使用Caddy作为Web服务器，通过Caddy实现了缓存与速率限制
