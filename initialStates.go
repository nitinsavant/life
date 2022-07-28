package main

func (u universe) loadBlinker() {
	u.state[1][2] = 1
	u.state[1][3] = 1
	u.state[1][4] = 1
	return
}

func (u universe) loadToad() {
	u.state[2][2] = 1
	u.state[2][3] = 1
	u.state[2][4] = 1
	u.state[3][1] = 1
	u.state[3][2] = 1
	u.state[3][3] = 1
	return
}

func (u universe) loadBeacon() {
	u.state[1][1] = 1
	u.state[1][2] = 1
	u.state[2][1] = 1
	u.state[3][4] = 1
	u.state[4][3] = 1
	u.state[4][4] = 1
	return
}

func (u universe) loadGlider() {
	u.state[1][2] = 1
	u.state[2][3] = 1
	u.state[3][1] = 1
	u.state[3][2] = 1
	u.state[3][3] = 1
	return
}

func (u universe) loadGliderGun() {
	u.state[1][25] = 1

	u.state[2][23] = 1
	u.state[2][25] = 1

	u.state[3][13] = 1
	u.state[3][14] = 1
	u.state[3][21] = 1
	u.state[3][22] = 1
	u.state[3][35] = 1
	u.state[3][36] = 1

	u.state[4][12] = 1
	u.state[4][16] = 1
	u.state[4][21] = 1
	u.state[4][22] = 1
	u.state[4][35] = 1
	u.state[4][36] = 1

	u.state[5][1] = 1
	u.state[5][2] = 1
	u.state[5][11] = 1
	u.state[5][17] = 1
	u.state[5][21] = 1
	u.state[5][22] = 1

	u.state[6][1] = 1
	u.state[6][2] = 1
	u.state[6][11] = 1
	u.state[6][15] = 1
	u.state[6][17] = 1
	u.state[6][18] = 1
	u.state[6][23] = 1
	u.state[6][25] = 1

	u.state[7][11] = 1
	u.state[7][17] = 1
	u.state[7][25] = 1

	u.state[8][12] = 1
	u.state[8][16] = 1

	u.state[9][13] = 1
	u.state[9][14] = 1
	return
}

