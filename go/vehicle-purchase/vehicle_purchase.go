package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var needLicense bool

	if kind == "car" {
		needLicense = true
	}
	if kind == "bike" {
		needLicense = false
	}
	if kind == "truck" {
		needLicense = true
	}
	return needLicense
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}
	return option2 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64

	if age < 3 {
		price = 0.8 * originalPrice
	}
	if age >= 10 {
		price = 0.5 * originalPrice
	}
	if (age >= 3) && (age < 10) {
		price = 0.7 * originalPrice
	}

	return price
}
