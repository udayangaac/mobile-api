package repositories

import "github.com/udayangaac/mobile-api/internal/entities"

type MobileAppUserRepo interface {
	AddMobileUser(mobileUser entities.MobileAppUser) (isUpdate bool)
	GetMobileUserByEmail(email string) (mobileUser entities.MobileAppUser, err error)
}
