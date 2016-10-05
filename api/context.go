package api

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/netlify/netlify-subscriptions/conf"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

const (
	dbKey        = "db"
	startTimeKey = "start_time"
	versionKey   = "app_version"
	configKey    = "app_config"
	loggerKey    = "app_logger"
	reqIDKey     = "request_id"
	adminFlagKey = "admin_flag"
	tokenKey     = "token"
)

func setStartTime(ctx context.Context, startTime time.Time) context.Context {
	return context.WithValue(ctx, startTimeKey, &startTime)
}
func getStartTime(ctx context.Context) *time.Time {
	obj := ctx.Value(startTimeKey)
	if obj == nil {
		return nil
	}
	return obj.(*time.Time)
}

func setVersion(ctx context.Context, version string) context.Context {
	return context.WithValue(ctx, versionKey, version)
}
func getVersion(ctx context.Context) string {
	obj := ctx.Value(versionKey)
	if obj == nil {
		return ""
	}
	return obj.(string)
}

func setConfig(ctx context.Context, config *conf.Config) context.Context {
	return context.WithValue(ctx, configKey, config)
}
func getConfig(ctx context.Context) *conf.Config {
	obj := ctx.Value(configKey)
	if obj == nil {
		return nil
	}
	return obj.(*conf.Config)
}

func setLogger(ctx context.Context, log *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}
func getLogger(ctx context.Context) *logrus.Entry {
	obj := ctx.Value(loggerKey)
	if obj == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}
	return obj.(*logrus.Entry)
}

func setRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, reqIDKey, reqID)
}
func getRequestID(ctx context.Context) string {
	obj := ctx.Value(reqIDKey)
	if obj == nil {
		return ""
	}
	return obj.(string)
}

func setAdminFlag(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, adminFlagKey, isAdmin)
}

func isAdmin(ctx context.Context) bool {
	obj := ctx.Value(adminFlagKey)
	if obj == nil {
		return false
	}
	return obj.(bool)
}

func setDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}
func getDB(ctx context.Context) *gorm.DB {
	return ctx.Value(dbKey).(*gorm.DB)
}

func getClaims(ctx context.Context) *JWTClaims {
	return ctx.Value(tokenKey).(*jwt.Token).Claims.(*JWTClaims)
}

func setToken(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}
