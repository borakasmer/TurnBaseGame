package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"turnBaseGame/druid"
	IHero "turnBaseGame/fightHero"
	"turnBaseGame/fighter"
	"turnBaseGame/wizard"
)

func fight(attacker IHero.FightHero, target IHero.FightHero) {
	var _wizardAttacker, wizardOk = attacker.(*wizard.Wizard)
	var _fighterAttacker, fighterOk = attacker.(*fighter.Fighter)
	var _druidAttacker, druidOk = attacker.(*druid.Druid)

	damageValue := attacker.Hit()
	target.TakeDamage(damageValue)

	var attackerMana, attackerStamina int
	if wizardOk {
		_wizardAttacker.SpendMana()
		attackerMana = _wizardAttacker.GetMana()
	} else if fighterOk {
		_fighterAttacker.SpendStamina()
		attackerStamina = _fighterAttacker.GetStamina()
	} else if druidOk {
		_druidAttacker.SpendMana()
		attackerMana = _druidAttacker.GetMana()
	}

	/*
		switch attacker.(type) {
		case *fighter.Fighter:
			_fighterAttacker.SpendStamina()
			attackerStamina = _fighterAttacker.GetStamina()
		case *wizard.Wizard:
			_wizardAttacker.SpendMana()
			attackerMana = _wizardAttacker.GetMana()
		case *druid.Druid:
			_druidAttacker.SpendMana()
			attackerMana = _druidAttacker.GetMana()
	*/

	nameAttacker, _ := attacker.GetInfo()
	nameTarget, bloodTarget := target.GetInfo()

	if wizardOk {
		fmt.Println(
			"(lvl-"+strconv.Itoa(_wizardAttacker.Level)+")"+nameAttacker,
			"'"+_wizardAttacker.Magic+"'",
			"ile",
			"saldırdı => ",
			nameTarget+"("+strconv.Itoa(bloodTarget+damageValue)+")'ya",
			damageValue,
			"hasar verdi.",
		)
		if damageValue > 0 {
			fmt.Println(
				nameAttacker,
				attackerMana,
				"manası kaldı.",
			)
		} else {
			fmt.Println(
				nameAttacker,
				"Bir sonraki Turn Manası yenilenecektir.",
			)
		}
		fmt.Println(
			nameTarget+"'nın",
			"kalan canı ",
			bloodTarget,
		)
		if damageValue == 0 {
			rand.Seed(time.Now().UnixNano())
			rndMana := rand.Intn(100-50) + 50 + 1
			_wizardAttacker.Mana = rndMana
		}
	} else if fighterOk {
		fmt.Println(
			"(lvl-"+strconv.Itoa(_fighterAttacker.Level)+")"+nameAttacker,
			"'"+_fighterAttacker.Weapon+"'",
			"ile",
			"saldırdı =>",
			nameTarget+"("+strconv.Itoa(bloodTarget+damageValue)+")'ya",
			damageValue,
			"hasar verdi.",
		)
		if damageValue > 0 {
			fmt.Println(
				nameAttacker,
				attackerStamina,
				"'ın staminası kaldı.",
			)
		} else {
			fmt.Println(
				nameAttacker,
				"Bir sonraki Turn Staminası yenilenecektir.",
			)
		}
		fmt.Println(
			nameTarget+"'nın",
			"kalan canı ",
			bloodTarget,
		)
		if damageValue == 0 {
			rand.Seed(time.Now().UnixNano())
			rndStamina := rand.Intn(100-50) + 50 + 1
			_fighterAttacker.Stamina = rndStamina
		}
	} else if druidOk {
		fmt.Println(
			"(lvl-"+strconv.Itoa(_druidAttacker.Level)+")"+nameAttacker,
			"'"+_druidAttacker.Animal+"'",
			"ile",
			"saldırdı => ",
			nameTarget+"("+strconv.Itoa(bloodTarget+damageValue)+")'ya",
			damageValue,
			"hasar verdi.",
		)
		if damageValue > 0 {
			fmt.Println(
				nameAttacker,
				attackerMana,
				"manası kaldı.",
			)
		} else {
			fmt.Println(
				nameAttacker,
				"Bir sonraki Turn Manası yenilenecektir.",
			)
		}
		fmt.Println(
			nameTarget+"'nın",
			"kalan canı ",
			bloodTarget,
		)
		if damageValue == 0 {
			rand.Seed(time.Now().UnixNano())
			rndMana := rand.Intn(100-50) + 50 + 1
			_druidAttacker.Mana = rndMana
		}
	}
}

