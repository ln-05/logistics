package hander

import (
	"api_gateway/api/request"
	__ "api_gateway/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CreateLogisticsOrder 创建物流运单
func CreateLogisticsOrder(c *gin.Context) {
	var req request.CreateLogisticsOrderRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.CreateLogisticsOrder(c, &__.CreateLogisticsOrderRequest{
		CreatorId:       req.CreatorId,
		SenderName:      req.SenderName,
		SenderMobile:    req.SenderMobile,
		SenderAddress:   req.SenderAddress,
		ReceiverName:    req.ReceiverName,
		ReceiverMobile:  req.ReceiverMobile,
		ReceiverAddress: req.ReceiverAddress,
		CargoType:       req.CargoType,
		CargoWeight:     req.CargoWeight,
		CargoVolume:     req.CargoVolume,
		CargoQuantity:   req.CargoQuantity,
		TransportType:   req.TransportType,
		Remark:          req.Remark,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": response,
	})
}

// GetWaybill 运单查询接口
func GetWaybill(c *gin.Context) {
	var req request.GetWaybillRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.GetWaybill(c, &__.GetWaybillRequest{
		WaybillId:      req.WaybillId,
		CreatorId:      req.CreatorId,
		Status:         req.Status,
		SenderMobile:   req.SenderMobile,
		ReceiverMobile: req.ReceiverMobile,
		TransportType:  req.TransportType,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		Page:           req.Page,
		PageSize:       req.PageSize,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// UpdateWaybillStatus 运单状态更新接口
func UpdateWaybillStatus(c *gin.Context) {
	var req request.UpdateWaybillStatusRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.UpdateWaybillStatus(c, &__.UpdateWaybillStatusRequest{
		WaybillId:        req.WaybillId,
		NewStatus:        req.NewStatus,
		OperatorId:       req.OperatorId,
		Remark:           req.Remark,
		VehicleId:        req.VehicleId,
		DriveId:          req.DriveId,
		EstimatedArrival: req.EstimatedArrival,
		ActualArrival:    req.ActualArrival,
		AbnormalReason:   req.AbnormalReason,
		Freight:          req.Freight,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// UpdateWaybillInfo 运单信息修改接口
func UpdateWaybillInfo(c *gin.Context) {
	var req request.UpdateWaybillInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.UpdateWaybillInfo(c, &__.UpdateWaybillInfoRequest{
		WaybillId:       req.WaybillId,
		OperatorId:      req.OperatorId,
		SenderName:      req.SenderName,
		SenderMobile:    req.SenderMobile,
		SenderAddress:   req.SenderAddress,
		ReceiverName:    req.ReceiverName,
		ReceiverMobile:  req.ReceiverMobile,
		ReceiverAddress: req.ReceiverAddress,
		CargoType:       req.CargoType,
		CargoWeight:     req.CargoWeight,
		CargoVolume:     req.CargoVolume,
		CargoQuantity:   req.CargoQuantity,
		TransportType:   req.TransportType,
		Remark:          req.Remark,
		Freight:         req.Freight,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// CancelWaybill 运单取消接口
func CancelWaybill(c *gin.Context) {
	var req request.CancelWaybillRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.CancelWaybill(c, &__.CancelWaybillRequest{
		WaybillId:    req.WaybillId,
		OperatorId:   req.OperatorId,
		CancelReason: req.CancelReason,
		Remark:       req.Remark,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// GetWaybillTrack 运单轨迹查询接口
func GetWaybillTrack(c *gin.Context) {
	var req request.GetWaybillTrackRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.GetWaybillTrack(c, &__.GetWaybillTrackRequest{
		WaybillId: req.WaybillId,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// CalculateFreight 运单费用计算接口
func CalculateFreight(c *gin.Context) {
	var req request.CalculateFreightRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.CalculateFreight(c, &__.CalculateFreightRequest{
		CargoType:            req.CargoType,
		CargoWeight:          req.CargoWeight,
		CargoVolume:          req.CargoVolume,
		CargoQuantity:        req.CargoQuantity,
		SenderAddress:        req.SenderAddress,
		ReceiverAddress:      req.ReceiverAddress,
		TransportType:        req.TransportType,
		ServiceType:          req.ServiceType,
		RequiredDeliveryTime: req.RequiredDeliveryTime,
		IsUrgent:             req.IsUrgent,
		NeedInsurance:        req.NeedInsurance,
		InsuranceValue:       req.InsuranceValue,
		NeedReceipt:          req.NeedReceipt,
		NeedPackaging:        req.NeedPackaging,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// BindWaybillResource 运单资源绑定接口
func BindWaybillResource(c *gin.Context) {
	var req request.BindWaybillResourceRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.BindWaybillResource(c, &__.BindWaybillResourceRequest{
		WaybillId:  req.WaybillId,
		VehicleId:  req.VehicleId,
		DriverId:   req.DriverId,
		OperatorId: req.OperatorId,
		Remark:     req.Remark,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// GetWaybillResources 查询运单资源接口
func GetWaybillResources(c *gin.Context) {
	var req request.GetWaybillResourcesRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.GetWaybillResources(c, &__.GetWaybillResourcesRequest{
		WaybillId: req.WaybillId,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}

// ReportException 异常上报接口
func ReportException(c *gin.Context) {
	var req request.ReportExceptionRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	response, err := client.ReportException(c, &__.ReportExceptionRequest{
		WaybillId:      req.WaybillId,
		ExceptionType:  req.ExceptionType,
		Description:    req.Description,
		ReporterId:     req.ReporterId,
		ReporterType:   req.ReporterType,
		Location:       req.Location,
		AttachmentUrls: req.AttachmentUrls,
		DamageLevel:    req.DamageLevel,
		EstimatedLoss:  req.EstimatedLoss,
		ContactPhone:   req.ContactPhone,
		Remark:         req.Remark,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": int(response.Code),
		"msg":  response.Message,
		"data": response,
	})
}
