package main

// TODO: Server Website over http , make a simple get API ..
func main() {
	// WebPage := http.FileServer(http.Dir("./front-end"))
	// http.Handle("/", WebPage)
	// http.ListenAndServe(":8080", nil)
	// // API's
}

var GameState struct {
	PlayerPocket int
}

// TODO:Get a random card from deck.js
// func GetCard(w http.ResponseWriter,r *http.Request){

// }
//TODO:Check who won
// func mid_check(w http.responsewriter,r *http.request){
// 	this function checks who won mid game aka after a hit
// }
//TODO:Check who won after game .. aka who is closer ?
// func Check(w http.ResponseWriter,r* http.Request){
// 	// who is closer to 21 wins  !
// if player wins bet*2
// if player loses PlayerBank-bet
// }
//TODO:this might be used as a authenticator in the future if the project get's bigger
// func SetGame(){
//Send Player there Bank balance
// }
//TODO:this function uses bunch of other functions to start the game, and it give out 2 cards to both player,bot
// func Start(){

// }
type Card struct {
	Name  string `json:"name"`
	Suit  string `json:"suit"`
	Value int    `json:"value"`
}
