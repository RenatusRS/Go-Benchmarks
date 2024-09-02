package structcopy

import "reflect"

// Define the struct type with 15 additional fields
type MyStruct struct {
	Name       string
	Age        int
	Email      string
	Address    string
	Phone      string
	Country    string
	City       string
	ZipCode    string
	Username   string
	Password   string
	Company    string
	Position   string
	Website    string
	LinkedIn   string
	Twitter    string
	Github     string
	Biography  string
	Department string
	Skills     string
}

// Overwrite oldStruct with newStruct values using reflection
func overwriteStructReflect[T any](oldStruct, newStruct *T) {
	oldVal := reflect.ValueOf(oldStruct).Elem()
	newVal := reflect.ValueOf(newStruct).Elem()

	for i := 0; i < oldVal.NumField(); i++ {
		oldField := oldVal.Field(i)
		newField := newVal.Field(i)

		// Check if the new field has a non-zero value
		if !reflect.DeepEqual(newField.Interface(), reflect.Zero(newField.Type()).Interface()) {
			oldField.Set(newField)
		}
	}
}

// Overwrite oldStruct with newStruct values using regular if statements
func overwriteStructIf(oldStruct, newStruct *MyStruct) {
	if newStruct.Name != "" {
		oldStruct.Name = newStruct.Name
	}
	if newStruct.Age != 0 {
		oldStruct.Age = newStruct.Age
	}
	if newStruct.Email != "" {
		oldStruct.Email = newStruct.Email
	}
	if newStruct.Address != "" {
		oldStruct.Address = newStruct.Address
	}
	if newStruct.Phone != "" {
		oldStruct.Phone = newStruct.Phone
	}
	if newStruct.Country != "" {
		oldStruct.Country = newStruct.Country
	}
	if newStruct.City != "" {
		oldStruct.City = newStruct.City
	}
	if newStruct.ZipCode != "" {
		oldStruct.ZipCode = newStruct.ZipCode
	}
	if newStruct.Username != "" {
		oldStruct.Username = newStruct.Username
	}
	if newStruct.Password != "" {
		oldStruct.Password = newStruct.Password
	}
	if newStruct.Company != "" {
		oldStruct.Company = newStruct.Company
	}
	if newStruct.Position != "" {
		oldStruct.Position = newStruct.Position
	}
	if newStruct.Website != "" {
		oldStruct.Website = newStruct.Website
	}
	if newStruct.LinkedIn != "" {
		oldStruct.LinkedIn = newStruct.LinkedIn
	}
	if newStruct.Twitter != "" {
		oldStruct.Twitter = newStruct.Twitter
	}
	if newStruct.Github != "" {
		oldStruct.Github = newStruct.Github
	}
	if newStruct.Biography != "" {
		oldStruct.Biography = newStruct.Biography
	}
	if newStruct.Department != "" {
		oldStruct.Department = newStruct.Department
	}
	if newStruct.Skills != "" {
		oldStruct.Skills = newStruct.Skills
	}
}
