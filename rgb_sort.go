package main

// RGBSort the most direct way of sorting color,
// compares it's  r g b alpha values respectively
func RGBSort(a, b Colour) bool {
	return a.Red > b.Red || a.Green > b.Green || a.Blue > b.Blue || a.Alpha > b.Alpha
}
