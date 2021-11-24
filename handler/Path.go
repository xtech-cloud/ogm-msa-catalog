package handler

import (
	"context"
	"ogm-catalog/model"

    "github.com/asim/go-micro/v3/logger"
	proto "github.com/xtech-cloud/ogm-msp-catalog/proto/catalog"
)

type Path struct{}


func (this *Path) Add(_ctx context.Context, _req *proto.PathAddRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Path.Add, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewPathDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Path) Update(_ctx context.Context, _req *proto.PathUpdateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Path.Update, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewPathDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Path) Remove(_ctx context.Context, _req *proto.PathDeleteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Path.Remove, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewPathDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Path) Get(_ctx context.Context, _req *proto.PathGetRequest, _rsp *proto.PathGetResponse) error {
	logger.Infof("Received Path.Get, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewPathDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Path) List(_ctx context.Context, _req *proto.PathListRequest, _rsp *proto.PathListResponse) error {
	logger.Infof("Received Path.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewPathDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Path) Search(_ctx context.Context, _req *proto.PathSearchRequest, _rsp *proto.PathListResponse) error {
	logger.Infof("Received Path.Search, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewPathDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 


