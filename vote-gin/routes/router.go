package routes

import (
	"net/http"
	"vote-gin/api/v1"
	"vote-gin/middleware"
	"vote-gin/model"
	"vote-gin/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger
var err error

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	sugar = utils.Logger.Sugar()
	defer sugar.Sync()
	defer model.Close() // 关闭数据库连接
	gin.SetMode(utils.AppMode)
	// 强制日志颜色化
	gin.ForceConsoleColor()

	r := gin.New()
	// 设置信任网络
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)
	// r.HTMLRender = createMyRender()

	// 添加中间件
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//r.Static("/static", "./web/front/dist/static")
	//r.Static("/admin", "./web/admin/dist")
	//r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "front", nil)
	})

	r.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin", nil)
	})
	// 后台管理路由接口
	auth := r.Group("api/v1")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		// 投票管理
		auth.POST("vote/opt/add", v1.AddVoteOpt)
	}

	// 前台管理路由接口
	router := r.Group("api/v1")
	{
		// 登录控制
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 投票相关
		// router.GET("vote", v1.GetVotes)
		router.GET("vote/info/:id", v1.GetVoteInfo)
	}

	if err = r.Run(utils.HttpPort); err != nil {
		sugar.Errorf("%d端口启动失败: %w", utils.HttpPort, err)
		panic("http server port is used")
	}
}
