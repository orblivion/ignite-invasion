package main

func generateAlienName() AlienName {
	// Pick 4 different name segments, order matters, that's 3000 alien names.
	// If we need more, we can append numbers

	// Capitalize the first letter
	_ = []string{
		"goom",
		"kor",
		"mon",
		"zor",
		"xan",
		"blax",
		"thu",
		"blar",
		"yaf",
	}

	a := ""
	return AlienName(&a)
}
