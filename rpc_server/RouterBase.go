package rpc_server

import "rpc_server/server_base"

type RouterBase struct {
}

func (r *RouterBase) PreHandler(req serverbase.IRequest, res serverbase.IResponse) {}

func (r *RouterBase) Handler(req serverbase.IRequest, res serverbase.IResponse) {}

func (r *RouterBase) AfterHandler(req serverbase.IRequest, res serverbase.IResponse) {}