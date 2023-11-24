// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: person/v1/greeter.proto

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

const OperationPersonServiceCreatePerson = "/person.v1.PersonService/CreatePerson"
const OperationPersonServiceDeletePerson = "/person.v1.PersonService/DeletePerson"
const OperationPersonServiceExistsPerson = "/person.v1.PersonService/ExistsPerson"
const OperationPersonServiceGetPersonById = "/person.v1.PersonService/GetPersonById"
const OperationPersonServiceUpdatePerson = "/person.v1.PersonService/UpdatePerson"

type PersonServiceHTTPServer interface {
	CreatePerson(context.Context, *CreatePersonRequest) (*Person, error)
	DeletePerson(context.Context, *DeletePersonRequest) (*DeletePersonResponse, error)
	ExistsPerson(context.Context, *DeletePersonRequest) (*DeletePersonResponse, error)
	GetPersonById(context.Context, *GetPersonIdRequest) (*Person, error)
	UpdatePerson(context.Context, *UpdatePersonRequest) (*Person, error)
}

func RegisterPersonServiceHTTPServer(s *http.Server, srv PersonServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/person/{personId}", _PersonService_GetPersonById0_HTTP_Handler(srv))
	r.POST("/person", _PersonService_CreatePerson0_HTTP_Handler(srv))
	r.PUT("/person/{personId}", _PersonService_UpdatePerson0_HTTP_Handler(srv))
	r.DELETE("/person/{personId}", _PersonService_DeletePerson0_HTTP_Handler(srv))
	r.GET("/person/{personId}", _PersonService_ExistsPerson0_HTTP_Handler(srv))
}

func _PersonService_GetPersonById0_HTTP_Handler(srv PersonServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPersonIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonServiceGetPersonById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPersonById(ctx, req.(*GetPersonIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Person)
		return ctx.Result(200, reply)
	}
}

func _PersonService_CreatePerson0_HTTP_Handler(srv PersonServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePersonRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonServiceCreatePerson)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePerson(ctx, req.(*CreatePersonRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Person)
		return ctx.Result(200, reply)
	}
}

func _PersonService_UpdatePerson0_HTTP_Handler(srv PersonServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePersonRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonServiceUpdatePerson)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePerson(ctx, req.(*UpdatePersonRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Person)
		return ctx.Result(200, reply)
	}
}

func _PersonService_DeletePerson0_HTTP_Handler(srv PersonServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePersonRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonServiceDeletePerson)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePerson(ctx, req.(*DeletePersonRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePersonResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonService_ExistsPerson0_HTTP_Handler(srv PersonServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePersonRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonServiceExistsPerson)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExistsPerson(ctx, req.(*DeletePersonRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePersonResponse)
		return ctx.Result(200, reply)
	}
}

type PersonServiceHTTPClient interface {
	CreatePerson(ctx context.Context, req *CreatePersonRequest, opts ...http.CallOption) (rsp *Person, err error)
	DeletePerson(ctx context.Context, req *DeletePersonRequest, opts ...http.CallOption) (rsp *DeletePersonResponse, err error)
	ExistsPerson(ctx context.Context, req *DeletePersonRequest, opts ...http.CallOption) (rsp *DeletePersonResponse, err error)
	GetPersonById(ctx context.Context, req *GetPersonIdRequest, opts ...http.CallOption) (rsp *Person, err error)
	UpdatePerson(ctx context.Context, req *UpdatePersonRequest, opts ...http.CallOption) (rsp *Person, err error)
}

type PersonServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPersonServiceHTTPClient(client *http.Client) PersonServiceHTTPClient {
	return &PersonServiceHTTPClientImpl{client}
}

func (c *PersonServiceHTTPClientImpl) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...http.CallOption) (*Person, error) {
	var out Person
	pattern := "/person"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPersonServiceCreatePerson))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PersonServiceHTTPClientImpl) DeletePerson(ctx context.Context, in *DeletePersonRequest, opts ...http.CallOption) (*DeletePersonResponse, error) {
	var out DeletePersonResponse
	pattern := "/person/{personId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPersonServiceDeletePerson))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PersonServiceHTTPClientImpl) ExistsPerson(ctx context.Context, in *DeletePersonRequest, opts ...http.CallOption) (*DeletePersonResponse, error) {
	var out DeletePersonResponse
	pattern := "/person/{personId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPersonServiceExistsPerson))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PersonServiceHTTPClientImpl) GetPersonById(ctx context.Context, in *GetPersonIdRequest, opts ...http.CallOption) (*Person, error) {
	var out Person
	pattern := "/person/{personId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPersonServiceGetPersonById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PersonServiceHTTPClientImpl) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...http.CallOption) (*Person, error) {
	var out Person
	pattern := "/person/{personId}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPersonServiceUpdatePerson))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
