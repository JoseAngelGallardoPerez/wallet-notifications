package push_token

import (
	"errors"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Confialink/wallet-notifications/internal/db/dao"
	"github.com/Confialink/wallet-notifications/internal/db/models"
	"github.com/Confialink/wallet-notifications/internal/validators"
)

var _ = Describe("PushToken service", func() {
	db, dbMock, _ := sqlmock.New()
	gormDbMock, _ := gorm.Open("mysql", db)
	//gormDbMock.LogMode(true)
	service := NewService(dao.NewPushToken(gormDbMock), time.Duration(1))
	in := &validators.AddPushToken{
		DeviceId:  "random_device_id",
		Name:      "random_name",
		Os:        "random",
		PushToken: "random_push",
	}
	userId := "random-uid"

	Context("AddOrUpdate()", func() {

		selectQuery := "SELECT (.+) FROM `push_tokens`  WHERE \\(push_token = \\?\\) (.+)"
		When("the DB cannot find an existing record and returns an error", func() {
			It("returns an error", func() {
				dbMock.ExpectQuery(selectQuery).
					WithArgs(in.PushToken).
					WillReturnError(errors.New("some db error"))

				_, err := service.AddOrUpdate(in, userId)

				Expect(err).Should(HaveOccurred())
				Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
			})
		})

		When("the push token already exists in the DB", func() {
			Context("the push token already exists in the DB", func() {
				It("updates the existing record successful", func() {
					sqlRows := sqlmock.NewRows([]string{"push_token", "os", "name", "uid", "device_id"}).AddRow(in.PushToken, "android", "random-name-2", "random-uid-2", "random-device-id-2")

					dbMock.ExpectQuery(selectQuery).
						WithArgs(in.PushToken).
						WillReturnRows(sqlRows)

					dbMock.ExpectBegin()
					dbMock.ExpectExec("UPDATE `push_tokens` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
					dbMock.ExpectCommit()

					res, err := service.AddOrUpdate(in, userId)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
					Expect(res.DeviceId).Should(Equal(in.DeviceId))
					Expect(res.Name).Should(Equal(in.Name))
					Expect(res.Os).Should(Equal(in.Os))
					Expect(res.PushToken).Should(Equal(in.PushToken))
				})

				It("returns an error", func() {
					sqlRows := sqlmock.NewRows([]string{"push_token", "os", "name", "uid", "device_id"}).AddRow(in.PushToken, "android", "random-name-2", "random-uid-2", "random-device-id-2")

					dbMock.ExpectQuery(selectQuery).
						WithArgs(in.PushToken).
						WillReturnRows(sqlRows)

					dbMock.ExpectBegin()
					dbMock.ExpectExec("UPDATE `push_tokens` (.+)").WillReturnError(errors.New("some db error"))

					_, err := service.AddOrUpdate(in, userId)
					Expect(err).Should(HaveOccurred())
					Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
				})
			})
		})

		When("DB does not have provided push token", func() {
			Context("adds a new push token to the DB", func() {
				Context("successful insert", func() {
					It("does successful and OS is Other", func() {
						dbMock.ExpectQuery(selectQuery).
							WithArgs(in.PushToken).
							WillReturnError(gorm.ErrRecordNotFound)

						dbMock.ExpectBegin()
						dbMock.ExpectExec("INSERT INTO `push_tokens` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
						dbMock.ExpectCommit()

						res, err := service.AddOrUpdate(in, userId)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
						Expect(res.DeviceId).Should(Equal(in.DeviceId))
						Expect(res.Name).Should(Equal(in.Name))
						Expect(res.Os).Should(Equal(in.Os))
						Expect(res.PushToken).Should(Equal(in.PushToken))
						Expect(res.Os).Should(Equal(models.OsTypeOther))
					})
					It("does successful and OS is `android`", func() {
						dbMock.ExpectQuery(selectQuery).
							WithArgs(in.PushToken).
							WillReturnError(gorm.ErrRecordNotFound)

						dbMock.ExpectBegin()
						dbMock.ExpectExec("INSERT INTO `push_tokens` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
						dbMock.ExpectCommit()

						in.Os = models.OsTypeAndroid
						res, err := service.AddOrUpdate(in, userId)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
						Expect(res.DeviceId).Should(Equal(in.DeviceId))
						Expect(res.Name).Should(Equal(in.Name))
						Expect(res.Os).Should(Equal(in.Os))
						Expect(res.PushToken).Should(Equal(in.PushToken))
						Expect(res.Os).Should(Equal(models.OsTypeAndroid))
					})
					It("does successful and OS is `ios`", func() {
						dbMock.ExpectQuery(selectQuery).
							WithArgs(in.PushToken).
							WillReturnError(gorm.ErrRecordNotFound)

						dbMock.ExpectBegin()
						dbMock.ExpectExec("INSERT INTO `push_tokens` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
						dbMock.ExpectCommit()

						in.Os = models.OsTypeIos
						res, err := service.AddOrUpdate(in, userId)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
						Expect(res.DeviceId).Should(Equal(in.DeviceId))
						Expect(res.Name).Should(Equal(in.Name))
						Expect(res.Os).Should(Equal(in.Os))
						Expect(res.PushToken).Should(Equal(in.PushToken))
						Expect(res.Os).Should(Equal(models.OsTypeIos))
					})
				})

				It("returns an error", func() {
					dbMock.ExpectQuery(selectQuery).
						WithArgs(in.PushToken).
						WillReturnError(gorm.ErrRecordNotFound)

					dbMock.ExpectBegin()
					dbMock.ExpectExec("INSERT INTO `push_tokens` (.+)").WillReturnError(errors.New("some db error"))

					_, err := service.AddOrUpdate(in, userId)
					Expect(err).Should(HaveOccurred())
					Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
				})
			})
		})
	})

	Context("RemovePushToken()", func() {
		slqQuery := "DELETE FROM `push_tokens`  WHERE \\(push_token = \\?\\)"
		When("we remove a push token from DB", func() {
			It("should not return an error", func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(slqQuery).
					WithArgs(in.PushToken).
					WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()

				Expect(service.RemovePushToken(in.PushToken)).ShouldNot(HaveOccurred())
				Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
			})

			It("should return an error", func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(slqQuery).
					WithArgs(in.PushToken).
					WillReturnError(errors.New("some db error"))

				Expect(service.RemovePushToken(in.PushToken)).Should(HaveOccurred())
				Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
			})
		})
	})

	Context("NotExpiredTokensByUser()", func() {
		selectQuery := "SELECT (.+) FROM `push_tokens` WHERE \\(uid = \\?\\ AND updated_at > \\?\\) (.+)"
		When("we request tokens from the DB", func() {
			It("should not return an error", func() {
				data := []string{"push_token", "os", "name", "uid", "device_id"}
				sqlRows := sqlmock.NewRows(data).AddRow(in.PushToken, "android", "random-name-2", "random-uid-2", "random-device-id-2")
				dbMock.ExpectQuery(selectQuery).
					WithArgs(userId, sqlmock.AnyArg()).
					WillReturnRows(sqlRows)

				res, err := service.NotExpiredTokensByUser(userId)

				Expect(len(res)).Should(Equal(1))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
			})

			It("should return an error", func() {
				dbMock.ExpectQuery(selectQuery).
					WithArgs(userId, sqlmock.AnyArg()).
					WillReturnError(errors.New("some db error"))

				_, err := service.NotExpiredTokensByUser(userId)

				Expect(err).Should(HaveOccurred())
				Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
			})
		})
	})

	Context("DeleteExpiredTokens()", func() {
		slqQuery := "DELETE FROM `push_tokens` WHERE \\(updated_at < \\?\\)"
		When("we remove old push tokens from DB", func() {
			It("should not return an error", func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(slqQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()

				Expect(service.DeleteExpiredTokens()).ShouldNot(HaveOccurred())
				Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
			})

			It("should return an error", func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(slqQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(errors.New("some db error"))

				Expect(service.DeleteExpiredTokens()).Should(HaveOccurred())
				Expect(dbMock.ExpectationsWereMet()).Should(BeNil())
			})
		})
	})
})
