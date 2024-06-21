package service

import (
	"TaamResan/internal/address"
	"context"
	"errors"
	"fmt"
)

type AddressService struct {
	addressOps *address.Ops
}

func NewAddressService(addressOps *address.Ops) *AddressService {
	return &AddressService{
		addressOps: addressOps,
	}
}

var (
	ErrFetchingAddress = errors.New("can not fetch address")
	ErrCreatingAddress = errors.New("can not create address")
	ErrUpdatingAddress = errors.New("can not update address")
	ErrDeletingAddress = errors.New("can not delete address")
)

func (s *AddressService) CreateAddress(ctx context.Context, address *address.Address) error {

	err := s.addressOps.Create(ctx, address)
	if err != nil {
		return fmt.Errorf(ErrCreatingAddress.Error()+": %w", err)
	}

	return nil
}

func (s *AddressService) UpdateAddress(ctx context.Context, address *address.Address) error {
	err := s.addressOps.Update(ctx, address)
	if err != nil {
		return fmt.Errorf(ErrUpdatingAddress.Error()+": %w", err)
	}

	return nil
}

func (s *AddressService) DeleteAddress(ctx context.Context, address *address.Address) error {
	err := s.addressOps.Delete(ctx, address)
	if err != nil {
		return fmt.Errorf(ErrDeletingAddress.Error()+": %w", err)
	}

	return nil
}

func (s *AddressService) GetAddressByID(ctx context.Context, id uint) (*address.Address, error) {
	foundAddress, err := s.addressOps.GetAddressByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf(ErrUpdatingAddress.Error()+": %w", err)
	}

	return foundAddress, nil
}

func (s *AddressService) GetAll(ctx context.Context) ([]*address.Address, error) {
	fetchedAddresses, err := s.addressOps.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrFetchingAddress.Error()+": %w", err)
	}

	return fetchedAddresses, nil
}
