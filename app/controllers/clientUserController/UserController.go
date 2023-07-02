package clientUserController

import (
	"context"
	"fmt"
	"github.com/amirsobhani/melk_back/app/Repository/otpClientRepository"
	"github.com/amirsobhani/melk_back/app/Repository/userClientRepository"
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/app/requests"
	"github.com/amirsobhani/melk_back/app/requests/clientOtp"
	"github.com/amirsobhani/melk_back/infastructure"
	"github.com/amirsobhani/melk_back/queue"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func Signup(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
		})
	}

	if errors := requests.ValidateStruct(*user); errors != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: errors,
			Status:  fiber.StatusBadRequest,
		})
	}

	if userClientRepository.CheckExists(user.Mobile) {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: "user exists please sign-in",
			Status:  fiber.StatusBadRequest,
		})
	}

	countOtp := otpClientRepository.CountValidation(user.Mobile)

	if countOtp > 4 {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: "too many request please try later",
			Status:  fiber.StatusTooManyRequests,
		})
	}

	token := otpClientRepository.GenerateOtp(user)

	return infastructure.Output(c, &infastructure.OutputStruct{
		Data:    token,
		Message: "otp token has been send",
	})
}

func Check(c *fiber.Ctx) error {
	//userId, err := infastructure.VerifyJWT(c)

	//defer func() {
	//	if r := recover(); r != nil {
	//		log.Fatalf("Runtime error: %v", r)
	//	}
	//}()

	q, err := queue.ChannelRabbitMQ.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = queue.ChannelRabbitMQ.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		fmt.Println(err)
	}

	return nil

}

func OtpValidator(c *fiber.Ctx) error {

	otp := new(clientOtp.OtpValidator)

	if err := c.BodyParser(otp); err != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
		})
	}

	if errors := requests.ValidateStruct(*otp); errors != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: errors,
			Status:  fiber.StatusBadRequest,
		})
	}

	check := otpClientRepository.CheckValidOtp(otp.Token, otp.Mobile)

	if check.ID == 0 {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: "otp token not valid",
			Status:  fiber.StatusBadRequest,
		})
	}

	otpClientRepository.RemoveAllMobileOtp(otp.Mobile)

	var user models.User

	if userClientRepository.CheckExists(otp.Mobile) {
		user = userClientRepository.FindByMobile(otp.Mobile)

	} else {
		var userTempData = check.TempData.Data()

		user = userClientRepository.Create(userTempData)
	}

	token, err := infastructure.GenerateJWT(user.ID)

	if err != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}

	return infastructure.Output(c, &infastructure.OutputStruct{
		Data:    token,
		Message: "user successful login",
	})
}
