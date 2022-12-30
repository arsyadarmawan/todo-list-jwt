package cart

import (
	_domain "task/model/domain"
	"time"
)

type Cart struct {
	Id         int
	User_id    int
	Product_id int
	Created_At time.Time
	Updated_at time.Time
}

func (c *Cart) ToDomain() _domain.Cart {
	return _domain.Cart{
		Id:         int(c.Id),
		User_id:    int(c.User_id),
		Product_id: int(c.Product_id),
		Created_At: c.Created_At,
		Updated_at: c.Updated_at,
	}
}
