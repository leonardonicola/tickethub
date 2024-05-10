package gateway

import (
	"errors"
	"os"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	ticketDTO "github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	user "github.com/leonardonicola/tickethub/internal/modules/user/domain"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/customer"
	"github.com/stripe/stripe-go/v78/paymentlink"
	"github.com/stripe/stripe-go/v78/product"
)

var (
	logger *config.Logger
)

type CreatedCustomer struct {
	ID    string
	Name  string
	Email string
}

type StripeGateway []struct{}

func SetupStripe() error {
	logger = config.NewLogger()
	str, ok := os.LookupEnv("STRIPE_SECRET_KEY")

	if !ok {
		return errors.New("Coulnd't find stripe key!")
	}
	stripe.Key = str
	stripe.SetAppInfo(&stripe.AppInfo{
		Name:    "TicketHub",
		Version: "0.0.1",
		URL:     "https://github.com/leonardonicola/tickethub",
	})
	logger.Info("Stripe configured successfully")
	return nil
}

func GetStripeGateway() *StripeGateway {
	return &StripeGateway{}
}

func (s *StripeGateway) CreateCustomer(user *user.User) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Name:  &user.Name,
		Email: &user.Email,
	}
	res, err := customer.New(params)
	if err != nil {
		logger.Errorf("Error while creating customer: %v", err)
		return nil, HandleStripeErrors(err)
	}

	logger.Debug("Stripe customer created successfully")
	return res, nil
}

func (s *StripeGateway) GetCustomer(id string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{}
	res, err := customer.Get(id, params)
	if err != nil {
		logger.Errorf("Error while getting STRIPE customer: %v", err)
		return nil, HandleStripeErrors(err)
	}
	return res, nil
}

func (s *StripeGateway) DeleteCustomer(id string) error {
	_, err := customer.Del(id, nil)
	if err != nil {
		logger.Errorf("Error while deleting a STRIPE customer: %v", err)
		return HandleStripeErrors(err)
	}
	return nil
}

func (s *StripeGateway) CreateProduct(ticket *dto.CreateTicketOutputDTO) (*dto.TicketProduct, error) {
	params := &stripe.ProductParams{
		Name:        &ticket.Name,
		Description: &ticket.Description,
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:    stripe.String(string(stripe.CurrencyBRL)),
			UnitAmount:  stripe.Int64(ticket.Price),
			TaxBehavior: stripe.String(string(stripe.PriceTaxBehaviorExclusive)),
		},
	}
	res, err := product.New(params)
	if err != nil {
		logger.Errorf("Error while creating STRIPE product: %v", err)
		return nil, HandleStripeErrors(err)
	}

	logger.Debug("STRIPE product created successfully")
	return &dto.TicketProduct{
		PriceID:     res.DefaultPrice.ID,
		TicketID:    ticket.ID,
		Description: res.Description,
		Name:        res.Name,
		Metadata:    res.Metadata,
		StripeID:    res.ID,
	}, nil
}

func (s *StripeGateway) GetProduct(id string) (*dto.TicketProduct, error) {
	params := &stripe.ProductParams{}

	res, err := product.Get(id, params)

	if err != nil {
		logger.Errorf("Error while getting STRIPE product: %v", err)
		return nil, HandleStripeErrors(err)
	}

	return &dto.TicketProduct{
		PriceID:     res.DefaultPrice.ID,
		Description: res.Description,
		StripeID:    res.ID,
		Name:        res.Name,
		Metadata:    res.Metadata,
	}, nil
}

func (s *StripeGateway) CreatePaymentLink(tickets []*ticketDTO.CreatePaymentDTO, metadata map[string]string) (string, error) {
	params := &stripe.PaymentLinkParams{
		Metadata:  metadata,
		LineItems: []*stripe.PaymentLinkLineItemParams{},
	}

	for _, ticket := range tickets {
		stripeProduct := &stripe.PaymentLinkLineItemParams{
			Price:    stripe.String(ticket.PriceID),
			Quantity: stripe.Int64(ticket.Quantity),
		}
		params.LineItems = append(params.LineItems, stripeProduct)
	}
	res, err := paymentlink.New(params)
	if err != nil {
		// logger.Errorf("Error while creating STRIPE payment link: %v", err)
		return "", HandleStripeErrors(err)
	}

	return res.URL, nil
}
