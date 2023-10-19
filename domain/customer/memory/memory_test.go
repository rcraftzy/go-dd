package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/roberto-carlos-tg/tavern/domain/customer"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name         string
		id           uuid.UUID
		exepectedErr error
	}

	cust, err := customer.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:         "no customer by id",
			id:           uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			exepectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:         "customer by id",
			id:           id,
			exepectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.exepectedErr) {
				t.Errorf("expected error %v, got %v", tc.exepectedErr, err)
			}
		})
	}
}
