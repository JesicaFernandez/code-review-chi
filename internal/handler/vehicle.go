package handler

import (
	"app/internal"
	"net/http"
	"github.com/bootcamp-go/web/response"
	"app/platform/web/request"
	"errors"
	"strconv"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

type BodyRequestVehicleJSON struct {
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission	string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

type BodyRequestVehicleBatchJSON struct {
	Vehicles []BodyRequestVehicleJSON `json:"vehicles"`
}

type BodyRequestVehicleMaxSpeedJSON struct {
    MaxSpeed float64 `json:"max_speed"`
}

type BodyRequestVehicleFuelTypeJSON struct {
	FuelType string `json:"fuel_type"`
}

type BodyRequestVehicleBrandAndYearRangeJSON struct {
	Brand string `json:"brand"`
	StartYear int `json:"start_year"`
	EndYear int `json:"end_year"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError,  "500 Internal Server Error")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
/* Respuestas:
- 201 Created: Vehículo creado exitosamente.
- 400 Bad Request: Datos del vehículo mal formados o incompletos.
- 409 Conflict: Identificador del vehículo ya existente.*/
func (h *VehicleDefault) Create() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        
		// request
        var body BodyRequestVehicleJSON

		// validate request body is correctly formed
        if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Vehicle data incorrectly formed")
			return
		}

        // process

		vehicles, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			return
		}
		//increment by 1 the last id of the vehicle
		vehiclesCount := len(vehicles)

		// Incrementar el ID en uno
		id := vehiclesCount + 1

        // - create vehicle
		vehicle := internal.Vehicle{
		Id: id, // Set the incremented ID
		// Set the vehicle attributes
		VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		// Validate the vehicle data
		if err := h.sv.ValidateVehicleData(vehicle); err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Vehicle data incorrectly formed")
			return
		}

		// Validate if the vehicle already exists and otherwise create it
		if err := h.sv.CreateVehicle(vehicle); err != nil {
			switch {
				case errors.Is(err, internal.ErrVehicleAlreadyExists):
					response.JSON(w, http.StatusConflict, "409 Conflict")
				default:
					response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
				}
			return
			}

        // response
		// create variable data with the vehicle data in JSON format
        data := VehicleJSON{
            ID:              vehicle.Id,
            Brand:           vehicle.Brand,
            Model:           vehicle.Model,
            Registration:    vehicle.Registration,
            Color:           vehicle.Color,
            FabricationYear: vehicle.FabricationYear,
            Capacity:        vehicle.Capacity,
            MaxSpeed:        vehicle.MaxSpeed,
            FuelType:        vehicle.FuelType,
            Transmission:    vehicle.Transmission,
            Weight:          vehicle.Weight,
            Height:          vehicle.Height,
            Length:          vehicle.Length,
            Width:           vehicle.Width,
        }
		// return the response with the status code 201 and the data in JSON format
        response.JSON(w, http.StatusCreated, map[string]interface{}{
            "message": "201 Created: Vehículo creado exitosamente.",
            "data":    data,
        })
	}
}
			
// GetByColorAndYear is a method that returns a handler for the route GET /vehicles?color={color}&year={year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract color and year from the URL query parameters
		color := chi.URLParam(r, "color")
		yearStr := chi.URLParam(r, "year")

		// Convert yearStr to an integer
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid year parameter")
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			return
		}

		// response
		// create variable data with the vehicle data in JSON format
		data := make(map[int]VehicleJSON)
		// Iterate over the vehicles map and fill the data variable with the vehicle data in JSON format
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		// Verify if the data variable is empty and return a 404 Not Found response if it is empty
		if len(data) == 0 {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos con esos criterios.")
			return
		}
		// return the response with the status code 200 and the data in JSON format
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetAverageSpeedByBrand is a method that returns a handler for the route GET /vehicles/average_speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract brand from the URL query parameters
		brand := chi.URLParam(r, "brand")

		// process
		// calculate the average speed of the vehicles by brand
		averageSpeed, err := h.sv.FindAverageSpeedByBrand(brand)
		// Verify if an error occurred and return a 404 Not Found response if it did
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNoVehiclesWithBrand):
				response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos con esa marca.")
			default:
				response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			}
			return
		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageSpeed,
		})
	}
}

// CreateBatch is a method that returns a handler for the route POST /vehicles/batch
func (h *VehicleDefault) CreateBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body BodyRequestVehicleBatchJSON

		if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid request body")
			return
		}

		// process

		lenVericles, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			return
		}
		//increment by 1 the last id of the vehicle
		vehiclesCount := len(lenVericles)

		// - create vehicles
		vehicles := make([]internal.Vehicle, len(body.Vehicles))
		for key, value := range body.Vehicles {
			
			id := vehiclesCount + key + 1

			// Verificar si el ID ya existe
			/*
			_, err := h.sv.FindById(id)
			if err != nil {
				response.JSON(w, http.StatusConflict, "409 Conflict: Duplicate ID found")
				return
			}*/

			vehicles[key] = internal.Vehicle{
				Id: id, // Set the incremented ID
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           value.Brand,
					Model:           value.Model,
					Registration:    value.Registration,
					Color:           value.Color,
					FabricationYear: value.FabricationYear,
					Capacity:        value.Capacity,
					MaxSpeed:        value.MaxSpeed,
					FuelType:        value.FuelType,
					Transmission:    value.Transmission,
					Dimensions: internal.Dimensions{
						Height: value.Height,
						Length: value.Length,
						Width:  value.Width,
					},
				},
			}
		}

		if err := h.sv.CreateVehicles(vehicles); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.JSON(w, http.StatusConflict, "409 Conflict")
			case errors.Is(err, internal.ErrInvalidVehicle):
				response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid request body")
			default:
				response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusCreated, map[string]interface{}{
			"message": "201 Created: Vehículos creados exitosamente.",
			"data":    data,
		})
	}
}

// UpdateMaxSpeed is a method that returns a handler for the route PUT /vehicles/{id}/update_speed
func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract id from the URL query parameters
		idStr := chi.URLParam(r, "id")

		// Convert idStr to an integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid request body")
			return
		}

		// Extract maxSpeed from the request body
		var body BodyRequestVehicleMaxSpeedJSON

		if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Velocidad mal formada o fuera de rango")
			return
		}

		// process
		// - update max speed
		if err := h.sv.UpdateSpeed(id, body.MaxSpeed); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontró el vehículo")
			case errors.Is(err, internal.ErrInvalidVehicle):
				response.JSON(w, http.StatusBadRequest, "400 Bad Request: Velocidad mal formada o fuera de rango")
			default:
				response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "200 OK: Velocidad del vehículo actualizada exitosamente.",
		})
	}
}

// GetByFuelType is a method that returns a handler for the route GET /vehicles/fuel_type/{type}
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract type from the URL query parameters
		fuelType := chi.URLParam(r, "type")

		// process
		// - get vehicles by fuel type
		vehicles, err := h.sv.FindByFuelType(fuelType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos con ese tipo de combustible.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    data,
		})
	}
}

// Delete is a method that returns a handler for the route DELETE /vehicles/{id}
func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract id from the URL query parameters
		idStr := chi.URLParam(r, "id")

		// Convert idStr to an integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid id parameter")
			return
		}

		// validate id in vehicle list
		if _, err := h.sv.FindById(id); err != nil {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontró el vehículo")
			return
		}

		// process
		// - delete vehicle
		if err := h.sv.DeleteVehicle(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, "404 Not Found")
			default:
				response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, map[string]interface{}{
			"message": "Vehículo eliminado exitosamente.",
		})
	}
}

// GetByTransmissionType is a method that returns a list of vehicles according to their transmission type (manual, automatic, etc.) 
func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract type from the URL query parameters
		transmissionType := chi.URLParam(r, "type")

		// process
		// - get vehicles by transmission type
		vehicles, err := h.sv.FindByTransmissionType(transmissionType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos con ese tipo de transmisión.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    data,
		})
	}
}

// UpdateFuelType is a method that returns a handler for the route PUT /vehicles/{id}/update_fuel
func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract id from the URL query parameters
		idStr := chi.URLParam(r, "id")

		// Convert idStr to an integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusConflict, "409 Conflict: Invalid id parameter")
			return
		}

		// Extract fuelType from the request body
		var body BodyRequestVehicleFuelTypeJSON

		if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Tipo de combustible mal formado o no admitido")
			return
		}

		// process
		// - update fuel type
		if err := h.sv.UpdateFuel(id, body.FuelType); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, "404 Not Found")
			case errors.Is(err, internal.ErrInvalidVehicle):
				response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid request body")
			default:
				response.JSON(w, http.StatusInternalServerError, "500 Internal Server Error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "200 OK: Tipo de combustible del vehículo actualizado exitosamente.",
		})
	}
}

// GetByDimensions is a method that returns a list of vehicles based on a range of dimensions (length, width).
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract min_length from the URL query parameters
		minLengthStr := r.URL.Query().Get("min_length")
		// Extract max_length from the URL query parameters
		maxLengthStr := r.URL.Query().Get("max_length")
		// Extract min_width from the URL query parameters
		minWidthStr := r.URL.Query().Get("min_width")
		// Extract max_width from the URL query parameters
		maxWidthStr := r.URL.Query().Get("max_width")

		// Convert minLengthStr to an integer
		minLength, err := strconv.Atoi(minLengthStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid min_length parameter")
			return
		}
		// Convert maxLengthStr to an integer
		maxLength, err := strconv.Atoi(maxLengthStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid max_length parameter")
			return
		}
		// Convert minWidthStr to an integer
		minWidth, err := strconv.Atoi(minWidthStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid min_width parameter")
			return
		}
		// Convert maxWidthStr to an integer
		maxWidth, err := strconv.Atoi(maxWidthStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid max_width parameter")
			return
		}

		// Convert minLength to a float64
		minLengthFloat64 := float64(minLength)
		// Convert maxLength to a float64
		maxLengthFloat64 := float64(maxLength)
		// Convert minWidth to a float64
		minWidthFloat64 := float64(minWidth)
		// Convert maxWidth to a float64
		maxWidthFloat64 := float64(maxWidth)

		// process
		// - get vehicles by dimensions
		vehicles, err := h.sv.FindByDimensions(minLengthFloat64, maxLengthFloat64, minWidthFloat64, maxWidthFloat64)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos con esas dimensiones.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:		value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByWeight is a method that returns a list of vehicles based on a range of weight.
func (h *VehicleDefault) GetByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// Extract min_weight from the URL query parameters
		minWeight := r.URL.Query().Get("min")
		// Extract max_weight from the URL query parameters
		maxWeight := r.URL.Query().Get("max")

		// convert minWeight to a float64
		minWeightFloat64, err := strconv.ParseFloat(minWeight, 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid min_weight parameter")
			return
		}

		// convert maxWeight to a float64
		maxWeightFloat64, err := strconv.ParseFloat(maxWeight, 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid max_weight parameter")
			return
		}

		// process
		// - get vehicles by weight
		vehicles, err := h.sv.FindByWeight(minWeightFloat64, maxWeightFloat64)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos en ese rango de peso.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:		 value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByBrandAndYear is a method that returns a list of vehicles based on a brand and a range of years.
func (h *VehicleDefault) GetByBrandAndRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// request
		// Extract brand from the URL query parameters
		brand := chi.URLParam(r, "brand")
		// Extract start_year from the URL query parameters
		startYearStr := chi.URLParam(r, "start_year")
		// Extract end_year from the URL query parameters
		endYearStr := chi.URLParam(r, "end_year")

		// convert startYear to an integer
		startYear, err := strconv.Atoi(startYearStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid start_year parameter")
			return
		}

		// convert endYear to an integer
		endYear, err := strconv.Atoi(endYearStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "400 Bad Request: Invalid end_year parameter")
			return
		}


		// process
		// - get vehicles by brand and year
		vehicles, err := h.sv.FindByBrandAndYearRange(brand, startYear, endYear)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos con esos criterios.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:		 value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		// data is empty
		if len(data) == 0 {
			response.JSON(w, http.StatusNotFound, "404 Not Found: No se encontraron vehículos con esos criterios.")
			return
		}
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    data,
		})
	}
}