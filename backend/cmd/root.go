package cmd

import (
	"Traveloka/common"
	"Traveloka/middleware"
	"fmt"
	"log"

	userStorage "Traveloka/modules/user/storage/postgreSQL"
	ginuser "Traveloka/modules/user/transport/ginuser/postgreSQL"
	"Traveloka/plugin/simple"
	"Traveloka/plugin/tokenprovider/jwt"
	"net/http"
	"os"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/plugin/storage/sdkgorm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("traveloka"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("Traveloka", common.PluginDBMain)),
		// goservice.WithInitRunnable(jwt.NewTokenJWTProvider(common.PluginJWT)),
		goservice.WithInitRunnable(simple.NewSimplePlugin("simple")),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start social TODO service",
	Run: func(cmd *cobra.Command, args []string) {

		service := newService()

		serviceLogger := service.Logger("service")
		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recovery())

			type CanGetValue interface {
				GetValue() string
			}

			log.Println(service.MustGet("simple").(CanGetValue).GetValue())
			serviceDb := service.MustGet(common.PluginDBMain).(*gorm.DB)
			authStore := userStorage.NewSQLStore(serviceDb)
			tokenprovider := jwt.NewTokenJWTProvider("jwt", os.Getenv("JWT_SECRET_KEY"))
			middlewareAuth := middleware.RequiredAuth(authStore, tokenprovider)

			engine.Static("/static", "./static")
			v1 := engine.Group("/v1")
			{
				users := v1.Group("/users")
				{
					users.POST("/register", ginuser.Register(serviceDb))
					users.POST("/login", ginuser.Login(serviceDb, tokenprovider))
					users.GET("/profile", middlewareAuth, ginuser.Profile())
				}
			}

			engine.GET("/ping", func(c *gin.Context) {
				go func() {
					defer common.Recovery()
					fmt.Println([]int{}[0])
				}()
				c.JSON(http.StatusOK, gin.H{
					"message": "Hello World",
				})
			})
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)

		}

	},
}

func Excucte() {
	rootCmd.AddCommand(outEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
