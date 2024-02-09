package report

import (
	"fmt"
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// ListReport func
// @Summary     ListReport function
// @Description  Get list of report
// @Tags Report
// @Params type query stirng false "earned or redeemed"
// @Params member_no query stirng false "member number"
// @Params dateStart query stirng false "15022023"
// @Params dateEnd query stirng false "15022023"
// @Params operator query stirng false "gt, lt, e, gte, lte"
// @Params operator_value query stirng false "operator value"
// @Params loyalty_program_id query stirng false "loyalty program id"
// @Success 200
// @Failure 400
// @Failure 404
// @Router       /report [get]
func ListReport(c *fiber.Ctx) error {
	reportType := c.Query("type")
	memberNumber := c.Query("member_no")
	dateStart := c.Query("dateStart")
	dateEnd := c.Query("dateEnd")
	operator := c.Query("operator")
	operatorValue := c.QueryInt("operator_value")
	loyaltyProgramId := c.Query("loyalty_program_id")

	db := services.DB

	switch operator {
	case "gt":
		operator = ">"
	case "lt":
		operator = "<"
	case "e":
		operator = "="
	case "gte":
		operator = "<="
	case "lte":
		operator = ">="
	}

	report := []model.PointTransaction{}
	q := db.Joins(`User`, `MemberGetMember`, `MemberActivity`, `Transaction`).
		Where("type = ?", reportType).
		Or("`User`.`member_number` LIKE %?%", memberNumber).
		Or("created_at >= ?", dateStart).
		Or("created_at <= ?", dateEnd).
		Or(fmt.Sprintf("value %s %d", operator, operatorValue))

	if reportType == "earned" {
		q = q.Joins(`LoyaltyProgram`).
			Where("`LoyaltyProgram`.`id` = ?", loyaltyProgramId)
	}

	q = q.Find(&report)

	if q.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Report not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&report)
}
