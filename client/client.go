package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/iamir0nman/train/trainService"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := trainService.NewTrainServiceClient(conn)

	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Purchase Ticket")
		fmt.Println("2. Get Reciept")
		fmt.Println("3. Get Users in section")
		fmt.Println("4. Cancel Ticket")
		fmt.Println("5. Modify Ticket")
		fmt.Println("q. Quit")

		choice := inputHelper("Enter your choice: ")

		switch choice {
		case "1":
			purchaseTicket(client)
		case "2":
			getReceipt(client)
		case "3":
			getUsersBySection(client)
		case "4":
			cancelTicket(client)
		case "5":
			modifyTicket(client)
		case "q":
			fmt.Println("Exiting the program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}

		fmt.Println("Press 'Enter' to continue...")
		fmt.Scanln()
	}
}

func inputHelper(label string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(label)
	scanner.Scan()
	return scanner.Text()
}

func modifyTicket(client trainService.TrainServiceClient) {
	email := inputHelper("Enter email: ")
	section := inputHelper("Enter section [A or B]: ")

	modifyUserSeatReq := &trainService.Ticket{
		User:    &trainService.User{Email: email},
		Section: section,
	}
	modifyUserSeatResp, err := client.ModifyUserSeat(context.Background(), modifyUserSeatReq)
	if err != nil {
		log.Fatalf("ModifyUserSeat failed: %v", err)
	}
	log.Printf("ModifyUserSeat response: %v", modifyUserSeatResp)
}

func cancelTicket(client trainService.TrainServiceClient) {
	email := inputHelper("Enter email: ")

	cancelTicketReq := &trainService.User{Email: email}
	cancelTicketResp, err := client.CancelTicket(context.Background(), cancelTicketReq)
	if err != nil {
		log.Fatalf("CancelTicket failed: %v", err)
	}
	log.Printf("CancelTicket response: %v", cancelTicketResp)
}

func getUsersBySection(client trainService.TrainServiceClient) {
	section := inputHelper("Enter section [A or B]: ")

	getUsersBySectionReq := &trainService.Ticket{Section: section}
	getUsersBySectionStream, err := client.GetUsersBySection(context.Background(), getUsersBySectionReq)
	if err != nil {
		log.Fatalf("GetUsersBySection failed: %v", err)
	}
	for {
		user, err := getUsersBySectionStream.Recv()
		if user == nil {
			log.Println("No bookings found in this section")
		}
		if err != nil {
			break
		}
		log.Printf("User in SectionA: %v", user)
	}
}

func getReceipt(client trainService.TrainServiceClient) {
	email := inputHelper("Enter email: ")

	getReceiptReq := &trainService.User{Email: email}
	getReceiptResp, err := client.GetReceipt(context.Background(), getReceiptReq)
	if err != nil {
		log.Fatalf("GetReceipt failed: %v", err)
	}
	log.Printf("GetReceipt response: %v", getReceiptResp)
}

func purchaseTicket(client trainService.TrainServiceClient) {
	from := inputHelper("Enter source station: ")
	to := inputHelper("Enter destination station: ")
	firstName := inputHelper("Enter first name: ")
	lastName := inputHelper("Enter last name: ")
	email := inputHelper("Enter email: ")
	section := inputHelper("Enter section [A or B]: ")

	purchaseTicketReq := &trainService.Ticket{
		From: from,
		To:   to,
		User: &trainService.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
		Price:   20,
		Section: section,
	}
	purchaseTicketResp, err := client.PurchaseTicket(context.Background(), purchaseTicketReq)
	if err != nil {
		log.Fatalf("PurchaseTicket failed: %v", err)
	}
	log.Printf("PurchaseTicket response: %v", purchaseTicketResp)
}
