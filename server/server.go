package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/shamskhalil/gApp/gen/contactpb"
	"google.golang.org/grpc"
)

type customError struct{}

func (ce *customError) Error() string {
	return "Invalid index"
}

type Person struct {
	Name  string
	Phone string
}

var contactStore []Person

type contactServiceApiServer struct {
	contactpb.UnimplementedContactServiceApiServer
}

func (c *contactServiceApiServer) GetAll(req *contactpb.GetOneContactRequest, stream contactpb.ContactServiceApi_GetAllServer) error {
	for i := 0; i < len(contactStore); i++ {
		person := contactStore[i]
		response := &contactpb.GetOneContactResponse{
			Name:  person.Name,
			Phone: person.Phone,
		}
		stream.Send(response)
		time.Sleep(time.Second * 5)
	}
	return nil
}

func (c *contactServiceApiServer) GetOne(ctx context.Context, req *contactpb.GetOneContactRequest) (*contactpb.GetOneContactResponse, error) {
	index := req.Index
	if index <= int64(len(contactStore)-1) && index >= 0 {
		person := contactStore[index]
		return &contactpb.GetOneContactResponse{
			Name:  person.Name,
			Phone: person.Phone,
		}, nil
	}
	return nil, &customError{}
}

func (c *contactServiceApiServer) Add(ctx context.Context, req *contactpb.AddContactRequest) (*contactpb.AddContactResponse, error) {
	name := req.Name
	phone := req.Phone
	person := Person{Name: name, Phone: phone}
	contactStore = append(contactStore, person)
	log.Printf("Storing Contact >> %+v\n", person)
	return &contactpb.AddContactResponse{Msg: fmt.Sprintf("Contact %s Added Successfully!", name)}, nil
}

func main() {
	//listener
	listener, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("Error starting listener %v\n", err)
	}
	//grpc server
	grpcServer := grpc.NewServer()
	//register Service
	contactpb.RegisterContactServiceApiServer(grpcServer, &contactServiceApiServer{})
	//serve
	fmt.Println("Grpc server listening on port 5000!")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error serving grpc services %v\n", err)
	}

}
