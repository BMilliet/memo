package src

type ViewBuilderInterface interface {
	NewListView(title string, op []ListItem, height int, crud bool) ListItem
	NewTextFieldView(title, placeHolder string) TextReturnObject
	NewTextAreaFieldView(title, placeHolder string) TextReturnObject
}

type ViewBuilder struct{}

func NewViewBuilder() *ViewBuilder {
	return &ViewBuilder{}
}

func (b *ViewBuilder) NewListView(title string, op []ListItem, height int, crud bool) ListItem {
	endValue := ListItem{}
	ListView(title, op, height, &endValue, crud)
	return endValue
}

func (b *ViewBuilder) NewTextFieldView(title, placeHolder string) TextReturnObject {
	endValue := TextReturnObject{}
	TextFieldView(title, placeHolder, &endValue)
	return endValue
}

func (b *ViewBuilder) NewTextAreaFieldView(title, placeHolder string) TextReturnObject {
	endValue := TextReturnObject{}
	TextAreaFieldView(title, placeHolder, &endValue)
	return endValue
}
