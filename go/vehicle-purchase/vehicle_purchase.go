package purchase

import "sort"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return (kind == "car" || kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	ss := make([]string, 2)
	ss[0] = option1
	ss[1] = option2
	sort.Strings(ss)
	return ss[0] + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	resellPercent := 0.0
	if age < 3.0 {
		resellPercent = .8
	} else if age >= 10.0 {
		resellPercent = .5
	} else if age >= 3.0 && age < 10.0 {
		resellPercent = .7
	}
	return originalPrice * resellPercent
}
