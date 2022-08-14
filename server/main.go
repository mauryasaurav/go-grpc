package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"go-grpc/server/db"
	userPB "go-grpc/server/proto"
	repo "go-grpc/server/repository"
	"go-grpc/server/utils/helpers"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
)

type Server struct {
	userPB.UnimplementedUserServiceServer
	conn *sql.DB
}

func main() {

	/* Loading TOML file */
	config, err := helpers.LoadEnvFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	/* Connecting To The Postgres Database */
	conn, err := db.PostgresConnect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] initializing postgres db: %+v\n", err)
		return
	}

	/* added default port address*/
	port := ":5000"
	if portEnv := config.Get("env.port").(string); len(portEnv) != 0 {
		port = portEnv
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register services
	userPB.RegisterUserServiceServer(grpcServer, &Server{
		conn: conn,
	})

	log.Printf("GRPC server listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// GetUser -
func (s *Server) GetUser(ctx context.Context, in *userPB.GetUserRequest) (*userPB.GetUserResponse, error) {
	user, err := repo.FindUser(ctx, s.conn, in.Id)
	if err != nil {
		return nil, err
	}

	return &userPB.GetUserResponse{
		User: &userPB.User{
			Id:        in.Id,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}, nil
}

// CreateUser -
func (s *Server) CreateUser(ctx context.Context, in *userPB.CreateUserRequest) (*userPB.CreateUserResponse, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	createUser := repo.User{
		ID:        uuid.String(),
		Name:      in.User.Name,
		Email:     in.User.Email,
		Password:  in.User.Password,
		Phone:     in.User.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = createUser.Insert(ctx, s.conn, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &userPB.CreateUserResponse{
		User: &userPB.User{
			Id:       uuid.String(),
			Name:     createUser.Name,
			Email:    createUser.Email,
			Password: createUser.Password,
			Phone:    createUser.Phone,
		},
	}, nil
}

// DeleteUser -
func (s *Server) DeleteUser(ctx context.Context, in *userPB.DeleteUserRequest) (*userPB.DeleteUserResponse, error) {
	_, err := repo.Users(repo.UserWhere.ID.EQ(in.Id)).DeleteAll(ctx, s.conn)
	fmt.Println("err", err)
	if err != nil {
		return nil, err
	}

	return &userPB.DeleteUserResponse{}, nil
}

// UpdateUser -
func (s *Server) UpdateUser(ctx context.Context, in *userPB.UpdateUserRequest) (*userPB.UpdateUserResponse, error) {
	updateUser := repo.M{
		"updated_at": time.Now(),
		"name":       in.User.Name,
		"phone":      in.User.Phone,
		"email":      in.User.Email,
		"password":   in.User.Password,
	}
	_, err := repo.Users(repo.UserWhere.ID.EQ(in.Id)).UpdateAll(ctx, s.conn, updateUser)
	if err != nil {
		return nil, err
	}

	return &userPB.UpdateUserResponse{}, nil
}

// ListModels -
func (s *Server) ListUsers(ctx context.Context, req *userPB.ListUsersRequest) (*userPB.ListUsersResponse, error) {
	users, err := repo.Users().All(ctx, s.conn)
	if err != nil {
		return nil, err
	}

	var results []*userPB.User
	for _, v := range users {
		results = append(results, &userPB.User{
			Id:        v.ID,
			Name:      v.Name,
			Phone:     v.Phone,
			Password:  v.Password,
			CreatedAt: v.CreatedAt.String(),
			Email:     v.Email,
			UpdatedAt: v.UpdatedAt.String(),
		})
	}

	return &userPB.ListUsersResponse{
		Users: results,
	}, nil
}
