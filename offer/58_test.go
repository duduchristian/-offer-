package offer

import (
	"fmt"
	"testing"
)

func TestMyCalendar_Book(t *testing.T) {
	calendar := NewMyCalendar()
	fmt.Println(calendar.Book(10, 20))
	fmt.Println(calendar.Book(15, 25))
	fmt.Println(calendar.Book(20, 30))
}
