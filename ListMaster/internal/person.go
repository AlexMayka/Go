package internal

import "fmt"

// Person представляет собой элемент двусвязного списка с информацией о человеке.
//
// Поля:
//   - FirstName: Имя человека.
//   - LastName: Фамилия человека.
//   - Age: Возраст человека.
//   - NextPerson: Указатель на следующий элемент в списке.
//   - LastPerson: Указатель на предыдущий элемент в списке.
type Person struct {
	FirstName  string
	LastName   string
	Age        int
	NextPerson *Person
	LastPerson *Person
}

// Options предоставляет параметры для обновления данных в объекте Person.
//
// Поля:
//   - FirstName: Новое имя человека (если нужно изменить).
//   - LastName: Новая фамилия человека (если нужно изменить).
//   - Age: Новый возраст человека (если нужно изменить).
type Options struct {
	FirstName string
	LastName  string
	Age       int
}

// Create создает новый объект Person.
//
// Параметры:
//   - firstName: Имя человека.
//   - lastName: Фамилия человека.
//   - age: Возраст человека.
//
// Возвращает:
//   - Указатель на созданный объект Person.
func Create(firstName string, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age, NextPerson: nil}
}

// UpdateData обновляет поля объекта Person на основе заданных опций.
//
// Параметры:
//   - opts: Объект Options с данными для обновления.
//
// Возвращает:
//   - Ошибку, если в опциях нет валидных полей для обновления.
func (p *Person) UpdateData(opts Options) error {
	if opts.FirstName == "" && opts.LastName == "" && opts.Age == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	if opts.FirstName != "" {
		p.FirstName = opts.FirstName
	}

	if opts.LastName != "" {
		p.LastName = opts.LastName
	}

	if opts.Age != 0 {
		p.Age = opts.Age
	}
	return nil
}

// String возвращает строковое представление объекта Person.
//
// Формат:
//
//	"<FirstName> <LastName> is <Age> years old".
//
// Пример:
//
//	Alex Mayka is 33 years old.
func (p Person) String() string {
	return fmt.Sprintf("%s %s is %d years old", p.FirstName, p.LastName, p.Age)
}

// GetValue возвращает значения всех полей объекта Person.
//
// Возвращает:
//   - Имя (FirstName).
//   - Фамилия (LastName).
//   - Возраст (Age).
func (p Person) GetValue() (string, string, int) {
	return p.FirstName, p.LastName, p.Age
}
