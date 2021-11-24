package handler

import (
	"context"
	"ogm-catalog/model"

    "github.com/asim/go-micro/v3/logger"
	proto "github.com/xtech-cloud/ogm-msp-catalog/proto/catalog"
)

type Volume struct{}


func (this *Volume) Create(_ctx context.Context, _req *proto.VolumeCreateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Volume.Create, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewVolumeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Volume) Update(_ctx context.Context, _req *proto.VolumeUpdateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Volume.Update, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewVolumeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Volume) Delete(_ctx context.Context, _req *proto.VolumeDeleteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Volume.Delete, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewVolumeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Volume) Get(_ctx context.Context, _req *proto.VolumeGetRequest, _rsp *proto.VolumeGetResponse) error {
	logger.Infof("Received Volume.Get, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewVolumeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Volume) List(_ctx context.Context, _req *proto.VolumeListRequest, _rsp *proto.VolumeListResponse) error {
	logger.Infof("Received Volume.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewVolumeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Volume) Search(_ctx context.Context, _req *proto.VolumeSearchRequest, _rsp *proto.VolumeListResponse) error {
	logger.Infof("Received Volume.Search, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewVolumeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Volume) Export(_ctx context.Context, _req *proto.VolumeExportRequest, _rsp *proto.VolumeExportResponse) error {
	logger.Infof("Received Volume.Export, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewVolumeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 


