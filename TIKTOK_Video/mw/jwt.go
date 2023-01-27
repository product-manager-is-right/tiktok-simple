package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"log"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Id       int64
	Username string
	Password string
}

var JwtMiddleware *jwt.HertzJWTMiddleware

const IdentityKey = "userId"

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "tiktok",
		Key:         []byte("jwt sign key"),
		SendCookie:  true,
		CookieName:  "jwt-cookie",
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		TokenLookup: "query: token, cookie: jwt",
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return int64(claims[IdentityKey].(float64))
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			if !strings.HasPrefix(string(c.Request.RequestURI()), "/douyin/feed/") {
				c.JSON(http.StatusOK, utils.H{
					"status_code": 1,
					"status_msg":  "jwt authorize fail",
				})
			}
			c.Abort()
		},
	})
	if err != nil {
		log.Fatal("jwt init fail")
	}
}
