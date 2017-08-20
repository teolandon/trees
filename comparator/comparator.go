package comparator

// Comparator compares two values, returns an int
// to indicate ordering.
type Comparator func(a, b interface{}) int
