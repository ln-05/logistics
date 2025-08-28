package hander

import (
	"context"
	"fmt"
	"logistics_srv/basic/global"
	"logistics_srv/model"
	__ "logistics_srv/proto"
	"logistics_srv/util"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func generateWaybillID() string {
	return uuid.New().String()
}

func generateOrderID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (c *UserServer) CreateLogisticsOrder(_ context.Context, in *__.CreateLogisticsOrderRequest) (*__.CreateLogisticsOrderResponse, error) {
	creatorID64, err := strconv.ParseInt(in.CreatorId, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("无效的创建者ID")
	}
	creatorID := int32(creatorID64)

	waybillID := generateWaybillID()
	orderID := generateOrderID()

	tx := global.DB.Begin()

	// 计算距离和预计送达时间
	distance := util.CalculateDistance(in.SenderAddress, in.ReceiverAddress)
	estimatedDays := util.CalculateEstimatedDays(distance, in.TransportType, "标准")
	estimatedArrival := time.Now().AddDate(0, 0, estimatedDays)

	// 计算基础运费
	baseFreight := util.CalculateBaseFreight(float64(in.CargoWeight), float64(in.CargoVolume), distance, in.TransportType)

	newWaybill := model.Waybill{
		ID:               waybillID,
		OrderID:          orderID,
		CreatorID:        creatorID,
		SenderName:       in.SenderName,
		SenderMobile:     in.SenderMobile,
		SenderAddress:    in.SenderAddress,
		ReceiverName:     in.ReceiverName,
		ReceiverMobile:   in.ReceiverMobile,
		ReceiverAddress:  in.ReceiverAddress,
		CargoType:        in.CargoType,
		CargoWeight:      float64(in.CargoWeight),
		CargoVolume:      float64(in.CargoVolume),
		CargoQuantity:    in.CargoQuantity,
		TransportType:    in.TransportType,
		Status:           "pending", // 默认为pending，表示待分配
		Freight:          baseFreight,
		EstimatedArrival: estimatedArrival,
		Remark:           in.Remark,
	}

	if err = tx.Create(&newWaybill).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建运单失败")
	}

	// 使用工具函数生成状态描述
	statusDescription := util.GetStatusName("", "pending")

	newLog := model.WaybillStatusLog{
		WaybillID:  waybillID,
		OldStatus:  "",
		NewStatus:  "pending",
		OperatorID: creatorID,
		Remark:     statusDescription,
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建运单日志失败")
	}

	tx.Commit()

	return &__.CreateLogisticsOrderResponse{
		WaybillId: waybillID,
		Message:   "创建成功",
	}, nil
}

