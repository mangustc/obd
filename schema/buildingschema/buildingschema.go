package buildingschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type BuildingDB struct {
	BuildingID       int    `json:"BuildingID" server:"y"`
	BuildingName     string `json:"BuildingName"`
	BuildingAddress  string `json:"BuildingAddress"`
	BuildingIsHidden bool   `json:"BuildingIsHidden"`
}

type BuildingInsert struct {
	BuildingName    string `json:"BuildingName"`
	BuildingAddress string `json:"BuildingAddress"`
}

type BuildingUpdate struct {
	BuildingID      int    `json:"BuildingID" server:"y"`
	BuildingName    string `json:"BuildingName"`
	BuildingAddress string `json:"BuildingAddress"`
}

type BuildingDelete struct {
	BuildingID int `json:"BuildingID" server:"y"`
}

type BuildingGet struct {
	BuildingID int `json:"BuildingID" server:"y"`
}

type BuildingsGet struct{}

func GetBuildingInputOptionsFromBuildingsDB(buildingsDB []*BuildingDB) []*schema.InputOption {
	notHiddenBuildingsDB := GetNotHiddenBuildingsDB(buildingsDB)
	inputOptions := []*schema.InputOption{}
	for _, buildingDB := range notHiddenBuildingsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s", buildingDB.BuildingName),
			InputOptionValue: fmt.Sprintf("%d", buildingDB.BuildingID),
		})
	}
	return inputOptions
}

func GetNotHiddenBuildingsDB(buildingsDB []*BuildingDB) []*BuildingDB {
	notHiddenBuildingsDB := []*BuildingDB{}
	for _, buildingDB := range buildingsDB {
		if !buildingDB.BuildingIsHidden {
			notHiddenBuildingsDB = append(notHiddenBuildingsDB, buildingDB)
		}
	}
	return notHiddenBuildingsDB
}

func ValidateBuildingDB(buildingDB *BuildingDB) (err error) {
	if buildingDB == nil {
		panic("Object is nil")
	}
	if buildingDB.BuildingID <= 0 || buildingDB.BuildingName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateBuildingInsert(buildingInsert *BuildingInsert) (err error) {
	if buildingInsert == nil {
		panic("Object is nil")
	}
	if buildingInsert.BuildingName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateBuildingUpdate(buildingUpdate *BuildingUpdate) (err error) {
	if buildingUpdate == nil {
		panic("Object is nil")
	}
	if buildingUpdate.BuildingID <= 0 || buildingUpdate.BuildingName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateBuildingDelete(buildingDelete *BuildingDelete) (err error) {
	if buildingDelete == nil {
		panic("Object is nil")
	}
	if buildingDelete.BuildingID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateBuildingGet(buildingGet *BuildingGet) (err error) {
	if buildingGet == nil {
		panic("Object is nil")
	}
	if buildingGet.BuildingID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateBuildingsGet(buildingGets *BuildingsGet) (err error) {
	if buildingGets == nil {
		panic("Object is nil")
	}
	return nil
}
