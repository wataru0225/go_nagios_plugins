package main

import (
	"github.com/wataru0225/go_nagios_plugins/lib/check_sqlserver"
	"github.com/wataru0225/go_nagios_plugins/utils"
)

// /path/to/sqlserver_check_by_query -u username -p password -h www.example.com -i instance -P port
func main() {
	database := checkmssql.Set()

	db := database.Connect()
	if db == nil {
		panic("failed")
	}

	execQuery := "SELECT COUNT(*) FROM [DB_NAME].[dbo].[TABLE_NAME]"

	rows, err := db.Query(execQuery)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			panic(err.Error())
		}
	}

	if count < 1 {
		msg := "CRITICAL ERROR"
		utils.Critical(msg).Exit()
	} else {
		msg := "OK STATUS"
		utils.Ok(msg).Exit()
	}
}
