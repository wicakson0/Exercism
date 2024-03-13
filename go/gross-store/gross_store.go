package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitStore := make(map[string]int)

	unitStore["quarter_of_a_dozen"] = 3
	unitStore["half_of_a_dozen"] = 6
	unitStore["dozen"] = 12
	unitStore["small_gross"] = 120
	unitStore["gross"] = 144
	unitStore["great_gross"] = 1728

	return unitStore
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exist := units[unit]

	if exist {
		bill[item] += value
		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExist := bill[item]
	if !itemExist {
		return false
	}

	unitValue, unitExist := units[unit]
	if !unitExist {
		return false
	}

	if bill[item] < unitValue {
		return false
	}

	bill[item] -= unitValue

	if bill[item] == 0 {
		delete(bill, item)
		return true
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemValue, itemExist := bill[item]

	if !itemExist {
		return 0, false
	}

	return itemValue, itemExist
}
