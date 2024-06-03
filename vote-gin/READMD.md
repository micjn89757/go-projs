## 环境
vue
gin
sqlx
go 1.21.6
mysql
redis

## 项目结构
```shell
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  LICENSE
│  main.go //主程序入口
│  README.md
├─app
|   └─v1  存放handler
├─config // 项目配置入口   
├─database  // 数据库备份文件（初始化）
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│      router.go // 路由入口    
├─upload   // 图片等
├─utils // 项目公用工具库
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web // 前端开发源码（VUECLI项目源文件)
    ├─admin             
    └─front
```


一个包下一个Logger

├ alt+195
─ alt+196
└ alt+192