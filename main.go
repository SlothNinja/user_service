package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"bitbucket.org/SlothNinja/restful"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	DEV               = "DEV"
	TRUE              = "true"
	PORT              = "PORT"
	DefaultPort       = "8080"
	sessionName       = "sngsession"
	sessionSecret     = "verySecretiveSecret1234!"
	rootPath          = "/"
	homePath          = "/"
	userPath          = "/user"
	staticPath        = "/"
	staticDir         = "dist"
	day               = time.Hour * 24
	week              = day * 7
	envDevClientCreds = "DEV_CLIENT_CREDENTIALS"
	envClientCreds    = "CLIENT_CREDENTIALS"
)

func main() {
	setGinMode()
	r := newRouter(newCookieStore())

	addRoutes(rootPath, r, getOAuth2Config())

	port := getPort()

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func getPort() string {
	port := os.Getenv(PORT)
	if port != "" {
		return port
	}
	log.Printf("Defaulting to port %s", DefaultPort)
	return DefaultPort
}

// staticHandler for local development since app.yaml is ignored
// static files are handled via app.yaml routes when deployed
func staticRoutes(r *gin.Engine) {
	if isDev() {
		r.StaticFile("/", "dist/index.html")
		r.StaticFile("/app.js", "dist/app.js")
		r.StaticFile("/favicon.ico", "dist/favicon.ico")
		r.Static("/img", "dist/img")
	}
}

func newRouter(store cookie.Store) *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		restful.CTXHandler(),
		restful.TemplateHandler(r),
		sessions.Sessions(sessionName, store),
	)
	staticRoutes(r)
	return r
}

func setGinMode() {
	if isDev() {
		gin.SetMode(gin.DebugMode)
		return
	}

	gin.SetMode(gin.ReleaseMode)
}

func newCookieStore() cookie.Store {
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		Path:   rootPath,
		MaxAge: int(week.Truncate(time.Second)),
	})
	return store
}

func getOAuth2Config() *oauth2.Config {
	if isDev() {
		return oauth2Config(envDevClientCreds, "email", "profile", "openid")
	}

	return oauth2Config(envClientCreds, "email", "profile", "openid")
}

func jsonKey(env string) []byte {
	path := os.Getenv(env)
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return key
}

func isDev() bool {
	return os.Getenv("DEV") == "true"
}

func oauth2Config(env string, scope ...string) *oauth2.Config {
	conf, err := google.ConfigFromJSON(jsonKey(env), scope...)
	if err != nil {
		panic(err)
	}
	return conf
}
