package input

import "time"

type ChildProfileCreateInput struct {
	FirstName string
	LastName  string
	Gender    string
	Birthday  time.Time
	FamilyID  string
}

type ChildProfileUpdateInput struct {
	ID        string
	FirstName string
	LastName  string
	Gender    string
	Birthday  time.Time
	FamilyID  string
}

type ChildProfileDeleteInput struct {
	ID       string
	FamilyID string
}