// 运单查询接口 (支持单条查询和条件查询，带分页)
func (c *UserServer) GetWaybill(_ context.Context, in *__.GetWaybillRequest) (*__.GetWaybillResponse, error) {
	var waybills []model.Waybill
	var total int64

	// 构建查询条件
	query := global.DB.Model(&model.Waybill{})

	// 单条查询：根据运单编号查询
	if in.WaybillId != "" {
		query = query.Where("id = ?", in.WaybillId)
	}

	// 条件查询
	if in.CreatorId != "" {
		creatorID, _ := strconv.ParseInt(in.CreatorId, 10, 32)
		query = query.Where("creator_id = ?", int32(creatorID))
	}

	if in.Status != "" {
		query = query.Where("status = ?", in.Status)
	}

	if in.SenderMobile != "" {
		query = query.Where("sender_mobile = ?", in.SenderMobile)
	}

	if in.ReceiverMobile != "" {
		query = query.Where("receiver_mobile = ?", in.ReceiverMobile)
	}

	if in.TransportType != "" {
		query = query.Where("transport_type = ?", in.TransportType)
	}

	// 时间范围查询
	if in.StartDate != "" {
		startDate, _ := time.Parse("2006-01-02", in.StartDate)
		query = query.Where("create_at >= ?", startDate)
	}

	if in.EndDate != "" {
		endDate, _ := time.Parse("2006-01-02", in.EndDate)
		// 结束日期包含当天，所以加一天
		endDate = endDate.AddDate(0, 0, 1)
		query = query.Where("create_at < ?", endDate)
	}

	// 获取总数
	query.Count(&total)

	// 分页处理
	page := in.Page
	pageSize := in.PageSize

	// 设置默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 执行分页查询
	query.Offset(int(offset)).Limit(int(pageSize)).Order("create_at DESC").Find(&waybills)

	// 转换为 proto 格式
	var protoWaybills []*__.WaybillInfo
	for _, waybill := range waybills {
		protoWaybill := &__.WaybillInfo{
			Id:              waybill.ID,
			OrderId:         waybill.OrderID,
			CreatorId:       waybill.CreatorID,
			SenderName:      waybill.SenderName,
			SenderMobile:    waybill.SenderMobile,
			SenderAddress:   waybill.SenderAddress,
			ReceiverName:    waybill.ReceiverName,
			ReceiverMobile:  waybill.ReceiverMobile,
			ReceiverAddress: waybill.ReceiverAddress,
			CargoType:       waybill.CargoType,
			CargoWeight:     waybill.CargoWeight,
			CargoVolume:     waybill.CargoVolume,
			CargoQuantity:   waybill.CargoQuantity,
			TransportType:   waybill.TransportType,
			VehicleId:       waybill.VehicleID,
			DriveId:         waybill.DriveID,
			Status:          waybill.Status,
			Freight:         waybill.Freight,
			Remark:          waybill.Remark,
			AbnormalReason:  waybill.AbnormalReason,
			CreateAt:        waybill.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:        waybill.UpdateAt.Format("2006-01-02 15:04:05"),
		}

		// 处理可能为空的时间字段
		if !waybill.EstimatedArrival.IsZero() {
			protoWaybill.EstimatedArrival = waybill.EstimatedArrival.Format("2006-01-02 15:04:05")
		}

		if !waybill.ActualArrival.IsZero() {
			protoWaybill.ActualArrival = waybill.ActualArrival.Format("2006-01-02 15:04:05")
		}

		protoWaybills = append(protoWaybills, protoWaybill)
	}

	// 计算总页数
	totalPages := int32((total + int64(pageSize) - 1) / int64(pageSize))

	return &__.GetWaybillResponse{
		Code:       0,
		Message:    "查询成功",
		Waybills:   protoWaybills,
		Total:      int32(total),
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// 运单状态更新接口
func (c *UserServer) UpdateWaybillStatus(_ context.Context, in *__.UpdateWaybillStatusRequest) (*__.UpdateWaybillStatusResponse, error) {
	// 开启事务
	tx := global.DB.Begin()

	// 查询当前运单信息
	var waybill model.Waybill
	tx.Where("id = ?", in.WaybillId).First(&waybill)

	// 记录旧状态
	oldStatus := waybill.Status

	// 处理时间字段
	estimatedTime, _ := time.Parse("2006-01-02 15:04:05", in.EstimatedArrival)
	actualTime, _ := time.Parse("2006-01-02 15:04:05", in.ActualArrival)

	// 直接更新所有字段
	updateData := map[string]interface{}{
		"status":            in.NewStatus,
		"vehicle_id":        in.VehicleId,
		"drive_id":          in.DriveId,
		"freight":           in.Freight,
		"abnormal_reason":   in.AbnormalReason,
		"estimated_arrival": estimatedTime,
		"actual_arrival":    actualTime,
	}

	// 更新运单
	tx.Model(&waybill).Updates(updateData)

	// 使用工具函数生成状态变更描述
	statusDescription := util.GetStatusName(oldStatus, in.NewStatus)
	remark := statusDescription
	if in.Remark != "" {
		remark = fmt.Sprintf("%s - %s", statusDescription, in.Remark)
	}

	// 创建状态变更日志
	statusLog := model.WaybillStatusLog{
		WaybillID:  in.WaybillId,
		OldStatus:  oldStatus,
		NewStatus:  in.NewStatus,
		OperatorID: in.OperatorId,
		Remark:     remark,
	}

	tx.Create(&statusLog)

	// 提交事务
	tx.Commit()

	return &__.UpdateWaybillStatusResponse{
		Code:       0,
		Message:    "状态更新成功",
		WaybillId:  in.WaybillId,
		OldStatus:  oldStatus,
		NewStatus:  in.NewStatus,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// 运单信息修改接口
func (c *UserServer) UpdateWaybillInfo(_ context.Context, in *__.UpdateWaybillInfoRequest) (*__.UpdateWaybillInfoResponse, error) {
	// 开启事务
	tx := global.DB.Begin()

	// 查询当前运单信息
	var waybill model.Waybill
	if err := tx.Where("id = ?", in.WaybillId).First(&waybill).Error; err != nil {
		tx.Rollback()
		return &__.UpdateWaybillInfoResponse{
			Code:    404,
			Message: "运单不存在",
		}, nil
	}

	// 权限校验：只有创建人可以修改
	if waybill.CreatorID != in.OperatorId {
		tx.Rollback()
		return &__.UpdateWaybillInfoResponse{
			Code:    403,
			Message: "无权限修改此运单",
		}, nil
	}

	// 状态校验：避免已发货后随意修改
	if waybill.Status == "in_transit" || waybill.Status == "completed" {
		tx.Rollback()
		return &__.UpdateWaybillInfoResponse{
			Code:    400,
			Message: "运单已发货或已完成，不允许修改",
		}, nil
	}

	// 如果地址、重量、体积或运输方式发生变化，重新计算运费和预计送达时间
	var newFreight = in.Freight
	var newEstimatedArrival = waybill.EstimatedArrival

	// 检查是否需要重新计算
	if in.SenderAddress != waybill.SenderAddress || in.ReceiverAddress != waybill.ReceiverAddress || in.CargoWeight != waybill.CargoWeight || in.CargoVolume != waybill.CargoVolume || in.TransportType != waybill.TransportType {

		// 重新计算距离和运费
		distance := util.CalculateDistance(in.SenderAddress, in.ReceiverAddress)
		newFreight = util.CalculateBaseFreight(in.CargoWeight, in.CargoVolume, distance, in.TransportType)

		// 重新计算预计送达时间
		estimatedDays := util.CalculateEstimatedDays(distance, in.TransportType, "标准")
		newEstimatedArrival = time.Now().AddDate(0, 0, estimatedDays)
	}

	// 直接更新所有字段
	updateData := map[string]interface{}{
		"sender_name":       in.SenderName,
		"sender_mobile":     in.SenderMobile,
		"sender_address":    in.SenderAddress,
		"receiver_name":     in.ReceiverName,
		"receiver_mobile":   in.ReceiverMobile,
		"receiver_address":  in.ReceiverAddress,
		"cargo_type":        in.CargoType,
		"cargo_weight":      in.CargoWeight,
		"cargo_volume":      in.CargoVolume,
		"cargo_quantity":    in.CargoQuantity,
		"transport_type":    in.TransportType,
		"remark":            in.Remark,
		"freight":           newFreight,
		"estimated_arrival": newEstimatedArrival,
	}

	// 更新运单信息
	if err := tx.Model(&waybill).Updates(updateData).Error; err != nil {
		tx.Rollback()
		return &__.UpdateWaybillInfoResponse{
			Code:    500,
			Message: "更新运单信息失败",
		}, nil
	}

	// 提交事务
	tx.Commit()

	return &__.UpdateWaybillInfoResponse{
		Code:       0,
		Message:    "运单信息更新成功",
		WaybillId:  in.WaybillId,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// 运单取消接口
func (c *UserServer) CancelWaybill(_ context.Context, in *__.CancelWaybillRequest) (*__.CancelWaybillResponse, error) {
	// 开启事务
	tx := global.DB.Begin()

	// 查询当前运单信息
	var waybill model.Waybill
	if err := tx.Where("id = ?", in.WaybillId).First(&waybill).Error; err != nil {
		tx.Rollback()
		return &__.CancelWaybillResponse{
			Code:    404,
			Message: "运单不存在",
		}, nil
	}

	// 权限校验：只有创建人可以取消
	if waybill.CreatorID != in.OperatorId {
		tx.Rollback()
		return &__.CancelWaybillResponse{
			Code:    403,
			Message: "无权限取消此运单",
		}, nil
	}

	// 状态校验：判断是否可以取消
	if waybill.Status == "completed" {
		tx.Rollback()
		return &__.CancelWaybillResponse{
			Code:    400,
			Message: "运单已完成（已签收），不允许取消",
		}, nil
	}

	if waybill.Status == "in_transit" {
		tx.Rollback()
		return &__.CancelWaybillResponse{
			Code:    400,
			Message: "运单已发货（在途中），不允许取消",
		}, nil
	}

	if waybill.Status == "canceled" {
		tx.Rollback()
		return &__.CancelWaybillResponse{
			Code:    400,
			Message: "运单已经是取消状态",
		}, nil
	}

	// 记录旧状态
	oldStatus := waybill.Status

	// 更新运单状态为取消，并记录取消原因
	updateData := map[string]interface{}{
		"status":          "canceled",
		"abnormal_reason": in.CancelReason,
	}

	// 如果有备注，也更新备注字段
	if in.Remark != "" {
		updateData["remark"] = in.Remark
	}

	// 更新运单
	if err := tx.Model(&waybill).Updates(updateData).Error; err != nil {
		tx.Rollback()
		return &__.CancelWaybillResponse{
			Code:    500,
			Message: "取消运单失败",
		}, nil
	}

	// 使用工具函数生成状态变更描述
	statusDescription := util.GetStatusName(oldStatus, "canceled")

	// 创建状态变更日志
	statusLog := model.WaybillStatusLog{
		WaybillID:  in.WaybillId,
		OldStatus:  oldStatus,
		NewStatus:  "canceled",
		OperatorID: in.OperatorId,
		Remark:     fmt.Sprintf("%s，取消原因：%s", statusDescription, in.CancelReason),
	}

	if err := tx.Create(&statusLog).Error; err != nil {
		tx.Rollback()
		return &__.CancelWaybillResponse{
			Code:    500,
			Message: "创建取消日志失败",
		}, nil
	}

	// 提交事务
	tx.Commit()

	return &__.CancelWaybillResponse{
		Code:       0,
		Message:    "运单取消成功",
		WaybillId:  in.WaybillId,
		OldStatus:  oldStatus,
		NewStatus:  "canceled",
		CancelTime: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// 运单轨迹查询接口
func (c *UserServer) GetWaybillTrack(_ context.Context, in *__.GetWaybillTrackRequest) (*__.GetWaybillTrackResponse, error) {
	// 查询运单基本信息
	var waybill model.Waybill
	if err := global.DB.Where("id = ?", in.WaybillId).First(&waybill).Error; err != nil {
		return &__.GetWaybillTrackResponse{
			Code:    404,
			Message: "运单不存在",
		}, nil
	}

	// 查询运单状态变更日志，按时间正序排列
	var statusLogs []model.WaybillStatusLog
	if err := global.DB.Where("waybill_id = ?", in.WaybillId).Order("operate_at ASC").Find(&statusLogs).Error; err != nil {
		return &__.GetWaybillTrackResponse{
			Code:    500,
			Message: "查询轨迹信息失败",
		}, nil
	}

	// 构建轨迹节点列表
	var trackNodes []*__.WaybillTrackNode
	for _, log := range statusLogs {
		// 查询操作人信息
		var operator model.User
		global.DB.Where("id = ?", log.OperatorID).First(&operator)

		// 根据状态推断位置信息
		location, address := util.GetLocationByStatus(log.NewStatus, waybill.SenderAddress, waybill.ReceiverAddress)

		// 获取状态中文名称
		statusName := util.GetStatusName(log.OldStatus, log.NewStatus)

		trackNode := &__.WaybillTrackNode{
			Id:             log.ID,
			OldStatus:      log.OldStatus,
			NewStatus:      log.NewStatus,
			StatusName:     statusName,
			OperateTime:    log.OperateAt.Format("2006-01-02 15:04:05"),
			Remark:         log.Remark,
			OperatorId:     log.OperatorID,
			OperatorName:   operator.UserName,
			OperatorMobile: operator.Mobile,
			Location:       location,
			Address:        address,
		}

		trackNodes = append(trackNodes, trackNode)
	}

	// 获取最后更新时间
	var lastUpdateTime string
	if len(statusLogs) > 0 {
		lastUpdateTime = statusLogs[len(statusLogs)-1].OperateAt.Format("2006-01-02 15:04:05")
	}

	return &__.GetWaybillTrackResponse{
		Code:            0,
		Message:         "查询成功",
		WaybillId:       in.WaybillId,
		SenderName:      waybill.SenderName,
		SenderAddress:   waybill.SenderAddress,
		ReceiverName:    waybill.ReceiverName,
		ReceiverAddress: waybill.ReceiverAddress,
		CurrentStatus:   waybill.Status,
		CreateTime:      waybill.CreateAt.Format("2006-01-02 15:04:05"),
		TrackNodes:      trackNodes,
		TotalNodes:      int32(len(trackNodes)),
		LastUpdateTime:  lastUpdateTime,
	}, nil
}

// 运单费用计算接口
func (c *UserServer) CalculateFreight(_ context.Context, in *__.CalculateFreightRequest) (*__.CalculateFreightResponse, error) {
	// 计算距离（模拟地址解析和距离计算）
	distance := util.CalculateDistance(in.SenderAddress, in.ReceiverAddress)

	// 计算各种系数
	weightFactor := util.CalculateWeightFactor(in.CargoWeight)
	volumeFactor := util.CalculateVolumeFactor(in.CargoVolume)
	distanceFactor := util.CalculateDistanceFactor(distance)
	serviceFactor := util.CalculateServiceFactor(in.TransportType, in.ServiceType, in.IsUrgent)

	// 计算基础运费
	baseFreight := util.CalculateBaseFreight(in.CargoWeight, in.CargoVolume, distance, in.TransportType)

	// 构建费用明细
	var freightItems []*__.FreightItem

	// 基础运费明细
	freightItems = append(freightItems, &__.FreightItem{
		Name:        "基础运费",
		UnitPrice:   util.GetBaseUnitPrice(in.TransportType),
		Quantity:    in.CargoWeight,
		Unit:        "kg",
		Amount:      baseFreight,
		Description: "按重量和距离计费",
	})

	// 距离附加费
	var distanceAdditionalFee float64
	if distance > 500 {
		distanceAdditionalFee = (distance - 500) * 0.15
		freightItems = append(freightItems, &__.FreightItem{
			Name:        "距离附加费",
			UnitPrice:   0.15,
			Quantity:    distance - 500,
			Unit:        "km",
			Amount:      distanceAdditionalFee,
			Description: "超过500km部分",
		})
	}

	// 服务类型附加费
	var serviceAdditionalFee float64
	if in.ServiceType == "加急" {
		serviceAdditionalFee = baseFreight * 0.3
		freightItems = append(freightItems, &__.FreightItem{
			Name:        "加急服务费",
			UnitPrice:   0.3,
			Quantity:    1,
			Unit:        "倍",
			Amount:      serviceAdditionalFee,
			Description: "加急服务附加费",
		})
	} else if in.ServiceType == "特快" {
		serviceAdditionalFee = baseFreight * 0.5
		freightItems = append(freightItems, &__.FreightItem{
			Name:        "特快服务费",
			UnitPrice:   0.5,
			Quantity:    1,
			Unit:        "倍",
			Amount:      serviceAdditionalFee,
			Description: "特快服务附加费",
		})
	}

	// 体积附加费（当体积重量大于实际重量时）
	volumeWeight := in.CargoVolume * 200 // 1立方米 = 200kg
	var volumeAdditionalFee float64
	if volumeWeight > in.CargoWeight {
		volumeAdditionalFee = (volumeWeight - in.CargoWeight) * util.GetBaseUnitPrice(in.TransportType) * 0.5
		freightItems = append(freightItems, &__.FreightItem{
			Name:        "体积附加费",
			UnitPrice:   util.GetBaseUnitPrice(in.TransportType) * 0.5,
			Quantity:    volumeWeight - in.CargoWeight,
			Unit:        "kg",
			Amount:      volumeAdditionalFee,
			Description: "体积重量超出部分",
		})
	}

	// 燃油附加费
	fuelSurcharge := baseFreight * 0.05
	freightItems = append(freightItems, &__.FreightItem{
		Name:        "燃油附加费",
		UnitPrice:   0.05,
		Quantity:    1,
		Unit:        "倍",
		Amount:      fuelSurcharge,
		Description: "当前燃油价格调整",
	})

	// 增值服务费用
	var insuranceFee, receiptFee, packagingFee float64

	if in.NeedInsurance && in.InsuranceValue > 0 {
		insuranceFee = in.InsuranceValue * 0.003 // 保险费率0.3%
		freightItems = append(freightItems, &__.FreightItem{
			Name:        "货物保险费",
			UnitPrice:   0.003,
			Quantity:    in.InsuranceValue,
			Unit:        "元",
			Amount:      insuranceFee,
			Description: "货物保险服务",
		})
	}

	if in.NeedReceipt {
		receiptFee = 10.0
		freightItems = append(freightItems, &__.FreightItem{
			Name:        "回单服务费",
			UnitPrice:   10.0,
			Quantity:    1,
			Unit:        "份",
			Amount:      receiptFee,
			Description: "回单签收服务",
		})
	}

	if in.NeedPackaging {
		packagingFee = float64(in.CargoQuantity * 5.0)
		freightItems = append(freightItems, &__.FreightItem{
			Name:        "包装服务费",
			UnitPrice:   5.0,
			Quantity:    float64(in.CargoQuantity),
			Unit:        "件",
			Amount:      packagingFee,
			Description: "专业包装服务",
		})
	}

	// 计算总费用
	additionalFreight := distanceAdditionalFee + serviceAdditionalFee + volumeAdditionalFee + fuelSurcharge + insuranceFee + receiptFee + packagingFee
	totalFreight := baseFreight + additionalFreight

	// 计算预计送达时间
	estimatedDays := util.CalculateEstimatedDays(distance, in.TransportType, in.ServiceType)
	estimatedDeliveryTime := time.Now().AddDate(0, 0, estimatedDays).Format("2006-01-02 18:00:00")

	return &__.CalculateFreightResponse{
		Code:                  0,
		Message:               "费用计算成功",
		TotalFreight:          totalFreight,
		BaseFreight:           baseFreight,
		AdditionalFreight:     additionalFreight,
		DiscountAmount:        0, // 暂无优惠
		FreightItems:          freightItems,
		CalculatedDistance:    distance,
		WeightFactor:          weightFactor,
		VolumeFactor:          volumeFactor,
		DistanceFactor:        distanceFactor,
		ServiceFactor:         serviceFactor,
		EstimatedDeliveryTime: estimatedDeliveryTime,
		EstimatedDays:         int32(estimatedDays),
		FreightRule:           "按重量、体积、距离综合计费，含燃油附加费",
		ValidityPeriod:        "报价有效期7天",
		CalculationTime:       time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// 运单资源绑定接口
func (c *UserServer) BindWaybillResource(_ context.Context, in *__.BindWaybillResourceRequest) (*__.BindWaybillResourceResponse, error) {
	tx := global.DB.Begin()

	// 查询运单
	var waybill model.Waybill
	if err := tx.Where("id = ?", in.WaybillId).First(&waybill).Error; err != nil {
		tx.Rollback()
		return &__.BindWaybillResourceResponse{Code: 404, Message: "运单不存在"}, nil
	}

	// 更新资源绑定
	updateData := map[string]interface{}{}
	if in.VehicleId != "" {
		updateData["vehicle_id"] = in.VehicleId
	}
	if in.DriverId > 0 {
		updateData["drive_id"] = in.DriverId
	}

	// 如果绑定了资源且运单状态为pending，更新为assigned
	if len(updateData) > 0 && waybill.Status == "pending" {
		updateData["status"] = "assigned"
	}

	// 执行更新
	if err := tx.Model(&waybill).Updates(updateData).Error; err != nil {
		tx.Rollback()
		return &__.BindWaybillResourceResponse{Code: 500, Message: "绑定失败"}, nil
	}

	if err := tx.Commit().Error; err != nil {
		return &__.BindWaybillResourceResponse{Code: 500, Message: "提交失败"}, nil
	}

	return &__.BindWaybillResourceResponse{
		Code:      0,
		Message:   "资源绑定成功",
		WaybillId: in.WaybillId,
		VehicleId: in.VehicleId,
		DriverId:  in.DriverId,
		BindTime:  time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// 查询运单资源接口
func (c *UserServer) GetWaybillResources(_ context.Context, in *__.GetWaybillResourcesRequest) (*__.GetWaybillResourcesResponse, error) {
	// 查询运单信息
	var waybill model.Waybill
	if err := global.DB.Where("id = ?", in.WaybillId).First(&waybill).Error; err != nil {
		return &__.GetWaybillResourcesResponse{
			Code:    404,
			Message: "运单不存在",
		}, nil
	}

	response := &__.GetWaybillResourcesResponse{
		Code:      0,
		Message:   "查询成功",
		WaybillId: in.WaybillId,
		VehicleId: waybill.VehicleID,
		DriverId:  waybill.DriveID,
		Status:    waybill.Status,
	}

	// 如果有绑定司机，查询司机详细信息
	if waybill.DriveID > 0 {
		var driver model.User
		if err := global.DB.Where("id = ?", waybill.DriveID).First(&driver).Error; err == nil {
			response.DriverName = driver.UserName
			response.DriverMobile = driver.Mobile
		}
	}

	// 设置绑定时间（使用运单的更新时间作为绑定时间）
	response.BindTime = waybill.UpdateAt.Format("2006-01-02 15:04:05")

	return response, nil
}

// 异常上报接口
func (c *UserServer) ReportException(_ context.Context, in *__.ReportExceptionRequest) (*__.ReportExceptionResponse, error) {
	// 开启事务
	tx := global.DB.Begin()

	// 验证运单是否存在
	var waybill model.Waybill
	if err := tx.Where("id = ?", in.WaybillId).First(&waybill).Error; err != nil {
		tx.Rollback()
		return &__.ReportExceptionResponse{
			Code:    404,
			Message: "运单不存在",
		}, nil
	}

	// 生成异常单号
	exceptionID := uuid.NewString()

	// 计算预计解决时间（根据异常类型设定不同的解决时间）
	var resolveHours int
	switch in.ExceptionType {
	case "damage":
		resolveHours = 24 // 损坏类异常24小时内解决
	case "delay":
		resolveHours = 12 // 延误类异常12小时内解决
	case "lost":
		resolveHours = 72 // 丢失类异常72小时内解决
	case "address_error":
		resolveHours = 6 // 地址错误6小时内解决
	case "refused":
		resolveHours = 48 // 拒收类异常48小时内解决
	default:
		resolveHours = 24 // 默认24小时
	}

	expectedResolveTime := time.Now().Add(time.Duration(resolveHours) * time.Hour)

	// 创建异常记录
	exception := model.WaybillException{
		ID:            exceptionID,
		WaybillID:     in.WaybillId,
		ExceptionType: in.ExceptionType,
		Description:   in.Description,
		ReporterID:    in.ReporterId,
		ReporterType:  in.ReporterType,
		Location:      in.Location,
		DamageLevel:   in.DamageLevel,
		EstimatedLoss: in.EstimatedLoss,
		ContactPhone:  in.ContactPhone,
		Remark:        in.Remark,
		Status:        "reported",
		ReportTime:    time.Now(),
	}

	// 保存异常记录
	if err := tx.Create(&exception).Error; err != nil {
		tx.Rollback()
		return &__.ReportExceptionResponse{
			Code:    500,
			Message: fmt.Sprintf("创建异常记录失败: %v", err),
		}, nil
	}

	// 保存附件信息
	if len(in.AttachmentUrls) > 0 {
		for _, url := range in.AttachmentUrls {
			attachment := model.ExceptionAttachment{
				ExceptionID: exceptionID,
				FileName:    "attachment",
				FileURL:     url,
				FileType:    "image", // 默认为图片类型
				UploadTime:  time.Now(),
			}
			if err := tx.Create(&attachment).Error; err != nil {
				tx.Rollback()
				return &__.ReportExceptionResponse{
					Code:    500,
					Message: "保存附件信息失败",
				}, nil
			}
		}
	}

	// 如果是严重异常，可能需要更新运单状态
	if in.ExceptionType == "lost" || (in.ExceptionType == "damage" && in.DamageLevel == "severe") {
		// 可以考虑将运单状态更新为异常状态，这里暂时不做处理
		// tx.Model(&waybill).Update("status", "exception")
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return &__.ReportExceptionResponse{
			Code:    500,
			Message: "提交事务失败",
		}, nil
	}

	return &__.ReportExceptionResponse{
		Code:                0,
		Message:             "异常上报成功",
		ExceptionId:         exceptionID,
		WaybillId:           in.WaybillId,
		ReportTime:          exception.ReportTime.Format("2006-01-02T15:04:05Z"),
		ExpectedResolveTime: expectedResolveTime.Format("2006-01-02T15:04:05Z"),
		Status:              "reported",
	}, nil
}
