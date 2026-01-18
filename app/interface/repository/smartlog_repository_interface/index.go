package smartlog_repository_interface

import "last-project/app/models"

type SmartLog_Repository_Interface interface {
	Create(smart *models.Smart_Log) error
}
