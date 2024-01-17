package input

type FamilyCreateInput struct {
	Name string
}

type FamilyUpdateInput struct {
	ID   string
	Name string
}

type FamilyDeleteInput struct {
	ID string
}
