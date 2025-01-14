package set

import (
    "testing"
)


type Person struct {
    FirstName string
    LastName  string
}


func TestSet(t *testing.T) {
    // Test with int
    intSet := make(Set[int])
    intSet.Add(1)
    intSet.Add(2)
    if !intSet.Contains(1) || !intSet.Contains(2) {
        t.Errorf("Expected set to contain 1 and 2")
    }
    intSet.Remove(1)
    if intSet.Contains(1) {
        t.Errorf("Expected set to not contain 1")
    }
    intSet.Clear()
    if intSet.Size() != 0 {
        t.Errorf("Expected set to be empty")
    }

    // Test with string
    stringSet := make(Set[string])
    stringSet.Add("hello")
    stringSet.Add("world")
    if !stringSet.Contains("hello") || !stringSet.Contains("world") {
        t.Errorf("Expected set to contain 'hello' and 'world'")
    }
    stringSet.Remove("hello")
    if stringSet.Contains("hello") {
        t.Errorf("Expected set to not contain 'hello'")
    }
    stringSet.Clear()
    if stringSet.Size() != 0 {
        t.Errorf("Expected set to be empty")
    }

    // Test with custom type
    personSet := make(Set[Person])
    personSet.Add(Person{FirstName: "John", LastName: "Doe"})
    personSet.Add(Person{FirstName: "Jane", LastName: "Smith"})
    if !personSet.Contains(Person{FirstName: "John", LastName: "Doe"}) ||
       !personSet.Contains(Person{FirstName: "Jane", LastName: "Smith"}) {
        t.Errorf("Expected set to contain John Doe and Jane Smith")
    }
    personSet.Remove(Person{FirstName: "John", LastName: "Doe"})
    if personSet.Contains(Person{FirstName: "John", LastName: "Doe"}) {
        t.Errorf("Expected set to not contain John Doe")
    }
    personSet.Clear()
    if personSet.Size() != 0 {
        t.Errorf("Expected set to be empty")
    }
    // Test creating a new set with values 
    testingSet := NewSet("Testing", "Set", "with", "values")
    if testingSet.Size() != 4 {
        t.Errorf("Expected set to contain 4 values")
    }
}
