package groupschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type GroupDB struct {
	GroupID         int    `json:"GroupID"`
	GroupNumber     string `json:"GroupNumber"`
	GroupYear       int    `json:"GroupYear"`
	GroupCourseName string `json:"GroupCourseName"`
	GroupIsHidden   bool   `json:"GroupIsHidden"`
}

type GroupInsert struct {
	GroupNumber     string `json:"GroupNumber"`
	GroupYear       int    `json:"GroupYear"`
	GroupCourseName string `json:"GroupCourseName"`
}

type GroupUpdate struct {
	GroupID         int    `json:"GroupID"`
	GroupNumber     string `json:"GroupNumber"`
	GroupYear       int    `json:"GroupYear"`
	GroupCourseName string `json:"GroupCourseName"`
}

type GroupDelete struct {
	GroupID int `json:"GroupID"`
}

type GroupGet struct {
	GroupID int `json:"GroupID"`
}

type GroupsGet struct{}

func GetGroupInputOptionsFromGroupsDB(groupsDB []*GroupDB) []*schema.InputOption {
	notHiddenGroupsDB := GetNotHiddenGroupsDB(groupsDB)
	inputOptions := []*schema.InputOption{}
	for _, groupDB := range notHiddenGroupsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s %d", groupDB.GroupNumber, groupDB.GroupYear),
			InputOptionValue: fmt.Sprintf("%d", groupDB.GroupID),
		})
	}
	return inputOptions
}

func GetNotHiddenGroupsDB(groupsDB []*GroupDB) []*GroupDB {
	notHiddenGroupsDB := []*GroupDB{}
	for _, groupDB := range groupsDB {
		if !groupDB.GroupIsHidden {
			notHiddenGroupsDB = append(notHiddenGroupsDB, groupDB)
		}
	}
	return notHiddenGroupsDB
}

func ValidateGroupDB(groupDB *GroupDB) (err error) {
	if groupDB == nil {
		panic("Object is nil")
	}
	if groupDB.GroupID <= 0 || groupDB.GroupNumber == "" || groupDB.GroupYear == 0 ||
		groupDB.GroupCourseName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateGroupInsert(groupInsert *GroupInsert) (err error) {
	if groupInsert == nil {
		panic("Object is nil")
	}
	if groupInsert.GroupNumber == "" || groupInsert.GroupYear == 0 ||
		groupInsert.GroupCourseName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateGroupUpdate(groupUpdate *GroupUpdate) (err error) {
	if groupUpdate == nil {
		panic("Object is nil")
	}
	if groupUpdate.GroupID == 0 || groupUpdate.GroupNumber == "" || groupUpdate.GroupYear == 0 ||
		groupUpdate.GroupCourseName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateGroupDelete(groupDelete *GroupDelete) (err error) {
	if groupDelete == nil {
		panic("Object is nil")
	}
	if groupDelete.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateGroupGet(groupGet *GroupGet) (err error) {
	if groupGet == nil {
		panic("Object is nil")
	}
	if groupGet.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateGroupsGet(groupGets *GroupsGet) (err error) {
	if groupGets == nil {
		panic("Object is nil")
	}
	return nil
}
