package services

import "92HW/utils"

func AddOrigin(userID, origin string) error {
	return utils.AddTrustedOrigin(userID, origin)
}

func RemoveOrigin(userID, origin string) error {
	return utils.RemoveTrustedOrigin(userID, origin)
}

func GetOrigins(userID string) ([]string, error) {
	return utils.GetTrustedOrigins(userID)
}
