module github.com/Confialink/wallet-notifications

go 1.13

replace github.com/Confialink/wallet-notifications/rpc/proto/notifications => ./rpc/proto/notifications

require (
	github.com/Confialink/wallet-accounts/rpc/accounts v0.0.0-20210218063536-4e2d21b26af2
	github.com/Confialink/wallet-messages/rpc/messages v0.0.0-20210218075414-cc2d184a57a0
	github.com/Confialink/wallet-notifications/rpc/proto/notifications v0.0.0-00010101000000-000000000000
	github.com/Confialink/wallet-permissions/rpc/permissions v0.0.0-20210218072732-21caf4a66e86
	github.com/Confialink/wallet-pkg-acl v0.0.0-20210218070839-a03813da4b89
	github.com/Confialink/wallet-pkg-discovery/v2 v2.0.0-20210217105157-30e31661c1d1
	github.com/Confialink/wallet-pkg-env_config v0.0.0-20210217112253-9483d21626ce
	github.com/Confialink/wallet-pkg-env_mods v0.0.0-20210217112432-4bda6de1ee2c
	github.com/Confialink/wallet-pkg-errors v1.0.2
	github.com/Confialink/wallet-pkg-gomail v0.0.0-20210217103943-f8af5fb0d369
	github.com/Confialink/wallet-pkg-http v0.0.0-20210217113129-11c7677deade
	github.com/Confialink/wallet-pkg-list_params v0.0.0-20210217104359-69dfc53fe9ee
	github.com/Confialink/wallet-pkg-service_names v0.0.0-20210217112604-179d69540dea
	github.com/Confialink/wallet-settings/rpc/proto/settings v0.0.0-20210218070334-b4153fc126a0
	github.com/Confialink/wallet-users/rpc/proto/users v0.0.0-20210218071418-0600c0533fb2
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/appleboy/go-fcm v0.1.5
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/inconshreveable/log15 v0.0.0-20200109203555-b30bc20e4fd1
	github.com/jasonlvhit/gocron v0.0.1
	github.com/jinzhu/gorm v1.9.15
	github.com/kevinburke/go-types v0.0.0-20201208005256-aee49f568a20 // indirect
	github.com/kevinburke/go.uuid v1.2.0 // indirect
	github.com/kevinburke/rest v0.0.0-20210106062955-7f83c79a0622
	github.com/kevinburke/twilio-go v0.0.0-20210106065930-582a176721b4
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/ttacon/builder v0.0.0-20170518171403-c099f663e1c2 // indirect
	github.com/ttacon/libphonenumber v1.1.0 // indirect
)
