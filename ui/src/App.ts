
import Game from "./components/Game";
import {Round} from "./components/Round"

let game: Game = new Game();
game.connect()

let round: Round = new Round();
round.increment();