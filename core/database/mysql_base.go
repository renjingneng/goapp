// @Description  基于gorm的mysql
// @Author  renjingneng
// @CreateTime  2020/8/18 10:49
package database

import (
	"database/sql"
	"strings"

	"gorm.io/gorm"
)

type MysqlBase struct {
	Tablename string
	//模式 normal/master，normal模式下读写库分离，master模式下读写均在写库
	Mode   string
	Dbname string
	DbptrW *gorm.DB
	DbptrR *gorm.DB
}

func NewMysqlBase(Dbname string) *MysqlBase {
	return &MysqlBase{
		Dbname: Dbname,
		DbptrW: GetEntityFromMysqlContainer(Dbname, "W"),
		DbptrR: GetEntityFromMysqlContainer(Dbname, "R"),
	}
}

func NewMysqlBaseWithTablename(Dbname string, Tablename string) *MysqlBase {
	return &MysqlBase{
		Dbname:    Dbname,
		Tablename: Tablename,
		Mode:      "normal",
		DbptrW:    GetEntityFromMysqlContainer(Dbname, "W"),
		DbptrR:    GetEntityFromMysqlContainer(Dbname, "R"),
	}
}

func (b *MysqlBase) SetTablename(tablename string) {
	b.Tablename = tablename
}
func (b *MysqlBase) SetMode(mode string) {
	b.Mode = mode
}

func (b *MysqlBase) FetchRowShort(condition map[string]string) map[string]interface{} {
	querySQL, values := b.buildQuerySQL("*", condition, map[string]string{"limit": "0,1", "order": "id desc"})
	db := b.chooseDb(1)
	rows, err := db.Raw(querySQL, values...).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	result := b.fetchResult(rows)
	if len(result) > 0 {
		return result[0]
	} else {
		return nil
	}

}

func (b *MysqlBase) FetchRowLong(fields string, condition map[string]string, other map[string]string) map[string]interface{} {
	other["limit"] = "0,1"
	querySQL, values := b.buildQuerySQL(fields, condition, other)
	db := b.chooseDb(1)
	rows, err := db.Raw(querySQL, values...).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	result := b.fetchResult(rows)
	if len(result) > 0 {
		return result[0]
	} else {
		return nil
	}
}

func (b *MysqlBase) FetchRowsShort(condition map[string]string) []map[string]interface{} {
	querySQL, values := b.buildQuerySQL("*", condition, map[string]string{"limit": "0,1", "order": "id desc"})
	db := b.chooseDb(1)
	rows, err := db.Raw(querySQL, values...).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	result := b.fetchResult(rows)
	return result
}

func (b *MysqlBase) FetchRowsLong(fields string, condition map[string]string, other map[string]string) []map[string]interface{} {
	querySQL, values := b.buildQuerySQL(fields, condition, other)
	db := b.chooseDb(1)
	rows, err := db.Raw(querySQL, values...).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	result := b.fetchResult(rows)
	return result
}

func (b *MysqlBase) FetchRowsBySql(sql string) []map[string]interface{} {
	db := b.chooseDb(1)
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		return nil
	}
	defer rows.Close()
	result := b.fetchResult(rows)
	return result
}

// Insert 插入数据并返回RowsAffected
//
// @Author  renjingneng
//
// @CreateTime  2020/8/21 19:59
func (b *MysqlBase) Insert(data map[string]interface{}) (int64, error) {
	var keys, placeholder = "", ""
	var values []interface{}
	for k, v := range data {
		keys += "," + k
		placeholder += ",?"
		values = append(values, v)
	}
	keys = strings.TrimLeft(keys, ",")
	placeholder = strings.TrimLeft(placeholder, ",")
	insertSQL := "INSERT INTO " + b.Tablename + "(" + keys + ") values(" + placeholder + ")"
	db := b.chooseDb(0)
	//err:=db.Exec(insertSQL,values...).Error
	res := db.Exec(insertSQL, values...)
	return res.RowsAffected, res.Error
}

// Update 更新数据并返回RowsAffected
//
// @Author  renjingneng
//
// @CreateTime  2020/8/21 19:59
func (b *MysqlBase) Update(data map[string]interface{}, condition map[string]string) (int64, error) {
	var placeholder = ""
	var values []interface{}
	for k, v := range data {
		placeholder += "," + k + "=?"
		values = append(values, v)
	}
	placeholder = strings.TrimLeft(placeholder, ",")
	where, vals := b.buildCondition(condition)
	updateSQL := "UPDATE " + b.Tablename + " SET " + placeholder + " WHERE " + where
	db := b.chooseDb(0)
	values = append(values, vals...)
	res := db.Exec(updateSQL, values...)
	return res.RowsAffected, res.Error
}

func (b *MysqlBase) chooseDb(isRead int) *gorm.DB {
	if b.Mode == "normal" {
		if isRead == 1 {
			return b.DbptrR
		} else {
			return b.DbptrW
		}
	} else {
		return b.DbptrW
	}
}

func (b *MysqlBase) buildQuerySQL(fields string, condition map[string]string, other map[string]string) (string, []interface{}) {
	if fields == "" {
		fields = "*"
	}
	if _, ok := other["order"]; !ok {
		other["order"] = ""
	} else {
		other["order"] = " ORDER BY " + other["order"]
	}
	if _, ok := other["group"]; !ok {
		other["group"] = ""
	} else {
		other["group"] = " GROUP BY " + other["group"]
	}
	if _, ok := other["limit"]; !ok {
		other["limit"] = ""
	} else {
		other["limit"] = " LIMIT " + other["limit"]
	}
	var where, values = b.buildCondition(condition)

	querySQL := " SELECT " + fields + " FROM " + b.Tablename + " WHERE " + where + other["group"] + other["order"] + other["limit"]
	return querySQL, values
}

func (b *MysqlBase) buildCondition(condition map[string]string) (string, []interface{}) {
	var where = " "
	var values []interface{}
	var i = 1
	for k, v := range condition {
		if i == 1 {
			where += k + "=" + "?"
		} else {
			where += " AND " + k + "=" + "?"
		}
		values = append(values, v)
		i++
	}
	return where, values
}

func (b *MysqlBase) fetchResult(rows *sql.Rows) []map[string]interface{} {
	var result []map[string]interface{}
	//获取记录列
	if columns, err := rows.Columns(); err != nil {
		return nil
	} else {
		//拼接记录Map
		values := make([]sql.RawBytes, len(columns))
		scans := make([]interface{}, len(columns))
		for i := range values {
			scans[i] = &values[i]
		}
		for rows.Next() {
			_ = rows.Scan(scans...)
			each := map[string]interface{}{}
			for i, col := range values {
				each[columns[i]] = string(col)
			}
			result = append(result, each)
		}
		if err := rows.Err(); err != nil {
			return nil
		}

	}
	return result
}
