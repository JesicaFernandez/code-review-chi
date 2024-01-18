package service

import (
	"app/internal"
	"errors"
	"fmt"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindById is a method that returns a vehicle by id
func (s *VehicleDefault) FindById(id int) (v internal.Vehicle, err error) {
	v, err = s.rp.FindById(id)
	return
}

// FindLastId is a method that returns the last vehicle registered
func (s *VehicleDefault) FindLastId() (id int, err error) {
	return s.rp.FindLastId()
}
// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// CreateVehicle is a method that registers a vehicle
func (s *VehicleDefault) CreateVehicle(v internal.Vehicle) (err error) {
	err = s.rp.CreateVehicle(v)
	return
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(color, year)
	return
}

// FindAverageSpeedByBrand is a method that returns the average speed of vehicles of a specific brand
func (s *VehicleDefault) FindAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed, err = s.rp.FindAverageSpeedByBrand(brand)
	if err != nil {
		fmt.Errorf("the average speed of vehicles of the brand %s is %f", brand, averageSpeed)
	}
	return
}

// CreateVehicles is a method that registers several vehicles at the same time
func (s *VehicleDefault) CreateVehicles(v []internal.Vehicle) (err error) {
	err = s.rp.CreateVehicles(v)
	return
}

// UpdateSpeed is a method that updates the maximum speed of a specific vehicle
func (s *VehicleDefault) UpdateSpeed(id int, speed float64) (err error) {
	err = s.rp.UpdateSpeed(id, speed)
	return
}

// FindByFuelType is a method that returns a list of vehicles according to the type of fuel
func (s *VehicleDefault) FindByFuelType(fuelType string) (v []internal.Vehicle, err error) {
	v, err = s.rp.FindByFuelType(fuelType)
	return
}

// DeleteVehicle is a method that deletes a vehicle
func (s *VehicleDefault) DeleteVehicle(id int) (err error) {
	err = s.rp.DeleteVehicle(id)
	return
}

// FindByTransmissionType is a method that returns a list of vehicles according to their transmission type (manual, automatic, etc.)
func (s *VehicleDefault) FindByTransmissionType(transmissionType string) (v []internal.Vehicle, err error) {
	v, err = s.rp.FindByTransmissionType(transmissionType)
	return
}

// UpdateFuel is a method that updates the fuel type of a specific vehicle
func (s *VehicleDefault) UpdateFuel(id int, fuelType string) (err error) {
	err = s.rp.UpdateFuel(id, fuelType)
	return
}

// FindByDimensions is a method that returns a list of vehicles according to their dimensions (length, width)
func (s *VehicleDefault) FindByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v []internal.Vehicle, err error) {
	v, err = s.rp.FindByDimensions(minLength, maxLength, minWidth, maxWidth)
	return
}

// FindByWeight is a method that returns a list of vehicles according to their weight (minWeight, maxWeight)
func (s *VehicleDefault) FindByWeight(minWeight, maxWeight float64) (v []internal.Vehicle, err error) {
	v, err = s.rp.FindByWeight(minWeight, maxWeight)
	return
}

// FindByBrandAndYearRange is a method that returns a list of vehicles of a specific brand and year
func (s *VehicleDefault) FindByBrandAndYearRange(brand string, startYear, endYear int) (v []internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandAndYearRange(brand, startYear, endYear)
	return
}

// ValidateVehicleData is a method that validates the data of a vehicle
func (s *VehicleDefault) ValidateVehicleData(vehicle internal.Vehicle) error {
    
	// Validate is not empty
    if vehicle.Brand == "" {
		return errors.New("the brand cannot be empty")
	}
	if vehicle.Model == "" {
		return errors.New("the model cannot be empty")
	}
	if vehicle.Registration == "" {
		return errors.New("the registration cannot be empty")
	}
	if vehicle.FabricationYear == 0 {
		return errors.New("the year cannot be empty")
	}
	if vehicle.Color == "" {
		return errors.New("the color cannot be empty")
	}
	if vehicle.MaxSpeed == 0 {
		return errors.New("the max speed cannot be empty")
	}
	if vehicle.FuelType == "" {
		return errors.New("the fuel type cannot be empty")
	}
	if vehicle.Transmission == "" {
		return errors.New("the transmission cannot be empty")
	}
	if vehicle.Height == 0 {
		return errors.New("the height cannot be empty")
	}
	if vehicle.Width == 0 {
		return errors.New("the width cannot be empty")
	}
	if vehicle.Weight == 0.0 {
		return errors.New("the weight cannot be empty")
	}

    return nil
}