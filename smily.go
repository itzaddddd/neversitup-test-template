package main

import "strings"

func countSmileys(listSmile []string) (count int) {
	var eyeElem = map[string]bool{
		":": true,
		";": true,
	}
	var noseElem = map[string]bool{
		"-": true,
		"~": true,
	}
	var mouseElem = map[string]bool{
		")": true,
		"D": true,
	}

	const (
		Eye   = "eye"
		Nose  = "nose"
		Mouse = "mouse"
	)

	isElement := func(elem string, s string) (res bool) {
		var mapElem map[string]bool
		switch elem {
		case Eye:
			mapElem = eyeElem
		case Nose:
			mapElem = noseElem
		case Mouse:
			mapElem = mouseElem
		default:
			return
		}

		if _, ok := mapElem[s]; ok {
			res = true
		}
		return

	}

	for _, smile := range listSmile {
		lenSmile := len(smile)
		smileElems := strings.Split(smile, "")

		if lenSmile == 2 {
			eyeElem, mouseElem := smileElems[0], smileElems[1]
			if isElement(Eye, eyeElem) && isElement(Mouse, mouseElem) {
				count++
			}
		} else {
			eyeElem, noseElem, mouseElem := smileElems[0], smileElems[1], smileElems[2]
			if isElement(Eye, eyeElem) && isElement(Nose, noseElem) && isElement(Mouse, mouseElem) {
				count++
			}

		}
	}
	return
}
