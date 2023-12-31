package services

import (
	"errors"

	"main/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"main/app/models"
	"main/platform/amazon"
)

func PutItem(user *models.UserData) (string, error) {
	user.Password = utils.GeneratePassword(user.Password)
	err := amazon.PutItem(user)
	if err != nil {
		return "", err
	}
	return "user create Success", nil
}

func GetItem(user *models.UserData) (string, error) {
	item, err := amazon.GetItem(user)
	if err != nil {
		return "", err
	}
	if item["email"].(*types.AttributeValueMemberS).Value == user.Email &&
		utils.ComparePasswords(item["password"].(*types.AttributeValueMemberS).Value, user.Password) {
		return "user login Success", nil
	}
	return "password", errors.New("password is wrong")
}

func UpdateItem(alertSettings *models.AlertSetting) (string, error) {
	err := amazon.UpdateItem(alertSettings)
	if err != nil {
		return "", err
	}
	err = DeleteAlertMessages(alertSettings.Email)
	if err != nil {
		return "", err
	}
	return "alert setting update Success", nil
}
