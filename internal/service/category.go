package service

import (
	"context"
	"github.com/rafiere/course_fullcycle_comunicacao-entre-sistemas_grpc-with-go/internal/database"
	"github.com/rafiere/course_fullcycle_comunicacao-entre-sistemas_grpc-with-go/internal/pb"
)

// Para dizermos que estamos criando o serviço, precisaremos
// implementar o "UnimplementedCategoryServiceServer".
type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {

	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

/* Vamos implementar o serviço de criação de categoria. */
func (c *CategoryService) CreateCategory(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {

	category, err := c.CategoryDB.Create(request.Name, request.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}
