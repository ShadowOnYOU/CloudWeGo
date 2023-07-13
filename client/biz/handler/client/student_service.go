// Code generated by hertz generator.

package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"

	student2 "client/biz/model/client"
	student "client/kitex_gen/client"
	"client/kitex_gen/client/studentservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /add-student-info [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req student.Student
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var opts []client.Option
	opts = append(opts, client.WithHostPorts("127.0.0.1:9999"))
	opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10, MaxIdlePerAddress: 1000}))

	studentClient := studentservice.MustNewClient("student", opts...)

	resp, err := studentClient.Register(context.Background(), &req)

	if resp.Success {
		c.String(consts.StatusOK, resp.Message)
	} else {
		c.String(consts.StatusBadRequest, errors.New("error: ID Already Existed").Error())
	}
}

// Query .
// @router /query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	var req student2.QueryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(student.Student)

	var opts []client.Option
	opts = append(opts, client.WithHostPorts("127.0.0.1:9999"))
	opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10, MaxIdlePerAddress: 1000}))

	studentClient := studentservice.MustNewClient("student", opts...)

	fmt.Printf("Got a Request for the information of ID: %v.\n", req.ID)

	var request = student.QueryReq{req.ID}
	resp, err = studentClient.Query(context.Background(), &request)

	fmt.Println("Query ended.")

	if err != nil {
		fmt.Println(err)
		c.String(consts.StatusBadGateway, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}
