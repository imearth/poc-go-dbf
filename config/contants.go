package config

import (
	"os"
)

const (
	// Project Name
	PROJECT_NAME = "TESTERP"

	// File Paths
	SOURCE_FILE_PATH             = "X:/Express1/" + PROJECT_NAME
	DESTINATION_BACKUP_FILE_PATH = "X:/BACKUP/Express/" + PROJECT_NAME

	// Customer - Address
	ARMAS_EXPRESS_FILE_PATH      = "X:/Express1/" + PROJECT_NAME + "/ARMAS"
	ARMAS_EXPRESS_TEMP_FILE_PATH = "X:/Express1/" + PROJECT_NAME + "/temp_ARMAS"
	ARMAS_BACKUP_FILE            = "ARMAS.DBF"

	// Sales Invoice
	ARTRN_EXPRESS_FILE_PATH      = "X:/Express1/" + PROJECT_NAME + "/ARTRN"
	ARTRN_EXPRESS_TEMP_FILE_PATH = "X:/Express1/" + PROJECT_NAME + "/temp_ARTRN"
	ARTRN_BACKUP_FILE            = "ARTRN.DBF"

	// Sales Invoice Service Item
	STCRD_EXPRESS_FILE_PATH      = "X:/Express1/" + PROJECT_NAME + "/STCRD"
	STCRD_EXPRESS_TEMP_FILE_PATH = "X:/Express1/" + PROJECT_NAME + "/temp_STCRD"
	STCRD_BACKUP_FILE            = "STCRD.DBF"

	// Job Master Data
	JBMAS_EXPRESS_FILE_PATH      = "X:/Express1/" + PROJECT_NAME + "/JBMAS"
	JBMAS_EXPRESS_TEMP_FILE_PATH = "X:/Express1/" + PROJECT_NAME + "/temp_JBMAS"
	JBMAS_BACKUP_FILE            = "JBMAS.DBF"

	// Bank Account
	BKTRN_EXPRESS_FILE_PATH      = "X:/Express1/" + PROJECT_NAME + "/BKTRN"
	BKTRN_EXPRESS_TEMP_FILE_PATH = "X:/Express1/" + PROJECT_NAME + "/temp_BKTRN"
	BKTRN_BACKUP_FILE            = "BKTRN.DBF"

	// Execute path for re-index
	RE_INDEX_EXEC_PATH = "X:/Express1/adm32"
	BASE_EXEC_PATH     = "X:/Express1"
)

// GetQueueConfigs loads the queue-related configurations from environment variables
func GetQueueConfigs() (string, string) {
	// Load from environment variables
	appENV := os.Getenv("APP_ENV")
	expressQueue := appENV + ".EXPRESS_QUEUE"
	expressToERP := appENV + ".EXPRESS_TO_ERP"

	return expressQueue, expressToERP
}
