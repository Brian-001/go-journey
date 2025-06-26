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