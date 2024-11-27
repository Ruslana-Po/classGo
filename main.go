package mypackage

import (
	"fmt"
	"hash/fnv"
)

// Item представляет элемент хэш-таблицы
type Item struct {
	key   string
	value string
	next  *Item
}

// HashTable представляет хэш-таблицу
type HashTable struct {
	sizeArr int
	tabl    []*Item
}

// NewHashTable создает новую хэш-таблицу с заданным размером
func NewHashTable(size int) *HashTable {
	return &HashTable{
		sizeArr: size,
		tabl:    make([]*Item, size),
	}
}

// Hash вычисляет хэш для заданного ключа
func (ht *HashTable) Hash(itemKey string) int {
	h := fnv.New32a()
	h.Write([]byte(itemKey))
	return int(h.Sum32()) % ht.sizeArr
}

// IsFull проверяет, заполнена ли хэш-таблица
func (ht *HashTable) IsFull() bool {
	count := 0
	for i := 0; i < ht.sizeArr; i++ {
		if ht.tabl[i] != nil {
			count++
		}
	}
	return count >= ht.sizeArr
}

// AddHash добавляет элемент в хэш-таблицу
func (ht *HashTable) AddHash(key, value string) {
	index := ht.Hash(key)

	// Проверка на наличие уже такого ключа
	current := ht.tabl[index]
	for current != nil {
		if current.key == key {
			fmt.Printf("Ключ '%s' уже существует. Значение не добавлено.\n", key)
			return
		}
		current = current.next
	}

	// Проверка на есть ли место
	if ht.IsFull() {
		fmt.Println("Хэш-таблица переполнена. Невозможно добавить новый элемент.")
		return
	}

	// Добавление элемента
	newItem := &Item{key: key, value: value, next: ht.tabl[index]}
	ht.tabl[index] = newItem
}

// KeyItem получает значение по ключу
func (ht *HashTable) KeyItem(key string) {
	index := ht.Hash(key)
	current := ht.tabl[index]
	for current != nil {
		if current.key == key {
			fmt.Printf("key: %s value: %s\n", key, current.value)
			return
		}
		current = current.next
	}
	fmt.Println("Такого ключа нет.")
}

// DelValue удаляет элемент по ключу
func (ht *HashTable) DelValue(key string) {
	index := ht.Hash(key)
	var prev *Item
	current := ht.tabl[index]
	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.tabl[index] = current.next
			} else {
				prev.next = current.next
			}
			return
		}
		prev = current
		current = current.next
	}
	fmt.Println("Такого ключа нет.")
}

// Print выводит содержимое хэш-таблицы
func (ht *HashTable) Print() {
	for i := 0; i < ht.sizeArr; i++ {
		current := ht.tabl[i]
		for current != nil {
			fmt.Printf("key: %s value: %s\n", current.key, current.value)
			current = current.next
		}
	}
}
