package todo

type TodoTitleError struct{}

func (TodoTitleError) Error() string {
	return "title can't be empty"
}
