package adapters

import (
	"CaliYa/config"
	"CaliYa/core/domain/ports"
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type twilioClient struct {
	client *twilio.RestClient
}

func NewTwilioClient(config config.Config) ports.MessageSender {

	cli := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.Twilio.Username,
		Password: config.Twilio.Password,
	})

	return &twilioClient{
		client: cli,
	}
}

func (t *twilioClient) SendCodeToConfirmPhone(to, name, code string) error {

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("whatsapp:" + to)
	params.SetFrom("whatsapp:" + config.Get().Twilio.Phone)
	params.SetBody(fmt.Sprintf("Hola %s \n\n Su codigo de confirmacion es: %s \n\n Gracias por elegir a	 caliya.app", name, code))

	resp, err := t.client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return nil
}
