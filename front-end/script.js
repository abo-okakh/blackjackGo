let BankEL= document.getElementById("player_bank")
let Bet_EL= document.getElementById("bet_amount")
let Bot_Total_EL = document.getElementById("bot_total")
let Player_Total_EL = document.getElementById("player_total")
let Hit_EL = document.getElementById('hit')
let Stand_EL = document.getElementById('stand')

Stand_EL.addEventListener('click',()=>{
    console.log("Stand!")
})
Hit_EL.addEventListener('click',()=>{
    console.log("Hit!")
    console.log("bet is",Bet_EL.value)
})