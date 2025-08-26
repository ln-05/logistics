package router

import (
	"api_gateway/api/hander"

	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	g := r.Group("/logistics")
	{
		user := g.Group("/user")
		{
			user.POST("/sendsms", hander.Sendsms)
			user.POST("/login", hander.Login)
		}
		waybill := g.Group("/waybill")
		{
			waybill.POST("/CreateLogisticsOrder", hander.CreateLogisticsOrder)
			waybill.POST("/GetWaybill", hander.GetWaybill)
			waybill.POST("/UpdateWaybillStatus", hander.UpdateWaybillStatus)
			waybill.POST("/UpdateWaybillInfo", hander.UpdateWaybillInfo)
			waybill.POST("/CancelWaybill", hander.CancelWaybill)
			waybill.POST("/GetWaybillTrack", hander.GetWaybillTrack)
			waybill.POST("/CalculateFreight", hander.CalculateFreight)
		}

	}

}
