package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"dozen":              12,
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	fmt.Println("input", bill, units, item, unit)
	if _, ok := units[unit]; !ok {

		fmt.Println("res false", bill, units, item, unit)
		return false
	} else {
		bill[item] += units[item]
	}
	if _, ok := bill[item]; ok {
		bill[item] += units[unit]
	}
	fmt.Println("res true", bill, units, item, unit)
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	fmt.Println("input", bill, units, item, unit)
	if _, ok := bill[item]; !ok {
		fmt.Println("res false", bill, units, item, unit)
		return false
	}
	if _, ok := units[unit]; !ok {
		fmt.Println("res false", bill, units, item, unit)
		return false
	}
	// bill[item] -= units[unit]
	newq := bill[item] - units[unit]
	if newq == 0 {
		delete(bill, item)
		fmt.Println("res true", bill, units, item, unit)
		return true
	} else if newq < 0 {
		fmt.Println("res false", bill, units, item, unit)
		return false
	} else {
		bill[item] = newq
		fmt.Println("res true", bill, units, item, unit)
		return true
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if val, ok := bill[item]; ok {
		return val, true
	} else {
		return 0, false
	}
}
