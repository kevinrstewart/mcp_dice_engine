package engine

import (
	"fmt"
	"kevinrstewart/mcp_dice_enginge/src/models"
)

type Analysis struct {
	records map[int]int
	count   int
}

func NewAnalysis() Analysis {
	return Analysis{
		records: make(map[int]int),
	}
}

func (a Analysis) String() string {
	m := make(map[int]string, len(a.records))

	for k, v := range a.records {
		m[k] = fmt.Sprintf("%d - %.6f", k, float32(v)/float32(a.count))
	}

	return fmt.Sprintf("%v", m)
}

func (a *Analysis) Record(result int) {
	val, ok := a.records[result]
	if !ok {
		a.records[result] = 1
	} else {
		a.records[result] = val + 1
	}

	a.count += 1
}

type Result int

func GenerateResult(attack int, attackOptions *models.AttackOptions, defence int, defenceOptions *models.DefenceOptions) int {

	a, ad := DetermineAttackSuccesses(attack, attackOptions, defenceOptions)

	d := DetermineDefenceSuccesses(defence, ad, attackOptions, defenceOptions)

	if a-d > 0 {
		return a - d
	}

	return 0
}

func DetermineAttackSuccesses(num int, attackOptions *models.AttackOptions, defenceOptions *models.DefenceOptions) (int, []models.McpDie) {
	dice := models.ExplodeCrits(models.RollDice(num), models.GenerateAttackCritOptions(attackOptions, defenceOptions))

	count := 0
	for _, roll := range dice {
		if roll.Value > 4 {
			count += 1
		}
	}

	return count, dice
}

func DetermineDefenceSuccesses(num int, attackDice []models.McpDie, attackOptions *models.AttackOptions, defenceOptions *models.DefenceOptions) int {
	dice := models.ExplodeCrits(models.RollDice(num), models.GenerateDefenceCritOptions(attackOptions, defenceOptions))

	count := 0
	for _, roll := range dice {
		if roll.Value > 6 || roll.Value == 3 {
			count += 1
		}
	}

	return count
}

func SmartPierce(attack []models.McpDie, attackOptions *models.AttackOptions, defence []models.McpDie, defenceOptions *models.DefenceOptions) {
	// based on the options figure out the most logical thing to pierce
}

func AnalyzeResults(results []int) Analysis {
	analysis := NewAnalysis()

	for _, res := range results {
		analysis.Record(res)
	}

	return analysis
}
