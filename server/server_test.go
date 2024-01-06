package main

import (
	"context"
	"testing"

	"github.com/iamir0nman/train/trainService"
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
			if !tc.expectedErr && resp != server.tickets[0] {
				t.Errorf("Expected ticket: %v,\n got ticket: %v", server.tickets[0], resp)
			}
		})
	}
}
