package service

import (
	"context"

	pb "origadmin/backend/api/v1/admin"
)

type LogService struct {
	pb.UnimplementedLogServer
}

func NewLogService() *LogService {
	return &LogService{}
}

func (s *LogService) CreateLog(ctx context.Context, req *pb.CreateLogRequest) (*pb.CreateLogReply, error) {
	return &pb.CreateLogReply{}, nil
}
func (s *LogService) UpdateLog(ctx context.Context, req *pb.UpdateLogRequest) (*pb.UpdateLogReply, error) {
	return &pb.UpdateLogReply{}, nil
}
func (s *LogService) DeleteLog(ctx context.Context, req *pb.DeleteLogRequest) (*pb.DeleteLogReply, error) {
	return &pb.DeleteLogReply{}, nil
}
func (s *LogService) GetLog(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogReply, error) {
	return &pb.GetLogReply{}, nil
}
func (s *LogService) ListLog(ctx context.Context, req *pb.ListLogRequest) (*pb.ListLogReply, error) {
	return &pb.ListLogReply{}, nil
}
