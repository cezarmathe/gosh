package prompt

// Item represents how a prompt item should behave
type Item interface {
	Generate() Item
	PrintPrefix()
	Print()
}
