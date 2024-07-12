package manageService

import (
	"fmt"
	"leetcodebot/internal/domain/repositories"
	"leetcodebot/internal/domain/entities"
	serviceInterfaces "leetcodebot/internal/domain/services/interfaces"
)

type ManageService struct {
	rep repositoryInterfaces.GroupRepository
}

func GetManageService(rep repositoryInterfaces.GroupRepository) serviceInterfaces.ManageService {
	return ManageService{rep: rep}
}

func (s ManageService) AddGroup(group entities.GroupInfo) (string, error) {
	inDB, err := s.rep.Find(entities.GroupQuery{
		ChatId: group.ChatId,
		ThreadId: group.ThreadId,
	})
	if err != nil {
		return "Произошла ошибка при добавлении группы в базу данных", err
	}
	if inDB {
		return "Группа уже есть в рассылке", nil
	}
	s.rep.Add(group)
	return fmt.Sprintf("Группа %v добавлена в рассылку", group.Title), err
}

func (s ManageService) GetGroupStatus(group entities.GroupInfo) (string, error) {
	inDB, err := s.rep.Find(entities.GroupQuery{
		ChatId: group.ChatId,
		ThreadId: group.ThreadId,
	})
	if err != nil {
		return "Произошла ошибка при получении статуса группы", err
	}
	if inDB {
		return "Группа уже есть в рассылке", nil
	}
	return "Группы еще нет в рассылке", nil
}

func (s ManageService) RemoveGroup(group entities.GroupInfo) (string, error) {
	inDB, err := s.rep.Find(entities.GroupQuery{
		ChatId: group.ChatId,
		ThreadId: group.ThreadId,
	})
	if err != nil {
		return "Произошла ошибка при удалении группы из базы данных", err
	}
	if !inDB {
		return "Группы еще нет в рассылке", nil
	}
	s.rep.Delete(entities.GroupQuery{
		ChatId: group.ChatId,
		ThreadId: group.ThreadId,
	})
	return fmt.Sprintf("Группа %v удалена из рассылки", group.Title), err
}

func (s ManageService) GetAllGroups() ([]entities.Group, error) {
	list, err := s.rep.GetAll()
	if err != nil {
		return nil, err
	}
	return list, err
}