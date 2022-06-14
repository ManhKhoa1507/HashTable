package main

import (
	"fmt"
	"hash_table/table"
	"math"
)

var (
	TABLE_SIZE = 0
)

type HashTableValue struct {
	key   int
	value interface{}
}

func main() {
	fmt.Println("-----------Hash table using linked list-------------")

	// Init hash_map_data
	hashMapData := table.List{}

	// Scan number of elements and hash_list
	numberOfElement := getHashMapInput(&hashMapData)
	TABLE_SIZE = numberOfElement + 1

	//Init hash table with m index with m is prime and m >= n
	for isPrime(TABLE_SIZE) != true {
		TABLE_SIZE++
	}
	// Create Hashtable and call hash function
	hashTableArray := createTable()
	HashFunction(&hashMapData, hashTableArray)

	// Print hash table
	PrintHashTable(hashTableArray)

	// Search value using keyword
	keyword, value := getKeywordAndValue()
	indexResult := IndexSearch(keyword, value, hashTableArray)

	// Print hashTable[indexResult]
	if indexResult != -1 {
		PrintSearch(indexResult, hashTableArray)
	}
}

// Get input from screen then insert to linked list
// n: number of elements
// l: key list (linked list)
func getHashMapInput(hashMapData *table.List) int {
	// Read input from screen
	fmt.Println("------- Get input from screen-------")
	var numberOfElement int
	fmt.Println("Number of elements: ")
	fmt.Scanf("%d", &numberOfElement)

	for i := 0; i < numberOfElement; i++ {
		var (
			value   string
			tempKey string
		)

		fmt.Println("Element ", i)
		fmt.Scanf("%s %s", &tempKey, &value)

		key := createKey(tempKey)

		// fmt.Println("key: ", key)
		hashMapData.Insert(key, value)
	}
	return numberOfElement
}

// func get keyword and value for searching
func getKeywordAndValue() (string, string) {
	var keyword, value string

	fmt.Println("Enter keyword: ")
	fmt.Scanf("%s", &keyword)
	fmt.Println("Enter value: ")
	fmt.Scanf("%s", &value)

	return keyword, value
}

// Print hash table
func PrintHashTable(hashTable []HashTableValue) {
	fmt.Println("--------Hash table-------")
	lenTable := len(hashTable)

	for i := 0; i < lenTable; i++ {
		fmt.Println(i, hashTable[i].key, hashTable[i].value)
	}
}

// Func print search result
func PrintSearch(position int, hashTable []HashTableValue) {
	fmt.Println("Value: ", hashTable[position])
}

// Create key for elements
func createKey(value string) int {
	lenString := len(value)
	key := 0
	runes := []rune(value)

	for i := 0; i < lenString; i++ {
		key += int(runes[i])
	}
	return key
}

// Check prime number
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	sqRoot := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqRoot; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Create hash table using linked list with m elements
func createTable() []HashTableValue {
	// Init table array
	fmt.Println("-----Create table-----")
	var a = make([]HashTableValue, 0, TABLE_SIZE)

	for i := 0; i < TABLE_SIZE; i++ {

		// Create temp element to append to array
		element := new(HashTableValue)
		element.key = 0
		element.value = 0

		a = append(a, *element)
	}
	return a
}

// Hash function
func HashFunction(hashMapData *table.List, hashTable []HashTableValue) {
	listLen := hashMapData.LengthOfList()
	collisionCount := 0

	for i := 0; i < listLen; i++ {

		// Get node's value
		node := hashMapData.GetElementAtPosition(i)
		value := node.Value
		key := node.Key

		// Calc position
		positionA := hashFuncA(key, TABLE_SIZE)

		// If not collision add to positionA
		isEmpty := checkEmptyPosition(hashTable, positionA)
		if isEmpty == true {
			addToIndex(positionA, key, value, hashTable)
		} else {
			collisionCount++
			// collision solve
			modCount := 1

			for collisionSolve(positionA, key, value, hashTable, modCount) != true {

				// Increase modCount if cann't insert
				// fmt.Println("\n modCount: ", modCount)
				modCount++
			}
		}

	}
	fmt.Println("Collision count: ", collisionCount)
}

// Conllision Solve
func collisionSolve(positionA int, key int, value interface{}, hashTable []HashTableValue, modCount int) bool {

	// Calc b := hashFuncB, then calc the collision
	positionB := hashFuncB(key, TABLE_SIZE)
	collisionPosition := hashFuncIfCollision(positionA, positionB, modCount, TABLE_SIZE)
	// fmt.Println("Collision pos: ", collisionPosition)

	// If not empty return false
	// Else add value to collisionPosition
	if checkEmptyPosition(hashTable, collisionPosition) == false {
		return false
	} else {
		addToIndex(collisionPosition, key, value, hashTable)
		return true
	}
}

// Add value to index
func addToIndex(position int, keyAdd int, valueAdd interface{}, hashTable []HashTableValue) {
	// Check position to add
	// Add if empty

	if checkEmptyPosition(hashTable, position) == true {
		hashTable[position].key = keyAdd
		hashTable[position].value = valueAdd
	}
}

// Check position is empty or not
// Return false if not empty
// Return true if empty
func checkEmptyPosition(a []HashTableValue, position int) bool {
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

// hash if collision = (a+modCount*b)%mod
func hashFuncIfCollision(a int, b int, modCount int, mod int) int {
	return (a + modCount*b) % mod
}

// Search index using key
func IndexSearch(keyword string, value string, hashTable []HashTableValue) int {
	fmt.Println("------------Searching-------------")
	index := -1

	// Create key for searching
	key := createKey(keyword)

	// Calc position an check value at hashTable[position]
	positionA := hashFuncA(key, TABLE_SIZE)
	isValue := checkValue(key, value, positionA, hashTable)
	// isNil := checkNil(positionA, hashTable)

	// If found value and key not collision return hashTable[position]
	if isValue == true {
		index = positionA
	} else {
		// If collision, measn value != hashtable[position].value
		index = collisionSearch(positionA, key, value, hashTable)
	}
	return index
}

// Collision searching
func collisionSearch(positionA int, key int, value string, hashTable []HashTableValue) int {
	// If collision, measn value != hashtable[position].value
	fmt.Println("---------Collision seaching------------")
	positionB := hashFuncB(key, TABLE_SIZE)
	modCount := 0
	isValue := false
	index := -1

	// If not value find another position
	// Increase modCount and check value
	for isValue != true {

		// Calc another position
		modCount++
		collisionPosition := hashFuncIfCollision(positionA, positionB, modCount, TABLE_SIZE)
		isValue = checkValue(key, value, collisionPosition, hashTable)
		isNil := checkNil(collisionPosition, hashTable)

		// If not found key or value return -1
		if isNil {
			fmt.Println("Key not found")
			break
		}

		// If found right value break
		if isValue {
			index = collisionPosition
			break
		}
	}

	// Return index position
	// If not found return -1
	return index
}

// CheckValue if hashTable[postion].key == key and hashTable[postion].value == value
// Return true if found, else return false
func checkValue(key int, value string, position int, hashTable []HashTableValue) bool {
	if hashTable[position].value == value && hashTable[position].key == key {
		fmt.Println("Found value at position: ", position)
		return true
	}
	return false
}

// Check i hashTable[position].value == 0 or hashTable[position].key == 0
// Return true if =0
// Else return 0
func checkNil(position int, hashTable []HashTableValue) bool {
	if hashTable[position].key == 0 || hashTable[position].value == 0 {
		return true
	}
	return false
}
