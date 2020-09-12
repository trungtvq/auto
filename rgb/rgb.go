package rgb

import "strconv"

type Hex string

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

const rangeColor = 30

func (r *RGB) Is(color *RGB) bool {
	if color.Blue < r.Blue+rangeColor && color.Blue+rangeColor > r.Blue {
		if color.Red < r.Red+rangeColor && color.Red+rangeColor > r.Red {
			if color.Green < r.Green+rangeColor && color.Green+rangeColor > r.Green {
				return true
			}
		}
	}
	return false
}

func (h Hex) toRGB() (RGB, error) {
	return Hex2RGB(h)
}

func Hex2RGB(hex Hex) (RGB, error) {
	var rgb RGB
	values, err := strconv.ParseUint(string(hex), 16, 32)

	if err != nil {
		return RGB{}, err
	}

	rgb = RGB{
		Red:   uint8(values >> 16),
		Green: uint8((values >> 8) & 0xFF),
		Blue:  uint8(values & 0xFF),
	}

	return rgb, nil
}
