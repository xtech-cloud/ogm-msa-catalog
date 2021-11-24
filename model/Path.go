package model

import (
	"time"
)

type Path struct {
	UUID      string `gorm:"column:uuid;type:char(32);not null;unique;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


func (Path) TableName() string {
	return "ogm_catalog_Path"
}

type PathDAO struct {
	conn *Conn
}

func NewPathDAO(_conn *Conn) *PathDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &PathDAO{
		conn: conn,
	}
}

func (this *PathDAO) Count() (int64, error) {
	var count int64
	err := this.conn.DB.Model(&Path{}).Count(&count).Error
	return count, err
}

func (this *PathDAO) Insert(_entity *Path) error {
	return this.conn.DB.Create(_entity).Error
}

func (this *PathDAO) Update(_entity *Path) error {
    // 只更新非零值
	return this.conn.DB.Updates(_entity).Error
}

func (this *PathDAO) List(_offset int64, _count int64) (int64, []*Path, error) {
	var entityAry []*Path
    count := int64(0)
	db := this.conn.DB.Model(&Path{})
    // db = db.Where("key = ?", value)
    if err := db.Count(&count).Error; nil != err {
        return 0, nil, err
    }
    db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}

func (this *PathDAO) Delete(_uuid string) error {
	return this.conn.DB.Where("uuid = ?", _uuid).Delete(&Path{}).Error
}
