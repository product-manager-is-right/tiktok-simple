package mw

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model"
	"TIKTOK_User/util"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

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
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			username := c.Query("username")
			password := c.Query("password")

			users, err := mysql.CheckUser(username, util.MD5(password))
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, jwt.ErrFailedAuthentication
			}
			c.Set("user_id", users[0].Id)
			return users[0], nil
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": 0,
				"status_mgs":  "login success",
				"user_id":     c.Value("user_id").(int64),
				"token":       token,
			})
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return int64(claims[IdentityKey].(float64))
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
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
			if e == jwt.ErrFailedAuthentication {
				return "账号或错误，请重新输入!"
			}
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": 1,
				"status_msg":  message,
			})
			// 鉴权失败，中断handler
			c.Abort()
		},
	})
	if err != nil {
		log.Fatal("jwt init fail")
	}
}
