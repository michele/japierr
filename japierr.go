import "fmt"

type Error struct {
	Status string `json:"status"`
	Source Source `json:"source"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type Source interface{}

type Pointer struct {
	Pointer string `json:"pointer"`
}

type Parameter struct {
	Parameter string `json:"parameter"`
}

type Errors struct {
	Errors []*Error `json:"errors"`
}

func New() Errors {
	return Errors{
		Errors: []*Error{},
	}
}

func (errs *Errors) Append(e *Error) {
	errs.Errors = append(err.Errors, e)
}

func (errs Errors) IsEmpty() bool {
	return len(errs.Errors) == 0
}

func (errs Errors) HasErrors() bool {
	return len(errs.Errors) > 0
}

func NewInvalidAttribute(attr, detail string) *Error {
	return &Error{
		Status: "422",
		Source: Pointer{
			Pointer: fmt.Sprintf("data/attributes/%s", attr),
		},
		Title:  "Invalid Attribute",
		Detail: detail,
	}
}

