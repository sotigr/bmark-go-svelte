package main

import (
	"context"
	"main/api"
	"main/emitter"
	"main/lib/system"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	storage "cloud.google.com/go/storage"
)

func main_middleware(bucket *storage.BucketHandle, c_emitter *emitter.HTTPEmitter) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Set("bucket", bucket)
		c.Set("emitter", c_emitter)

		c.Next()

	}
}

func main() {
	env := system.GetEnv()
	cn := 0

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	bkt := client.Bucket(os.Getenv("BUCKET_NAME"))

	server := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Authorization", "User-Agent", "Content-Type"}

	server.Use(cors.New(corsConfig))
	// Logging middleware
	if env == system.ENV_DEV {
		server.Use(gin.Logger())
	}

	c_emitter := emitter.HTTPEmitter{}.Init(server, "/sys/event-buffer")

	lobby_pass := string("123")
	c_emitter.CreateLobby("test_lobby", &lobby_pass)

	server.Use(main_middleware(bkt, c_emitter))

	server.GET("/test", func(c *gin.Context) {

		cn += 1
		c_emitter.Send("test_lobby", "test_event", "datass")
		// c_emitter.Send_all("test_event", "hello from test "+strconv.Itoa(cn))

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/stop", func(c *gin.Context) {
		emitter.CloseEmitter(c_emitter)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	server.GET("/list", api.List)
	server.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/app")
	})
	server.StaticFS("/app", gin.Dir("public", false))
	// Run server

	server.Run("0.0.0.0:80")

}
