package structcopy

import (
	"testing"
)

// Benchmark for reflection method
func BenchmarkOverwriteStructReflect(b *testing.B) {
	oldStruct := &MyStruct{
		Name:       "John",
		Age:        30,
		Email:      "john@example.com",
		Address:    "123 Main St",
		Phone:      "123-456-7890",
		Country:    "USA",
		City:       "New York",
		ZipCode:    "10001",
		Username:   "john_doe",
		Password:   "password123",
		Company:    "Tech Corp",
		Position:   "Developer",
		Website:    "www.johndoe.com",
		LinkedIn:   "john-linkedin",
		Twitter:    "john-twitter",
		Github:     "john-github",
		Biography:  "Software Engineer",
		Department: "IT",
		Skills:     "Go, Python, JavaScript",
	}
	newStruct := &MyStruct{
		Name:       "Johnny",
		Age:        0,
		Email:      "",
		Address:    "456 Side St",
		Phone:      "",
		Country:    "Canada",
		City:       "",
		ZipCode:    "",
		Username:   "",
		Password:   "",
		Company:    "",
		Position:   "",
		Website:    "",
		LinkedIn:   "",
		Twitter:    "",
		Github:     "",
		Biography:  "",
		Department: "",
		Skills:     "",
	}

	for i := 0; i < b.N; i++ {
		overwriteStructReflect(oldStruct, newStruct)
	}
}

// Benchmark for regular if statements method
func BenchmarkOverwriteStructIf(b *testing.B) {
	oldStruct := &MyStruct{
		Name:       "John",
		Age:        30,
		Email:      "john@example.com",
		Address:    "123 Main St",
		Phone:      "123-456-7890",
		Country:    "USA",
		City:       "New York",
		ZipCode:    "10001",
		Username:   "john_doe",
		Password:   "password123",
		Company:    "Tech Corp",
		Position:   "Developer",
		Website:    "www.johndoe.com",
		LinkedIn:   "john-linkedin",
		Twitter:    "john-twitter",
		Github:     "john-github",
		Biography:  "Software Engineer",
		Department: "IT",
		Skills:     "Go, Python, JavaScript",
	}
	newStruct := &MyStruct{
		Name:       "Johnny",
		Age:        0,
		Email:      "",
		Address:    "456 Side St",
		Phone:      "",
		Country:    "Canada",
		City:       "",
		ZipCode:    "",
		Username:   "",
		Password:   "",
		Company:    "",
		Position:   "",
		Website:    "",
		LinkedIn:   "",
		Twitter:    "",
		Github:     "",
		Biography:  "",
		Department: "",
		Skills:     "",
	}

	for i := 0; i < b.N; i++ {
		overwriteStructIf(oldStruct, newStruct)
	}
}
