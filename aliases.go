package actually

// Actual is an alias of Got.
func Actual(g any) *testingA {
	return Got(g)
}

// Actual is an alias of Got.
func (a *testingA) Actual(g any) *testingA {
	return a.Got(g)
}

// Want is an alias of Expect.
func Want(e any) *testingA {
	return Expect(e)
}

// Want is an alias of Expect.
func (a *testingA) Want(e any) *testingA {
	return a.Expect(e)
}
