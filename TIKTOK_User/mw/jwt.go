package mw

import (
	"GoProject/dal/mysql"
	"GoProject/model"
	"GoProject/util"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "user_name"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "tiktok",
		Key:         []byte("jwt sign key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		TokenLookup: "query: token",
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			username := c.PostForm("username")
			password := c.PostForm("password")

			users, err := mysql.CheckUser(username, util.MD5(password))
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, errors.New("user already exists or wrong password")
			}
			c.Set("user_id", users[0].Id)
			return users[0], nil
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": 1,
				"status_mgs":  "login success",
				"user_id":     c.Value("user_id").(int64),
				"token":       token,
			})
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.User{
				Username: claims[IdentityKey].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": 1,
				"status_msg":  "jwt authorize fail",
			})
		},
	})
	if err != nil {
		log.Fatal("jwt init fail")
	}
}
