package repositories

import (
	"context"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/entities"
	"github.com/udayangaac/mobile-api/internal/errors_custom"
	"github.com/udayangaac/mobile-api/internal/lib/orm"
)

type mobileAppUserMySqlRepo struct {
	DB *gorm.DB
}

func NewMobileAppUser() MobileAppUserRepo {
	return &mobileAppUserMySqlRepo{
		DB: orm.DB,
	}
}

func (m mobileAppUserMySqlRepo) AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser, mobileUserConfiguration entities.MobileUserConfiguration) (user entities.MobileAppUser, err error) {
	// log.Info(log_traceable.GetMessage(ctx, fmt.Sprintf("%v", mobileUser)))
	rowAffected := m.DB.Create(&mobileUser).RowsAffected

	if rowAffected == 0 {
		err = errors_custom.ErrDuplicateUserEntry
		return
	}
	mobileUserConfiguration.UserId = mobileUser.ID
	err = m.DB.Create(&mobileUserConfiguration).Error
	err = m.DB.Where("id=?", mobileUser.ID).First(&user).Error
	return
}

func (m mobileAppUserMySqlRepo) GetMobileUserByEmail(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error) {
	err = m.DB.Where("email=?", email).First(&mobileUser).Error
	return
}

func (m mobileAppUserMySqlRepo) UserLogout(ctx context.Context) (err error) {
	return
}

func (m mobileAppUserMySqlRepo) PushNotification(ctx context.Context, userId int, lat float64, lon float64) (PushNotification entities.Notification, err error) {
	err = m.DB.Where("id=?", userId).First(&PushNotification).Error
	return
}

func (m mobileAppUserMySqlRepo) PullNotification(ctx context.Context, userId int, lat float64, lon float64) (Notification entities.Notification, err error) {
	err = m.DB.Where("id=?", userId).First(&Notification).Error
	return

}

func (m mobileAppUserMySqlRepo) LocationTrack(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileAppUser{}).Where("id = ?", userId).Update("location_service_status", status).Error
	return

}

func (m mobileAppUserMySqlRepo) UserProfilePicture(ctx context.Context, userId int16) (mobileUser entities.MobileAppUser, err error) {
	err = m.DB.Where("email=?").First(&mobileUser).Error
	return

}

func (m mobileAppUserMySqlRepo) SoundSettingChange(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileAppUser{}).Where("id = ?", userId).Update("sound_status", status).Error
	log.Info(err)
	return
}

func (m mobileAppUserMySqlRepo) PushNotificationSetting(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileAppUser{}).Where("id = ?", userId).Update("push_notification_status", status).Error
	log.Info(err)
	return
}

func (m mobileAppUserMySqlRepo) SetLoginStatus(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileAppUser{}).Where("id = ?", userId).Update("login_status", status).Error
	return
}

func (m mobileAppUserMySqlRepo) NotificationTypesList(ctx context.Context, userId int) (notificationTypes interface{}, err error) {
	log.Info(userId)
	nts := []entities.AdvertisementsCategories{}

	if userId == 0 {
		err = m.DB.Select([]string{"id", "category_name"}).Where("status=1").Find(&nts).Error
		return nts, err
	} else {
		rows, err := m.DB.Raw("SELECT ac.id, ac.category_name FROM advertisements_categories ac INNER JOIN user_advertisement_categories uac on ac.id = uac.advertisement_cat_id WHERE uac.user_id = ?", userId).Rows()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			nt := entities.AdvertisementsCategories{}
			rows.Scan(&nt.ID, &nt.CategoryName)
			nts = append(nts, nt)
		}

	}
	return nts, nil
}

func (m mobileAppUserMySqlRepo) BankList(ctx context.Context, userId int) (notificationTypes interface{}, err error) {
	log.Info(userId)
	ub := []entities.Banks{}

	if userId == 0 {
		err = m.DB.Select([]string{"id", "name"}).Where("status=1").Find(&ub).Error
		return ub, err
	} else {
		rows, err := m.DB.Raw("SELECT ub.id, ub.name FROM banks ub INNER JOIN mobile_user_bank mub on ub.id = mub.bank_id WHERE mub.mobile_user_id = ?", userId).Rows()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			nt := entities.Banks{}
			rows.Scan(&nt.ID, &nt.Name)
			ub = append(ub, nt)
		}

	}
	return ub, nil
}

