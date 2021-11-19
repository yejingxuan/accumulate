package dto

import (
	"github.com/gin-gonic/gin"
)

const (
	Success       = 0
	Error         = 500   //服务内部错误
	ErrBadRequest = 20400 //请求参数错误

	ErrMd5Calculate      = 21001 //计算MD5失败
	ErrFileCompress      = 21002 //文件解压失败
	ErrUnmarshalMata     = 21003 //解析metadata.json文件失败
	ErrSaveOp            = 21004 //算子保存失败
	ErrFindOp            = 21005 //找不到该算子
	ErrUnmarshalFileInfo = 21006 //文件元信息解析失败
	ErrFindFileInfo      = 21007 //找不到文件信息
	ErrMoveFile          = 21008 //文件移动失败
	ErrCheckMata         = 21009 //校验metadata.json文件失败
	ErrDelOp             = 21010 //删除算子失败
	ErrFindMata          = 21011 //读取metadata.json文件失败
	ErrSaveFile          = 21012 //保存文件失败
	ErrUploadFile        = 21013 //下载文件失败
	ErrUpdateOp          = 21014 //更新算子失败
	ErrEncrypt           = 21015 //加密失败
	ErrUncheckOut        = 21016 //算子未测试通过
	ErrOnline            = 21017 //算子上线失败

	ErrCreateDockerFile = 22001 //创建dockerfile失败
	ErrBuildImage       = 22002 //创建镜像失败
	ErrPushImage        = 22003 //推送镜像失败
	ErrMatchImage       = 22004 //匹配镜像失败
	ErrGetBuildLog      = 22005 //获取镜像制作日志失败
	ErrUpdateStatus     = 22006 //更新算子状态失败
	ErrLoadImage        = 22007 //解析镜像失败
	ErrSaveImage        = 22008 //保存镜像失败
	ErrRunSvr           = 22009 //服务启动失败
)

var msgCode = map[int]string{
	Success:       "ok",
	Error:         "服务内部错误",
	ErrBadRequest: "请求参数错误",

	ErrMd5Calculate:      "计算MD5失败",
	ErrFileCompress:      "文件解压失败",
	ErrSaveOp:            "算子保存失败",
	ErrUnmarshalMata:     "解析metadata.json文件失败",
	ErrFindOp:            "找不到该算子",
	ErrUnmarshalFileInfo: "文件元信息解析失败",
	ErrFindFileInfo:      "找不到文件信息",
	ErrMoveFile:          "文件移动失败",
	ErrCheckMata:         "校验metadata.json文件失败",
	ErrDelOp:             "删除算子失败",
	ErrFindMata:          "读取metadata.json文件失败,请检查压缩包的文件目录",
	ErrSaveFile:          "保存文件失败",
	ErrUploadFile:        "下载文件失败",
	ErrUpdateOp:          "更新算子失败",
	ErrEncrypt:           "加密失败",
	ErrUncheckOut:        "算子未测试通过",
	ErrOnline:            "算子上线失败",

	ErrCreateDockerFile: "创建dockerfile失败",
	ErrBuildImage:       "创建镜像失败",
	ErrPushImage:        "推送镜像失败",
	ErrMatchImage:       "匹配镜像失败",
	ErrGetBuildLog:      "获取镜像制作日志失败",
	ErrUpdateStatus:     "更新算子状态失败",
	ErrLoadImage:        "解析镜像失败",
	ErrSaveImage:        "保存镜像失败",
	ErrRunSvr:           "服务启动失败",
}

// GetMsg get msg by code
func GetMsg(code int) string {
	var (
		msg    string
		exists bool
	)

	if msg, exists = msgCode[code]; exists {
		return msg
	}
	return "unknown"
}

// GetMsgErr get error msg by code
/*func GetMsgErr(code int) error {
	msg := GetMsg(code)
	return errors.New(msg)
}*/

// Response api response
type Response struct {
	Code  int         `json:"code" comment:"111"` // msg
	Msg   string      `json:"msg"`                // code
	Data  interface{} `json:"data" form:"111"`    // data
	Count int         `json:"count,omitempty"`    // data count
}

// JSON gin resp to json
func JSON(c *gin.Context, code int, data ...interface{}) {
	resp := Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: data[0],
	}
	if len(data) == 2 {
		resp.Count = data[1].(int)
	}
	c.JSON(200, resp)
	c.Set("statuscode", code)
	return
}

// JSON gin resp to json
func JSONWithError(c *gin.Context, err error) {
	resp := Response{
		Code: Error,
		Msg:  err.Error(),
		Data: err.Error(),
	}
	c.JSON(200, resp)
	c.Set("statuscode", Error)
	return
}
