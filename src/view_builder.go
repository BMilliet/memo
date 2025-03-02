package src

type ViewBuilderInterface interface {
	NewTodoListView(op []*Todo, height int) []string
	NewListView(title string, op []ListItem, height int) ListItem
	NewTextFieldView(title, placeHolder string) string
	NewTextAreaFieldView(title, placeHolder string) string
}

type ViewBuilder struct{}

func NewViewBuilder() *ViewBuilder {
	return &ViewBuilder{}
}

func (b *ViewBuilder) NewListView(title string, op []ListItem, height int) ListItem {
	endValue := ListItem{}
	ListView(title, op, height, &endValue)
	return endValue
}

func (b *ViewBuilder) NewTodoListView(op []*Todo, height int) []string {
	endValue := []string{}
	TodoListView(op, height, &endValue)
	return endValue
}

func (b *ViewBuilder) NewTextFieldView(title, placeHolder string) string {
	endValue := ""
	TextFieldView(title, placeHolder, &endValue)
	return endValue
}

func (b *ViewBuilder) NewTextAreaFieldView(title, placeHolder string) string {
	endValue := ""
	TextAreaFieldView(title, placeHolder, &endValue)
	return endValue
}
