package request

// CreateLogisticsOrderRequest 创建物流运单请求
type CreateLogisticsOrderRequest struct {
	CreatorId       string  `json:"creator_id" form:"creator_id" binding:"required"`              // 创建人ID
	SenderName      string  `json:"sender_name" form:"sender_name" binding:"required"`            // 发货人姓名
	SenderMobile    string  `json:"sender_mobile" form:"sender_mobile" binding:"required"`        // 发货人电话
	SenderAddress   string  `json:"sender_address" form:"sender_address" binding:"required"`      // 发货地址
	ReceiverName    string  `json:"receiver_name" form:"receiver_name" binding:"required"`        // 收货人姓名
	ReceiverMobile  string  `json:"receiver_mobile" form:"receiver_mobile" binding:"required"`    // 收货人电话
	ReceiverAddress string  `json:"receiver_address" form:"receiver_address" binding:"required"`  // 收货地址
	CargoType       string  `json:"cargo_type" form:"cargo_type" binding:"required"`              // 货物类型
	CargoWeight     float32 `json:"cargo_weight" form:"cargo_weight" binding:"required,gt=0"`     // 货物重量(kg)
	CargoVolume     float32 `json:"cargo_volume" form:"cargo_volume" binding:"required,gt=0"`     // 货物体积(m³)
	CargoQuantity   int32   `json:"cargo_quantity" form:"cargo_quantity" binding:"required,gt=0"` // 货物数量
	TransportType   string  `json:"transport_type" form:"transport_type" binding:"required"`      // 运输方式
	Remark          string  `json:"remark" form:"remark"`                                         // 备注信息
}

// GetWaybillRequest 运单查询请求 (支持单条查询和条件查询)
type GetWaybillRequest struct {
	WaybillId      string `json:"waybill_id" form:"waybill_id"`           // 运单编号 (可选，用于单条查询)
	CreatorId      string `json:"creator_id" form:"creator_id"`           // 创建人ID (可选)
	Status         string `json:"status" form:"status"`                   // 运单状态 (可选)
	SenderMobile   string `json:"sender_mobile" form:"sender_mobile"`     // 发货人电话 (可选)
	ReceiverMobile string `json:"receiver_mobile" form:"receiver_mobile"` // 收货人电话 (可选)
	TransportType  string `json:"transport_type" form:"transport_type"`   // 运输方式 (可选)
	StartDate      string `json:"start_date" form:"start_date"`           // 开始日期 (可选, 格式: 2024-01-01)
	EndDate        string `json:"end_date" form:"end_date"`               // 结束日期 (可选, 格式: 2024-01-31)
	Page           int32  `json:"page" form:"page"`                       // 页码 (默认1)
	PageSize       int32  `json:"page_size" form:"page_size"`             // 每页数量 (默认10, 最大100)
}

// UpdateWaybillStatusRequest 运单状态更新请求
type UpdateWaybillStatusRequest struct {
	WaybillId        string  `json:"waybill_id" form:"waybill_id" binding:"required"`   // 运单编号 (必填)
	NewStatus        string  `json:"new_status" form:"new_status" binding:"required"`   // 新状态 (必填: pending/assigned/in_transit/completed/canceled)
	OperatorId       int32   `json:"operator_id" form:"operator_id" binding:"required"` // 操作人ID (必填)
	Remark           string  `json:"remark" form:"remark"`                              // 操作备注 (可选)
	VehicleId        string  `json:"vehicle_id" form:"vehicle_id"`                      // 运输车辆ID (可选，状态为assigned时可填)
	DriveId          int32   `json:"drive_id" form:"drive_id"`                          // 司机ID (可选，状态为assigned时可填)
	EstimatedArrival string  `json:"estimated_arrival" form:"estimated_arrival"`        // 预计到达时间 (可选，格式: 2024-01-01 15:04:05)
	ActualArrival    string  `json:"actual_arrival" form:"actual_arrival"`              // 实际到达时间 (可选，状态为completed时可填)
	AbnormalReason   string  `json:"abnormal_reason" form:"abnormal_reason"`            // 异常原因 (可选，状态为canceled时可填)
	Freight          float64 `json:"freight" form:"freight"`                            // 运费金额 (可选)
}

