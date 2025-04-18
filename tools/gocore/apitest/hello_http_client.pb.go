// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.3
// - protoc            v4.24.2
// source: apitest/hello.proto

package apitest

import (
	context "context"
	ecode "github.com/sunmi-OS/gocore/v2/api/ecode"
	calloption "github.com/sunmi-OS/gocore/v2/tools/protoc-gen/calloption"
	http_request "github.com/sunmi-OS/gocore/v2/utils/http-request"
)

// HelloHTTPClient is the client API for Hello service.
type HelloHTTPClient interface {
	CreateHello(context.Context, *CreateHelloReq, ...calloption.CallOption) (*TResponse[CreateHelloResp], error)
	UpdateHello(context.Context, *UpdateHelloReq, ...calloption.CallOption) (*TResponse[UpdateHelloResp], error)
	DeleteHello(context.Context, *DeleteHelloReq, ...calloption.CallOption) (*TResponse[DeleteHelloResp], error)
	GetHello(context.Context, *GetHelloReq, ...calloption.CallOption) (*TResponse[GetHelloResp], error)
	ListHello(context.Context, *ListHelloReq, ...calloption.CallOption) (*TResponse[ListHelloResp], error)
}

type HelloHTTPClientImpl struct {
	hh *http_request.HttpClient
}

func NewHelloHTTPClient(hh *http_request.HttpClient) HelloHTTPClient {
	return &HelloHTTPClientImpl{hh: hh}
}

func (c *HelloHTTPClientImpl) CreateHello(ctx context.Context, req *CreateHelloReq, opts ...calloption.CallOption) (*TResponse[CreateHelloResp], error) {
	resp := &TResponse[CreateHelloResp]{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/v3/createHello")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(int(resp.Code), resp.Msg)
	}
	return resp, err
}

func (c *HelloHTTPClientImpl) UpdateHello(ctx context.Context, req *UpdateHelloReq, opts ...calloption.CallOption) (*TResponse[UpdateHelloResp], error) {
	resp := &TResponse[UpdateHelloResp]{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/v3/updateHello")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(int(resp.Code), resp.Msg)
	}
	return resp, err
}

func (c *HelloHTTPClientImpl) DeleteHello(ctx context.Context, req *DeleteHelloReq, opts ...calloption.CallOption) (*TResponse[DeleteHelloResp], error) {
	resp := &TResponse[DeleteHelloResp]{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/v3/deleteHello")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(int(resp.Code), resp.Msg)
	}
	return resp, err
}

func (c *HelloHTTPClientImpl) GetHello(ctx context.Context, req *GetHelloReq, opts ...calloption.CallOption) (*TResponse[GetHelloResp], error) {
	resp := &TResponse[GetHelloResp]{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/v3/getHello")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(int(resp.Code), resp.Msg)
	}
	return resp, err
}

func (c *HelloHTTPClientImpl) ListHello(ctx context.Context, req *ListHelloReq, opts ...calloption.CallOption) (*TResponse[ListHelloResp], error) {
	resp := &TResponse[ListHelloResp]{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/v3/listHello")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(int(resp.Code), resp.Msg)
	}
	return resp, err
}
