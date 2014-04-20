package fizzbuzz

import "strconv"

func DivisibleByThree(me int) bool {
    if me % 3 == 0 {
        return true
    }

    return false
}

func DivisibleByFive(me int) bool {
    if me  % 5 == 0 {
        return true
    }

    return false
}

func DivisibleByBoth(me int) bool {
    return DivisibleByThree(me) && DivisibleByFive(me)
}

func FizzBuzz(me int) string {
    if DivisibleByBoth(me) {
        return "FizzBuzz!"
    }

    if DivisibleByThree(me) {
        return "Fizz!"
    }

    if DivisibleByFive(me) {
        return "Buzz!"
    }

    return strconv.Itoa(me)
}