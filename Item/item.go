package item

type Item struct {
	Id          int
	Title       string
	Description string
}

var todoList []Item

func Add(item Item) {
	todoList = append(todoList, item)
}

func Remove(item Item) {
	for i := 0; i < len(todoList); i++ {
		if todoList[i].Id == item.Id {
			firstLap := todoList[:i-1]
			secondLap := todoList[i+1:]
			var newTodoList []Item
			newTodoList = append(newTodoList, firstLap...)
			newTodoList = append(newTodoList, secondLap...)
			todoList = newTodoList
			break
		}
	}
}

func Edit(item Item) {
	for i := 0; i < len(todoList); i++ {
		if todoList[i].Id == item.Id {
			todoList[i] = item
			break
		}
	}
}

func GetList() []Item {
	return todoList
}
