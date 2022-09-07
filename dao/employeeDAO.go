package dao

import (
	"context"
	"errors"
	"log"

	"employee-management/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmployeeDAO struct {
	Server     string
	Database   string
	Collection string
}

var Collection *mongo.Collection
var ctx = context.TODO()

func (e *EmployeeDAO) Connect() {

	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(e.Database).Collection(e.Collection)
}

func (e *EmployeeDAO) Insert(employee model.Employee) error {

	_, err := Collection.InsertOne(ctx, employee)

	if err != nil {
		return errors.New("Unable to create new record")
	}

	return nil
}

func (e *EmployeeDAO) FindByEmpId(empId string) ([]*model.Employee, error) {

	var employees []*model.Employee

	cur, err := Collection.Find(ctx, bson.D{primitive.E{Key: "employee_id", Value: empId}})

	if err != nil {
		return employees, errors.New("Unable to query db")
	}

	for cur.Next(ctx) {
		var e model.Employee

		err := cur.Decode(&e)

		if err != nil {
			return employees, err
		}

		employees = append(employees, &e)
	}

	if err := cur.Err(); err != nil {
		return employees, err
	}

	cur.Close(ctx)

	if len(employees) == 0 {
		return employees, mongo.ErrNoDocuments
	}

	return employees, nil
}

func (e *EmployeeDAO) DeleteEmployee(empId string) error {

	filter := bson.D{primitive.E{Key: "employee_id", Value: empId}}

	res, err := Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no employee were deleted")
	}

	return nil
}

func (epd *EmployeeDAO) UpdateEmployee(empId string, emp model.Employee) error {

	filter := bson.D{primitive.E{Key: "employee_id", Value: empId}}

	update := bson.D{primitive.E{Key: "$set", Value: emp}}

	e := &model.Employee{}
	return Collection.FindOneAndUpdate(ctx, filter, update).Decode(e)

}
