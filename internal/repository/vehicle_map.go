package repository

import (	
	"app/internal"
	"errors"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		v[key] = value
	}

	return
}

// FindById is a method that returns a vehicle by id
func (r *VehicleMap) FindById(id int) (v internal.Vehicle, err error) {
v, ok := r.db[id]
	if !ok {
		err = errors.New("vehicle not found")
	}
	return
}

// FindLastId is a method that returns the last vehicle registered
func (r *VehicleMap) FindLastId() (id int, err error) {
	
	for key := range r.db {
		id = key
	}
	return
}
// CreateVehicle is a method that registers a vehicle
func (r *VehicleMap) CreateVehicle(v internal.Vehicle) (err error) {
	r.db[v.Id] = v
	return
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {

	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	return
}

// FindAverageSpeedByBrand is a method that returns the average speed of vehicles of a specific brand
func (r *VehicleMap) FindAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	var sum float64
	var count int

	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			sum += value.MaxSpeed
			count++
		}
	}

	// calculate average speed
	averageSpeed = sum / float64(count)
	return
}

// CreateVehicles is a method that registers several vehicles at the same time
func (r *VehicleMap) CreateVehicles(v []internal.Vehicle) (err error) {
	for _, value := range v {
		r.db[value.Id] = value
	}
	return
}

// UpdateSpeed is a method that updates the maximum speed of a specific vehicle
func (r *VehicleMap) UpdateSpeed(id int, speed float64) (err error) {
	// Check if the vehicle with the given ID exists
    vehicle, exists := r.db[id]
	if !exists {
        return internal.ErrVehicleNotFound
    }

    // Update the maximum speed
    vehicle.MaxSpeed = speed
	r.db[id] = vehicle
    return
}

// FindByFuelType is a method that returns a list of vehicles according to the type of fuel
func (r *VehicleMap) FindByFuelType(fuelType string) (v []internal.Vehicle, err error) {
	// copy db
	for _, value := range r.db {
		if value.FuelType == fuelType {
			v = append(v, value)
		}
	}
	return
}

// DeleteVehicle is a method that deletes a vehicle
func (r *VehicleMap) DeleteVehicle(id int) (err error) {
	delete(r.db, id)
	return
}

// FindByTransmissionType is a method that returns a list of vehicles according to their transmission type (manual, automatic, etc.)
func (r *VehicleMap) FindByTransmissionType(transmissionType string) (v []internal.Vehicle, err error) {
	// copy db
	for _, value := range r.db {
		if value.Transmission == transmissionType {
			v = append(v, value)
		}
	}
	return
}

// UpdateFuel is a method that updates the fuel type of a specific vehicle
func (r *VehicleMap) UpdateFuel(id int, fuelType string) (err error) {
	// Check if the vehicle with the given ID exists
	vehicle, exists := r.db[id]
	if !exists {
		return internal.ErrVehicleNotFound
	}

	// Update the fuel type
	vehicle.FuelType = fuelType
	r.db[id] = vehicle
	return nil
}

// FindByDimensions is a method that returns a list of vehicles according to their dimensions (length, width)
func (r *VehicleMap) FindByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v []internal.Vehicle, err error) {
	// copy db
	for _, value := range r.db {
		if value.Length >= minLength && value.Length <= maxLength && value.Width >= minWidth && value.Width <= maxWidth {
			v = append(v, value)
		}
	}
	return
}

// FindByWeight is a method that returns a list of vehicles according to their weight (minWeight, maxWeight)
func (r *VehicleMap) FindByWeight(minWeight, maxWeight float64) (v []internal.Vehicle, err error) {
	// for each vehicle in the db, check if the weight is between minWeight and maxWeight
	for _, value := range r.db {
		// if the weight is between minWeight and maxWeight, append the vehicle to the list of vehicles
		if value.Weight >= minWeight && value.Weight <= maxWeight {
			// append the vehicle to the list of vehicles
			v = append(v, value)
		}
	}
	return
}

// FindByBrandAndYearRange is a method that returns a list of vehicles of a specific brand manufactured in a range of years
func (r *VehicleMap) FindByBrandAndYearRange(brand string, startYear, endYear int) (v []internal.Vehicle, err error) {

	for _, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			v = append(v, value)
		}
	}
	return
}
