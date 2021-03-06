// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package golang

import (
	"github.com/gin-gonic/gin"
	"github.com/searKing/golang/third_party/github.com/grpc-ecosystem/grpc-gateway/grpc"
	"github.com/searKing/sole/web/golang/app/configs/values"
	"github.com/searKing/sole/web/golang/app/modules/date"
	"github.com/searKing/sole/web/golang/app/modules/debug"
	"github.com/searKing/sole/web/golang/app/modules/doc/swagger"
	"github.com/searKing/sole/web/golang/app/modules/health"
	"github.com/searKing/sole/web/golang/app/modules/index"
	"github.com/searKing/sole/web/golang/app/modules/webapp"
	"github.com/searKing/sole/web/golang/app/shared/middlewares"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// SetRoutes registers this handler's routes.
func (h *Handler) SetRoutes(ginRouter gin.IRouter, grpcRouter *grpc.Gateway) {
	// bind grpcGateway as default
	middlewares.MiddlewaresRouter(ginRouter)
	index.Router(ginRouter)
	debug.Router(ginRouter, "")
	health.Router(ginRouter)
	// webapp static files
	webapp.Router(ginRouter)
	// doc
	swagger.Router(ginRouter)
	// API
	apiRouter := ginRouter.Group(values.APIPathPrefix)
	index.Router(apiRouter)
	debug.Router(apiRouter, values.APIPathPrefix)
	health.Router(apiRouter)

	date.Router(grpcRouter)

	//// NOTE: It might be required to set Router.HandleMethodNotAllowed to false to avoid problems.
	//r.HandleMethodNotAllowed = false
	//r.NotFound = Routes(h.c, values.PathPrefix)
}
