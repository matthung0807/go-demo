package main

import (
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				NewCartServiceImpl,
				fx.As(new(CartService)),
			),
			fx.Annotate(
				NewPaymentServiceImpl,
				fx.As(new(PaymentService)),
			),
			fx.Annotate(
				NewShippingServiceImpl,
				fx.As(new(ShippingService)),
			),
			fx.Annotate(
				NewOrderRepoImpl,
				fx.As(new(OrderRepo)),
			),
			fx.Annotate(
				NewProductRepoImpl,
				fx.As(new(ProductRepo)),
			),
			fx.Annotate(
				NewMemberRepoImpl,
				fx.As(new(MemberRepo)),
			),
			NewHttpClient,
		),
		fx.Invoke(func(cartService CartService) {
			cartService.Checkout()
		}),
	).Run()
}

// shopping cart
type CartService interface {
	Checkout()
}

type CartServiceImpl struct {
	orderRepo       OrderRepo
	paymentService  PaymentService
	shippingService ShippingService
}

func NewCartServiceImpl(
	orderRepo OrderRepo,
	paymentService PaymentService,
	shippingService ShippingService,
) *CartServiceImpl {
	return &CartServiceImpl{
		orderRepo:       orderRepo,
		paymentService:  paymentService,
		shippingService: shippingService,
	}
}

func (svc CartServiceImpl) Checkout() {
	fmt.Println("checkout - start")
	svc.orderRepo.GetOrder()
	svc.paymentService.Pay()
	svc.shippingService.Ship()
	fmt.Println("checkout - end")
}

// payment
type PaymentService interface {
	Pay()
}

type PaymentServiceImpl struct {
	httpClient *http.Client
}

func NewPaymentServiceImpl(httpClient *http.Client) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		httpClient: httpClient,
	}
}

func (svc PaymentServiceImpl) Pay() {
	fmt.Println("make payment")
}

// shipping
type ShippingService interface {
	Ship()
}

type ShippingServiceImpl struct {
	productRepo ProductRepo
	memberRepo  MemberRepo
}

func NewShippingServiceImpl(
	productService ProductRepo,
	memberService MemberRepo,
) *ShippingServiceImpl {
	return &ShippingServiceImpl{
		productRepo: productService,
		memberRepo:  memberService,
	}
}

func (svc ShippingServiceImpl) Ship() {
	fmt.Println("ship - start")
	svc.productRepo.GetProduct()
	svc.memberRepo.GetMember()
	fmt.Println("ship - end")
}

// product
type ProductRepo interface {
	GetProduct()
}

type ProductRepoImpl struct {
}

func NewProductRepoImpl() *ProductRepoImpl {
	return &ProductRepoImpl{}
}

func (svc ProductRepoImpl) GetProduct() {
	fmt.Println("get product")
}

// order
type OrderRepo interface {
	GetOrder()
}

type OrderRepoImpl struct {
}

func NewOrderRepoImpl() *OrderRepoImpl {
	return &OrderRepoImpl{}
}

func (svc OrderRepoImpl) GetOrder() {
	fmt.Println("get order")
}

// member
type MemberRepo interface {
	GetMember()
}

type MemberRepoImpl struct {
}

func NewMemberRepoImpl() *MemberRepoImpl {
	return &MemberRepoImpl{}
}

func (svc MemberRepoImpl) GetMember() {
	fmt.Println("get member")
}

func NewHttpClient() *http.Client {
	return http.DefaultClient
}
