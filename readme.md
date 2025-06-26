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

