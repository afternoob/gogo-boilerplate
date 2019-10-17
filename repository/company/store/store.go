package store

import (
	"context"

	domain "github.com/afternoob/gogo-boilerplate/domain/company"
	repoCompany "github.com/afternoob/gogo-boilerplate/repository/company"
	"github.com/devit-tel/goerror"
	"github.com/devit-tel/jaegerstart"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	mongo          *mongo.Client
	dbName         string
	collectionName string
}

func New(mongoEndpoint, dbName, collectionName string) *Store {
	clientOptions := options.Client().ApplyURI(mongoEndpoint)

	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	return &Store{
		dbName:         dbName,
		collectionName: collectionName,
		mongo:          db,
	}
}

func (s *Store) collectionCompany() *mongo.Collection {
	return s.mongo.Database(s.dbName).Collection(s.collectionName)
}

func (s *Store) Get(ctx context.Context, companyId string) (*domain.Company, goerror.Error) {
	if span := jaegerstart.StartNewSpan(ctx, "REPO_COMPANY_Get"); span != nil {
		defer span.Finish()
	}

	company := &domain.Company{}
	if err := s.collectionCompany().FindOne(ctx, bson.D{{"_id", companyId}}).Decode(company); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repoCompany.ErrCompanyNotFound.WithInput(companyId)
		}

		return nil, repoCompany.ErrUnableGetCompany.WithInput(companyId).WithCause(err)
	}

	return company, nil
}

func (s *Store) Save(ctx context.Context, company *domain.Company) goerror.Error {
	if span := jaegerstart.StartNewSpan(ctx, "REPO_COMPANY_Save"); span != nil {
		defer span.Finish()
	}

	_, err := s.collectionCompany().InsertOne(ctx, company)
	if err != nil {
		return repoCompany.ErrUnableSaveCompany.WithInput(company).WithCause(err)
	}

	return nil
}
