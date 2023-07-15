package graph

import (
	"context"
	"time"

	"github.com/1rvyn/graphql-service/graph/model"
)

// TODO: Make this use the GORM database to actually add real users and fufil the project requirements
func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	// Placeholder code
	return &model.Employee{
		ID:           "1",
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		Password:     input.Password,
		Email:        input.Email,
		Dob:          input.Dob,
		DepartmentID: input.DepartmentID,
		Position:     input.Position,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}, nil
}
func (r *mutationResolver) UpdateEmployee(ctx context.Context, id string, input model.UpdateEmployee) (*model.Employee, error) {
	var firstName, lastName, username, password, email, dob, departmentID, position string

	// Check each field to see if it's nil before assigning the value
	if input.FirstName != nil {
		firstName = *input.FirstName
	}
	if input.LastName != nil {
		lastName = *input.LastName
	}
	if input.Username != nil {
		username = *input.Username
	}
	if input.Password != nil {
		password = *input.Password
	}
	if input.Email != nil {
		email = *input.Email
	}
	if input.Dob != nil {
		dob = *input.Dob
	}
	if input.DepartmentID != nil {
		departmentID = *input.DepartmentID
	}
	if input.Position != nil {
		position = *input.Position
	}

	// Placeholder code
	return &model.Employee{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Username:     username,
		Password:     password,
		Email:        email,
		Dob:          dob,
		DepartmentID: departmentID,
		Position:     position,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}, nil
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (*bool, error) {
	// Placeholder code
	result := true
	return &result, nil
}

func (r *queryResolver) Employees(ctx context.Context) ([]*model.Employee, error) {
	// Placeholder code
	return []*model.Employee{
		{
			ID:           "1",
			FirstName:    "John",
			LastName:     "Doe",
			Username:     "johndoe",
			Password:     "password",
			Email:        "john.doe@example.com",
			Dob:          "2000-01-01",
			DepartmentID: "1",
			Position:     "Developer",
			CreatedAt:    time.Now().String(),
			UpdatedAt:    time.Now().String(),
		},
	}, nil
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*model.Employee, error) {
	// Placeholder code
	return &model.Employee{
		ID:           "1",
		FirstName:    "John",
		LastName:     "Doe",
		Username:     "johndoe",
		Password:     "password",
		Email:        "john.doe@example.com",
		Dob:          "2000-01-01",
		DepartmentID: "1",
		Position:     "Developer",
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
