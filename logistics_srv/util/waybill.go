package util

import (
	"fmt"
	"strings"
)

// GetStatusName 获取状态中文名称
func GetStatusName(oldStatus, newStatus string) string {
	statusMap := map[string]string{
		"pending":    "待分配",
		"assigned":   "已分配",
		"in_transit": "运输中",
		"completed":  "已完成",
		"canceled":   "已取消",
	}

	if oldStatus == "" {
		return fmt.Sprintf("运单创建 - %s", statusMap[newStatus])
	}

	oldName := statusMap[oldStatus]
	newName := statusMap[newStatus]
	return fmt.Sprintf("%s → %s", oldName, newName)
}

// GetLocationByStatus 根据状态推断位置信息
func GetLocationByStatus(status, senderAddress, receiverAddress string) (location, address string) {
	switch status {
	case "pending":
		return "待分配", senderAddress
	case "assigned":
		return "分拣中心", senderAddress
	case "in_transit":
		return "运输途中", "货物正在运输中"
	case "completed":
		return "已送达", receiverAddress
	case "canceled":
		return "已取消", senderAddress
	default:
		return "未知状态", ""
	}
}

// CalculateDistance 计算距离（模拟）
func CalculateDistance(senderAddress, receiverAddress string) float64 {
	// 简化的距离计算逻辑，实际应该调用地图API
	cityDistanceMap := map[string]map[string]float64{
		"北京": {"上海": 1200, "广州": 2100, "深圳": 2200, "杭州": 1100},
		"上海": {"北京": 1200, "广州": 1300, "深圳": 1400, "杭州": 180},
		"广州": {"北京": 2100, "上海": 1300, "深圳": 120, "杭州": 1200},
		"深圳": {"北京": 2200, "上海": 1400, "广州": 120, "杭州": 1300},
	}

	// 简单的城市匹配
	senderCity := extractCity(senderAddress)
	receiverCity := extractCity(receiverAddress)

	if distances, ok := cityDistanceMap[senderCity]; ok {
		if distance, ok := distances[receiverCity]; ok {
			return distance
		}
	}

	// 默认距离
	return 800.0
}

// extractCity 提取城市名称（简化）
func extractCity(address string) string {
	if strings.Contains(address, "北京") {
		return "北京"
	} else if strings.Contains(address, "上海") {
		return "上海"
	} else if strings.Contains(address, "广州") {
		return "广州"
	} else if strings.Contains(address, "深圳") {
		return "深圳"
	} else if strings.Contains(address, "杭州") {
		return "杭州"
	}
	return "其他"
}

// CalculateWeightFactor 计算重量系数
func CalculateWeightFactor(weight float64) float64 {
	if weight <= 5 {
		return 1.0
	} else if weight <= 20 {
		return 1.1
	} else if weight <= 50 {
		return 1.2
	} else {
		return 1.3
	}
}

// CalculateVolumeFactor 计算体积系数
func CalculateVolumeFactor(volume float64) float64 {
	if volume <= 0.1 {
		return 1.0
	} else if volume <= 0.5 {
		return 1.1
	} else if volume <= 1.0 {
		return 1.2
	} else {
		return 1.3
	}
}

// CalculateDistanceFactor 计算距离系数
func CalculateDistanceFactor(distance float64) float64 {
	if distance <= 200 {
		return 1.0
	} else if distance <= 500 {
		return 1.1
	} else if distance <= 1000 {
		return 1.2
	} else {
		return 1.3
	}
}

// CalculateServiceFactor 计算服务系数
func CalculateServiceFactor(transportType, serviceType string, isUrgent bool) float64 {
	factor := 1.0

	// 运输方式系数
	switch transportType {
	case "陆运":
		factor *= 1.0
	case "空运":
		factor *= 2.0
	case "海运":
		factor *= 0.8
	}

	// 服务类型系数
	switch serviceType {
	case "标准":
		factor *= 1.0
	case "加急":
		factor *= 1.3
	case "特快":
		factor *= 1.5
	}

	// 加急系数
	if isUrgent {
		factor *= 1.2
	}

	return factor
}

// CalculateBaseFreight 计算基础运费
func CalculateBaseFreight(weight, volume, distance float64, transportType string) float64 {
	unitPrice := GetBaseUnitPrice(transportType)

	// 按重量计算
	weightFreight := weight * unitPrice

	// 按体积重量计算
	volumeWeight := volume * 200 // 1立方米 = 200kg
	volumeFreight := volumeWeight * unitPrice

	// 取较大值
	baseFreight := weightFreight
	if volumeFreight > weightFreight {
		baseFreight = volumeFreight
	}

	// 距离调整
	if distance > 500 {
		baseFreight *= 1.2
	}

	return baseFreight
}

// GetBaseUnitPrice 获取基础单价
func GetBaseUnitPrice(transportType string) float64 {
	switch transportType {
	case "陆运":
		return 6.5
	case "空运":
		return 15.0
	case "海运":
		return 4.0
	default:
		return 6.5
	}
}

// calculateEstimatedDays 计算预计天数
func CalculateEstimatedDays(distance float64, transportType, serviceType string) int {
	baseDays := 1

	// 根据距离调整
	if distance > 1000 {
		baseDays = 3
	} else if distance > 500 {
		baseDays = 2
	}

	// 根据运输方式调整
	switch transportType {
	case "陆运":
		// 基础天数
	case "空运":
		baseDays = 1 // 空运更快
	case "海运":
		baseDays += 2 // 海运更慢
	}

	// 根据服务类型调整
	switch serviceType {
	case "加急":
		if baseDays > 1 {
			baseDays -= 1
		}
	case "特快":
		baseDays = 1
	}

	return baseDays
}
