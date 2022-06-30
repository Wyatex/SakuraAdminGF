package hook

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

type opLog struct {
	pathMode    int
	unAuthLog   bool
	includePath g.SliceStr
	excludePath g.SliceStr
}

var (
	ctx         = gctx.New()
	opLogConfig = opLog{}
)

func init() {
	pathMode := g.Cfg().MustGet(ctx, "operationLog.pathMode").Int()
	if pathMode != 1 && pathMode != 2 {
		panic(gerror.New("operationLog.pathMode只能为1或2"))
	}
	opLogConfig = opLog{
		pathMode:    pathMode,
		unAuthLog:   g.Cfg().MustGet(ctx, "operationLog.unAuthLog").Bool(),
		includePath: g.Cfg().MustGet(ctx, "operationLog.includePath").Strings(),
		excludePath: g.Cfg().MustGet(ctx, "operationLog.excludePath").Strings(),
	}
}

// OperationLog 日志记录钩子
func OperationLog(r *ghttp.Request) {
	// 异步处理
	go func() {
		// 判断是否需要记录日志
		if !ifNeedLog(r) {
			return
		}

		// 生成记录
		status := 0
		errMsg := ""
		if r.GetError() != nil {
			status = 1
			errMsg = r.GetError().Error()
		}
		opLog := &model.SysOpLogInput{
			Method:       r.Request.Method,
			UserId:       r.GetCtxVar("uid").Uint64(), // 0说明未登录用户
			Ip:           r.Request.RemoteAddr,
			Path:         r.URL.Path,
			Status:       status,
			Request:      gjson.New(r.GetRequestMap()).MustToJsonString(),
			Response:     r.GetCtxVar("resJson").String(),
			ErrorMessage: errMsg,
		}

		// 写入数据库
		err := service.SysOpLog().Add(r.GetCtx(), opLog)
		if err != nil {
			g.Log().Error(r.GetCtx(), err)
		}
	}()
}

func ifNeedLog(r *ghttp.Request) bool {
	// 请求用户id
	UserId := r.GetCtxVar("uid").Uint64()
	if !opLogConfig.unAuthLog && UserId == 0 {
		// 如果当前请求的用户未登录而且配置不记录未登录用户的操作
		return false
	}

	// 请求路径
	urlPath := r.URL.Path
	// 去除后斜杠
	if strings.HasSuffix(urlPath, "/") {
		urlPath = gstr.SubStr(urlPath, 0, len(urlPath)-1)
	}

	// 判断路径是否需要日志记录
	if opLogConfig.pathMode == 1 {
		isInclude := false
		// 包含路径处理
		for _, includePath := range opLogConfig.includePath {
			tmpPath := includePath
			if strings.HasSuffix(tmpPath, "/*") {
				// 前缀匹配
				tmpPath = gstr.SubStr(tmpPath, 0, len(tmpPath)-2)
				if gstr.HasPrefix(urlPath, tmpPath) {
					isInclude = true
				}
			} else {
				// 全路径匹配
				if strings.HasSuffix(tmpPath, "/") {
					tmpPath = gstr.SubStr(tmpPath, 0, len(tmpPath)-1)
				}
				if urlPath == tmpPath {
					// 全路径匹配不拦截
					isInclude = true
				}
			}
		}
		// 匹配不存在，说明不记录
		return isInclude
	} else if opLogConfig.pathMode == 2 {
		// 排除路径处理
		for _, excludePath := range opLogConfig.excludePath {
			tmpPath := excludePath
			// 前缀匹配
			if strings.HasSuffix(tmpPath, "/*") {
				tmpPath = gstr.SubStr(tmpPath, 0, len(tmpPath)-2)
				if gstr.HasPrefix(urlPath, tmpPath) {
					// 前缀匹配直接退出
					return false
				}
			} else {
				// 全路径匹配
				if strings.HasSuffix(tmpPath, "/") {
					tmpPath = gstr.SubStr(tmpPath, 0, len(tmpPath)-1)
				}
				if urlPath == tmpPath {
					// 全路径匹配不拦截
					return false
				}
			}
		}
	}
	return true
}
