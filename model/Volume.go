package model

import (
	"time"
)

type Volume struct {
	UUID      string `gorm:"column:uuid;type:char(32);not null;unique;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


func (Volume) TableName() string {
	return "ogm_catalog_Volume"
}

type VolumeDAO struct {
	conn *Conn
}

func NewVolumeDAO(_conn *Conn) *VolumeDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &VolumeDAO{
		conn: conn,
	}
}

func (this *VolumeDAO) Count() (int64, error) {
	var count int64
	err := this.conn.DB.Model(&Volume{}).Count(&count).Error
	return count, err
}

func (this *VolumeDAO) Insert(_entity *Volume) error {
	return this.conn.DB.Create(_entity).Error
}

func (this *VolumeDAO) Update(_entity *Volume) error {
    // 只更新非零值
	return this.conn.DB.Updates(_entity).Error
}

func (this *VolumeDAO) List(_offset int64, _count int64) (int64, []*Volume, error) {
	var entityAry []*Volume
    count := int64(0)
	db := this.conn.DB.Model(&Volume{})
    // db = db.Where("key = ?", value)
    if err := db.Count(&count).Error; nil != err {
        return 0, nil, err
    }
    db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}

func (this *VolumeDAO) Delete(_uuid string) error {
	return this.conn.DB.Where("uuid = ?", _uuid).Delete(&Volume{}).Error
}
