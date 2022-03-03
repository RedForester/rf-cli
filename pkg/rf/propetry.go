package rf

type Property struct {
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Argument Argument `json:"argument"`
}

type Argument string

const (
	DatetimeDateArgument Argument = "DATETIME_DATE"
	NumberRealArgument   Argument = "NUMBER_REAL"
	TextSimpleArgument   Argument = "TEXT_SIMPLE"
)

type Category string

const (
	DateTimeCategory Category = "DATE_TIME"
	EnumCategory     Category = "ENUM"
	NumberCategory   Category = "NUMBER"
	TextCategory     Category = "TEXT"
	UserCategory     Category = "USER"
)
