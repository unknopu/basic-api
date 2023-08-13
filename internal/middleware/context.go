package middleware

import (
	"basic-api/internal/core/context"
	db "basic-api/internal/core/database"
	"bytes"
	ctx "context"
	"io/ioutil"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// CustomContext custom context
func CustomContext(db *gorm.DB, log *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			cc := &context.Context{
				Context: c,
				Db:      db,
				Log:     log.WithField("id", ""),
			}
			return next(cc)
		}
	}
}

func LogRecorder(mgdb *mongo.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, cancel := ctx.WithTimeout(ctx.Background(), 20*time.Second)
			defer cancel()

			var body []byte
			if c.Request().Header.Get("Content-Type") == echo.MIMEApplicationJSON {
				data, err := ioutil.ReadAll(c.Request().Body)
				if err == nil {
					body = data
					c.Request().Body = ioutil.NopCloser(bytes.NewReader(data))
				}
			}

			m := &db.LogModel{
				IP:       c.RealIP(),
				Method:   c.Request().Method,
				Uri:      c.Request().RequestURI,
				BodyJson: body,
			}
			m.Stamp()
			_, _ = mgdb.Collection("kg_log").InsertOne(ctx, m)

			return next(c)
		}
	}
}
