package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Person struct {
	Name string
	Age  int
}

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  38,
	}

	actual := Person{
		Name: "Deniis",
		Age:  38,
	}

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Error(diff)
	}

}
