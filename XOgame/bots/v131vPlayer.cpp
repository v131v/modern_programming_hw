#pragma once
#include "../common.hpp"
#include "../player.cpp"
#include "../game.cpp"
#include <string>

using namespace std;

class v131vPlayer : public Player {
	private:
	string name;
	public:
	v131vPlayer() : name("v131v") {}

	string getName() override {
		return name;
	}

	Turn makeTurn(X0 playerKind, const X0** field) override {
		if (field[1][1] == X0::undefined) return Turn(1, 1);

		int turnNum = 0;
		int fieldSize = 3;

		X0** model = new X0 * [fieldSize];
		for (int i = 0; i < fieldSize; i++) {
			model[i] = new X0[fieldSize]{};

			for (int j = 0; j < fieldSize; j++) {
				model[i][j] = field[i][j];
				if (model[i][j] != X0::undefined) {
					turnNum++;
				}
			}
		}

		if (turnNum < 3) {
			for (int i = 0; i < 3; i += 2) {
				for (int j = 0; j < 3; j += 2) {
					if (field[i][j] == -playerKind) {
						return Turn(2 - i, 2 - j);
					}
				}
			}

			return Turn(0, 0);
		}

		for (int i = 0; i < 3; i++) {
			for (int j = 0; j < 3; j++) {
				if (model[i][j] == X0::undefined) {
					model[i][j] = playerKind;

					if (checkWin(playerKind, model, fieldSize)) {
						return Turn(i, j);
					}

					model[i][j] = X0::undefined;
				}
			}
		}

		for (int i = 0; i < 3; i++) {
			for (int j = 0; j < 3; j++) {
				if (model[i][j] == X0::undefined) {

					model[i][j] = X0(-playerKind);

					if (checkWin(X0(-playerKind), model, fieldSize)) {
						return Turn(i, j);
					}

					model[i][j] = X0::undefined;
				}
			}
		}

		for (int x1 = 0; x1 < 3; x1++) {
			for (int y1 = 0; y1 < 3; y1++) {
				if (model[x1][y1] == X0::undefined) {
					model[x1][y1] = playerKind;

					int cnt = 0;
					for (int i = 0; i < 3; i++) {
						for (int j = 0; j < 3; j++) {
							if (model[i][j] == X0::undefined) {
								model[i][j] = playerKind;

								if (checkWin(playerKind, model, fieldSize)) {
									cnt++;
								}

								model[i][j] = X0::undefined;
							}
						}
					}

					if (cnt > 1) {
						return Turn(x1, y1);
					}
				}
			}
		}

		for (int i = 0; i < 3; i++) {
			for (int j = 0; j < 3; j++) {
				if (field[i][j] == X0::undefined) return Turn(i, j);
			}
		}
	}

	bool checkWin(X0 kindOf, X0** field, int fieldSize) {
		for (int index = 0; index < fieldSize; index++) {
			bool win = true;

			//horizontal
			for (int jndex = 0; jndex < fieldSize; jndex++) {
				if (field[index][jndex] != kindOf) {
					win = false;
					break;
				}
			}
			if (win) return true;
			win = true;

			// vertical
			for (int jndex = 0; jndex < fieldSize; jndex++) {
				if (field[jndex][index] != kindOf) {
					win = false;
					break;
				}
			}
			if (win) return true;
			win = true;
		}

		// left-to-right
		bool win = true;
		for (int index = 0; index < fieldSize; index++) {
			if (field[index][index] != kindOf) {
				win = false;
				break;
			}
		}
		if (win) return true;
		win = true;

		// right-to-left
		for (int index = 0; index < fieldSize; index++) {
			if (field[index][fieldSize - index - 1] != kindOf) {
				win = false;
				break;
			}
		}
		return win;
	}

};