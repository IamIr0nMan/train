package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/iamir0nman/train/trainService"
	"google.golang.org/grpc"
)

type TrainServer struct {
	*trainService.UnimplementedTrainServiceServer
	tickets   []*trainService.Ticket
	seatCount map[string]int
}

func main() {
	server := &TrainServer{
		tickets: []*trainService.Ticket{},
		seatCount: map[string]int{
			"A": 20,
			"B": 20,
		},
	}

	grpcServer := grpc.NewServer()
	trainService.RegisterTrainServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server started on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *TrainServer) PurchaseTicket(ctx context.Context, req *trainService.Ticket) (*trainService.Ticket, error) {
	if s.seatCount[req.Section] > 0 {
		s.tickets = append(s.tickets, req)
		s.seatCount[req.Section]--
		return req, nil
	}
	return nil, fmt.Errorf("no available seats in section %s", req.Section)
}

func (s *TrainServer) GetReceipt(ctx context.Context, req *trainService.User) (*trainService.Ticket, error) {
	for _, ticket := range s.tickets {
		if ticket.User.Email == req.Email {
			return ticket, nil
		}
	}
	return nil, fmt.Errorf("ticket not found for user with email: %s", req.Email)
}

func (s *TrainServer) GetUsersBySection(req *trainService.Ticket, stream trainService.TrainService_GetUsersBySectionServer) error {
	for _, ticket := range s.tickets {
		if ticket.Section == req.Section {
			if err := stream.Send(ticket); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *TrainServer) CancelTicket(ctx context.Context, req *trainService.User) (*trainService.Ticket, error) {
	for i, ticket := range s.tickets {
		if ticket.User.Email == req.Email {
			s.tickets = append(s.tickets[:i], s.tickets[i+1:]...)
			s.seatCount[ticket.Section]++
			return ticket, nil
		}
	}
	return nil, fmt.Errorf("ticket not found for user with email: %s", req.Email)
}

func (s *TrainServer) ModifyUserSeat(ctx context.Context, req *trainService.Ticket) (*trainService.Ticket, error) {
	for i, ticket := range s.tickets {
		if ticket.User.Email == req.User.Email {
			s.tickets[i].Section = req.Section
			return s.tickets[i], nil
		}
	}
	return nil, fmt.Errorf("ticket not found for user with email: %s", req.User.Email)
}
