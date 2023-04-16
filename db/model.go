package db

import (
	"time"
)

const TableNameFlight = "flight"

// Flight mapped from table <flight>
type Flight struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(200);not null;uniqueIndex:uniq_name,priority:1" json:"name"`
	Status     int32     `gorm:"column:status;type:tinyint;not null;index:idx_status,priority:1" json:"status"`
	StartTime  time.Time `gorm:"column:start_time;type:timestamp;not null;index:idx_st,priority:1;default:CURRENT_TIMESTAMP" json:"start_time"`
	ModifyTime time.Time `gorm:"column:modify_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"modify_time"`
	Pid        int64     `gorm:"column:pid;type:int;not null;index:idx_pid,priority:1" json:"pid"`
	LayerID    int64     `gorm:"column:layer_id;type:int;not null;index:idx_lid,priority:1" json:"layer_id"`
	Fliter     string    `gorm:"column:fliter;type:text" json:"fliter"`
	TrafficMap string    `gorm:"column:traffic_map;type:longtext" json:"traffic_map"`
	EndTime    time.Time `gorm:"column:end_time;type:timestamp;not null;index:idx_et,priority:1" json:"end_time"`
	Kind       int32     `gorm:"column:kind;type:tinyint" json:"kind"`
}

// TableName Flight's table name
func (Flight) TableName() string {
	return TableNameFlight
}

const TableNameLayer = "layer"

// Layer mapped from table <layer>
type Layer struct {
	ID           int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(200);not null;uniqueIndex:name,priority:1;uniqueIndex:uniq_key_name,priority:1" json:"name"`
	Status       int32     `gorm:"column:status;type:tinyint;not null" json:"status"` // 0未生效 1生效中
	Pid          int64     `gorm:"column:pid;type:int;not null;index:idx_pid,priority:1" json:"pid"`
	Owner        string    `gorm:"column:owner;type:varchar(200);not null;default:default" json:"owner"`
	ModifyTime   time.Time `gorm:"column:modify_time;type:timestamp;not null;index:idx_modify_time,priority:1;default:CURRENT_TIMESTAMP" json:"modify_time"`
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	HashStrategy string    `gorm:"column:hash_strategy;type:varchar(50)" json:"hash_strategy"`
	HashUnit     string    `gorm:"column:hash_unit;type:varchar(50);not null;default:normal" json:"hash_unit"`
}

// TableName Layer's table name
func (Layer) TableName() string {
	return TableNameLayer
}

const TableNameProduct = "product"

// Product mapped from table <product>
type Product struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`                         // pid
	Token       string    `gorm:"column:token;type:varchar(50);not null;uniqueIndex:uniq_token,priority:1" json:"token"` // token名字
	ProductName string    `gorm:"column:product_name;type:varchar(50)" json:"product_name"`                              // 产品名字
	Status      int32     `gorm:"column:status;type:tinyint;not null" json:"status"`                                     // 产品状态 0为不生效 1为生效
	ModifyTime  time.Time `gorm:"column:modify_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"modify_time"`
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	Description string    `gorm:"column:description;type:varchar(2048)" json:"description"` // 描述
}

// TableName Product's table name
func (Product) TableName() string {
	return TableNameProduct
}

const TableNameTrafficMap = "traffic_map"

// TrafficMap mapped from table <traffic_map>
type TrafficMap struct {
	ID               int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreateTime       time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	TrafficMap       string    `gorm:"column:traffic_map;type:text" json:"traffic_map"`
	LayerID          int64     `gorm:"column:layer_id;type:int;index:idx_layer_id,priority:1" json:"layer_id"`
	BucketTrafficMap string    `gorm:"column:bucket_traffic_map;type:mediumtext" json:"bucket_traffic_map"`
	Pid              int64     `gorm:"column:pid;type:int;index:idx_pid,priority:1" json:"pid"`
}

// TableName TrafficMap's table name
func (TrafficMap) TableName() string {
	return TableNameTrafficMap
}

const TableNameVersion = "version"

// Version mapped from table <version>
type Version struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(200)" json:"name"`
	FlightID   int64     `gorm:"column:flight_id;type:int;index:idx_flight_id,priority:1;index:idx_pid_flight_id,priority:2" json:"flight_id"`
	Status     int32     `gorm:"column:status;type:tinyint" json:"status"`
	Summary    string    `gorm:"column:summary;type:longtext" json:"summary"`
	ModifyTime time.Time `gorm:"column:modify_time;type:timestamp;not null;index:idx_modify_time,priority:1;default:CURRENT_TIMESTAMP" json:"modify_time"`
	Pid        int64     `gorm:"column:pid;type:int;index:idx_pid_flight_id,priority:1" json:"pid"`
	Config     string    `gorm:"column:config;type:mediumtext" json:"config"`
}

// TableName Version's table name
func (Version) TableName() string {
	return TableNameVersion
}
