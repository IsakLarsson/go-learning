package gross

const (
	quarterDozen = 3
	halfDozen    = 6
	dozen        = 12
	small_gross  = 120
	gross        = 144
	great_gross  = 1728
)

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": quarterDozen, "half_of_a_dozen": halfDozen, "dozen": dozen, "small_gross": small_gross, "gross": gross, "great_gross": great_gross,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if !exists {
		return false
	}
	_, billItemExists := bill[item]
	if billItemExists {
		bill[item] += Units()[unit]
	} else {
		bill[item] = Units()[unit]
	}

	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemCount, exists := bill[item]
	ammount, unitExists := units[unit]
	if !exists || !unitExists || itemCount-ammount < 0 {
		return false
	} else if itemCount-ammount == 0 {

		delete(bill, item)
		return true
	}
	bill[item] = itemCount - ammount
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	if !exists {
		return 0, false
	}
	return quantity, true
}
