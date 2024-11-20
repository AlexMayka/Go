package internal

import (
	"fmt"
)

// ListPerson представляет собой двусвязный список, содержащий объекты типа Person.
// Он включает ссылки на первый и последний элементы списка и хранит общую длину списка.
// Поля:
//   - FirstElem: Указатель на следующий элемент в списке.
//   - LastElem: Указатель на предыдущий элемент в списке.
//   - Len: Количество элементов в списке.
type ListPerson struct {
	FirstElem *Person
	LastElem  *Person
	Len       int
}

// AddPerson добавляет нового человека в конец списка.
//
// Параметры:
//   - firstName: имя человека.
//   - lastName: фамилия человека.
//   - age: возраст человека.
func (l *ListPerson) AddPerson(firstName string, lastName string, age int) {

	newPerson := Create(firstName, lastName, age)

	if l.FirstElem == nil {
		l.FirstElem = newPerson
		l.LastElem = newPerson
		newPerson.NextPerson = nil
		newPerson.LastPerson = nil

		l.Len = 1
		return
	}

	newPerson.LastPerson = l.LastElem
	l.LastElem.NextPerson = newPerson
	l.LastElem = newPerson
	newPerson.NextPerson = nil

	l.Len += 1
}

// GetByIndex возвращает элемент списка по индексу.
//
// Параметры:
//   - index: индекс элемента (от 0 до len-1). Для последнего элемента используйте -1.
//
// Возвращает:
//   - Указатель на объект Person.
//   - Ошибку, если индекс выходит за пределы списка.
func (l *ListPerson) GetByIndex(index int) (*Person, error) {
	if index < -1 || index >= l.Len {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}

	if index == -1 {
		return l.LastElem, nil
	}

	if index <= l.Len/2 {
		current := l.FirstElem
		for i := 0; i < index; i++ {
			current = current.NextPerson
		}
		return current, nil
	} else {
		current := l.LastElem
		for i := l.Len - 1; i > index; i-- {
			current = current.LastPerson
		}
		return current, nil
	}
}

// Filter фильтрует элементы списка на основе заданных критериев.
//
// Параметры:
//   - criteria: объект Options, содержащий фильтры (имя, фамилия, возраст).
//
// Возвращает:
//   - Новый список ListPerson, содержащий отфильтрованные элементы.
//   - Ошибку, если фильтрация невозможна.
func (l *ListPerson) Filter(criteria Options) (ListPerson, error) {
	result := ListPerson{}
	for elem := l.FirstElem; elem != nil; elem = elem.NextPerson {
		if (criteria.FirstName == "" || criteria.FirstName == elem.FirstName) &&
			(criteria.LastName == "" || criteria.LastName == elem.LastName) &&
			(criteria.Age == 0 || criteria.Age == elem.Age) {
			result.AddPerson(elem.FirstName, elem.LastName, elem.Age)
		}
	}
	return result, nil
}

// DeletePerson удаляет элемент из списка по индексу.
//
// Параметры:
//   - index: индекс элемента для удаления.
//
// Возвращает:
//   - Ошибку, если индекс выходит за пределы списка.
func (l *ListPerson) DeletePerson(index int) error {

	elemByIndex, err := l.GetByIndex(index)
	if err != nil {
		return err
	}

	if elemByIndex.LastPerson != nil {
		elemByIndex.LastPerson.NextPerson = elemByIndex.NextPerson
	} else {
		l.FirstElem = elemByIndex.NextPerson
	}

	if elemByIndex.NextPerson != nil {
		elemByIndex.NextPerson.LastPerson = elemByIndex.LastPerson
	} else {
		l.LastElem = elemByIndex.LastPerson
	}

	l.Len--
	return nil
}

// Swap меняет местами два элемента списка по индексам.
//
// Параметры:
//   - firstIndex: индекс первого элемента.
//   - secondIndex: индекс второго элемента.
//
// Возвращает:
//   - Ошибку, если один из индексов выходит за пределы списка.
func (l *ListPerson) Swap(firstIndex int, secondIndex int) error {
	if firstIndex == secondIndex {
		return nil
	}

	FirstElem, err := l.GetByIndex(firstIndex)
	if err != nil {
		return err
	}

	secondElem, err := l.GetByIndex(secondIndex)
	if err != nil {
		return err
	}

	l.swapNodes(FirstElem, secondElem)

	return nil
}

// swapNodes обменивает местами два узла списка.
//
// Параметры:
//   - firstNode: указатель на первый узел.
//   - secondNode: указатель на второй узел.
func (l *ListPerson) swapNodes(firstNode, secondNode *Person) {
	if firstNode.NextPerson == secondNode {
		if firstNode.LastPerson != nil {
			firstNode.LastPerson.NextPerson = secondNode
		} else {
			l.FirstElem = secondNode
		}
		if secondNode.NextPerson != nil {
			secondNode.NextPerson.LastPerson = firstNode
		} else {
			l.LastElem = firstNode
		}
		secondNode.LastPerson = firstNode.LastPerson
		firstNode.NextPerson = secondNode.NextPerson
		firstNode.LastPerson = secondNode
		secondNode.NextPerson = firstNode
		return
	}

	if firstNode.LastPerson != nil {
		firstNode.LastPerson.NextPerson = secondNode
	} else {
		l.FirstElem = secondNode
	}
	if firstNode.NextPerson != nil {
		firstNode.NextPerson.LastPerson = secondNode
	}

	if secondNode.LastPerson != nil {
		secondNode.LastPerson.NextPerson = firstNode
	} else {
		l.FirstElem = firstNode
	}
	if secondNode.NextPerson != nil {
		secondNode.NextPerson.LastPerson = firstNode
	}

	firstNode.LastPerson, secondNode.LastPerson = secondNode.LastPerson, firstNode.LastPerson
	firstNode.NextPerson, secondNode.NextPerson = secondNode.NextPerson, firstNode.NextPerson

}

// Clear очищает список, удаляя все элементы.
func (l *ListPerson) Clear() {
	l.FirstElem = nil
	l.LastElem = nil
	l.Len = 0
}

// PrintList выводит все элементы списка в консоль.
func (l *ListPerson) PrintList() {
	current := l.FirstElem
	for current != nil {
		fmt.Printf("%s %s, %d\n", current.FirstName, current.LastName, current.Age)
		current = current.NextPerson
	}
}
