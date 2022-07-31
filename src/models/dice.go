package models

import (
	"math/rand"
	"sort"
)

const (
	Skull  = "skull"
	Blank  = "blank"
	Shield = "shield"
	Hit    = "hit"
	Wild   = "wild"
	Crit   = "crit"
)

var faces map[int]string

// Look, don't use init functions in actual programming. Running code on import is terrible in
// library construction, and you shoud just call this explicitly. However, since this is just a
// toy project I'm going to do it anyways, you aren't my dad
func init() {
	faces = make(map[int]string)

	// Also this is shitty but I hate Go's implementation of enums
	faces[0] = Skull
	faces[1] = Blank
	faces[2] = Blank
	faces[3] = Shield
	faces[4] = Hit
	faces[5] = Hit
	faces[6] = Wild
	faces[7] = Crit
}

// McpDie stores the result of a die roll
type McpDie struct {
	Value int `default:"-1"`
}

func (d *McpDie) Roll() {
	d.Value = rand.Int() % len(faces)
}

func (d McpDie) String() string {
	return faces[d.Value]
}

func RollDice(num int) []McpDie {
	result := make([]McpDie, num)
	for i := 0; i < num; i++ {
		result[i] = McpDie{}
		result[i].Roll()
	}

	return result
}

func ExplodeCrits(rolls []McpDie, critOptions CritOptions) []McpDie {

	if critOptions.CritsToSkulls {
		for i, roll := range rolls {
			if roll.String() == Crit {
				rolls[i].Value = 0
			}
		}
	}

	if critOptions.NoExplosion {
		return rolls
	}

	count := 0
	skullCrits := 0

	for _, roll := range rolls {
		if roll.String() == Crit {
			count += 1
		} else if skullCrits < critOptions.SkullsAsCrits && roll.String() == Skull {
			skullCrits += 1
			count += 1
		}
	}

	if count > 0 {
		rolls = append(rolls, RollDice(count)...)
	}

	return rolls
}

func SortDice(rolls []McpDie) ([]McpDie, error) {
	result := make([]McpDie, len(rolls))
	copy(result, rolls)

	for _, die := range result {
		if die.Value == -1 {
			// the intention is the dice should be rolled before you sort them, but we can roll it here
			die.Roll()
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})

	return result, nil
}

type AttackOptions struct {
	DoubleExplodeCrits bool
	SkullsToCrits      bool
}

type DefenceOptions struct {
	CritsToSkulls    bool
	NoCritExplosions bool
	SkullsToCrits    bool
}

type CritOptions struct {
	NoExplosion   bool
	CritsToSkulls bool
	SkullsAsCrits int
	DoubleExplode bool
}

func GenerateAttackCritOptions(attack *AttackOptions, defence *DefenceOptions) CritOptions {
	critOptions := CritOptions{}
	return critOptions
}

func GenerateDefenceCritOptions(attack *AttackOptions, defence *DefenceOptions) CritOptions {
	critOptions := CritOptions{}

	return critOptions
}
