package user

import (
	"github.com/gin-gonic/gin"
	"go-metrics-demo/pkg/metrics"
	"go-metrics-demo/pkg/metrics/custom"
	"time"
)

// Login godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /api/user/login [post]
func (u *userController) Login(ctx *gin.Context) {
	labels := &custom.ExternalCallDurationLabel{
		Method: "Login",
		Status: "success",
	}
	start := time.Now()
	time.Sleep(2 * time.Second)
	elapsed := time.Since(start).Seconds()
	metrics.RecordExternalCallDuration(labels, elapsed)
	ctx.JSON(200, "success")
}
