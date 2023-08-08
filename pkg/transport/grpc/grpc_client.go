package grpc_client

import (
	"context"
	"fmt"

	loggerbin "github.com/salesforceanton/grpc-logger-bin/pkg/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LoggerbinClient struct {
	conn   *grpc.ClientConn
	client loggerbin.LoggerbinServiceClient
}

func NewClient(port int) (*LoggerbinClient, error) {
	// Create connection trought grpc
	var conn *grpc.ClientConn

	addr := fmt.Sprintf(":%d", port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Construct client instance defined in generated files from .proto
	return &LoggerbinClient{
		conn:   conn,
		client: loggerbin.NewLoggerbinServiceClient(conn),
	}, nil
}

func (c *LoggerbinClient) CloseConnection() error {
	return c.conn.Close()
}

func (c *LoggerbinClient) SendLog(ctx context.Context, req loggerbin.LogItem) error {
	action, err := loggerbin.ToPbAction(req.Action)
	if err != nil {
		return err
	}

	entity, err := loggerbin.ToPbEntity(req.Entity)
	if err != nil {
		return err
	}

	_, err = c.client.AddLog(ctx, &loggerbin.LogRequest{
		Action:    action,
		Entity:    entity,
		EntityId:  req.EntityID,
		Timestamp: timestamppb.New(req.Timestamp),
	})

	return err
}
