package main

import (
	"fmt"
	"hash_table/table"
	"math"
)

var (
	TABLE_SIZE = 0
)

type hash_table_value struct {
	key   int
	value interface{}
}

func main() {
	fmt.Println("-----------Hash table using linked list-------------")

	// Init hash_map_data
	hash_map_data := table.List{}

	// Scan number of elements
	numberOfElement := getInput(&hash_map_data)
	TABLE_SIZE = numberOfElement + 1

	//Init hash table with m index with m is prime and m >= n
	for isPrime(TABLE_SIZE) != true {
		TABLE_SIZE++
	}
	// Create Hashtable
	hash_table_array := createTable()
	HashFunction(&hash_map_data, hash_table_array)
}

// Get input from screen then insert to linked list
// n: number of elements
// l: key list (linked list)
func getInput(hash_map_data *table.List) int {
	// Read input from screen
	fmt.Println("------- Get input from screen-------")
	var numberOfElement int
	fmt.Scanf("%d", &numberOfElement)

	for i := 0; i < numberOfElement; i++ {
		var value int
		var temp_key string

		fmt.Println("Element ", i)
		fmt.Scanf("%v %v", &temp_key, &value)

		key := createKey(temp_key, i+1)

		// fmt.Println("key: ", key)
		hash_map_data.Insert(key, value)
	}
	return numberOfElement
}

// Create key for elements
func createKey(value string, position int) int {
	key := sumString(value) * position
	return key
}

// Calculate sum of string
func sumString(str string) int {
	len_string := len(str)
	sum := 0
	runes := []rune(str)

	for i := 0; i < len_string; i++ {
		sum += int(runes[i])
	}
	return sum
}

// Check prime number
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Create hash table using linked list with m elements
func createTable() []hash_table_value {
	// Init table array
	fmt.Println("-----Create table-----")
	var a = make([]hash_table_value, 0, TABLE_SIZE)

	for i := 0; i < TABLE_SIZE; i++ {

		// Create temp element to append to array
		element := new(hash_table_value)
		element.key = 0
		element.value = 0

		a = append(a, *element)
	}
	return a
}

// Hash function
func HashFunction(hash_map_data *table.List, hash_table []hash_table_value) {
	list_len := hash_map_data.LengthOfList()
	collision_count := 0

	for i := 0; i < list_len; i++ {

		// Get node's value
		node := hash_map_data.GetElementAtPosition(i)
		value := node.Value
		key := node.Key

		// Calc position
		position_a := hashFuncA(key, TABLE_SIZE)

		// If not collision add to position_a
		isEmpty := checkEmptyPosition(hash_table, position_a)
		if isEmpty == true {
			addToIndex(position_a, key, value, hash_table)
		} else {
			collision_count++
			// collision solve
			mod_count := 1

			for collisionSolve(position_a, key, value, hash_table, mod_count) != true {

				// Increase mod_count if cann't insert
				// fmt.Println("\n mod_count: ", mod_count)
				mod_count++
			}
		}

	}
	fmt.Println(hash_table)
	fmt.Println("Collision count: ", collision_count)
}

// Conllision Solve
func collisionSolve(position_a int, key int, value interface{}, hash_table []hash_table_value, mod_count int) bool {

	// Calc b := hashFuncB, then calc the collision
	position_b := hashFuncB(key, TABLE_SIZE)
	collision_position := hashFuncIfCollision(position_a, position_b, mod_count, TABLE_SIZE)
	// fmt.Println("Collision pos: ", collision_position)

	// If not empty return false
	// Else add value to collision_position
	if checkEmptyPosition(hash_table, collision_position) == false {
		return false
	} else {
		addToIndex(collision_position, key, value, hash_table)
		return true
	}
}

// Add value to index
func addToIndex(position int, key_add int, value_add interface{}, hash_table []hash_table_value) {
	// Check position to add
	// Add if empty

	if checkEmptyPosition(hash_table, position) == true {
		hash_table[position].key = key_add
		hash_table[position].value = value_add
	}
}

// Check position is empty or not
// Return false if not empty
// Return true if empty
func checkEmptyPosition(a []hash_table_value, position int) bool {
	if a[position].key != 0 {
		return false
	} else {
		return true
	}
}

// hashFunc
// a = hash(key) = key % M
// M is prime to avoid collision
func hashFuncA(key int, mod int) int {
	return key % mod
}

// b = hash(key) = (mod-2) - key%(mod-2)
func hashFuncB(key int, mod int) int {
	return (mod - 2) - key%(mod-2)
}

// hash if collision = (a+mod_count*b)%mod
func hashFuncIfCollision(a int, b int, mod_count int, mod int) int {
	return (a + mod_count*b) % mod
}
