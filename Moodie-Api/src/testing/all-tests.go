package testing

import (
	"moody/moody-api/src/literals"
	"moody/moody-api/src/models"
	"moody/moody-api/src/service"
)

func TestAll() {
	literals.RedLog("All Test Are Bing Performed")

}

// test cases
// fat tests first , ie , overall functions then indicidual functions
// create user test
func CreatUserTest() bool {
	// case 1 true value
	literals.BlueLog("create-user-test,case:true value ")
	trueUserVlaue := models.UserRequestModel{
		Email: "mianrayan211@gmail.com",
	}
	service.CreateNewUser(trueUserVlaue)
	return true
}
