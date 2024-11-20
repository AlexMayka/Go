package tests

import (
	"ListMaster/internal"
	"testing"
)

func getTestData() []struct {
	firstName string
	lastName  string
	age       int
} {
	return []struct {
		firstName string
		lastName  string
		age       int
	}{
		{"Alex", "Mayka", 24},
		{"Ivan", "Romanov", 42},
		{"Vladimir", "Жуков", 32},
		{"Ваня", "Ivanov", 65},
		{"Евгений", "Песков", 12},
	}
}

// Проверяет добавление элементов в список ListPerson.
func TestAddList(t *testing.T) {
	tests := getTestData()

	t.Run("Add elements to list", func(t *testing.T) {
		testList := internal.ListPerson{}
		for _, tt := range tests {
			testList.AddPerson(tt.firstName, tt.lastName, tt.age)
		}

		if testList.Len != len(tests) {
			t.Fatalf("AddList failed: expected length %d, got %d", len(tests), testList.Len)
		}

		t.Log("AddList: success")
	})
}

// Тестирует метод GetByIndex, проверяя следующие сценарии:
//   - Получение корректных элементов по индексам.
//   - Получение последнего элемента с индексом -1.
//   - Обработка выхода за границы индекса.
func TestGetByIndex(t *testing.T) {
	tests := getTestData()

	t.Run("Get valid indices", func(t *testing.T) {
		testList := internal.ListPerson{}
		for _, tt := range tests {
			testList.AddPerson(tt.firstName, tt.lastName, tt.age)
		}

		for index, value := range tests {
			test, err := testList.GetByIndex(index)
			if err != nil {
				t.Fatalf("GetByIndex(%d) failed with error: %v", index, err)
			}
			if firstName, lastName, age := test.GetValue(); firstName != value.firstName || lastName != value.lastName || age != value.age {
				t.Fatalf("GetByIndex(%d) returned wrong values: got (%s, %s, %d), want (%s, %s, %d)",
					index, firstName, lastName, age, value.firstName, value.lastName, value.age)
			}
		}

		t.Log("Get valid indices: success")
	})

	t.Run("Get last element (-1)", func(t *testing.T) {
		testList := internal.ListPerson{}
		for _, tt := range tests {
			testList.AddPerson(tt.firstName, tt.lastName, tt.age)
		}

		last, err := testList.GetByIndex(-1)
		if err != nil {
			t.Fatalf("GetByIndex(-1) failed with error: %v", err)
		}
		if firstName, lastName, age := last.GetValue(); firstName != tests[len(tests)-1].firstName || lastName != tests[len(tests)-1].lastName || age != tests[len(tests)-1].age {
			t.Fatalf("GetByIndex(-1) returned wrong values: got (%s, %s, %d), want (%s, %s, %d)",
				firstName, lastName, age, tests[len(tests)-1].firstName, tests[len(tests)-1].lastName, tests[len(tests)-1].age)
		}

		t.Log("Get last element (-1): success")
	})

	t.Run("Out of range index", func(t *testing.T) {
		testList := internal.ListPerson{}
		for _, tt := range tests {
			testList.AddPerson(tt.firstName, tt.lastName, tt.age)
		}

		if _, err := testList.GetByIndex(len(tests)); err == nil {
			t.Fatalf("GetByIndex(out of range) did not return error")
		}

		t.Log("Out of range index: success")
	})
}

// Вспомогательная функция для проверки содержимого списка.
// Сравнивает фактические данные в списке с ожидаемыми.
func verifyList(t *testing.T, list *internal.ListPerson, expected []struct {
	firstName string
	lastName  string
	age       int
}) {
	current := list.FirstElem
	for i, expectedPerson := range expected {
		if current == nil {
			t.Fatalf("List ended prematurely at index %d", i)
		}
		firstName, lastName, age := current.GetValue()
		if firstName != expectedPerson.firstName || lastName != expectedPerson.lastName || age != expectedPerson.age {
			t.Fatalf("Mismatch at index %d: got (%s, %s, %d), want (%s, %s, %d)",
				i, firstName, lastName, age, expectedPerson.firstName, expectedPerson.lastName, expectedPerson.age)
		}
		current = current.NextPerson
	}

	if current != nil {
		t.Fatalf("List contains extra elements beyond expected length")
	}
}

