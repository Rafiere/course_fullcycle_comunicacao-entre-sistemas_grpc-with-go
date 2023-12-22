package service

import (
	"context"
	"github.com/rafiere/course_fullcycle_comunicacao-entre-sistemas_grpc-with-go/internal/database"
	"github.com/rafiere/course_fullcycle_comunicacao-entre-sistemas_grpc-with-go/internal/pb"
	"io"
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

// Esse serviço trabalhará com um stream de category, assim, poderemos criar e enviar as categories aos poucos.
func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}
