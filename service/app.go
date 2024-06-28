package service

import (
	"TaamResan/cmd/api/config"
	"TaamResan/internal/action_log"
	"TaamResan/internal/address"
	"TaamResan/internal/block_restaurant"
	"TaamResan/internal/cart"
	"TaamResan/internal/cart_item"
	"TaamResan/internal/category"
	"TaamResan/internal/category_food"
	"TaamResan/internal/food"
	"TaamResan/internal/order"
	"TaamResan/internal/restaurant"
	"TaamResan/internal/restaurant_staff"
	"TaamResan/internal/role"
	"TaamResan/internal/user"
	"TaamResan/internal/wallet"
	storage2 "TaamResan/pkg/adapters/storage"
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
	foodService            *FoodService
	categoryFoodService    *CategoryFoodService
	blockRestaurantService *BlockRestaurantService
	cartService            *CartService
	cartItemService        *CartItemService
	searchService          *SearchService
	orderService           *OrderService
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
	app.setFoodService()
	app.setCategoryFoodService()
	app.setBlockRestaurantService()
	app.setCartService()
	app.setCartItemService()
	app.setSearchService()
	app.setOrderService()

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

func (a *AppContainer) FoodService() *FoodService { return a.foodService }

func (a *AppContainer) CategoryFoodService() *CategoryFoodService { return a.categoryFoodService }

func (a *AppContainer) RestaurantStaffService() *RestaurantStaffService {
	return a.restaurantStaffService
}

func (a *AppContainer) ActionLogService() *ActionLogService {
	return a.actionLogService
}

func (a *AppContainer) BlockRestaurantService() *BlockRestaurantService {
	return a.blockRestaurantService
}

func (a *AppContainer) CartService() *CartService {
	return a.cartService
}

func (a *AppContainer) CartItemService() *CartItemService {
	return a.cartItemService
}

func (a *AppContainer) SearchService() *SearchService {
	return a.searchService
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
	a.restaurantService = NewRestaurantService(
		restaurant.NewOps(storage2.NewRestaurantRepo(a.dbConn)),
		action_log.NewOps(storage2.NewActionLogRepo(a.dbConn)),
	)
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

func (a *AppContainer) setFoodService() {
	if a.foodService != nil {
		return
	}
	a.foodService = NewFoodService(
		food.NewOps(storage2.NewFoodRepo(a.dbConn)),
		category.NewOps(storage2.NewCategoryRepo(a.dbConn)),
	)
}

func (a *AppContainer) setCategoryFoodService() {
	if a.categoryFoodService != nil {
		return
	}
	a.categoryFoodService = NewCategoryFoodService(category_food.NewOps(storage2.NewCategoryFoodRepo(a.dbConn)))
}

func (a *AppContainer) setBlockRestaurantService() {
	if a.blockRestaurantService != nil {
		return
	}
	a.blockRestaurantService = NewBlockRestaurantService(
		block_restaurant.NewOps(storage2.NewBlockRestaurantRepo(a.dbConn)),
		restaurant.NewOps(storage2.NewRestaurantRepo(a.dbConn)))
}

func (a *AppContainer) setCartService() {
	if a.cartService != nil {
		return
	}
	a.cartService = NewCartService(cart.NewOps(storage2.NewCartRepo(a.dbConn)))
}

func (a *AppContainer) setCartItemService() {
	if a.cartItemService != nil {
		return
	}
	a.cartItemService = NewCartItemService(cart_item.NewOps(storage2.NewCartItemRepo(a.dbConn)))
}

func (a *AppContainer) setSearchService() {
	if a.searchService != nil {
		return
	}

	a.searchService = NewSearchService(restaurant.NewOps(storage2.NewRestaurantRepo(a.dbConn)), food.NewOps(storage2.NewFoodRepo(a.dbConn)))
}

func (a *AppContainer) OrderService() *OrderService {
	return a.orderService
}

func (a *AppContainer) setOrderService() {
	if a.orderService != nil {
		return
	}

	orderOps := order.NewOps(storage2.NewOrderRepo(a.dbConn))
	cartOps := cart.NewOps(storage2.NewCartRepo(a.dbConn))
	foodOps := food.NewOps(storage2.NewFoodRepo(a.dbConn))
	walletOps := wallet.NewOps(storage2.NewWalletRepo(a.dbConn))
	restaurantStaffOps := restaurant_staff.NewOps(storage2.NewRestaurantStaffRepo(a.dbConn))

	a.orderService = NewOrderService(orderOps, cartOps, foodOps, walletOps, restaurantStaffOps)
}
