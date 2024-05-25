package result

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	socketio "github.com/googollee/go-socket.io"
	"github.com/syyongx/php2go"
)

var ResultService = ResultServiceStruct{}

type ResultServiceStruct struct {
	Socket socketio.Conn
}

type ResultMap struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data g.Map  `json:"data"`
}

//正确返回
func (this ResultServiceStruct) SuccessMsg(msg ...string) string {
	msgStr := "操作成功"
	if len(msg) > 0 {
		msgStr = msg[0]
	}
	return this.Result(200, msgStr, g.Map{})
}

//正确返回
func (this ResultServiceStruct) Success(data g.Map) string {
	return this.Result(200, "操作成功", data)
}

//错误返回
func (this ResultServiceStruct) Error(msg string) string {
	return this.Result(500, msg, g.Map{})
}

//警告返回
func (this ResultServiceStruct) Warning(msg string) string {
	return this.Result(501, msg, g.Map{})
}

//未登录
func (this ResultServiceStruct) Unauthorized(msg string) string {
	return this.Result(401, msg, g.Map{})
}

//幂等处理
func (this ResultServiceStruct) Idempotent() string {
	return this.Result(507, "请等待上次请求的响应", g.Map{})
}

//湖区响应结构体
func (this ResultServiceStruct) Result(code int, msg string, data g.Map) string {
	return gconv.String(data)
	//return gconv.String(ResultMap{
	//	Code: code,
	//	Msg:  msg,
	//	Data: this.Deal(data),
	//})
}

//处理字符串
func (this ResultServiceStruct) Deal(data g.Map) g.Map {

	for k, v := range data {
		switch v.(type) {
		case map[string]interface{}:
			v = this.Deal(gconv.Map(v))
			break
		case []map[string]interface{}:
			vGmaps := gconv.Maps(v)
			for key, value := range vGmaps {
				vGmaps[key] = this.Deal(value)
			}
			v = vGmaps
			break
		case []string, []int, []int64:
			vInterfaces := gconv.Interfaces(v)
			for kk, vv := range vInterfaces {
				vStr := gconv.String(vv)
				if php2go.MbStrlen(vStr) > 11 || vStr == "" {
					vv = vStr
				}
				vInterfaces[kk] = vv
			}
			v = vInterfaces
			break
		default:
			vStr := gconv.String(v)
			if php2go.MbStrlen(vStr) > 11 || vStr == "" {
				v = vStr
			}
		}
		data[k] = v
	}
	return data
}
