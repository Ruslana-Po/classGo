package mypackage

import (
	"bytes"
	"os"
	"testing"
)

// создает хэш-таблицу с правильным размером и инициализирует массив правильно.
func TestNewHashTable(t *testing.T) {
	ht := NewHashTable(10)
	if ht.sizeArr != 10 {
		t.Errorf("Expected sizeArr to be 10, got %d", ht.sizeArr)
	}
	if len(ht.tabl) != 10 {
		t.Errorf("Expected tabl length to be 10, got %d", len(ht.tabl))
	}
}

// Проверяет, что функция Hash возвращает корректный для заданного ключа
func TestHash(t *testing.T) {
	ht := NewHashTable(10)
	index := ht.Hash("test")
	if index < 0 || index >= ht.sizeArr {
		t.Errorf("Expected index to be between 0 and %d, got %d", ht.sizeArr-1, index)
	}
}

// Проверяет, что функция AddHash добавляет элементы в хэш-таблицу и не добавляет элементы с дублирующимися ключами.
func TestAddHash(t *testing.T) {
	ht := NewHashTable(4)
	ht.AddHash("key1", "value1")
	ht.AddHash("key2", "value2")
	ht.AddHash("key2", "value3")

	index1 := ht.Hash("key1")
	index2 := ht.Hash("key2")

	if ht.tabl[index1].key != "key1" || ht.tabl[index1].value != "value1" {
		t.Errorf("Expected key1 and value1 at index %d, got %v", index1, ht.tabl[index1])
	}

	if ht.tabl[index2].key != "key2" || ht.tabl[index2].value != "value2" {
		t.Errorf("Expected key2 and value2 at index %d, got %v", index2, ht.tabl[index2])
	}
	//похожее значение
	ht.AddHash("2key", "value3")

	index3 := ht.Hash("2key")

	if ht.tabl[index3].key != "2key" || ht.tabl[index3].value != "value3" {
		t.Errorf("Expected 2key and value3 at index %d, got %v", index2, ht.tabl[index2])
	}
}

func TestAddHash_Full(t *testing.T) {
	ht := NewHashTable(2) // Создаем хэш-таблицу размером 2

	// Добавляем два элемента, чтобы хэш-таблица заполнилась
	ht.AddHash("key1", "value1")
	ht.AddHash("key2", "value2")

	// Перехватываем вывод в stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Пытаемся добавить третий элемент, когда хэш-таблица уже полна
	ht.AddHash("key3", "value3")

	// Останавливаем перехват вывода
	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = old

	// Проверяем вывод
	output := buf.String()
	expected := "Хэш-таблица переполнена. Невозможно добавить новый элемент.\n"
	if output != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, output)
	}
}

// корректно находит и выводит значение по ключу
func TestKeyItem(t *testing.T) {
	ht := NewHashTable(10)
	ht.AddHash("key4", "value1")
	ht.AddHash("key2", "value2")
	ht.AddHash("ab", "value2")
	ht.AddHash("ba", "value2")
	ht.KeyItem("ab")
	ht.KeyItem("ba")
	ht.KeyItem("key4")
	ht.KeyItem("key2")
	ht.KeyItem("key3")
}

// корректно удаляет элементы по ключу
func TestDelValue(t *testing.T) {
	ht := NewHashTable(10)
	ht.AddHash("key1", "value1")
	ht.AddHash("key2", "value2")
	ht.AddHash("abc", "value2")
	ht.AddHash("bac", "value2")
	ht.AddHash("cba", "value2")

	ht.DelValue("cba")
	ht.DelValue("bac")
	ht.DelValue("abc")
	ht.DelValue("key1")
	ht.DelValue("key2")
	ht.DelValue("key3")

	//похожее значение
	ht.AddHash("2key", "value3")
	ht.DelValue("2key")

	index1 := ht.Hash("key1")
	index2 := ht.Hash("key2")
	index3 := ht.Hash("2key")
	if ht.tabl[index1] != nil {
		t.Errorf("Expected key1 to be deleted, got %v", ht.tabl[index1])
	}

	if ht.tabl[index2] != nil {
		t.Errorf("Expected key2 to be deleted, got %v", ht.tabl[index2])
	}

	if ht.tabl[index3] != nil {
		t.Errorf("Expected 2key to be deleted, got %v", ht.tabl[index2])
	}
}

// корректно определяет, заполнена ли хэш-таблица
func TestIsFull(t *testing.T) {
	ht := NewHashTable(2)
	if ht.IsFull() {
		t.Errorf("Expected hash table to be not full initially")
	}

	ht.AddHash("key1", "value1")
	if ht.IsFull() {
		t.Errorf("Expected hash table to be not full after adding one element")
	}

	ht.AddHash("key2", "value2")
	if !ht.IsFull() {
		t.Errorf("Expected hash table to be full after adding two elements")
	}
}
func TestPrint(t *testing.T) {
	// Создаем новую хэш-таблицу с размером 10
	ht := NewHashTable(10)

	// Добавляем элементы в хэш-таблицу
	ht.AddHash("key1", "value1")
	ht.AddHash("key2", "value2")

	// Вызываем метод Print, который будет записывать вывод в Pipe
	ht.Print()

}
