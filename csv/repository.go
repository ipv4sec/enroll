package csv

import "enroll/mysql"

func Save(csv *Csv) (int64, int64) {
	return mysql.Clinet.Save(&csv).RowsAffected, csv.Id
}