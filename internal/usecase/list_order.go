package usecase

import "github.com/lwazevedo/goexpert_clean_arch/internal/entity"

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	var out []OrderOutputDTO
	orders, err := l.OrderRepository.List()
	if err != nil {
		return out, nil
	}
	for _, o := range orders {
		order := OrderOutputDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.Price + o.Tax,
		}
		out = append(out, order)
	}
	return out, nil
}
