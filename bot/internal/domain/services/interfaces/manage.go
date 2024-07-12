package serviceInterfaces

import "leetcodebot/internal/domain/entities"

type ManageService interface {
	RemoveGroup(info entities.GroupInfo) (string, error)
	AddGroup(info entities.GroupInfo) (string, error)
	GetGroupStatus(info entities.GroupInfo) (string, error)
	GetAllGroups() ([]entities.Group, error)
}