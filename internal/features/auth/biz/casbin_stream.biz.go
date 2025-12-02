/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the auth module of OrigAdmin.
package biz

import (
	"context"
	"errors"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "origadmin/application/admin/api/v1/services/auth"
)

// CasbinRuleStream is a CasbinSource use case.
type CasbinRuleStream struct {
	ctx      context.Context
	cancel   context.CancelFunc
	receiver chan *pb.StreamRulesResponse
	client   *CasbinSourceServiceBiz
}

func (c CasbinRuleStream) Recv() (*pb.StreamRulesResponse, error) {
	select {
	case msg := <-c.receiver:
		c.client.log.Debugf("received message: %v", msg)
		if msg == nil {
			c.client.log.Debugf("stream closed")
			return nil, io.EOF
		}
		return msg, nil
	case <-c.ctx.Done():
		c.client.log.Debugf("no message received")
		return nil, c.ctx.Err()
	}
}

func (c CasbinRuleStream) Header() (metadata.MD, error) {
	return metadata.MD{}, errors.New("not implemented")
}

func (c CasbinRuleStream) Trailer() metadata.MD {
	return metadata.MD{}
}

func (c CasbinRuleStream) CloseSend() error {
	return errors.New("not implemented")
}

func (c CasbinRuleStream) Context() context.Context {
	return c.ctx
}

func (c CasbinRuleStream) SendMsg(m any) error {
	return errors.New("not implemented")
}

func (c CasbinRuleStream) RecvMsg(m any) error {
	return errors.New("not implemented")
}

func (c CasbinRuleStream) Start(request *pb.StreamRulesRequest) error {
	defer close(c.receiver)
	//c.client.log.Infof("sending request: %v", request)
	if request.WithPolicies {
		policies, err := c.client.ListPolicies(c.ctx, &pb.ListPoliciesRequest{})
		if err != nil {
			return err
		}
		//c.client.log.Infof("sending %d policies", len(policies.Rules))
		for _, rule := range policies.Rules {
			c.receiver <- newPolicyResponse(rule)
		}
	}

	if request.WithGroupings {
		groupings, err := c.client.ListGroupings(c.ctx, &pb.ListGroupingsRequest{})
		if err != nil {
			return err
		}
		//c.client.log.Infof("sending %d groupings", len(groupings.Rules))
		for _, grouping := range groupings.Rules {
			c.receiver <- newGroupingResponse(grouping)
		}
	}
	return nil
}

func NewCasbinRuleStream(ctx context.Context, client *CasbinSourceServiceBiz) *CasbinRuleStream {
	ctx, cancel := context.WithCancel(ctx)
	return &CasbinRuleStream{
		ctx:      ctx,
		cancel:   cancel,
		receiver: make(chan *pb.StreamRulesResponse, 1),
		client:   client,
	}
}

var _ grpc.ServerStreamingClient[pb.StreamRulesResponse] = (*CasbinRuleStream)(nil)