func turnFights(turn int, wizardChan chan *wizard.Wizard, fighterChan chan *fighter.Fighter, druidChan chan *druid.Druid, wg *sync.WaitGroup, lockObject *sync.Mutex) {
	select {
	case wiz, ok := <-wizardChan:
		if ok && wiz.Blood > 0 { //Ölenle ölünmez...		
			wiz.Magic = MapRandomKeyGet(wiz.Manas).(string)
			z := strings.Repeat("#", 100)
			fmt.Println(z)

			//SelectRandom Victom
			var rndVictom int
			for {
				rndVictom = GetRandomID(3)
				//If selected Victom different from the Fighter
				var _, wizardOk = fighterNumberList[rndVictom].(*wizard.Wizard)
				if !wizardOk {
					_, blood := fighterNumberList[rndVictom].GetInfo()
					if blood > 0 {
						break
					}
				}
			}
			randomVictom := fighterNumberList[rndVictom]
			//We Selected Random Victom Different From the Wizard

			fmt.Printf("TURN %d \n", turn)
			//fight(&wiz, fighterList["fighter"])

			//fmt.Printf(wiz.Name)
			//fight(fighterList["wizard"], randomVictom)
			lockObject.Lock()
			fight(wiz, randomVictom)
			lockObject.Unlock()
		} else {
			z := strings.Repeat("#", 100)
			fmt.Println(z)
			fmt.Printf("TURN %d Wizard is Dead..\n", turn)
		}
		var _druidAttack, _ = fighterList["druid"].(*druid.Druid)
		druidChan <- _druidAttack

	case fig, ok := <-fighterChan:
		if ok && fig.Blood > 0 { //If fighter is not dead
			fig.Weapon = MapRandomKeyGet(fig.Staminas).(string)
			//SelectRandom Victom
			var rndVictom int
			for {
				rndVictom = GetRandomID(3)
				//If selected Victom different from the Fighter
				var _, fighterOk = fighterNumberList[rndVictom].(*fighter.Fighter)
				if !fighterOk {
					_, blood := fighterNumberList[rndVictom].GetInfo()
					if blood > 0 {
						break
					}
				}
			}
			randomVictom := fighterNumberList[rndVictom]
			//We Selected Random Victom Different From the Fighter

			fmt.Printf("TURN %d \n", turn)
			//fight(&fig, fighterList["wizard"])

			//fmt.Printf(fig.Name)
			//fight(fighterList["fighter"], randomVictom)
			lockObject.Lock()
			fight(fig, randomVictom)
			lockObject.Unlock()

			/*fmt.Printf("STAMINA Fig : %d \n", fig.Stamina)
			var _fighterAttacker, _ = fighterList["fighter"].(*fighter.Fighter)
			stamina := _fighterAttacker.Stamina
			fmt.Printf("STAMINA LİST: %d \n", stamina)*/
		} else {
			z := strings.Repeat("#", 100)
			fmt.Println(z)
			fmt.Printf("TURN %d Fighter is Dead..\n", turn)
		}
		var _wizardAttack, _ = fighterList["wizard"].(*wizard.Wizard)
		wizardChan <- _wizardAttack

	case dru, ok := <-druidChan:
		if ok && dru.Blood > 0 { //If druid is not dead
			dru.Animal = MapRandomKeyGet(dru.Manas).(string)
			z := strings.Repeat("#", 100)
			fmt.Println(z)

			//SelectRandom Victom
			var rndVictom int
			for {
				rndVictom = GetRandomID(3)
				//If selected Victom different from the Druid
				var _, druidOk = fighterNumberList[rndVictom].(*druid.Druid)
				if !druidOk {
					_, blood := fighterNumberList[rndVictom].GetInfo()
					if blood > 0 {
						break
					}
				}
			}
			randomVictom := fighterNumberList[rndVictom]
			//We Selected Random Victom Different From the Druid

			fmt.Printf("TURN %d \n", turn)
			//fight(&dru, fighterList["fighter"])

			//fmt.Printf(dru.Name)
			//fight(fighterList["druid"], randomVictom)
			lockObject.Lock()
			fight(dru, randomVictom)
			lockObject.Unlock()
		} else {
			z := strings.Repeat("#", 100)
			fmt.Println(z)
			fmt.Printf("TURN %d Druid is Dead..\n", turn)
		}
	default:
		var _fighterAttack, _ = fighterList["fighter"].(*fighter.Fighter)

		//fmt.Printf("STAMINA Fighter Default : %d \n", _fighterAttack.Stamina)
		fighterChan <- _fighterAttack

		//fmt.Printf("CAN : %d \n",_fighterAttack.Blood)
		//fmt.Printf("STAMINA : %d \n",_fighterAttack.Stamina)
	}
	wg.Done()
}

