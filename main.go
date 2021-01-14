package main

import (
	"math/rand"
	"time"
)

func main() {

label_run:
	rand.Seed(time.Now().UnixNano())

	DATA_POINTER := 0
	DATA := []string{
		// 3630
		"11HALF-DUG GRAVE", "12GOBLIN GRAVEYARD",
		"11HOLLOW TOMB", "23STALACTITES AND STALAGMITES",
		"11MAZE OF TUNNELS", "11VAULTED CAVERN",
		"23HIGH GLASS GATES", "12ENTRANCE HALL TO THE PALACE",
		"31GRARG SENTRY POST", "12GUARD ROOM",
		"31MARSHY INLET", "23RUSTY GATES",
		"12GAMEKEEPER's COTTAGE", "31MISTY POOL",
		"11HIGH-WALLED GARDEN", "14INSCRIBED CAVERN",
		"34ORNATE FOUNTAIN", "11DANK CORRIDOR",
		"12LONG GALLERY", "12KITCHENS OF THE PALACE",
		"34OLD KILN", "44OVERGROWN TRACK",
		"31DISUSED WATERWHEEL", "33SLUICE GATES",
		"11GAP BETWEEN SOME BOULDERS", "41PERILOUS PATH",
		"31SILVER BELL IN THE ROCK", "12DUNGEONS OF THE PALACE",
		"11BANQUETING HALL", "42PALACE BATTLEMENTS",
		"44ISLAND SHORE", "31BEACHED KETCH",
		"13BARREN COUNTRYSIDE", "33SACKS ON THE UPPER FLOOR",
		"46FROZEN POND", "21MOUNTAIN HUT",
		"31ROW OF CASKS", "11WINE CELLAR",
		"12HALL OF TAPESTRIES", "11DUSTY LIBRARY",
		"13ROUGH WATER", "11PLOUGHED FIELD",
		"55OUTSIDE A WINDMILL", "42LOWER FLOOR OF THE MILL",
		"44ICY PATH", "41SCREE SLOPE",
		"12SILVER CHAMBER", "12WIZARD's LAIR",
		"11MOSAIC-FLOORED HALL", "12SILVER THRONE ROOM",
		"12MIDDLE OF THE LAKE", "42EDGE OF AN ICY LAKE",
		"41PITTED TRACK", "41HIGH PINNACLE",
		"55ABOVE A GALCIER", "21HUGE FALLEN OAK",
		"11TURRET ROOOM WITH A SLOT MACHINE", "11COBWEBBY ROOM",
		"31SAFE IN OGBAN's CHAMBER", "31CUPBOARD IN A CORNER",
		"11NARROW PASSAGE", "16CAVE",
		"11WOODMAN'S HUT", "42SIDE OF A WOODED VALLEY",
		"21STREAM IN A VALLEY BOTTOM", "11DEEP DARK WOOD",
		"11SHADY HOLLOW", "34ANCIENT STONE CIRCLE",
		"16STABLE", "14ATTIC BEDROOM",
		"11DAMP WELL BOTTOM", "32TOP OF A DEEP WELL",
		"31BURNT-OUT CAMPFIRE", "16ORCHARD",
		"62END OF A BRIDGE", "62END OF A BRIDGE",
		"61CROSSROADS", "41WINDING ROAD",
		"11VILLAGE OF RUSTIC HOUSES", "11WHITE COTTAGE",
		// 4030
		"3COINS", "1SHEET", "3BOOTS", "1HORSESHOE", "3APPLES", "1BUCKET", "4AXE", "1BOAT", "1PHIAL",
		"3REEDS", "1BONE", "1SHIELD", "3PLANKS", "1ROPE", "1RING", "1JUG", "1NET", "1SWORD",
		"1SILVER PLATE", "1UNIFORM", "1KEY", "3SEEDS", "1LAMP", "3BREAD", "1BROOCH", "3MATCHES",
		"2STONE OF DESTINY", "4APPLE", "BED", "CUPBOARD", "BRIDGE", "TREES", "SAIL", "KILN",
		"KETCH", "BRICKS", "WINDMILL", "SACKS", "OGBAN's BOAR", "WHEEL",
		"PONY", "GRAVESTONES", "POOL", "GATES", "HANDLE", "HUT", "VINE", "INSCRIPTIONS", "TROLL", "RUBBLE",
		"HOUND", "FOUNTAIN", "CIRCLE", "MOSAICS", "BOOKS", "CASKS", "WELL", "WALLS", "RATS", "SAFE",
		"COBWEBS", "COIN", "BELL", "UP SILVER PLATE", "STONES", "KITCHENS", "GOBLET", "WINE",
		"GRARGS", "DOOR", "AWAKE", "GUIDE", "PROTECT", "LEAD", "HELP", "CHEST", "WATER",
		"STABLES", "SLUICE GATES", "POT", "STATUE", "PINNACLE", "MUSIC", "MAGIC WORDS",
		"MISTY POOL", "WELL BOTTOM", "OLD KILN", "MOUNTAIN HUT",
		// 4140
		"IN", "A", "NEAR", "THE", "BY", "SOME", "ON", "AN", "", "", "AT", "A SMALL",
		"E", "ESW", "WE", "EW", "EW", "ESW", "ESW", "ES", "EW", "SW",
		"S", "N", "ES", "SW", "S", "NW", "N", "N", "ES", "NSW",
		"NS", "E", "NSW", "N", "NES", "EW", "W", "S", "NS", "N",
		"NES", "W", "NS", "D", "NES", "SW", "E", "NW", "NS", "S",
		"NS", "E", "NSEW", "WU", "UD", "NS", "E", "SW", "NSE", "NW",
		"NE", "EW", "NSW", "E", "WN", "S", "E", "NEW", "NW", "S",
		"ES", "SW", "NES", "EW", "SW", "NE", "EW", "ESW", "SW", "ND",
		" ", "E", "NEW", "EW", "NEW", "EW", "EW", "NEW", "NEW", "WU",
		// 4230
		"80", "70", "60", "69", "74", "72", "63", "52", "20", "11", "1", "14", "36", "54", "61", "21", "32", "10", "50",
		"29", "59", "34", "13", "80", "30", "81", "47", "74",
		"1", "2", "3", "4", "5", "9", "12", "13", "16", "17", "20", "21", "22",
	}

	READ := func() string {
		val := DATA[DATA_POINTER]
		DATA_POINTER++
		return val
	}
	READ_NUMERIC := func() int {
		return VAL(READ())
	}
	// 10
	EL, NO, NV, G := 39, 88, 57, 28

	// 3380
	C := make([]int, G+1)
	E_ := make([]string, 81)
	F := make([]int, 71)
	X_ := make([]string, 7)
	Y_ := make([]string, 7)
	G_ := make([]string, 3)
	var D_ string
	GOSUB_3330 := func() {
		DATA_POINTER = 0
		for I := 1; I <= 80; I++ {
			D_ = READ()
		}
	}
	var T_, B_ string
	var X6_, X1_, X2_, X3_, X4_, X5_, X7_, X8_, X9_, XB_ string
	GOSUB_4400 := func() {
		CLS()
		PRINT("")
		TAB(EL/2 - 9)
		PRINT("MYSTERY OF SILVER")
		TAB(EL/2 - 9)
		PRINT("    MOUNTAIN")
		PRINT("========================================")
		PRINT("")
		PRINT("")
	}
	var R int
	var R_ string
	GOSUB_4450 := func() {
		for I := 1; I <= 80; I++ {
			E_[I] = READ()
		}
		for I := 1; I <= G; I++ {
			C[I] = READ_NUMERIC()
		}
		for I := 1; I <= 13; I++ {
			A := READ_NUMERIC()
			F[A] = 1
		}
		// 4480
		F[41] = INT(RND_1()*900) + 100
		F[42] = INT(RND_1()*3) + 2
		F[44] = 4
		F[57] = 68
		F[58] = 54
		F[59] = 15
		F[52] = INT(RND_1() * 3)
		R = 77
		R_ = "GOOD LUCK ON YOUR QUEST!"
		G_[1] = ""
		for I := 1; I < 8; I++ {
			// 4520
			F_ := MID_(B_, 1+INT(RND_1()*4)*3, 1)
			G_[1] += F_
			var L_ string
			if F_ == "N" {
				L_ = "S"
			}
			if F_ == "S" {
				L_ = "N"
			}
			if F_ == "E" {
				L_ = "W"
			}
			if F_ == "W" {
				L_ = "E"
			}
			G_[2] = L_ + G_[2]
		}
	}
	var FL_ string
	GOSUB_4640 := func() {
		PRINT("")
		PRINT("PLEASE ENTER FILE NAME")
		FL_ = INPUT()
	}
	GOSUB_4670 := func() {
		PRINT("OK. SEARCHING FOR " + FL_)
		X := OPENIN(FL_)
		PRINT("OK. LOADING")
		for I := 1; I <= 80; I++ {
			E_[I] = INPUT_DISC(X)
		}
		for I := 1; I <= G; I++ {
			C[I] = INPUT_DISC_NUMERIC(X)
		}
		for I := 1; I <= 70; I++ {
			F[I] = INPUT_DISC_NUMERIC(X)
		}
		G_[1] = INPUT_DISC(X)
		G_[2] = INPUT_DISC(X)
		CLOSE(X)
	}
	GOSUB_4600 := func() {
		GOSUB_4640()
		GOSUB_4670()
		R = F[69]
		R_ = "OK. CARRY ON"
	}
	GOSUB_3380 := func() {
		GOSUB_3330()
		// 3400
		for I := 1; I <= NO; I++ {
			T_ = READ()
		}
		// 3410
		for I := 1; I <= 6; I++ {
			X_[I] = READ()
			Y_[I] = READ()
		}
		// 3420
		B_ = "NOOEOOSOOWOOUOODOOINVGETTAKEXAREAGIVSAYPICWEATIECLIRIGUSEOPE"
		B_ += "LIGFILPLAWATSWIEMPENTCROREMFEETURDIVBAILEATHRINSBLODROEATMOV"
		B_ += "INTRINCUTHOLBURPOISHOUNLWITDRICOUPAYMAKBRESTEGATREF"
		// 3450
		X6_ = "ZPV SFGMFDUFE UIF XJABSET HMBSF! IF JT EFBE"
		X1_ = "THE GHOST OF THE GOBLIN GUARDIAN"
		X2_ = "B MBSHF WJOF HSPXT JO TFDPOET!"
		X3_ = "A GRARG PATROL APPROACHES"
		X4_ = "MAGIC WORDS LIE AT THE CROSSROADS, THE FOUNTAIN AND THE "
		X5_ = "A PILE OF RUBBLE BLOCKS YOUR PATH"
		X7_ = "THE MOUNTAIN RUMBLES!"
		X8_ = "TOWERS FALL DOWN!"
		X9_ = "THE WIZARD HAS YOU IN HIS GLARE"
		XB_ = "HE LEADS YOU"
		// 3550
		GOSUB_4400()
		PRINT("DO YOU WANT TO")
		PRINT("")
		PRINT("   1. START A NEW GAME")
		PRINT("OR 2. CONTINUE A SAVED GAME")
		// 3580
	label_3580:
		PRINT("")
		PRINT("")
		PRINT("TYPE IN EITHER 1 OR 2")
		C := INPUT_NUMERIC()
		if C != 1 && C != 2 {
			goto label_3580
		}
		if C == 1 {
			GOSUB_4450()
		}
		if C == 2 {
			GOSUB_4600()
		}
	}

	GOSUB_3380()
label_30:
	GOSUB_4400()
	LL := 0
	GOSUB_3310 := func() {
		DATA_POINTER = 0
		for I := 1; I <= R; I++ {
			D_ = READ()
		}
	}
	GOSUB_3310()
	// 60
	P_ := X_[VAL(LEFT_(D_, 1))] + " " + Y_[VAL(MID_(D_, 2, 1))] + " "
	J_ := R_ + ". " + "YOU ARE " + P_ + RIGHT_(D_, LEN(D_)-2) + " "
	GOSUB_4830 := func() {
		LS := 1
		LP := 1
		for I := 1; I <= LEN(J_); I++ {
			if MID_(J_, I, 1) == " " && LL > EL {
				PRINT(MID_(J_, LP, LS-LP))
				LL = I - LS
				LP = LS + 1
			}
			if MID_(J_, I, 1) == " " {
				LS = I
			}
			LL = LL + 1
		}
		PRINT_(MID_(J_, LP, LEN(J_)-LP))
	}
	GOSUB_4830()
	GOSUB_3330()
	J_ = ""
	var O_ string
	GOSUB_3350 := func() {
		O_ = RIGHT_(O_, LEN(O_)-1)
	}
	for I := 1; I <= G-1; I++ {
		O_ = READ()
		P_ = Y_[VAL(LEFT_(O_, 1))]
		GOSUB_3350()
		if F[I] == 0 && C[I] == R {
			J_ = J_ + " " + P_ + " " + O_ + ","
		}
	}
	// 140
	if R == 29 && F[48] == 0 {
		J_ = J_ + " GRARGS FEASTING,"
	}
	if R == 29 && F[48] == 1 {
		J_ = J_ + " A SLEEPING GRARG,"
	}
	if R == 12 || R == 22 {
		J_ = J_ + " A PONY,"
	}
	if R == 64 {
		J_ = J_ + "A HERMIT,"
	}
	if R == 18 && E_[18] == "N" {
		J_ = J_ + " AN OAK DOOR,"
	}
	if R == 59 && F[68] == 1 {
		J_ = J_ + " OGBAN (DEAD),"
	}
	// 200
	if J_ != "" {
		J_ = ", YOU CAN SEE" + J_
	}
	J_ = J_ + " AND YOU CAN GO "
	GOSUB_4830()
	PRINT_(" ")
	for I := 1; I <= LEN(E_[R]); I++ {
		PRINT_(MID_(E_[R], I, 1) + ",")
	}
	PRINT("")
	PRINT("")
	R_ = "PARDON?"
	PRINT("========================================")
	PRINT("")
	PRINT("")
	PRINT("WHAT WILL YOU DO NOW ")
	I_ := INPUT()
	if I_ == "SAVE GAME" {
		// GOTO 4630
		F[69] = R
		GOSUB_4640()
		GOSUB_4760 := func() {
			X := OPENOUT(FL_)
			PRINT("OK. SAVING")
			for I := 1; I <= 80; I++ {
				PRINT_DISC(X, E_[I])
			}
			for I := 1; I <= G; I++ {
				PRINT_DISC_NUMERIC(X, C[I])
			}
			for I := 1; I <= 70; I++ {
				PRINT_DISC_NUMERIC(X, F[I])
			}
			PRINT_DISC(X, G_[1])
			PRINT_DISC(X, G_[2])
			CLOSE(X)
		}
		GOSUB_4760()
		PRINT("BYE...")
		STOP(4630)
	}
	// 290
	V_ := ""
	T_ = ""
	VB := 0
	B := 0
	for I := 1; I <= LEN(I_); I++ {
		if MID_(I_, I, 1) == " " && V_ == "" {
			V_ = LEFT_(I_, I-1)
		}
		if MID_(I_, I+1, 1) != " " && V_ != "" {
			T_ = MID_(I_, I+1, LEN(I_)-1)
			I = LEN(I_)
		}
	}
	if T_ == "" {
		V_ = I_
	}
	// 340
label_340:
	if LEN(V_) < 3 {
		V_ = V_ + "O"
		goto label_340
	}
	if V_ == "PLAY" {
		V_ = "BLO"
	}
	U_ := LEFT_(V_, 3)
	for I := 1; I <= NV; I++ {
		if MID_(B_, I*3-2, 3) == U_ {
			VB = I
			I = NV
		}
	}
	F[36] = 0
label_390:
	GOSUB_3330()
	// 400
	for I := 1; I <= NO; I++ {
		O_ = READ()
		if I <= G {
			GOSUB_3350()
		}
		if T_ == O_ {
			B = I
			I = NO
		}
	}
	// 430
	if B == 0 && F[36] == 0 && T_ > "" {
		T_ = T_ + "S"
		F[36] = 1
		goto label_390
	}
	// 440
	if VB == 0 {
		VB = NV + 1
	}
	if T_ == "" {
		R_ = "I NEED TWO WORDS"
	}
	if VB > NV {
		R_ = "TRY SOMETHING ELSE"
	}
	if VB > NV && B == 0 {
		R_ = "YOU CANNOT " + I_
	}
	if B > G || B == 0 {
		goto label_510
	}
	if VB == 8 || VB == 9 || VB == 14 || VB == 17 || VB == 44 || VB > 54 {
		goto label_510
	}
	if VB < NV && C[B] != 0 {
		R_ = "YOU DO NOT HAVE THE " + T_
		goto label_30
	}
	// 510
label_510:
	if R == 56 && F[35] == 0 && VB != 37 && VB != 53 {
		R_ = X1_ + " HAS GOT YOU!"
		goto label_30
	}
	if VB == 44 || VB == 47 || VB == 19 || VB == 57 || VB == 49 {
		goto label_540
	}
	if R == 48 && F[63] == 0 {
		R_ = X9_
		goto label_30
	}
	// 540
label_540:
	var H int
	GOSUB_4260 := func() {
		Z_ := ""
		for I := 1; I <= LEN(R_); I++ {
			C_ := MID_(R_, I, 1)
			if C_ < "A" {
				Z_ = Z_ + C_
				continue // avoids broken goto jump over var declaration
			}
			C := ASC(C_) - 1
			if C == 64 {
				C = 90
			}
			Z_ = Z_ + CHR_(C)
		}
		R_ = Z_
	}
	GOSUB_3360 := func() {
		PRINT("PRESS RETURN TO CONTINUE")
		INPUT()
	}
	GOSUB_2380 := func() {
		if (B == 16 || B == 6) && (R == 41 || R == 51) {
			R_ = "YOU CAPSIZED!"
			F[56] = 1
		}
		if H == 6516 && C[16] == 0 {
			R_ = "IT IS NOW FULL"
			F[34] = 1
		}
		if H == 656 {
			R_ = "IT LEAKS OUT!"
		}
	}
	GOSUB_800 := func() {
		D := VB
		if D == 5 {
			D = 1
		}
		if D == 6 {
			D = 3
		}
		if !((R == 75 && D == 2) || (R == 76 && D == 4)) || F[64] == 1 {
			goto label_850
		}
		R_ = "B USPMM TUPQT ZPV DSPTTJOH"
		GOSUB_4260()
		return
	label_850:
		if F[64] == 1 {
			F[64] = 0
		}
		if F[51] == 1 || F[29] == 1 {
			goto label_900
		}
		if F[55] == 1 {
			F[56] = 1
			R_ = "GRARGS HAVE GOT YOU!"
			return
		}
		if R == 29 && F[48] == 0 {
			R_ = "GRARGS WILL SEE YOU!"
			return
		}
		if R == 73 || R == 42 || R == 9 || R == 10 {
			R_ = X3_
			F[55] = 1
			return
		}
	label_900:
		if C[8] == 0 && ((R == 52 && D == 2) || (R == 31 && D != 3)) {
			R_ = "THE BOAT IS TOO HEAVY"
			return
		}
		if C[8] != 0 && ((R == 52 && D == 4) || (R == 31 && D == 3)) {
			R_ = "YOU CANNOT SWIM"
			return
		}
		if R == 52 && C[8] == 0 && D == 4 && F[30] == 0 {
			R_ = "NO POWER!"
			return
		}
		if R == 41 && D == 3 && F[31] == 0 {
			R_ = "UIF CPBU JT TJOLJOH!"
			GOSUB_4260()
			return
		}
		if R == 33 && D == 1 && F[32] == 0 {
			R_ = "OGBAN'S BOAR BLOCKS YOUR PATH"
			return
		}
		if ((R == 3 && D == 2) || (R == 4 && D == 4)) && F[45] == 0 {
			R_ = X5_
			return
		}
		if R == 35 && C[13] != R {
			R_ = "THE ICE IS BREAKING!"
			return
		}
		if R == 5 && (D == 2 || D == 4) {
			GOSUB_4310 := func() {
				J_ := "SSSSSSSS"
				NG := 0
			label_4320:
				MP := D / 2
				GOSUB_4400()
				PRINT("YOU ARE LOST IN THE")
				PRINT("     TUNNELS")
				PRINT("WHICH WAY? (N,S,W OR E)")
				if NG > 15 {
					PRINT("(OR G TO GIVE UP!)")
				}
				PRINT("")
				W_ := INPUT()
				J_ = RIGHT_(J_+RIGHT_(W_, 1), 8)
				if W_ == "G" {
					F[56] = 1
					return
				}
				if J_ != G_[MP] {
					NG = NG + 1
					goto label_4320
				}
			}
			GOSUB_4310()
		}
		if R == 4 && D == 4 {
			R_ = "PASSAGE IS TOO STEEP"
			return
		}
		if R == 7 && D == 2 && F[46] == 0 {
			R_ = "A HUGE HOUND BARS YOUR WAY"
			return
		}
		if (R == 38 || R == 37) && F[50] == 0 {
			R_ = "JU JT UPP EBSL"
			GOSUB_4260()
			return
		}
		if R == 49 && D == 2 && F[54] == 0 {
			R_ = "MYSTERIOUS FORCES HOLD YOU BACK"
			return
		}
		if R == 49 && D == 3 && F[68] == 0 {
			R_ = "YOU MET OGBAN!!!"
			F[56] = 1
			return
		}
		if R == 38 && F[65] == 0 {
			R_ = "RATS NIBBLE YOUR ANKLES"
			return
		}
		if R == 58 && (D == 1 || D == 4) && F[66] == 0 {
			R_ = "YOU GET CAUGHT IN THE WEBS!"
			return
		}
		if R == 48 && D == 4 && F[70] == 0 {
			R_ = "THE DOOR DOES NOT OPEN"
			return
		}
		if R == 40 && F[47] == 1 {
			F[68] = 1
		}
		if R == 37 && D == 4 && E_[37] == "EW" {
			R = 67
			R_ = "THE PASSAGE WAS STEEP!"
			return
		}
		if R == 29 && D == 3 {
			F[48] = 1
			F[20] = 0
		}
		if R == 8 && D == 2 {
			F[46] = 0
		}
		OM := R
		for I := 1; I <= LEN(E_[R]); I++ {
			K_ := MID_(E_[OM], I, 1)
			if (K_ == "N" || K_ == "U") && D == 1 {
				R = R - 10
			}
			if K_ == "E" && D == 2 {
				R = R + 1
			}
			if (K_ == "S" || K_ == "D") && D == 3 {
				R = R + 10
			}
			if K_ == "W" && D == 4 {
				R = R - 1
			}
		}
		R_ = "OK"
		if R == OM {
			R_ = "YOU CANNOT GO THAT WAY"
		}
		if (OM == 75 && D == 2) || (OM == 76 && D == 4) {
			R_ = "OK. YOU CROSSED"
		}
		if F[29] == 1 {
			F[39] = F[39] + 1
		}
		if F[39] > 5 && F[29] == 1 {
			R_ = "CPPUT IBWF XPSO PVU"
			GOSUB_4260()
			F[29] = 0
			C[3] = 81
		}
	}
	GOSUB_2550 := func() {
		if H == 4337 {
			VB = 2
			GOSUB_800()
		}
	}
	GOSUB_1290 := func() {
		if H == 6577 {
			R_ = "HOW?"
			return
		}
		if H == 4177 || H == 5177 {
			B = 16
			GOSUB_2380()
			return
		}
		if B == 38 {
			R_ = "TOO HEAVY!"
			return
		}
		if B == 4 && F[43] == 0 {
			R_ = "IT IS FIRMLY NAILED ON!"
			return
		}
		CO := 0
		for I := 1; I <= G-1; I++ {
			if C[I] == 0 {
				CO = CO + 1
			}
		}
		if CO > 13 {
			R_ = "YOU CANNOT CARRY ANY MORE"
			return
		}
		if B > G {
			R_ = "YOU CANNOT GET THE " + T_
			return
		}
		if B == 0 {
			return
		}
		if C[B] != R {
			R_ = "IT IS NOT HERE!"
		}
		if F[B] == 1 {
			R_ = "WHAT " + T_ + "?"
		}
		if C[B] == 0 {
			R_ = "YOU ALREADY HAVE IT"
		}
		if C[B] == R && F[B] == 0 {
			C[B] = 0
			R_ = "YOU HAVE THE " + T_
		}
		if B == 28 {
			C[5] = 81
		}
		if B == 5 {
			C[28] = 0
		}
		if C[4] == 0 && C[12] == 0 && C[15] == 0 {
			F[54] = 1
		}
		if B == 8 && F[30] == 1 {
			C[2] = 0
		}
		if B == 2 {
			F[30] = 0
		}
	}
	GOSUB_2870 := func() {
		if B == 10 {
			R_ = "B OJDF UVOF"
			GOSUB_4260()
		}
		if H == 5233 {
			R_ = "WHAT WITH?"
		}
		if B == 83 {
			R_ = "HOW, O MUSICAL ONE?"
		}
		if H == 5610 {
			F[35] = 1
			R_ = X1_ + " IS FREE!"
			E_[56] = "NS"
		}
	}
	GOSUB_1750 := func() {
		if R == 64 {
			R_ = "HE GIVES IT BACK!"
		}
		if H == 6425 {
			GOSUB_3210 := func() {
				R_ = "HE TAKES IT AND SAYS '" + STR_(F[42]) + " RINGS ARE NEEDED'"
				C[25] = 81
			}
			GOSUB_3210()
		}
		if R == 75 || R == 76 {
			R_ = "HE DOES NOT WANT IT"
		}
		if B == 62 && F[44] == 0 {
			R_ = "YOU HAVE RUN OUT!"
		}
		if (H == 7562 || H == 7662) && F[44] > 0 && C[1] == 0 {
			R_ = "HE TAKES IT"
			F[64] = 1
		}
		if F[64] == 1 {
			F[44] = F[44] - 1
			if F[44] == 0 {
				C[1] = 81
			}
		}
		if B == 1 {
			R_ = "HE TAKES THEM ALL!"
			C[1] = 81
			F[64] = 1
			F[44] = 0
		}
		if H == 2228 && C[5] == 81 {
			R_ = XB_ + "NORTH"
			C[28] = 81
			R = 12
		}
		if (H == 2228 && C[5] == 0) || H == 225 {
			R_ = XB_ + "NORTH"
			R = 12
		}
		if (H == 2228 && C[5] == 0) || H == 125 {
			R_ = XB_ + "SOUTH"
			R = 22
		}
		if R == 7 || R == 33 {
			R_ = "HE EATS IT!"
			C[B] = 81
		}
		if H == 711 {
			F[46] = 1
			R_ = "HE IS DISTRACTED"
		}
		if H == 385 || H == 3824 {
			R_ = "THEY SCURRY AWAY"
			C[B] = 81
			F[65] = 1
		}
	}
	GOSUB_2470 := func() {
		if B == 7 || B == 18 {
			R_ = "THWACk!"
		}
		if H == 5818 {
			R_ = "YOU CLEARED THE WEBS"
			F[66] = 1
		}
		if H == 187 {
			R_ = "THE DOOR BROKE!"
			E_[18] = "NS"
			E_[28] = "NS"
		}
		if H == 717 {
			R_ = "YOU BROKE THROUGH"
			E_[71] = "N"
		}
	}
	GOSUB_2730 := func() {
		if B == 0 || B > G {
			return
		}
		C[B] = R
		R_ = "DONE"
		if H == 418 || H == 518 {
			R_ = "YOU DROWNED!"
			F[56] = 1
		}
		if B == 8 && F[30] == 1 {
			C[2] = R
		}
		if B == 16 && F[34] == 1 {
			R_ = "YOU LOST THE WATER!"
			F[34] = 0
		}
		if B == 2 && F[30] == 1 {
			F[30] = 0
		}
	}
	GOSUB_3070 := func() {
		if (H == 4864 || H == 4819) && C[19] == 0 {
			R_ = X6_
			F[63] = 1
			GOSUB_4260()
		}
		if B == 27 {
			GOSUB_1290()
		}
	}
	GOSUB_2120 := func() {
		if H == 522 {
			R_ = "OK"
			F[30] = 1
		}
		if B == 1 || B == 62 || B == 5 || B == 28 || B == 11 || B == 24 {
			GOSUB_1750()
		}
		if H == 416 {
			R_ = "ZPV IBWF LFQU BGMPBU"
			F[31] = 1
			GOSUB_4260()
			return
		}
		if H == 4116 {
			R_ = "IT IS NOT BIG ENOUGH!"
			return
		}
		if B == 18 || B == 7 {
			GOSUB_2470()
		}
		if B == 13 {
			GOSUB_2730()
		}
		if B == 19 {
			GOSUB_3070()
		}
		if B == 10 {
			GOSUB_2870()
		}
		if B == 16 || B == 6 {
			GOSUB_2380()
		}
	}
	GOSUB_1470 := func() {
		R_ = "YOU SEE WHAT YOU MIGHT EXPECT!"
		if B > 0 {
			R_ = "NOTHING SPECIAL"
		}
		if B == 46 || B == 88 {
			GOSUB_2550()
		}
		if H == 8076 {
			R_ = "IT IS EMPTY"
		}
		if H == 8080 {
			R_ = "AHA!"
			F[1] = 0
		}
		if H == 7029 {
			R_ = "OK"
			F[2] = 0
		}
		if B == 20 {
			R_ = "NBUDIFT JO QPDLFU"
			GOSUB_4260()
			C[26] = 0
		}
		if H == 1648 {
			R_ = "THERE ARE SOME LETTERS '" + G_[2] + "'"
		}
		if H == 7432 {
			R_ = "UIFZ BSF BQQMF USFFT"
			GOSUB_4260()
			F[5] = 0
		}
		if H == 2134 || H == 2187 {
			R_ = "OK"
			F[16] = 0
		}
		if B == 35 {
			R_ = "IT IS FISHY!"
			F[17] = 0
		}
		if H == 3438 {
			R_ = "OK"
			F[22] = 0
		}
		if H == 242 {
			R_ = "A FADED INSCRIPTION"
		}
		if (H == 1443 || H == 1485) && F[33] == 0 {
			R_ = "B HMJNNFSJOH GSPN UIF EFQUIT"
			GOSUB_4260()
		}
		if (H == 1443 || H == 1485) && F[33] == 1 {
			R_ = "SOMETHING HERE..."
			F[12] = 0
		}
		if H == 2479 || H == 2444 {
			R_ = "THERE IS A HANDLE"
		}
		if B == 9 {
			R_ = "UIF MBCFM SFBET 'QPJTPO'"
			GOSUB_4260()
		}
		if H == 4055 {
			GOSUB_3290 := func() {
				T := R
				R = F[F[52]+57]
				GOSUB_3310()
				R = T
				R_ = X4_ + RIGHT_(D_, LEN(D_)-2)
			}
			GOSUB_3290()
		}
		if H == 2969 && F[48] == 1 {
			R_ = "VERY UGLY!"
		}
		if H == 7158 || H == 7186 {
			R_ = "THERE ARE LOOSE BRICKS"
		}
		if R == 49 {
			R_ = "VERY INTERESTING!"
		}
		if B == 52 || B == 82 || B == 81 {
			R_ = "INTERESTING!"
		}
		if H == 6978 {
			R_ = "THERE IS A WOODEN DOOR"
		}
		if H == 6970 {
			R_ = "YOU FOUND SOMETHING"
			F[4] = 0
		}
		if H == 2066 {
			R_ = "A LARGE CUPBOARD IN THE CORNER"
		}
		if H == 6865 || H == 6853 {
			R_ = "THERE ARE NINE STONES"
		}
		if H == 248 {
			R_ = "B GBEFE XPSE - 'N S I T'"
			GOSUB_4260()
		}
	}
	GOSUB_2310 := func() {
		if B > G {
			R_ = "IT DOES NOT BURN"
		}
		if B == 26 {
			R_ = "YOU LIT THEM"
		}
		if H == 3826 {
			R_ = "NOT BRIGHT ENOUGH"
		}
		if (B == 23 || H == 6970) && C[26] != 0 {
			R_ = "OP NBUDIFT"
			GOSUB_4260()
		}
		if B == 23 && C[26] == 0 {
			R_ = "A BRIGHT " + V_
			F[50] = 1
		}
		if H == 6970 && C[26] == 0 {
			F[43] = 1
			R_ = "IT HAS TURNED TO ASHES"
		}
	}
	GOSUB_2450 := func() {
		if B == 22 && F[37] == 1 && F[34] == 1 {
			R_ = X2_
			F[38] = 1
			GOSUB_4260()
		}
	}
	GOSUB_2950 := func() {
		if R == 4 && B == 50 {
			F[45] = 1
			R_ = "YOU REVEALED A STEEP PASSAGE"
		}
		if R == 3 && B == 50 {
			R_ = "YOU CANNOT MOVE RUBBLE FROM HERE"
		}
		if H == 7136 {
			R_ = "THEY ARE WEDGED IN"
		}
	}
	GOSUB_2990 := func() {
		if (B == 67 || B == 68) && C[9] == 0 && R == 49 {
			R_ = "OK"
			F[47] = 1
		}
	}
	H = VAL(STR_(R) + STR_(B))
	switch INT(float64(VB-1)/float64(13)) + 1 {
	case 1:
		goto label_560
	case 2:
		goto label_580
	case 3:
		goto label_600
	case 4:
		goto label_620
	case 5:
		goto label_640
	}
label_560:
	switch VB {
	case 1, 2, 3, 4, 5, 6:
		GOSUB_800()
	case 7:
		GOSUB_1220 := func() {
			GOSUB_3330()
			R_ = "OK"
			F[49] = 0
			PRINT_("YOU HAVE ")
			for I := 1; I <= G; I++ {
				O_ = READ()
				GOSUB_3350()
				if I == 1 && C[1] == 0 && F[44] == 1 {
					O_ = "COIN"
				}
				if I == G && C[5] == 0 {
					goto label_1270
				}
				if C[I] == 0 {
					PRINT_(O_ + ",")
					F[49] = 1
				}
			label_1270:
			}
			if F[49] == 0 {
				PRINT("NOTHING")
			}
			PRINT("")
			GOSUB_3360()
		}
		GOSUB_1220()
	case 8, 9:
		GOSUB_1290()
	case 10, 11:
		GOSUB_1470()
	case 12:
		GOSUB_1750()
	case 13:
		GOSUB_1890 := func() {
			R_ = "YOU SAID IT"
			if B == 84 {
				R_ = "YOU MUST SAY THEM ONE BY ONE!"
				return
			}
			if R != 47 || B < 71 || B > 75 || C[27] != 0 {
				return
			}
			if B == 71 && F[60] == 0 {
				R_ = X7_
				F[60] = 1
				return
			}
			if B == 72 && F[60] == 1 && F[61] == 0 {
				R_ = X8_
				F[61] = 1
				return
			}
			if B == (F[52]+73) && F[60] == 1 && F[61] == 1 {
				F[62] = 1
				return
			}
			R_ = "THE WRONG SACRED WORD!"
			F[56] = 1
		}
		GOSUB_1890()
	}
	goto label_650
label_580:
	switch VB - 13 {
	case 1:
		GOSUB_1960 := func() {
			if B == 5 || B == 10 {
				GOSUB_1290()
			}
		}
		GOSUB_1960()
	case 2:
		GOSUB_1980 := func() {
			if B == 3 {
				F[29] = 1
				R_ = "ZPV BSF JOWJTJCMF"
				F[55] = 0
				GOSUB_4260()
			}
			if B == 20 {
				F[51] = 1
				R_ = "ZPV BSF EJTHVJTFE"
				F[55] = 0
				GOSUB_4260()
			}
		}
		GOSUB_1980()
	case 3:
		GOSUB_2010 := func() {
			if B == 2 || B == 14 {
				R_ = "NOTHING TO TIE IT TO!"
				if H == 7214 {
					R_ = "IT IS TIED"
					C[14] = 72
					F[53] = 1
				}
				if H == 722 {
					R_ = "OK"
					F[40] = 1
					C[2] = 72
				}
			}
		}
		GOSUB_2010()
	case 4:
		GOSUB_2050 := func() {
			if H == 1547 && F[38] == 1 {
				R_ = "ALL RIGHT"
				R = 16
			}
			if B == 14 || B == 2 {
				R_ = "NOT ATTACHED TO ANYTHING!"
			}
			if H == 5414 && C[14] == 54 {
				R_ = "YOU ARE AT THE TOP"
			}
			if H == 7214 && F[53] == 1 {
				R_ = "GOING DOWN"
				R = 71
			}
			if H == 722 && F[40] == 1 {
				R = 71
				R_ = "IT IS TORN"
				C[2] = 81
				F[40] = 0
			}
			if H == 7114 && F[53] == 1 {
				C[14] = 71
				F[53] = 0
				R_ = "IT FALLS DOWN-BUMP!"
			}
		}
		GOSUB_2050()
	case 5:
		GOSUB_2870()
	case 6:
		GOSUB_2120()
	case 7:
		GOSUB_2220 := func() {
			if B == 76 || B == 38 {
				GOSUB_1470()
			}
			if H == 2030 {
				F[9] = 0
				R_ = "OK"
			}
			if H == 6030 {
				R_ = "OK"
				F[3] = 0
			}
			if H == 2444 || H == 1870 {
				R_ = "YOU ARE NOT STRONG ENOUGH"
			}
			if H == 3756 {
				R_ = "A PASSAGE!"
				E_[37] = "EW"
			}
			if H == 5960 {
				GOSUB_3260 := func() {
					PRINT("")
					R_ = "XIBU JT UIF DPEF"
					GOSUB_4260()
					PRINT(R_)
					CN := INPUT_NUMERIC()
					R_ = "WRONG!"
					if CN == F[41] {
						R_ = "IT OPENS"
						F[21] = 0
					}
				}
				GOSUB_3260()
			}
			if H == 6970 {
				R_ = "IT FALLS OFF ITS HINGES"
			}
			if H == 4870 {
				R_ = "IT IS LOCKED"
			}
		}
		GOSUB_2220()
	case 8:
		GOSUB_2310()
	case 9:
		GOSUB_2380()
	case 10:
		GOSUB_2420 := func() {
			if B != 22 || R != 15 {
				R_ = "DOES NOT GROW!"
				return
			}
			R_ = "OK"
			F[37] = 1
		}
		GOSUB_2420()
	case 11:
		GOSUB_2450()
	case 12:
		GOSUB_2470()
	case 13:
		GOSUB_2520 := func() {
			if B == 16 {
				B = 22
				GOSUB_2450()
			}
			if H == 499 {
				R_ = "WHERE?"
			}
		}
		GOSUB_2520()
	}
	goto label_650
label_600:
	switch VB - 26 {
	case 1:
		GOSUB_2550()
	case 2:
		GOSUB_2580 := func() {
			if R == 76 {
				VB = 4
				GOSUB_800()
				return
			}
			if R == 75 {
				VB = 2
				GOSUB_800()
			}
		}
		GOSUB_2580()
	case 3:
		GOSUB_2610 := func() {
			if B == 3 && F[29] == 1 {
				R_ = "TAKEN OFF"
				F[29] = 0
			}
			if B == 20 && F[51] == 1 {
				R_ = "OK"
				F[51] = 0
			}
			if B == 36 || B == 50 {
				GOSUB_2950()
			}
		}
		GOSUB_2610()
	case 4:
		GOSUB_2650 := func() {
			if H == 3859 || H == 3339 || H == 1241 || H == 2241 || H == 751 {
				R_ = "WITH WHAT?"
			}
		}
		GOSUB_2650()
	case 5:
		GOSUB_2670 := func() {
			if H == 2340 {
				R_ = "IT GOES ROUND"
			}
			if H == 2445 {
				R_ = "UIF HBUFT PQFO, UIF QPPM FNQUJFT"
				F[33] = 1
				GOSUB_4260()
			}
		}
		GOSUB_2670()
	case 6:
		GOSUB_2700 := func() {
			if R == 14 || R == 51 {
				R_ = "YOU HAVE DROWNED"
				F[56] = 1
			}
		}
		GOSUB_2700()
	case 7:
		GOSUB_2720 := func() {
			R_ = "HOW?"
		}
		GOSUB_2720()
	case 8:
		GOSUB_2730()
	case 9:
		GOSUB_2830 := func() {
			if B == 0 || B > G {
				return
			}
			R_ = "DID NOT GO FAR!"
			C[B] = R
			if H == 3317 {
				R_ = "ZPV DBVHIU UIF CPBS"
				F[32] = 1
				GOSUB_4260()
			}
		}
		GOSUB_2830()
	case 10:
		GOSUB_2800 := func() {
			if B == 62 && F[44] == 0 {
				R_ = "YOU DO NOT HAVE ANY"
			}
			if H == 5762 && C[1] == 0 && F[44] > 0 {
				GOSUB_3230 := func() {
					F[44] = F[44] - 1
					R_ = "A NUMBER APPEARS - " + STR_(F[41])
					if F[44] == 0 {
						C[1] = 81
					}
				}
				GOSUB_3230()
			}
		}
		GOSUB_2800()
	case 11:
		GOSUB_2870()
	case 12:
		GOSUB_2730()
	case 13:
		GOSUB_2920 := func() {
			if B == 0 || B > G {
				return
			}
			if B == 5 || B == 24 {
				R_ = "YUM YUM!"
				C[B] = 81
			}
		}
		GOSUB_2920()
	}
	goto label_650
label_620:
	switch VB - 39 {
	case 1:
		GOSUB_2950()
	case 2:
		GOSUB_2990()
	case 3:
		GOSUB_3010 := func() {
			if R != 27 || B != 63 {
				return
			}
		label_3020:
			PRINT("")
			PRINT("HOW MANY TIMES?")
			MR := INPUT_NUMERIC()
			if MR == 0 {
				PRINT("A NUMBER")
				goto label_3020
			}
			if MR == F[42] {
				R_ = "A ROCK DOOR OPENS"
				E_[27] = "EW"
				return
			}
			R_ = "ZPV IBWF NJTUSFBUFE UIF CFMM!"
			F[56] = 1
			GOSUB_4260()
		}
		GOSUB_3010()
	case 4:
		GOSUB_3050 := func() {
			if H == 5861 {
				H = 5818
				GOSUB_2470()
			}
		}
		GOSUB_3050()
	case 5:
		GOSUB_3070()
	case 6:
		GOSUB_2310()
	case 7:
		GOSUB_2990()
	case 8:
		GOSUB_3070()
	case 9:
		GOSUB_3130 := func() {
			if H == 4870 && C[21] == 0 {
				R_ = "THE KEY TURNS!"
				F[70] = 1
			}
		}
		GOSUB_3130()
	case 10:
		GOSUB_2120()
	case 11:
		GOSUB_3190 := func() {
			R_ = "ARE YOU THIRSTY?"
		}
		GOSUB_3190()
	case 12:
		GOSUB_1470()
	case 13:
		GOSUB_3100 := func() {
			if H == 7549 || H == 7649 {
				R_ = "WHAT WITH?"
			}
			if B == 1 || B == 62 {
				GOSUB_1750()
			}
		}
		GOSUB_3100()
	}
	goto label_650
label_640:
	switch VB - 52 {
	case 1:
		GOSUB_2870()
	case 2:
		GOSUB_3150 := func() {
			if H == 1870 {
				R_ = "HOW?"
			}
		}
		GOSUB_3150()
	case 3, 4:
		GOSUB_1290()
	case 5:
		GOSUB_3170 := func() {
			if R == 48 {
				R_ = "HOW?"
			}
		}
		GOSUB_3170()
	case 6:
		GOSUB_3200 := func() {
			// line 3200 is just a RETURN statement
		}
		GOSUB_3200()
	}
	// 650
label_650:
	if F[62] == 1 {
		goto label_730
	}
	if R == 41 {
		F[67] = F[67] + 1
		if F[67] == 10 {
			F[56] = 1
			R_ = "YOU SANK!"
		}
	}
	if R == 56 && F[35] == 0 && C[10] != 0 {
		R_ = X1_ + " GETS YOU!"
		F[56] = 1
	}
	if F[56] == 0 {
		goto label_30
	}
	// 690
	GOSUB_4400()
	PRINT(R_)
	PRINT("YOU HAVE FAILED IN YOUR QUEST!")
	PRINT("")
	PRINT("BUT YOU ARE GRANTED ANOTHER TRY")
	GOSUB_3360()
	goto label_run
	// 730
label_730:
	GOSUB_4400()
	PRINT("HOOOOORRRRRRAAAAAYYYYYY!")
	PRINT("")
	PRINT("YOU HAVE SUCCEEDED IN YOUR")
	PRINT("QUEST AND BROUGHT PEACE TO")
	PRINT("THE LAND")
	STOP(790)

}
