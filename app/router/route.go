package router

import (
	"kodegiri/app/controller/auth"
	"kodegiri/app/controller/community"
	"kodegiri/app/controller/loyaltyprogram"
	"kodegiri/app/controller/membership"
	"kodegiri/app/controller/redeemedpoint"
	"kodegiri/app/controller/report"
	"kodegiri/app/controller/tiermanagement"
	"kodegiri/app/controller/transaction"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Handle(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/info", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(
			fiber.Map{
				"status":  200,
				"message": "info api",
			},
		)
	})

	// auth route
	api.Post("/login", auth.Login)

	protected := api.Group("", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("s3cret!!443")},
	}))

	// tier management route
	protected.Get("/tier-management", tiermanagement.ListTierManagement)
	protected.Post("/tier-management", tiermanagement.StoreTierManagement)
	protected.Get("/tier-management/:id", tiermanagement.DetailTierManagement)
	protected.Put("/tier-management/:id", tiermanagement.UpdateTierManagement)
	protected.Delete("/tier-management/:id", tiermanagement.DeleteTierManagement)

	// membership route
	protected.Get("/membership", membership.ListMembership)
	protected.Get("/membership/:id", membership.DetailMembership)

	// transaction
	protected.Post("/transaction", transaction.StoreTransaction)

	// community
	protected.Post("/community/member-get-member", community.StoreMemberGetMember)
	protected.Post("/community/member-activity", community.StoreMemberActivity)

	// reedem point
	protected.Post("/reedem-point", redeemedpoint.ReedemedPointStore)

	// loyalty program
	protected.Get("/loyalty-program", loyaltyprogram.ListLoyaltyProgram)
	protected.Post("/loyalty-program", loyaltyprogram.StoreLoyaltyProgram)

	// Report
	protected.Get("/report", report.ListReport)
}
