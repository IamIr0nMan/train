package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/iamir0nman/train/trainService"
	"google.golang.org/grpc"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
}

type Ticket struct {
	From    string
	To      string
	User    User
	Price   float32
	Section string
}

type TrainServer struct {
	*trainService.UnimplementedTrainServiceServer
	tickets   []Ticket
	seatCount map[string]int
}

func main() {
	server := &TrainServer{
		tickets: []Ticket{},
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
		newTicket := Ticket{
			From: req.From,
			To:   req.To,
			User: User{
				FirstName: req.User.FirstName,
				LastName:  req.User.LastName,
				Email:     req.User.Email,
			},
			Price:   req.Price,
			Section: req.Section,
		}
		s.tickets = append(s.tickets, newTicket)
		s.seatCount[req.Section]--
		return req, nil
	}
	return nil, fmt.Errorf("no available seats in section %s", req.Section)
}

func (s *TrainServer) GetReceipt(ctx context.Context, req *trainService.User) (*trainService.Ticket, error) {
	for _, ticket := range s.tickets {
		if ticket.User.Email == req.Email {
			return &trainService.Ticket{
				From:    ticket.From,
				To:      ticket.To,
				User:    &trainService.User{FirstName: ticket.User.FirstName, LastName: ticket.User.LastName, Email: ticket.User.Email},
				Price:   ticket.Price,
				Section: ticket.Section,
			}, nil
		}
	}
	return nil, fmt.Errorf("ticket not found for user with email: %s", req.Email)
}

func (s *TrainServer) GetUsersBySection(req *trainService.Ticket, stream trainService.TrainService_GetUsersBySectionServer) error {
	for _, ticket := range s.tickets {
		if ticket.Section == req.Section {
			if err := stream.Send(&trainService.Ticket{
				From:    ticket.From,
				To:      ticket.To,
				User:    &trainService.User{FirstName: ticket.User.FirstName, LastName: ticket.User.LastName, Email: ticket.User.Email},
				Price:   ticket.Price,
				Section: ticket.Section,
			}); err != nil {
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
			return &trainService.Ticket{
				From:    ticket.From,
				To:      ticket.To,
				User:    &trainService.User{FirstName: ticket.User.FirstName, LastName: ticket.User.LastName, Email: ticket.User.Email},
				Price:   ticket.Price,
				Section: ticket.Section,
			}, nil
		}
	}
	return nil, fmt.Errorf("ticket not found for user with email: %s", req.Email)
}
