package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	bufbuild "github.com/tk42/sqlc-buf-pg-template/gen/proto/golang/github.com/tk42/sqlc-buf-pg-template"
	sqlc "github.com/tk42/sqlc-buf-pg-template/gen/sqlc"
	"google.golang.org/genproto/googleapis/type/datetime"
)

type ServiceServer struct {
	db *sqlc.Queries
}

func NewServiceServer(client *sql.DB) bufbuild.PetStoreServiceServer {
	db := sqlc.New(client)
	return &ServiceServer{
		db: db,
	}
}

func (s ServiceServer) PutPet(ctx context.Context, req *bufbuild.PutPetRequest) (*bufbuild.PutPetResponse, error) {
	// Write implementation here
	pet, err := s.db.CreatePetQuery(ctx, sqlc.CreatePetQueryParams{
		Name: req.Name,
		Memo: sql.NullString{
			String: req.PetType.String(),
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &bufbuild.PutPetResponse{
		Pet: &bufbuild.Pet{
			PetId:   fmt.Sprint(pet.ID),
			PetType: req.PetType,
			Name:    pet.Name,
			CreatedAt: &datetime.DateTime{
				Year:    int32(now.Year()),
				Month:   int32(now.Month()),
				Day:     int32(now.Day()),
				Hours:   int32(now.Hour()),
				Minutes: int32(now.Minute()),
				Seconds: int32(now.Second()),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil
}

func (s ServiceServer) DeletePet(ctx context.Context, req *bufbuild.DeletePetRequest) (*bufbuild.DeletePetResponse, error) {
	// Write implementation here
	id, err := strconv.ParseInt(req.PetId, 10, 64)
	if err != nil {
		return nil, err
	}
	if err := s.db.DeletePetQuery(ctx, id); err != nil {
		return nil, err
	}
	return &bufbuild.DeletePetResponse{}, nil
}

func (s ServiceServer) GetPet(ctx context.Context, req *bufbuild.GetPetRequest) (*bufbuild.GetPetResponse, error) {
	// Write implementation here
	id, err := strconv.ParseInt(req.PetId, 10, 64)
	if err != nil {
		return nil, err
	}
	pet, err := s.db.GetPetQuery(ctx, id)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &bufbuild.GetPetResponse{
		Pet: &bufbuild.Pet{
			PetId:   fmt.Sprint(pet.ID),
			PetType: bufbuild.PetType(bufbuild.PetType_value[pet.Memo.String]),
			Name:    pet.Name,
			CreatedAt: &datetime.DateTime{
				Year:    int32(now.Year()),
				Month:   int32(now.Month()),
				Day:     int32(now.Day()),
				Hours:   int32(now.Hour()),
				Minutes: int32(now.Minute()),
				Seconds: int32(now.Second()),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil
}
