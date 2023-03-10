package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
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
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		TokenLookup: "query: token, form: token",
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
			if e == jwt.ErrExpiredToken {
				return "token已过期，请重新登录!"
			}
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			if !strings.HasPrefix(string(c.Request.RequestURI()), "/douyin/feed/") {
				c.JSON(http.StatusOK, utils.H{
					"status_code": 1,
					"status_msg":  message,
				})
			}
			c.Abort()
		},
	})
	if err != nil {
		log.Fatal("jwt init fail")
	}
}
