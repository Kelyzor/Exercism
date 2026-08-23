package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	quantity := map[string]int {}
    quantity["quarter_of_a_dozen"] = 3
    quantity["half_of_a_dozen"] = 6
    quantity["dozen"] = 12
    quantity["small_gross"] = 120
    quantity["gross"] = 144
    quantity["great_gross"] = 1728
    return quantity
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int {}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
    
    if !exists {
        return false
    }

    bill[item] += value
    return true
    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	valueBill, existsBill := bill[item]
    valueUnits, existsUnits := units[unit]
    if !existsBill || !existsUnits || valueBill - valueUnits < 0 {
        return false
    }
    bill[item] -= valueUnits;
    if bill[item] == 0 {
        delete(bill, item)
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
    if !exists {
        return 0, false
    }
    return value, true
}
