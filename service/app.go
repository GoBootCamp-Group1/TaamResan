package service

import (
	"TaamResan/cmd/api/config"
	storage2 "TaamResan/internal/adapters/storage"
	"TaamResan/internal/address"
	"TaamResan/internal/user"
	"log"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg            config.Config
	dbConn         *gorm.DB
	userService    *UserService
	authService    *AuthService
	addressService *AddressService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB()
	storage2.Migrate(app.dbConn)

	app.setUserService()
	app.setAuthService()
	app.setAddressService()

	return app, nil
}

func (a *AppContainer) UserService() *UserService {
	return a.userService
}

func (a *AppContainer) AuthService() *AuthService {
	return a.authService
}

func (a *AppContainer) AddressService() *AddressService {
	return a.addressService
}

func (a *AppContainer) setUserService() {
	if a.userService != nil {
		return
	}
	a.userService = NewUserService(user.NewOps(storage2.NewUserRepo(a.dbConn)))
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage2.NewPostgresGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db
}

func (a *AppContainer) setAuthService() {
	if a.authService != nil {
		return
	}

	a.authService = NewAuthService(user.NewOps(storage2.NewUserRepo(a.dbConn)), []byte(a.cfg.Server.TokenSecret),
		a.cfg.Server.TokenExpMinutes,
		a.cfg.Server.RefreshTokenExpMinutes)
}

func (a *AppContainer) setAddressService() {
	if a.addressService != nil {
		return
	}

	a.addressService = NewAddressService(address.NewOps(storage2.NewAddressRepo(a.dbConn)))
}
