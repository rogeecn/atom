package err

import (
	"net/http"

	"github.com/rogeecn/gen"
)

var (
	BindBodyFailed  = gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "Body参数错误")
	BindQueryFailed = gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "Query参数错误")
	BindPathFailed  = gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "Path参数错误: %s")
)
