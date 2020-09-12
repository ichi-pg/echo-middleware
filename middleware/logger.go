package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/ichi-pg/echo-middleware/contexts"
	"github.com/ichi-pg/echo-middleware/env"
	"github.com/ichi-pg/echo-middleware/header"
	"github.com/ichi-pg/echo-middleware/presenter"
	"github.com/ichi-pg/echo-middleware/util"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const accessLog = "accessLog"

// Logger はリクエスト/レスポンス情報をログエントリーに追加します。
func Logger(ep presenter.ErrorPresenter) echo.MiddlewareFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})

	{
		log := logger.WithField("env", os.Environ())
		log.Info("initializeLog")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			req := c.Request()
			log := logger.WithFields(logrus.Fields{
				"labels": logrus.Fields{
					"projectID": os.Getenv(env.ProjectID),
					"platform":  req.Header.Get(header.Platform),
					"userID":    req.Header.Get(header.UserID),
					"requestID": req.Header.Get(echo.HeaderXRequestID),
					"loggerID":  uuid.NewV4().String(),
				},
			})

			{
				bodyBytes := []byte{}
				var bodyJSON logrus.Fields
				if req.Body != nil {
					bodyBytes, _ = ioutil.ReadAll(req.Body)
					json.Unmarshal(bodyBytes, &bodyJSON)
				}

				var bodyString string
				if bodyJSON == nil {
					bodyString = string(bodyBytes)
				}

				params := map[string]string{}
				for _, v := range c.ParamNames() {
					params[v] = c.Param(v)
				}

				headers := map[string]string{}
				for k, v := range req.Header {
					if len(v) > 0 {
						headers[k] = v[0]
					}
				}

				log = log.WithField("request", logrus.Fields{
					"schema":   req.URL.Scheme,
					"proto":    req.Proto,
					"host":     req.Host,
					"method":   req.Method,
					"path":     c.Path(),
					"params":   params,
					"query":    c.QueryParams(),
					"body":     bodyString,
					"json":     bodyJSON,
					"size":     req.ContentLength,
					"remoteIP": c.RealIP(),
					"header":   headers,
				})
				contexts.SetLogger(c, log)

				req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			{
				res := c.Response()
				bodyBuf := new(bytes.Buffer)
				res.Writer = util.BodyDumpResponseWriter(bodyBuf, res.Writer)

				err := next(c)
				if err != nil {
					ep.Response(c, err)
				}

				var bodyJSON map[string]interface{}
				json.Unmarshal(bodyBuf.Bytes(), &bodyJSON)

				var bodyString string
				if bodyJSON == nil {
					bodyString = string(bodyBuf.Bytes())
				}

				var errorString string
				if err != nil {
					errorString = err.Error()
				}

				log = log.WithField("response", logrus.Fields{
					"status":  res.Status,
					"body":    bodyString,
					"json":    bodyJSON,
					"size":    res.Size,
					"error":   errorString,
					"latency": int64(time.Now().Sub(start)),
					"header":  res.Header(),
				})

				switch res.Status / 100 {
				case 5:
					log.Error(accessLog)
				case 4:
					log.Warn(accessLog)
				case 2:
					log.Info(accessLog)
				}
				return nil
			}
		}
	}
}
