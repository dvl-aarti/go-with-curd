package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	EmployeeId  string             `bson:"employee_id,omitempty" json:"employee_id,omitempty"`
	Designation string             `bson:"designation,omitempty" json:"designation,omitempty"`
	EmailId     string             `bson:"email_id, omitempty" json:"email_id,omitempty"`
	DOB         primitive.DateTime `bson:"dob, omitempty" json:"dob,omitempty"`
	Skills      []string           `bson:"skills, omitempty" json:"skills,omitempty"`
}
