package server

import (
	"context"
	"fmt"
	subv1 "github.com/b3liv3r/protos-for-gym/gen/go/subscription"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SubscriptionRPCServer struct {
	subv1.UnimplementedSubscriptionServer
	srv service.Subscriber
}

func NewSubscriptionRPCServer(srv service.Subscriber) subv1.SubscriptionServer {
	return &SubscriptionRPCServer{srv: srv}
}

func (s *SubscriptionRPCServer) Create(ctx context.Context, req *subv1.CreateRequest) (*subv1.CreateResponse, error) {
	userID := int(req.GetUserId())
	message, err := s.srv.Create(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать подписку: %v", err)
	}
	return &subv1.CreateResponse{Message: message}, nil
}

// Метод UpdateType реализует gRPC метод UpdateType для обновления типа подписки и продления даты окончания.
func (s *SubscriptionRPCServer) UpdateType(ctx context.Context, req *subv1.UpdateTypeRequest) (*subv1.UpdateTypeResponse, error) {
	userID := int(req.GetUserId())
	subType := int(req.GetType())
	monthCount := int(req.GetMonth())
	message, err := s.srv.UpdateType(ctx, userID, subType, monthCount)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить тип подписки: %v", err)
	}
	return &subv1.UpdateTypeResponse{Message: message}, nil
}

// Метод Extend реализует gRPC метод Extend для продления даты окончания подписки.
func (s *SubscriptionRPCServer) Extend(ctx context.Context, req *subv1.ExtendRequest) (*subv1.ExtendResponse, error) {
	userID := int(req.GetUserId())
	monthCount := int(req.GetCount())
	message, err := s.srv.ExtendEndDate(ctx, userID, monthCount)
	if err != nil {
		return nil, fmt.Errorf("не удалось продлить дату окончания подписки: %v", err)
	}
	return &subv1.ExtendResponse{Message: message}, nil
}

// Метод GetData реализует gRPC метод GetData для получения информации о подписке.
func (s *SubscriptionRPCServer) GetData(ctx context.Context, req *subv1.GetDataRequest) (*subv1.GetDataResponse, error) {
	userID := int(req.GetUserId())
	subscriptionData, err := s.srv.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить информацию о подписке: %v", err)
	}
	startTime := timestamppb.New(subscriptionData.StartDate)
	endTime := timestamppb.New(subscriptionData.EndDate)
	return &subv1.GetDataResponse{
		Type:      int64(subscriptionData.SubscriptionType),
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}
