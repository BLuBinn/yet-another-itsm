package service

import (
	"context"
	"fmt"

	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/dtos"
	"yet-another-itsm/internal/repository"
	"yet-another-itsm/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type UserService interface {
	GetAllUsersInDepartment(ctx context.Context, departmentID string) ([]*dtos.User, error)
	GetUserByID(ctx context.Context, id string) (*dtos.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dtos.User, error)
	CreateUser(ctx context.Context, req *dtos.CreateUserRequest) (*dtos.User, error)
}

type userService struct {
	repo *repository.Queries
}

func NewUserService(repo *repository.Queries) UserService {
	return &userService{
		repo: repo,
	}
}

// GetAllUsersInDepartment gets all users in a department.
func (s *userService) GetAllUsersInDepartment(ctx context.Context, departmentID string) ([]*dtos.User, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	var uuid pgtype.UUID
	err = uuid.Scan(departmentID)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrInvalidDepartmentUUIDFormat, err)
	}

	log.Info().
		Str("service", "UserService").
		Str("endpoint", "GetAllUsersInDepartment").
		Str("department_id", departmentID).
		Str("user_id", userID).
		Msg("Getting all users in department")

	repoUsers, err := s.repo.GetAllUsersInDepartment(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetUsers, err)
	}

	var users []*dtos.User
	for _, repoUser := range repoUsers {
		dto := (&dtos.User{}).FromRepositoryModel(repoUser)
		users = append(users, dto)
	}

	return users, nil
}

// GetUserByID gets a user by ID.
func (s *userService) GetUserByID(ctx context.Context, id string) (*dtos.User, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	var uuid pgtype.UUID
	err = uuid.Scan(id)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrInvalidUUIDFormat, err)
	}

	log.Info().
		Str("service", "UserService").
		Str("endpoint", "GetUserByID").
		Str("id", id).
		Str("user_id", userID).
		Msg("Getting user by ID")

	repoUser, err := s.repo.GetUserByID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetUser, err)
	}

	dto := (&dtos.User{}).FromRepositoryModel(repoUser)
	return dto, nil
}

// GetUserByEmail gets a user by email.
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*dtos.User, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "UserService").
		Str("endpoint", "GetUserByEmail").
		Str("email", email).
		Str("user_id", userID).
		Msg("Getting user by email")

	repoUser, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetUser, err)
	}

	dto := (&dtos.User{}).FromRepositoryModel(repoUser)
	return dto, nil
}

// CreateUser creates a new user.
func (s *userService) CreateUser(ctx context.Context, req *dtos.CreateUserRequest) (*dtos.User, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetUserID, err)
	}

	log.Info().
		Str("service", "UserService").
		Str("endpoint", "CreateUser").
		Str("email", req.Mail).
		Str("user_id", userID).
		Msg("Creating new user")

	params := repository.CreateUserParams{
		AzureAdObjectID: req.AzureAdObjectID,
		Mail:            req.Mail,
		DisplayName:     req.DisplayName,
		IsActive:        pgtype.Bool{Bool: req.IsActive, Valid: true},
	}

	if req.HomeTenantID != "" {
		err = params.HomeTenantID.Scan(req.HomeTenantID)
		if err != nil {
			return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrInvalidHomeTenantUUIDFormat, err)
		}
	}
	if req.DepartmentID != "" {
		err = params.DepartmentID.Scan(req.DepartmentID)
		if err != nil {
			return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrInvalidDepartmentUUIDFormat, err)
		}
	}
	if req.ManagerID != "" {
		err = params.ManagerID.Scan(req.ManagerID)
		if err != nil {
			return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrInvalidManagerUUIDFormat, err)
		}
	}

	if req.GivenName != "" {
		params.GivenName = pgtype.Text{String: req.GivenName, Valid: true}
	}
	if req.SurName != "" {
		params.SurName = pgtype.Text{String: req.SurName, Valid: true}
	}
	if req.JobTitle != "" {
		params.JobTitle = pgtype.Text{String: req.JobTitle, Valid: true}
	}
	if req.OfficeLocation != "" {
		params.OfficeLocation = pgtype.Text{String: req.OfficeLocation, Valid: true}
	}

	repoUser, err := s.repo.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToCreateUser, err)
	}

	dto := (&dtos.User{}).FromRepositoryModel(repoUser)
	return dto, nil
}
