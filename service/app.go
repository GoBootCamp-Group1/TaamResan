package service

import (
	"TaamResan/cmd/api/config"
	"TaamResan/internal/action_log"
	storage2 "TaamResan/internal/adapters/storage"
	"TaamResan/internal/address"
	"TaamResan/internal/category"
	"TaamResan/internal/restaurant"
	"TaamResan/internal/restaurant_staff"
	"TaamResan/internal/role"
	"TaamResan/internal/user"
	"TaamResan/internal/wallet"
	"log"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg                    config.Config
	dbConn                 *gorm.DB
	userService            *UserService
	authService            *AuthService
	addressService         *AddressService
	roleService            *RoleService
	walletService          *WalletService
	restaurantService      *RestaurantService
	restaurantStaffService *RestaurantStaffService
	actionLogService       *ActionLogService
	categoryService        *CategoryService
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
	app.setRoleService()
	app.setWalletService()
	app.setRestaurantService()
	app.setRestaurantStaffService()
	app.setActionLogService()
	app.setCategoryService()

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

func (a *AppContainer) RoleService() *RoleService {
	return a.roleService
}

func (a *AppContainer) WalletService() *WalletService {
	return a.walletService
}

func (a *AppContainer) RestaurantService() *RestaurantService { return a.restaurantService }

func (a *AppContainer) CategoryService() *CategoryService { return a.categoryService }

func (a *AppContainer) RestaurantStaffService() *RestaurantStaffService {
	return a.restaurantStaffService
}

func (a *AppContainer) ActionLogService() *ActionLogService {
	return a.actionLogService
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

func (a *AppContainer) setRoleService() {
	if a.roleService != nil {
		return
	}
	a.roleService = NewRoleService(role.NewOps(storage2.NewRoleRepo(a.dbConn)))
}

func (a *AppContainer) setWalletService() {
	if a.walletService != nil {
		return
	}
	a.walletService = NewWalletService(wallet.NewOps(storage2.NewWalletRepo(a.dbConn)))
}

func (a *AppContainer) setRestaurantService() {
	if a.restaurantService != nil {
		return
	}
	a.restaurantService = NewRestaurantService(restaurant.NewOps(storage2.NewRestaurantRepo(a.dbConn)))
}

func (a *AppContainer) setRestaurantStaffService() {
	if a.restaurantStaffService != nil {
		return
	}
	a.restaurantStaffService = NewRestaurantStaffService(restaurant_staff.NewOps(storage2.NewRestaurantStaffRepo(a.dbConn)))
}

func (a *AppContainer) setActionLogService() {
	if a.actionLogService != nil {
		return
	}
	a.actionLogService = NewActionLogService(action_log.NewOps(storage2.NewActionLogRepo(a.dbConn)))
}

func (a *AppContainer) setCategoryService() {
	if a.categoryService != nil {
		return
	}
	a.categoryService = NewCategoryService(category.NewOps(storage2.NewCategoryRepo(a.dbConn)))
}
