package e

import "github.com/gin-gonic/gin"

func MsgFlags(code int) string {
	switch code {
	case SUCCESS:
		return "成功"
	case ERROR:
		return "失败"
	case INVALID_PARAMS:
		return "请求参数错误"
	case ERROR_EXIST_TAG:
		return "名称重复"
	case ERROR_EXIST_TAG_FAIL:
		return "获取已存在标签失败"
	case ERROR_NOT_EXIST_TAG:
		return "不存在"
	case ERROR_GET_TAGS_FAIL:
		return "获取所有标签失败"
	case ERROR_COUNT_TAG_FAIL:
		return "统计标签失败"
	case ERROR_ADD_TAG_FAIL:
		return "新增标签失败"
	case ERROR_EDIT_TAG_FAIL:
		return "修改标签失败"
	case ERROR_DELETE_TAG_FAIL:
		return "删除标签失败"
	case ERROR_EXPORT_TAG_FAIL:
		return "导出标签失败"
	case ERROR_IMPORT_TAG_FAIL:
		return "导入标签失败"
	case ERROR_NOT_EXIST_ARTICLE:
		return "该文章不存在"
	case ERROR_ADD_ARTICLE_FAIL:
		return "新增文章失败"
	case ERROR_DELETE_ARTICLE_FAIL:
		return "删除文章失败"
	case ERROR_CHECK_EXIST_ARTICLE_FAIL:
		return "检查文章是否存在失败"
	case ERROR_EDIT_ARTICLE_FAIL:
		return "修改文章失败"
	case ERROR_COUNT_ARTICLE_FAIL:
		return "统计文章失败"
	case ERROR_GET_ARTICLES_FAIL:
		return "获取多个文章失败"
	case ERROR_GET_ARTICLE_FAIL:
		return "获取单个文章失败"
	case ERROR_GEN_ARTICLE_POSTER_FAIL:
		return "生成文章海报失败"
	case ERROR_AUTH_CHECK_TOKEN_FAIL:
		return "Token鉴权失败"
	case ERROR_AUTH_CHECK_TOKEN_TIMEOUT:
		return "Token已超时"
	case ERROR_AUTH_TOKEN:
		return "Token生成失败"
	case ERROR_AUTH:
		return "Token错误"
	case ERROR_UPLOAD_SAVE_IMAGE_FAIL:
		return "保存图片失败"
	case ERROR_UPLOAD_CHECK_IMAGE_FAIL:
		return "检查图片失败"
	case ERROR_UPLOAD_CHECK_IMAGE_FORMAT:
		return "校验图片错误，图片格式或大小有问题"
	default:
		return ""
	}

}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg := MsgFlags(code)
	if msg != "" {
		return msg
	}
	return MsgFlags(ERROR)
}

// ErrorMsg 错误回复 code 状态  error错误信息
func ErrorMsg(c *gin.Context, code int, error string) {
	c.JSON(code, map[string]any{"code": code, "error": error})
	c.Abort()
}
