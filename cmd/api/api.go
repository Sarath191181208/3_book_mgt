package api

import (
	"sarath/3_book_mgt/internal/data"
	"sarath/3_book_mgt/internal/logger"
)

type Application struct{
  Logger logger.Logger
  Db *data.Models
}
