package api

import (
	"shop-test/model"
	"strings"
)

func JsonSuccess(message string, data interface{}) interface{} {
	return jsonSuccess(message, data)
}

func jsonSuccess(message string, data interface{}) interface{} {
	var code = 200
	if message == "deleted" {
		code = 204
	}
	obj := map[string]interface{}{
		"success": true,
		"code":    code,
		"message": message,
		"data":    data,
	}
	return obj
}

func jsonSuccessPaging(message string, data interface{}, p *model.Paging) interface{} {
	var code = 200
	if message == "deleted" {
		code = 204
	}
	obj := map[string]interface{}{
		"success": true,
		"code":    code,
		"message": message,
		"data":    data,
		"paging" : &model.Paging{
			TotalItem: p.TotalItem,
			Page: p.Page,
			PageSize: p.PageSize,
		},
	}
	return obj
}

func JsonSuccessPaging(message string, data interface{}, p *model.Paging) interface{} {
	return jsonSuccessPaging(message, data, p)
}

func JsonError(code int, message string) interface{} {
	return jsonError(code, message)
}

func jsonError(code int, message string) interface{} {
	obj := map[string]interface{}{
		"success": false,
		"code":    code,
		"errors": []string{
			message,
		},
	}
	return obj
}


func ProcessHandleSuccess(message string) interface{} {
	obj := map[string]interface{}{
		"success": true,
		"code":    200,
		"errors": []string{
			message,
		},
	}
	return obj
} 

func GetErrorCode(err error) int {
	errCase := strings.Index(err.Error(), "not found")
	if errCase != -1 {
		return 404
	}
	errCase = strings.Index(err.Error(), "unauthorized")
	if errCase != -1 {
		return 401
	}
	return 400
}
