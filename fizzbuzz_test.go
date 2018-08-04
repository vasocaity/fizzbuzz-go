package fizzbuzz

import "testing"

func TestFizzBizzOneSholudBeOne(t *testing.T) {
	v := fizzbuzz(1)
	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' bot have", v)
	}
}

func TestFizzBizzTwoShouldBeTwo(t *testing.T) {
	v := fizzbuzz(2)

	if "2" != v {
		t.Error("fizzbuzz of 2 should be '2' bot have", v)
	}
}

func TestFizzBizzThreeShouldBeFizz(t *testing.T) {
	v := fizzbuzz(3)

	if "fizz" != v {
		t.Error("fizzbuzz of 3 should be 'fizz' bot have", v)
	}
}

func TestFizzBizzFourShouldBeFour(t *testing.T) {
	v := fizzbuzz(4)

	if "4" != v {
		t.Error("fizzbuzz of 4 should be 'four' bot have", v)
	}
}

func TestFizzBizzFiveShouldBeBuzz(t *testing.T) {
	v := fizzbuzz(5)

	if "buzz" != v {
		t.Error("fizzbuzz of 5 should be 'buzz' bot have", v)
	}
}

func TestFizzBuzzShouldBeFizzBuzz(t *testing.T) {
	v := fizzbuzz(15)
	if "fizzbuzz" != v {
		t.Error("fizzbuzz of 15 should be 'fizzbuzz' bot have", v)
	}
	vari := fizzbuzz(30)
	if "fizzbuzz" != vari {
		t.Error("fizzbuzz of 30 should be 'fizzbuzz' bot have", v)
	}
}
