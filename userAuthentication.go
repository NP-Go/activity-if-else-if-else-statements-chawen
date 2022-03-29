package main

import "fmt"

func main() {
	value := 12
	if value > 10 && value < 13 {
		fmt.Println("12 bigger then 10, smaller then 13")
	}
}

package main

import "fmt"

func main() {

    if 9%2 == 0 {
        fmt.Println("EVEN CHECK")
    } else {
        fmt.Println("ODD CHECK")
    }
}
/*
NOTES:
+   Adds two operands   A + B gives 30
-   Subtracts second operand from the first A - B gives -10
*   Multiplies both operands    A * B gives 200
/   Divides the numerator by the denominator.   B / A gives 2
%   Modulus operator; gives the remainder after an integer division. LEARN MODULUS? Clock example
++  Increment operator. It increases the integer value by one.  A++ gives 11
--  Decrement operator. It decreases the integer value by one.
&&  Called Logical AND operator. If both the operands are non-zero, then condition becomes true.    (A && B) is false.
||  Called Logical OR Operator. If any of the two operands is non-zero, then condition becomes true.    (A || B) is true.
!   Called Logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true then Logical NOT operator will make false.
*/