var fighterList map[string]IHero.FightHero
var fighterNumberList map[int]IHero.FightHero

func GetRandomID(limit int) int {
	rand.Seed(time.Now().UnixNano())
	rndVictom := rand.Intn(limit) + 1
	//fmt.Printf("Random : %d \n", rndVictom)
	return rndVictom
}
func GetRandomBetweenID(minLimit int, maxlimit int) int {
	rand.Seed(time.Now().UnixNano())
	rndVictom := rand.Intn(maxlimit-minLimit) + minLimit + 1
	//fmt.Printf("Random : %d \n", rndVictom)
	return rndVictom
}

/*func GetRandomBloodMana() int {
	rand.Seed(time.Now().UnixNano())
	rndVictom := rand.Intn(100)
	return rndVictom
}*/

var isFighterDead bool = false
var isWizardDead bool = false
var isDruidDead bool = false

var TotalLive int = 3

func main() {

	merlin := wizard.CreateWizard(GetRandomBetweenID(50, 100), GetRandomBetweenID(50, 100))
	merlin.Name = "Bora"
	//merlin.Level = 25
	merlin.Level = GetRandomID(60)
	//merlin.Magic = MapRandomKeyGet(merlin.Manas).(string)

	barbar := fighter.CreateFighter(GetRandomBetweenID(50, 100), GetRandomBetweenID(50, 100))
	barbar.Name = "Barbar"
	//barbar.Level = 40
	barbar.Level = GetRandomID(60)
	barbar.Stamina = 20
	//barbar.Weapon = MapRandomKeyGet(barbar.Staminas).(string)

	forest := druid.CreateDruid(GetRandomBetweenID(50, 100), GetRandomBetweenID(50, 100))
	forest.Name = "Forest"
	//forest.Level = 60
	forest.Level = GetRandomID(60)
	//forest.Animal = MapRandomKeyGet(forest.Manas).(string)

	runtime.GOMAXPROCS(1)
	fighterChan := make(chan *fighter.Fighter)
	wizardChan := make(chan *wizard.Wizard)
	druidChan := make(chan *druid.Druid)

	fighterList = map[string]IHero.FightHero{"fighter": barbar, "wizard": merlin, "druid": forest}
	fighterNumberList = map[int]IHero.FightHero{1: barbar, 2: merlin, 3: forest}

	/*for i := 1; i < 5; i++ {
		go turnFights(i, wizardChan, fighterChan, druidChan)
	}*/

	/*var isFighterDead bool = false
	var isWizardDead bool = false
	var isDruidDead bool = false*/

	var stage int = 0

	wg := &sync.WaitGroup{}
	lockObject := &sync.Mutex{}
	//for !isFighterDead && !isWizardDead && !isDruidDead {
	for TotalLive > 1 { //Ölümüne savaş :)

		z := strings.Repeat("-", 100)
		fmt.Println(z)
		stage++
		fmt.Printf("STAGE %d \n", stage)
		fmt.Println(z)

		//loopTurnFights(wizardChan, fighterChan, druidChan)
		for i := 1; i < 5; i++ {
			wg.Add(1)
			go turnFights(i, wizardChan, fighterChan, druidChan, wg, lockObject)
		}

		<-time.After(time.Millisecond * 100)

		//Check Who is Standing!
		TotalLive = 0
		lockObject.Lock()
		if !fighterList["fighter"].IsDeath() {
			TotalLive += 1
		}
		if !fighterList["wizard"].IsDeath() {
			TotalLive += 1
		}
		if !fighterList["druid"].IsDeath() {
			TotalLive += 1
		}
		lockObject.Unlock()

		/*isFighterDead = fighterList["fighter"].IsDeath()
		isWizardDead = fighterList["wizard"].IsDeath()
		isDruidDead = fighterList["druid"].IsDeath()*/
	}
	wg.Wait()
	close(fighterChan)
	close(wizardChan)
	close(druidChan)
	//<-time.After(time.Second * 5)

	/*
		fmt.Println("TURN 1")
		fight(barbar, merlin)

		y := strings.Repeat("#", 100)
		fmt.Println(y)

		fmt.Println("TURN 2")
		fight(merlin, barbar)

		z := strings.Repeat("#", 100)
		fmt.Println(z)

		fmt.Println("TURN 3")
		fight(forest, barbar)
	*/
}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}

/*func loopTurnFights(wizardChan chan wizard.Wizard, fighterChan chan fighter.Fighter, druidChan chan druid.Druid) {
	for i := 1; i < 5; i++ {
		go turnFights(i, wizardChan, fighterChan, druidChan)
	}
}*/
