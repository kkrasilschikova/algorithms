package hashtables

/*
HashTable returns a hash table where collisions are resolved by chaining
*/
func HashTable(input []int, size int) map[int][]int {
	hashTable := make(map[int][]int)
	for _, v := range input {
		key := v % size
		hashTable[key] = append(hashTable[key], v)
	}
	return hashTable
}
