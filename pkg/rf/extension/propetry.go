package extension

type Property struct {
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Argument Argument `json:"argument"`
}

type Argument string

const (
	DatetimeDate Argument = "DATETIME_DATE"
	NumberReal   Argument = "NUMBER_REAL"
	TextSimple   Argument = "TEXT_SIMPLE"
)

type Category string

const (
	DateTime Category = "DATE_TIME"
	Enum     Category = "ENUM"
	Number   Category = "NUMBER"
	Text     Category = "TEXT"
	User     Category = "USER"
)
