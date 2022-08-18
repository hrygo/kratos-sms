// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.5
// source: sms/v1/simple_message.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSmsAsyncResultQuery = "/sms.v1.Sms/AsyncResultQuery"
const OperationSmsTemplateMessageSend = "/sms.v1.Sms/TemplateMessageSend"
const OperationSmsTextMessageSend = "/sms.v1.Sms/TextMessageSend"

type SmsHTTPServer interface {
	AsyncResultQuery(context.Context, *AsyncResultQueryRequest) (*AsyncResultQueryReply, error)
	TemplateMessageSend(context.Context, *TemplateMessageRequest) (*SendMessageReply, error)
	TextMessageSend(context.Context, *TextMessageRequest) (*SendMessageReply, error)
}

func RegisterSmsHTTPServer(s *http.Server, srv SmsHTTPServer) {
	r := s.Route("/")
	r.POST("/sms/send/text", _Sms_TextMessageSend0_HTTP_Handler(srv))
	r.POST("/sms/send/template", _Sms_TemplateMessageSend0_HTTP_Handler(srv))
	r.GET("/sms/async/result/{query_id}", _Sms_AsyncResultQuery0_HTTP_Handler(srv))
}

func _Sms_TextMessageSend0_HTTP_Handler(srv SmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TextMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSmsTextMessageSend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TextMessageSend(ctx, req.(*TextMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Sms_TemplateMessageSend0_HTTP_Handler(srv SmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TemplateMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSmsTemplateMessageSend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TemplateMessageSend(ctx, req.(*TemplateMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Sms_AsyncResultQuery0_HTTP_Handler(srv SmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AsyncResultQueryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSmsAsyncResultQuery)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AsyncResultQuery(ctx, req.(*AsyncResultQueryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AsyncResultQueryReply)
		return ctx.Result(200, reply)
	}
}

type SmsHTTPClient interface {
	AsyncResultQuery(ctx context.Context, req *AsyncResultQueryRequest, opts ...http.CallOption) (rsp *AsyncResultQueryReply, err error)
	TemplateMessageSend(ctx context.Context, req *TemplateMessageRequest, opts ...http.CallOption) (rsp *SendMessageReply, err error)
	TextMessageSend(ctx context.Context, req *TextMessageRequest, opts ...http.CallOption) (rsp *SendMessageReply, err error)
}

type SmsHTTPClientImpl struct {
	cc *http.Client
}

func NewSmsHTTPClient(client *http.Client) SmsHTTPClient {
	return &SmsHTTPClientImpl{client}
}

func (c *SmsHTTPClientImpl) AsyncResultQuery(ctx context.Context, in *AsyncResultQueryRequest, opts ...http.CallOption) (*AsyncResultQueryReply, error) {
	var out AsyncResultQueryReply
	pattern := "/sms/async/result/{query_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSmsAsyncResultQuery))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SmsHTTPClientImpl) TemplateMessageSend(ctx context.Context, in *TemplateMessageRequest, opts ...http.CallOption) (*SendMessageReply, error) {
	var out SendMessageReply
	pattern := "/sms/send/template"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSmsTemplateMessageSend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SmsHTTPClientImpl) TextMessageSend(ctx context.Context, in *TextMessageRequest, opts ...http.CallOption) (*SendMessageReply, error) {
	var out SendMessageReply
	pattern := "/sms/send/text"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSmsTextMessageSend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
