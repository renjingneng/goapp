package mysql

import (
	"github.com/renjingneng/goapp/core/database"
)

type Open struct {
	*database.MysqlBase
}

func NewOpen() *Open {
	return &Open{
		MysqlBase: database.NewMysqlBase("MysqlOpen"),
	}
}

func (open *Open) FetchRowNew() {

}
