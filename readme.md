## Golang (Go)

### Slices
A slice is a data structure that provides a more powerful and convinient way to work with sequences of elements compared to fixed size arrays.

A slice is backed by an underlying array and is defined in three compenents namely:

`Pointer`: Points in the start of the slice in the underlying array.

`Length`: The number of elements in the slice (`len(slice)`)

`Capacity`: The number of elements in the underlying array from the start of the slice (`cap(slice)`)

Example:

Imagine you have a box of 10 chocolates (an array), but you only want to share 3 of them with your friends. You can create a "slice" of that box, which is like picking out just those 3 chocolates to show or work with. If you later decide to share 2 more, you can extend the slice (as long as it fits within the original box).

In Go:

>An array is like the whole box: `[10]string{"choco1", "choco2", ..., "choco10"}`.

>A slice is a portion of that box: `slice := array[2:5] (elements from index 2 to 4)`.

Actual implementation

```Go
package main

import "fmt"

func main() {
    //Create a slice
    //A slice can be created directly without defining an array

    fruits := []string{"apple", "banana", "orange"}
    fmt.Println("Initial slice:", fruits) //Output [apple banana orange]
	fmt.Println("Length:", len(fruits)) //Output 3
	fmt.Println("Capacity:", cap(fruits)) //Output 3

	//Creating a slice from an array
	numbers :=[5]int{1, 2, 3, 4, 5} //An array
	sliceFromArray := numbers[1:4] // Slice form index 1 to 3 
	fmt.Println("Slice from array", sliceFromArray)//output [2 3 4]
	fmt.Println("Length", len(sliceFromArray)) // 3
	fmt.Println("Capacity", cap(sliceFromArray)) // 4 (from index 1 to end of array)


	//Appending to a slice
	fruits = append(fruits, "Mango", "grape")
	fmt.Println("After appending fruits:", fruits) // Output [apple banana orange mango grape]
	fmt.Println("New Length:", len(fruits)) // 5
	fmt.Println("New Capacity", cap(fruits)) //6  since current capacity 3 is insufficient to hold 5 elements, the new capacity is doubled

	//Modifying a slice
	fruits[1] = "Kiwi"
	fmt.Println("After modification", fruits) //[apple kiwi orange mango grape]

	//Slicing a slice
	subSlice := fruits[1:4] //Slicing from Index 1-3
	fmt.Println("Sub-slice", subSlice) //Output [Kiwi Orange Mango]
}
```

#### Make()
The `make()` function in Go is used to create and initialize certain built-in types, including slices, maps and channels.
When used with slices, `make()` allows you to create a slice with a specified length and capacity, backed by an underlying array.
This is particularly useful when you want to control the initial size and potential growth of the slice, avoiding unnecessary reallocations.

```Go
slice := make([]type, length, capacity)
```

Example:

```Go
package main

import "fmt"

func main(){
    slice := make([]string, 3, 5)
    fmt.Println("Length:", len(slice));
    fmt.Println("Capacity:", cap(slice));
    fmt.Println(slice)
}
```

#### Map
A map is a built-in data structure that stores key-value pairs.

Characteristics:

`Keys` must be of type that supports equality (eg strings, integers, structs with comparable fields). Slices, maps, and functions cannot be keys

`values` can be of any type

`Dynamic` Maps grow dynamically as you add key-value pairs

`Reference Type` Maps are reference types, meaning they are `nil` until initialized (eg using `make()` or a map literal)

Friendly example

Think of a map like a phonebook whic haas the name as the key and phone number as value. You can add, update or remove entries as needed.

Actual implementation

```Go
package main
import "fmt"

func main(){
	//Create map using make
	phoneBook := make(map[string] string)
	fmt.Println("Initial Map:", phoneBook) //Output map[]

	//Adding key value pairs
	phoneBook["Alice"] = "123-456-789"
	phoneBook["Bob"] = "987-654-321"
	fmt.Println("After adding entries:", phoneBook) //map[Alice:123-456-789 Bob:987-654-321]

	//Accessing a value by key
	aliceNumber := phoneBook["Alice"]
	fmt.Println("Alice number:", aliceNumber) //123-456-789

	//Checking if a key exists
	number, exists :=phoneBook["Charlie"]
	if exists {
		fmt.Println("Charlie's number", number)
	} else{
		fmt.Println("Charlie is not found in phonebook") //Charlie is not found in phonebook"
	}

	//Updating a value
	phoneBook["Alice"] = "555-555-555"
	fmt.Println("After updating Alice:", phoneBook) //Output map[Alice:555-555-555 Bob:987-654-321]

	//Deleting a key-value pair
	delete(phoneBook, "Bob")
	fmt.Println("After deleting Bob:", phoneBook)//map[Alice:555-555-555]

	//Length of the map
	fmt.Println("Length:", len(phoneBook)) // 1 (Only one entry)

	//Creating a map using a map literal
	scores := map[string]int{
		"Alice": 95,
		"Bob": 87,
	}
	fmt.Println("Scores map:", scores) //[Alice:95 Bob:87]

}
```

#### Struct
A struct is a user-defined data type that groups together a collection of fields, each with its own type and name.

Structs are useful for organizing complex data, like representing a real-world entity (eg, a person, a car, or book) with multiple attributes.

Characteristics

`Fields`: Each field in a struct has a name and a type eg, `Name String`, `Age int`.

`Custom Type`: You define a struct using the `type` keyword, and it can be used to create instances(variables) of that type.

`Value Type` Structs are value types, meaning they are copied when assigned or passed to functions(unless pointers are used)

`Flexible`Structs can contain fields different types, including other structs, slices, maps,etc

Friendly example:

Imagine you're organizing information about a student. A student has a name, age and a grade. Instead of using seperate variables for each peace of information, you can create a struct called `Student` to group these records together to a single record in databse.

For example:

A `Student` struct might have:
<ul>
	<li>Name: "Alice" (string)</li>
	<li>Age: 20 (int)</li>
	<li>Grade: "A" (string)</li>
</ul>

Actual implementation