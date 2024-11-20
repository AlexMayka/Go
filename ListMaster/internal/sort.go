package internal

// SortBy выполняет сортировку списка ListPerson на основе заданного атрибута.
//
// Параметры:
//   - attribute: Функция, которая принимает указатель на Person и возвращает значение атрибута для сравнения.
//   - reverse: Логическое значение. Если true, список сортируется в обратном порядке.
//
// Описание:
//
//	Сортировка выполняется методом пузырька, где сравнение элементов происходит
//	на основе значения, возвращаемого функцией attribute. Если reverse=true,
//	порядок сортировки изменяется на обратный.
func (l *ListPerson) SortBy(attribute func(p *Person) any, reverse bool) {
	for i := 0; i < l.Len-1; i++ {
		swapped := false
		current := l.FirstElem
		for current.NextPerson != nil {
			next := current.NextPerson

			if compare(attribute(current), attribute(next), reverse) {
				l.swapNodes(current, next)
				swapped = true
			}

			current = next
		}

		if !swapped {
			break
		}
	}
}

// Compare выполняет сравнение двух значений (a и b) на основе их типа.
//
// Параметры:
//   - a, b: Значения для сравнения. Должны быть одного типа (string или int).
//   - reverse: Если true, выполняется обратное сравнение (убывание).
//
// Возвращает:
//   - true, если значения должны быть переставлены.
//   - false, если значения находятся в правильном порядке.
func compare(a, b any, reverse bool) bool {
	switch va := a.(type) {
	case string:
		if reverse {
			return va < b.(string)
		}
		return va > b.(string)
	case int:
		if reverse {
			return va < b.(int)
		}
		return va > b.(int)
	default:
		return false
	}
}
