package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	// FindById is a method that returns a vehicle by id
	FindById(id int) (v Vehicle, err error)

	//FindLast is a method that returns the last vehicle registered
	FindLastId() (id int, err error)

	// Post is a method that registers a vehicle
	CreateVehicle(v Vehicle) (err error)

	// FindByColorAndYear is a method that returns a map of vehicles by color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)

	// FindAverageSpeedByBrand is a method that returns the average speed of vehicles of a specific brand
	FindAverageSpeedByBrand(brand string) (averageSpeed float64, err error)

	// CreateVehicles is a method that registers several vehicles at the same time
	CreateVehicles(v []Vehicle) (err error)

	// UpdateSpeed is a method that updates the maximum speed of a specific vehicle
	UpdateSpeed(id int, speed float64) (err error)

	// FindByFuelType is a method that returns a list of vehicles according to the type of fuel
	FindByFuelType(fuelType string) (v []Vehicle, err error)

	// DeleteVehicle is a method that deletes a vehicle
	DeleteVehicle(id int) (err error)
	
	// FindByTransmissionType is a method that returns a list of vehicles according to their transmission type (manual, automatic, etc.)
	FindByTransmissionType(transmissionType string) (v []Vehicle, err error)

	// UpdateFuel is a method that updates the fuel type of a specific vehicle
	UpdateFuel(id int, fuelType string) (err error)

	// FindByDimensions is a method that returns a list of vehicles according to their dimensions (length, width)
	FindByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v []Vehicle, err error)

	// FindByWeight is a method that returns a list of vehicles according to their weight (minWeight, maxWeight)
	FindByWeight(minWeight, maxWeight float64) (v []Vehicle, err error)

	// FindByBrandAndYearRange is a method that returns a list of vehicles according to their brand and year range (startYear, endYear)
	FindByBrandAndYearRange(brand string, startYear, endYear int) (v []Vehicle, err error)

}