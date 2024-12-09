package field

type Field [100]int

func New(fieldTemplate []int) (*Field, error) {
	if len(fieldTemplate) != 100 {
		return nil, ErrFieldStructError
	}

	var f Field
	for i := 0; i < 100; i++ {
		f[i] = fieldTemplate[i]
	}

	return &f, nil
}

func (f *Field) GetAt(i, j int) (int, error) {
	if i < 0 || j < 0 || i > 9 || j > 9 {
		return 0, ErrOutOfBounds
	}
	return f[i*10+j], nil
}