func (m mobileAppUserMySqlRepo) GetUserProfile(ctx context.Context, userId int) (userProfile entities.MobileAppUser, err error) {
	err = m.DB.Where("id=?", userId).First(&userProfile).Error
	if err != nil {
		return userProfile, err
	}
	rows, err := m.DB.Raw("SELECT ac.id, ac.category_name FROM advertisements_categories ac INNER JOIN user_advertisement_categories uac on ac.id = uac.advertisement_cat_id WHERE uac.deleted_at is null and  uac.user_id = ?", userId).Rows()
	if err != nil {
		log.Info(err.Error())
		return userProfile, err
	}
	for rows.Next() {
		nt := entities.AdvertisementsCategories{}
		rows.Scan(&nt.ID, &nt.CategoryName)
		userProfile.UserAdvertisementCategories = append(userProfile.UserAdvertisementCategories, nt)
	}

	return
}
                             
func (m mobileAppUserMySqlRepo) UpdateUserProfile(ctx context.Context, user entities.MobileAppUser, mobileUserConfiguration entities.MobileUserConfiguration ,userAdvertisementCategories []int,userBankList []int ,userId int) (err error) {
	user.MobileUserConfigurations = mobileUserConfiguration
	err = m.DB.Model(&user).Where("id = ?", userId).Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "hash_password": user.HashPassword, "dob": user.DOB, "gender": user.Gender, "employee_status": user.EmployeeStatus, "address": user.Address, "civil_status": user.CivilStatus, "job_company_name": user.JobCompanyName, "job_company_location": user.JobCompanyLocation, "kids": user.Kids, "login_status": user.MobileUserConfigurations.LoginStatus, "push_notification_status": user.MobileUserConfigurations.PushNotificationStatus, "sound_status": user.MobileUserConfigurations.SoundStatus, "location_service_status": user.MobileUserConfigurations.LocationServiceStatus, "any_status": user.MobileUserConfigurations.AnyStatus }).Error
	if err != nil {
		err = errors_custom.ErrDuplicateUserEntry
		return
	}
	err = m.DB.Model(&mobileUserConfiguration).Where("user_id = ? ", userId).Updates(map[string]interface{}{"login_status": user.MobileUserConfigurations.LoginStatus, "push_notification_status": user.MobileUserConfigurations.PushNotificationStatus, "sound_status": user.MobileUserConfigurations.SoundStatus, "location_service_status": user.MobileUserConfigurations.LocationServiceStatus, "any_status": user.MobileUserConfigurations.AnyStatus}).Error
	count := 0
	bankCount := 0
	m.DB.Model(&entities.UserAdvertisementCategories{}).Where("user_id = ?", userId).Count(&count)
	if count == 0 {
		for _, element := range userAdvertisementCategories {
			m.DB.Create(&entities.UserAdvertisementCategories{UserId: userId, AdvertisementCatId: element})
		}

		}else{
		m.DB.Where("user_id = ?", userId).Delete(&entities.UserAdvertisementCategories{})
		for _, element := range userAdvertisementCategories {
		 m.DB.Create(&entities.UserAdvertisementCategories{UserId: userId, AdvertisementCatId: element})
		}
	}

	m.DB.Model(&entities.MobileUserBank{}).Where("mobile_user_id = ?", userId).Count(&bankCount)
	log.Info("mobile user bank list", userBankList)
	if bankCount == 0{
		for _, element := range userBankList {
			m.DB.Create(&entities.MobileUserBank{MobileUserId: userId, BankId: element})
		}
	}   else {
		m.DB.Where("mobile_user_id = ?", userId).Delete(&entities.MobileUserBank{})
		for _, element := range userBankList {
			m.DB.Create(&entities.MobileUserBank{MobileUserId: userId, BankId: element})
		}
	}

	return
}
