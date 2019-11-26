package serverbase

type IConfManager interface {

	GetConf(key string) (string, error)

	JsonConfInit() error

}