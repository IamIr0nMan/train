package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/iamir0nman/train/trainService"
	"google.golang.org/grpc"
)

func TestPurchaseTicket(t *testing.T) {
	tests := []struct {
		name         string
		initialSeats map[string]int
		request      *trainService.Ticket
		expectedResp *trainService.Ticket
		expectedErr  bool
	}{
		{
			name: "Purchase ticket with available seats",
			initialSeats: map[string]int{
				"A": 1,
				"B": 0,
			},
			request: &trainService.Ticket{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
			expectedResp: &trainService.Ticket{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
			expectedErr: false,
		},
		{
			name: "Purchase ticket with no available seats",
			initialSeats: map[string]int{
				"A": 1,
				"B": 0,
			},
			request: &trainService.Ticket{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "B",
			},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name: "Purchase ticket without user information",
			initialSeats: map[string]int{
				"A": 10,
				"B": 10,
			},
			request: &trainService.Ticket{
				From:    "London",
				To:      "Paris",
				Price:   20,
				Section: "B",
			},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name: "Purchase ticket without station information",
			initialSeats: map[string]int{
				"A": 10,
				"B": 10,
			},
			request: &trainService.Ticket{
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
			expectedResp: nil,
			expectedErr:  true,
		},
	}

	server := &TrainServer{
		tickets: []*trainService.Ticket{},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server.seatCount = tc.initialSeats
			numBookedTickets := len(server.tickets)

			ctx := context.Background()
			resp, err := server.PurchaseTicket(ctx, tc.request)

			if !tc.expectedErr && numBookedTickets == len(server.tickets) {
				t.Error("Number of booked tickets didn't increase")
			}
			if tc.expectedErr && err == nil {
				t.Error("Expected an error, got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if tc.expectedResp != nil && (resp == nil || resp.Section != tc.expectedResp.Section || resp.From != tc.expectedResp.From || resp.To != tc.expectedResp.To) {
				t.Errorf("Expected ticket %v, got %v", tc.expectedResp, resp)
			}
		})
	}
}

func TestGetReceipt(t *testing.T) {
	tests := []struct {
		name         string
		request      *trainService.User
		expectedResp *trainService.Ticket
		expectedErr  bool
	}{
		{
			name:         "Invoke GetReceipt func with nil request",
			request:      nil,
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:         "Get receipt without email",
			request:      &trainService.User{},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:         "Get receipt for user who didn't book a ticket",
			request:      &trainService.User{Email: "test@example.com"},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:    "Get receipt for booked ticket",
			request: &trainService.User{Email: "deepak@example.com"},
			expectedResp: &trainService.Ticket{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
			expectedErr: false,
		},
	}

	server := &TrainServer{
		tickets: []*trainService.Ticket{
			{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
		},
		seatCount: map[string]int{
			"A": 10,
			"B": 10,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			resp, err := server.GetReceipt(ctx, tc.request)

			if tc.expectedErr && err == nil {
				t.Error("Expected an error, got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tc.expectedErr && !reflect.DeepEqual(tc.expectedResp, resp) {
				t.Errorf("Expected ticket: %v,\n got ticket: %v", tc.expectedResp, resp)
			}
		})
	}
}

func (m *mockStream) Send(ticket *trainService.Ticket) error {
	m.data = append(m.data, ticket)
	return nil
}

type mockStream struct {
	data []*trainService.Ticket
	grpc.ServerStream
}

func TestGetUsersBySection(t *testing.T) {
	tests := []struct {
		name         string
		request      *trainService.Ticket
		expectedResp []*trainService.Ticket
		expectedErr  bool
	}{
		{
			name:         "Invoke GetUsersBySection func with nil request",
			request:      nil,
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:         "Empty section field in request",
			request:      &trainService.Ticket{},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:         "Invalid section in request",
			request:      &trainService.Ticket{Section: "C"},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:    "Valid section and streaming tickets",
			request: &trainService.Ticket{Section: "A"},
			expectedResp: []*trainService.Ticket{
				{
					From: "London",
					To:   "Paris",
					User: &trainService.User{
						FirstName: "Deepak",
						LastName:  "Kumar",
						Email:     "deepak@example.com",
					},
					Price:   20,
					Section: "A",
				},
				{
					From: "London",
					To:   "Paris",
					User: &trainService.User{
						FirstName: "Test",
						LastName:  "User",
						Email:     "testuser@example.com",
					},
					Price:   20,
					Section: "A",
				},
			},
			expectedErr: false,
		},
	}

	server := &TrainServer{
		tickets: []*trainService.Ticket{
			{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
			{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Test",
					LastName:  "User",
					Email:     "testuser@example.com",
				},
				Price:   20,
				Section: "A",
			},
		},
		seatCount: map[string]int{
			"A": 10,
			"B": 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockStream := &mockStream{}
			err := server.GetUsersBySection(tc.request, mockStream)

			if tc.expectedErr && err == nil {
				t.Error("Expected an error, got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(mockStream.data) != len(tc.expectedResp) {
				t.Errorf("Expected %d tickets, got %d", len(tc.expectedResp), len(mockStream.data))
			}
			for i := range tc.expectedResp {
				if !reflect.DeepEqual(tc.expectedResp[i], mockStream.data[i]) {
					t.Errorf("Mismatch in streamed tickets. Expected: %v,\n got: %v", tc.expectedResp[i], mockStream.data[i])
				}
			}
		})
	}
}

func TestCancelTicket(t *testing.T) {
	tests := []struct {
		name         string
		request      *trainService.User
		expectedResp *trainService.Ticket
		expectedErr  bool
	}{
		{
			name:         "Invoke CancelTicket func with nil request",
			request:      nil,
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:         "Empty email field in request",
			request:      &trainService.User{},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:         "Ticket not found",
			request:      &trainService.User{Email: "test@example.com"},
			expectedResp: nil,
			expectedErr:  true,
		},
		{
			name:    "Ticket found and cancelled",
			request: &trainService.User{Email: "deepak@example.com"},
			expectedResp: &trainService.Ticket{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
			expectedErr: false,
		},
	}

	server := &TrainServer{
		tickets: []*trainService.Ticket{
			{
				From: "London",
				To:   "Paris",
				User: &trainService.User{
					FirstName: "Deepak",
					LastName:  "Kumar",
					Email:     "deepak@example.com",
				},
				Price:   20,
				Section: "A",
			},
		},
		seatCount: map[string]int{
			"A": 10,
			"B": 10,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			resp, err := server.CancelTicket(ctx, tc.request)

			if tc.expectedErr && err == nil {
				t.Error("Expected an error, got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tc.expectedErr && !reflect.DeepEqual(tc.expectedResp, resp) {
				t.Errorf("Expected ticket: %v,\n got ticket: %v", tc.expectedResp, resp)
			}
		})
	}
}
