package main

import (
	"context"
	"github.com/Printemps417/optionloader/examples/hello/kitex_gen/hellodemo"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *hellodemo.Request) (resp *hellodemo.Response, err error) {
	// TODO: Your code here...
	resp = &hellodemo.Response{Message: req.Message}
	return
}