// Проверяет обработку пустого списка:
//   - Сортировка пустого списка не должна изменять его состояние.
//   - Попытка вызвать методы Swap или GetByIndex должна возвращать ошибки.
func TestSortBy(t *testing.T) {
	tests := []struct {
		firstName string
		lastName  string
		age       int
	}{
		{"Alex", "Mayka", 33},
		{"Ivan", "Ivanov", 42},
		{"John", "Doe", 33},
		{"Sasha", "Ivanova", 23},
		{"Alex", "Ivanov", 23},
	}

	expected := []struct {
		firstName string
		lastName  string
		age       int
	}{
		{"Alex", "Ivanov", 23},
		{"Alex", "Mayka", 33},
		{"Ivan", "Ivanov", 42},
		{"John", "Doe", 33},
		{"Sasha", "Ivanova", 23},
	}

	list := internal.ListPerson{}
	for _, tt := range tests {
		list.AddPerson(tt.firstName, tt.lastName, tt.age)
	}

	list.SortBy(func(p *internal.Person) any {
		return p.FirstName + " " + p.LastName
	}, false)

	current := list.FirstElem
	for i, exp := range expected {
		if current == nil {
			t.Fatalf("List ended prematurely at index %d", i)
		}

		firstName, lastName, age := current.GetValue()
		if firstName != exp.firstName || lastName != exp.lastName || age != exp.age {
			t.Fatalf("Sort mismatch at index %d: got (%s, %s, %d), want (%s, %s, %d)",
				i, firstName, lastName, age, exp.firstName, exp.lastName, exp.age)
		}

		current = current.NextPerson
	}

	if current != nil {
		t.Fatalf("List contains extra elements after sorting")
	}

	t.Log("SortBy: success")
}

// Проверяет обработку пустого списка:
//   - Сортировка пустого списка не должна изменять его состояние.
//   - Попытка вызвать методы Swap или GetByIndex должна возвращать ошибки.
func TestEmptyList(t *testing.T) {
	emptyList := internal.ListPerson{}

	t.Run("SortBy on empty list", func(t *testing.T) {
		emptyList.SortBy(func(p *internal.Person) any {
			return p.FirstName
		}, false)

		if emptyList.FirstElem != nil {
			t.Fatal("SortBy failed: empty list should remain empty")
		}
		t.Log("SortBy on empty list: success")
	})

	t.Run("Swap on empty list", func(t *testing.T) {
		err := emptyList.Swap(0, 1)
		if err == nil {
			t.Fatal("Swap on empty list did not return error")
		}
		t.Log("Swap on empty list: success")
	})

	t.Run("GetByIndex on empty list", func(t *testing.T) {
		_, err := emptyList.GetByIndex(0)
		if err == nil {
			t.Fatal("GetByIndex on empty list did not return error")
		}
		t.Log("GetByIndex on empty list: success")
	})
}

// Тестирует метод Swap, проверяя:
//   - Перестановку элементов в различных позициях списка.
//   - Проверку корректности после перестановки.
func TestSwap(t *testing.T) {
	tests := getTestData()
	testList := internal.ListPerson{}
	for _, tt := range tests {
		testList.AddPerson(tt.firstName, tt.lastName, tt.age)
	}

	t.Run("Swap elem in list", func(t *testing.T) {
		for i := 0; i < testList.Len; i++ {
			for j := 0; j < testList.Len; j++ {
				current := internal.ListPerson{}
				for _, tt := range tests {
					current.AddPerson(tt.firstName, tt.lastName, tt.age)
				}

				err := current.Swap(i, j)
				if err != nil {
					t.Fatalf("Swap(%d, %d) failed with error: %v", i, j, err)
				}

				expected := make([]struct {
					firstName string
					lastName  string
					age       int
				}, len(tests))
				copy(expected, tests)
				expected[i], expected[j] = expected[j], expected[i] // Меняем местами i и j

				verifyList(t, &current, expected)

			}
		}
	})
}
