package serverbase

type ILogManager interface {

	SetLogLevel(level int)

	Println(logs string)
}
