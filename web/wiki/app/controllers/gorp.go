package controllers
import (
	"github.com/go/src/fmt"
	"github.com/revel/modules/db/app"
	"github.com/go-gorp/gorp"
	"github.com/subh007/goodl/web/wiki/app/models"
	_ "github.com/mattn/go-sqlite3"
	r "github.com/revel/revel"
	"database/sql"
)

var (
	Dbm *gorp.DbMap
)


func InitDB() {
	fmt.Print("invoked the init db")
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	tableMap := Dbm.AddTable(models.WikiUser{}).SetKeys(true, "UserId")
	setColumnSizes(tableMap, map[string]int{
		"UserName": 20,
		"Password": 20,
	})
	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}