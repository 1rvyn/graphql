package graph

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/1rvyn/graphql-service/database"
	"github.com/1rvyn/graphql-service/graph/model"
	"github.com/1rvyn/graphql-service/models"
	"gorm.io/gorm"
)

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	// Convert dob from string to time.Time
	dob, err := time.Parse("2006-01-02", input.Dob)
	if err != nil {
		return nil, err
	}

	// Convert departmentID from string to uint
	departmentID, err := strconv.Atoi(input.DepartmentID)
	if err != nil {
		return nil, err
	}

	// Check if an employee with the same username or email already exists
	var existingEmployee models.Employee
	result := database.Database.Db.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingEmployee)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// An unexpected error occurred
			return nil, result.Error
		}
		// No employee with the same username or email exists, continue to create the new employee
	} else {
		// An employee with the same username or email already exists
		return nil, fmt.Errorf("an employee with the same username or email already exists")
	}

	employee := &models.Employee{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		Password:     input.Password,
		Email:        input.Email,
		DOB:          dob,
		DepartmentID: uint(departmentID),
		Position:     input.Position,
	}

	if err := database.Database.Db.Create(employee).Error; err != nil {
		return nil, err
	}

	// Convert models.Employee to model.Employee before returning
	return &model.Employee{
		ID:           strconv.Itoa(int(employee.ID)),
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Username:     employee.Username,
		Password:     employee.Password,
		Email:        employee.Email,
		Dob:          employee.DOB.Format("2006-01-02"),
		DepartmentID: strconv.Itoa(int(employee.DepartmentID)),
		Position:     employee.Position,
		CreatedAt:    employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    employee.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, id string, input model.UpdateEmployee) (*model.Employee, error) {
	// Convert id from string to uint
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Find the employee to be updated
	var employee models.Employee
	if err := database.Database.Db.First(&employee, employeeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("employee not found")
		}
		return nil, err
	}

	// Update the fields if they are set
	if input.FirstName != nil {
		employee.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		employee.LastName = *input.LastName
	}
	if input.Username != nil {
		employee.Username = *input.Username
	}
	if input.Password != nil {
		employee.Password = *input.Password
	}
	if input.Email != nil {
		employee.Email = *input.Email
	}
	if input.Dob != nil {
		dob, err := time.Parse("2006-01-02", *input.Dob)
		if err != nil {
			return nil, err
		}
		employee.DOB = dob
	}
	if input.DepartmentID != nil {
		departmentID, err := strconv.Atoi(*input.DepartmentID)
		if err != nil {
			return nil, err
		}
		employee.DepartmentID = uint(departmentID)
	}
	if input.Position != nil {
		employee.Position = *input.Position
	}

	// Save the updated employee to the database
	if err := database.Database.Db.Save(&employee).Error; err != nil {
		return nil, err
	}

	// Convert models.Employee to model.Employee before returning
	return &model.Employee{
		ID:           strconv.Itoa(int(employee.ID)),
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Username:     employee.Username,
		Password:     employee.Password,
		Email:        employee.Email,
		Dob:          employee.DOB.Format("2006-01-02"),
		DepartmentID: strconv.Itoa(int(employee.DepartmentID)),
		Position:     employee.Position,
		CreatedAt:    employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    employee.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (*bool, error) {
	// Convert id from string to uint
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Find the employee to be deleted
	var employee models.Employee
	if err := database.Database.Db.First(&employee, employeeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("employee not found")
		}
		return nil, err
	}

	// Delete the employee from the database
	if err := database.Database.Db.Delete(&employee).Error; err != nil {
		return nil, err
	}

	// Return true if the deletion was successful
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
