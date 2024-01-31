package main

import "fmt"

// Vertex is a struct that represents a point on the earth, 
// defined by its latitude (Lat) and longitude (Lon).
type Vertex struct {
    Lat, Lon float64
}

// m is a map where the keys are strings and the values are Vertex objects.
var m map[string]Vertex

func main() {
    // Print the initial value of m. As it's not initialized, it should print "nil".
    fmt.Println(m)

    // Check if m is nil, which it should be at this point.
    if m == nil {
        fmt.Println("m is nil")
    }

    // Initialize m using the make function. This allocates and initializes 
    // a hash map data structure and returns a map value that points to it.
    m = make(map[string]Vertex)
    
    // Check if m is not nil, which it shouldn't be after the call to make.
    if m != nil {
        fmt.Println("m is NOT nil")
    }

    // Add a new element to the map with the key "ODDS" and a value of a Vertex object.
    m["ODDS"] = Vertex{
        40.68433, -74.39967,
    }

    // Print the map after adding the new element. It should print the map with the new element.
    fmt.Println(m)

	json := map[string]string{}
	json["name"] = "John"
	fmt.Println(json)
}