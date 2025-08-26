package model

import (
	"time"
)

type User struct {
	Id       int32     `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`
	Mobile   string    `gorm:"column:mobile;type:char(11);default:NULL;" json:"mobile"`
	UserName string    `gorm:"column:user_name;type:varchar(20);default:NULL;" json:"user_name"`
	NickName string    `gorm:"column:nick_name;type:varchar(20);default:NULL;" json:"nick_name"`
	IdCard   string    `gorm:"column:id_card;type:varchar(18);default:NULL;" json:"id_card"`
	Sex      string    `gorm:"column:sex;type:varchar(5);default:NULL;" json:"sex"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"update_at"`
	DeleteAt time.Time `gorm:"column:delete_at;type:datetime(3);default:NULL;" json:"delete_at"`
}

func (u *User) TableName() string {
	return "user"
}

// Waybill 运单表
type Waybill struct {
	ID               string    `gorm:"column:id;type:varchar(36);comment:运单编号;primaryKey;not null;" json:"id"`                                                                                   // 运单编号
	OrderID          string    `gorm:"column:order_id;type:varchar(32);comment:关联订单编号;not null;" json:"order_id"`                                                                                // 关联订单编号
	CreatorID        int32     `gorm:"column:creator_id;type:int;comment:创建人ID(关联用户表);not null;" json:"creator_id"`                                                                              // 创建人ID(关联用户表)
	SenderName       string    `gorm:"column:sender_name;type:varchar(50);comment:发货人姓名;not null;" json:"sender_name"`                                                                           // 发货人姓名
	SenderMobile     string    `gorm:"column:sender_mobile;type:char(11);comment:发货人电话;not null;" json:"sender_mobile"`                                                                          // 发货人电话
	SenderAddress    string    `gorm:"column:sender_address;type:varchar(255);comment:发货地址;not null;" json:"sender_address"`                                                                     // 发货地址
	ReceiverName     string    `gorm:"column:receiver_name;type:varchar(50);comment:收货人姓名;not null;" json:"receiver_name"`                                                                       // 收货人姓名
	ReceiverMobile   string    `gorm:"column:receiver_mobile;type:char(11);comment:收货人电话;not null;" json:"receiver_mobile"`                                                                      // 收货人电话
	ReceiverAddress  string    `gorm:"column:receiver_address;type:varchar(255);comment:收货地址;not null;" json:"receiver_address"`                                                                 // 收货地址
	CargoType        string    `gorm:"column:cargo_type;type:varchar(20);comment:货物类型;not null;" json:"cargo_type"`                                                                              // 货物类型
	CargoWeight      float64   `gorm:"column:cargo_weight;type:decimal(10, 2);comment:货物重量(kg);not null;" json:"cargo_weight"`                                                                   // 货物重量(kg)
	CargoVolume      float64   `gorm:"column:cargo_volume;type:decimal(10, 2);comment:货物体积(m³);default:0.00;" json:"cargo_volume"`                                                               // 货物体积(m³)
	CargoQuantity    int32     `gorm:"column:cargo_quantity;type:int;comment:货物数量;not null;default:1;" json:"cargo_quantity"`                                                                    // 货物数量
	TransportType    string    `gorm:"column:transport_type;type:varchar(20);comment:运输方式(陆运/空运/海运);not null;" json:"transport_type"`                                                            // 运输方式(陆运/空运/海运)
	VehicleID        string    `gorm:"column:vehicle_id;type:varchar(36);comment:运输车辆ID;default:NULL;" json:"vehicle_id"`                                                                        // 运输车辆ID
	DriveID          int32     `gorm:"column:drive_id;type:int;comment:司机ID(关联用户表);default:NULL;" json:"drive_id"`                                                                               // 司机ID(关联用户表)
	Status           string    `gorm:"column:status;type:varchar(20);comment:运单状态(pending-待分配/assigned-已分配/in_transit-在途/completed-已完成/canceled-已取消异常);not null;default:pending;" json:"status"` // 运单状态(pending-待分配/assigned-已分配/in_transit-在途/completed-已完成/canceled-已取消异常)
	EstimatedArrival time.Time `gorm:"column:estimated_arrival;type:datetime;comment:预计到达时间;default:NULL;" json:"estimated_arrival"`                                                             // 预计到达时间
	ActualArrival    time.Time `gorm:"column:actual_arrival;type:datetime;comment:实际到达时间;default:NULL;" json:"actual_arrival"`                                                                   // 实际到达时间
	Freight          float64   `gorm:"column:freight;type:decimal(10, 2);comment:运费金额;default:0.00;" json:"freight"`                                                                             // 运费金额
	Remark           string    `gorm:"column:remark;type:text;comment:备注信息;" json:"remark"`                                                                                                      // 备注信息
	AbnormalReason   string    `gorm:"column:abnormal_reason;type:text;comment:异常原因;" json:"abnormal_reason"`                                                                                    // 异常原因
	CreateAt         time.Time `gorm:"column:create_at;type:datetime(3);comment:创建时间;default:CURRENT_TIMESTAMP(3);" json:"create_at"`                                                            // 创建时间
	UpdateAt         time.Time `gorm:"column:update_at;type:datetime(3);comment:更新时间;default:CURRENT_TIMESTAMP(3);" json:"update_at"`                                                            // 更新时间
	DeleteAt         time.Time `gorm:"column:delete_at;type:datetime(3);comment:删除时间;default:NULL;" json:"delete_at"`                                                                            // 删除时间
}

// TableName 自定义表名
func (Waybill) TableName() string {
	return "waybill"
}

// WaybillStatusLog 运单状态变更日志表
type WaybillStatusLog struct {
	ID         int32     `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;comment:'日志ID'" json:"id"`
	WaybillID  string    `gorm:"column:waybill_id;type:varchar(36);not null;comment:'运单编号'" json:"waybill_id"`
	OldStatus  string    `gorm:"column:old_status;type:varchar(20);default:NULL;comment:'旧状态'" json:"old_status"`
	NewStatus  string    `gorm:"column:new_status;type:varchar(20);not null;comment:'新状态'" json:"new_status"`
	OperatorID int32     `gorm:"column:operator_id;type:int;not null;comment:'操作人ID(关联用户表)'" json:"operator_id"`
	OperateAt  time.Time `gorm:"column:operate_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);comment:'操作时间'" json:"operate_at"`
	Remark     string    `gorm:"column:remark;type:text;default:NULL;comment:'操作备注'" json:"remark"`
}

// TableName 自定义表名
func (WaybillStatusLog) TableName() string {
	return "waybill_status_log"
}
