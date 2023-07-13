package main

import (
	"context"
	"errors"
	"fmt"
	server0 "server/kitex_gen/server"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

var table = map[int32]server0.Student{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, info *server0.Student) (resp *server0.RegisterResp, err error) {
	_, ok := table[info.Id]
	if ok {
		fmt.Println("Registration Failed")
		return nil, errors.New("error: ID Already Existed")
	}

	table[info.Id] = *info
	resp = server0.NewRegisterResp()
	resp.Success = true
	resp.Message = "Information Added: " + info.Name
	fmt.Printf("Add %v %v %v to table\n", info.Id, info.College.String(), info.Name)

	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *server0.QueryReq) (resp *server0.Student, err error) {
	res, ok := table[req.Id]
	if !ok {
		fmt.Printf("you requested for %v\n", req.Id)
		return nil, errors.New("error: Queried ID Not Found")
	}

	return &res, nil
}