// UpdateWaybillInfoRequest 运单信息修改请求
type UpdateWaybillInfoRequest struct {
	WaybillId       string  `json:"waybill_id" form:"waybill_id" binding:"required"`   // 运单编号 (必填)
	OperatorId      int32   `json:"operator_id" form:"operator_id" binding:"required"` // 操作人ID (必填，用于权限校验)
	SenderName      string  `json:"sender_name" form:"sender_name"`                    // 发货人姓名
	SenderMobile    string  `json:"sender_mobile" form:"sender_mobile"`                // 发货人电话
	SenderAddress   string  `json:"sender_address" form:"sender_address"`              // 发货地址
	ReceiverName    string  `json:"receiver_name" form:"receiver_name"`                // 收货人姓名
	ReceiverMobile  string  `json:"receiver_mobile" form:"receiver_mobile"`            // 收货人电话
	ReceiverAddress string  `json:"receiver_address" form:"receiver_address"`          // 收货地址
	CargoType       string  `json:"cargo_type" form:"cargo_type"`                      // 货物类型
	CargoWeight     float64 `json:"cargo_weight" form:"cargo_weight"`                  // 货物重量(kg)
	CargoVolume     float64 `json:"cargo_volume" form:"cargo_volume"`                  // 货物体积(m³)
	CargoQuantity   int32   `json:"cargo_quantity" form:"cargo_quantity"`              // 货物数量
	TransportType   string  `json:"transport_type" form:"transport_type"`              // 运输方式
	Remark          string  `json:"remark" form:"remark"`                              // 备注信息
	Freight         float64 `json:"freight" form:"freight"`                            // 运费金额
}

// CancelWaybillRequest 运单取消请求
type CancelWaybillRequest struct {
	WaybillId    string `json:"waybill_id" form:"waybill_id" binding:"required"`       // 运单编号 (必填)
	OperatorId   int32  `json:"operator_id" form:"operator_id" binding:"required"`     // 操作人ID (必填，用于权限校验)
	CancelReason string `json:"cancel_reason" form:"cancel_reason" binding:"required"` // 取消原因 (必填)
	Remark       string `json:"remark" form:"remark"`                                  // 备注信息 (可选)
}

// GetWaybillTrackRequest 运单轨迹查询请求
type GetWaybillTrackRequest struct {
	WaybillId string `json:"waybill_id" binding:"required" form:"waybill_id"` // 运单编号 (必填)
}

// CalculateFreightRequest 运单费用计算请求
type CalculateFreightRequest struct {
	CargoType            string  `json:"cargo_type" form:"cargo_type" binding:"required"`              // 货物类型
	CargoWeight          float64 `json:"cargo_weight" form:"cargo_weight" binding:"required,gt=0"`     // 货物重量(kg)
	CargoVolume          float64 `json:"cargo_volume" form:"cargo_volume" binding:"required,gt=0"`     // 货物体积(m³)
	CargoQuantity        int32   `json:"cargo_quantity" form:"cargo_quantity" binding:"required,gt=0"` // 货物数量
	SenderAddress        string  `json:"sender_address" form:"sender_address" binding:"required"`      // 发货地址
	ReceiverAddress      string  `json:"receiver_address" form:"receiver_address" binding:"required"`  // 收货地址
	TransportType        string  `json:"transport_type" form:"transport_type" binding:"required"`      // 运输方式(陆运/空运/海运)
	ServiceType          string  `json:"service_type" form:"service_type" binding:"required"`          // 服务类型(标准/加急/特快)
	RequiredDeliveryTime string  `json:"required_delivery_time" form:"required_delivery_time"`         // 要求送达时间 (格式: 2024-01-15 18:00:00)
	IsUrgent             bool    `json:"is_urgent" form:"is_urgent"`                                   // 是否加急
	NeedInsurance        bool    `json:"need_insurance" form:"need_insurance"`                         // 是否需要保险
	InsuranceValue       float64 `json:"insurance_value" form:"insurance_value"`                       // 保险价值
	NeedReceipt          bool    `json:"need_receipt" form:"need_receipt"`                             // 是否需要回单
	NeedPackaging        bool    `json:"need_packaging" form:"need_packaging"`                         // 是否需要包装服务
}
