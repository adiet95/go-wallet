package database

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"
	"go-wallet/src/modules/payment"
	redisrepo "go-wallet/src/modules/redis"
	"go-wallet/src/modules/topup"
	"go-wallet/src/modules/transfer"
	"go-wallet/src/modules/users"
	"log"

	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "start seeding",
	RunE:  seeder,
}

func seeder(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}
	rdClient := RedisClient()

	adminPin, _ := libs.HashPassword("123987")
	adminPinSql := libs.ToNullString(adminPin)
	phone := libs.ToNullString("081388355300")
	phoneA := libs.ToNullString("081388355301")
	phoneB := libs.ToNullString("081388355302")
	phoneC := libs.ToNullString("081388355303")
	phoneD := libs.ToNullString("081388355304")
	phoneE := libs.ToNullString("081388355305")
	phoneF := libs.ToNullString("081388355306")

	adminRole := "admin"
	userRole := "user"

	adminUserId, _ := uuid.NewV4()
	userA, _ := uuid.NewV4()
	userB, _ := uuid.NewV4()
	userC, _ := uuid.NewV4()
	userD, _ := uuid.NewV4()
	userE, _ := uuid.NewV4()
	userF, _ := uuid.NewV4()

	userArr := []string{userA.String(), userB.String(), userC.String(), userD.String(), userE.String(), userF.String()}

	nameAdmin := libs.ToNullString("Admin")
	nameA := libs.ToNullString("UserA")
	nameB := libs.ToNullString("UserB")
	nameC := libs.ToNullString("UserC")
	nameD := libs.ToNullString("UserD")
	nameE := libs.ToNullString("UserE")
	nameF := libs.ToNullString("UserF")

	lastName := libs.ToNullString("Dummy")

	balanceDefault := libs.ToNullInt64(100000)

	var datas = []entity.User{
		{UserId: adminUserId.String(), PhoneNumber: phone, Pin: adminPinSql, Role: adminRole, Balance: balanceDefault, FirstName: nameAdmin, LastName: lastName},
		{UserId: userA.String(), PhoneNumber: phoneA, Pin: adminPinSql, Role: userRole, Balance: balanceDefault, FirstName: nameA, LastName: lastName},
		{UserId: userB.String(), PhoneNumber: phoneB, Pin: adminPinSql, Role: userRole, Balance: balanceDefault, FirstName: nameB, LastName: lastName},
		{UserId: userC.String(), PhoneNumber: phoneC, Pin: adminPinSql, Role: userRole, Balance: balanceDefault, FirstName: nameC, LastName: lastName},
		{UserId: userD.String(), PhoneNumber: phoneD, Pin: adminPinSql, Role: userRole, Balance: balanceDefault, FirstName: nameD, LastName: lastName},
		{UserId: userE.String(), PhoneNumber: phoneE, Pin: adminPinSql, Role: userRole, Balance: balanceDefault, FirstName: nameE, LastName: lastName},
		{UserId: userF.String(), PhoneNumber: phoneF, Pin: adminPinSql, Role: userRole, Balance: balanceDefault, FirstName: nameF, LastName: lastName},
	}

	redisRepo := redisrepo.NewRepo(rdClient)
	userRepo := users.NewRepo(db)

	paySvc := payment.NewService(redisRepo, userRepo)
	trfSvc := transfer.NewService(redisRepo, userRepo)
	topSvc := topup.NewService(userRepo, redisRepo)

	payReq1 := &models.PaymentRequest{
		Amount:  11000,
		Remarks: "Payment Request 01",
	}

	payReq2 := &models.PaymentRequest{
		Amount:  12000,
		Remarks: "Payment Request 02",
	}

	payReq3 := &models.PaymentRequest{
		Amount:  13000,
		Remarks: "Payment Request 03",
	}

	topReq1 := &models.TopUpRequest{
		Amount:  15000,
		Remarks: "Top Up Request 01",
	}
	topReq2 := &models.TopUpRequest{
		Amount:  16000,
		Remarks: "Top Up Request 02",
	}
	topReq3 := &models.TopUpRequest{
		Amount:  17000,
		Remarks: "Top Up Request 03",
	}

	for key, userId := range userArr {
		paySvc.PostPayment(payReq1, userId)
		topSvc.PostTopUp(topReq1, userId)

		paySvc.PostPayment(payReq2, userId)
		topSvc.PostTopUp(topReq2, userId)

		paySvc.PostPayment(payReq3, userId)
		topSvc.PostTopUp(topReq3, userId)

		if key+1 == len(userArr) {
			trfReq := &models.TransferRequest{
				Amount:     20000,
				TargetUser: userArr[0],
				Remarks:    "Transfer Request",
			}
			trfSvc.PostTransfer(trfReq, userId)
		} else {
			trfReq := &models.TransferRequest{
				Amount:     20000,
				TargetUser: userArr[key+1],
				Remarks:    "Transfer Request",
			}
			trfSvc.PostTransfer(trfReq, userId)
		}
	}

	if res := db.Create(&datas); res.Error != nil {
		return res.Error
	}
	log.Println("Seeding successful")
	return nil
}
