package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, ok := units[unit]
	if !ok {
		return false
	}

	bill[item] += value
	return ok

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	valueItem, okItem := bill[item]
	if !okItem {
		return false
	}
	valueUnit, okUnit := units[unit]

	if !okUnit {
		return false
	}

	if valueItem-valueUnit < 0 {
		return false
	}

	bill[item] -= valueUnit

	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, ok := bill[item]

	if !ok {
		return 0, false
	}

	return value, true
}
