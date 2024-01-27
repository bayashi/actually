package actually

// Actual is an alias of Got.
func Actual(g any) *TestingA {
	return Got(g)
}

// Actual is an alias of Got.
func (a *TestingA) Actual(g any) *TestingA {
	return a.Got(g)
}

// Want is an alias of Expect.
func Want(e any) *TestingA {
	return Expect(e)
}

// Want is an alias of Expect.
func (a *TestingA) Want(e any) *TestingA {
	return a.Expect(e)
}
