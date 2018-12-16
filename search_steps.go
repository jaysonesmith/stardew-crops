package main

import (
	"fmt"
	"strings"

	"github.com/jaysonesmith/stardew-crops/utils"
)

func (sc *ScenarioContext) SearchNoArgs() error {
	StardewCropsCommand([]string{})

	return nil
}

func (sc *ScenarioContext) Search(flag, value string) error {
	args := []string{
		"search",
		fmt.Sprintf(`--%s=%s`, flag, value),
	}
	sc.STDOut = StardewCropsCommand(args)

	return nil
}

func (sc *ScenarioContext) MatchBundleCrops(bundle string) error {
	expected := `[{"name":"blueberry","info":{"description":"A popular berry reported to have many health benefits. The blue skin has the highest nutrient concentration.","seed":"blueberry seeds","growth_time":13,"season":["summer"],"continual":true,"regrowth":4},"seed_prices":{"general_store":80},"bundles":["Summer Crops"],"recipes":["Blueberry Tart","Fruit Salad"],"notes":["Has a small chance for multiple fruit from each harvest"]},{"name":"hot pepper","info":{"description":"Fiery hot with a hint of sweetness.","seed":"pepper seeds","growth_time":5,"season":["summer"],"continual":true,"regrowth":3},"seed_prices":{"general_store":40,"jojamart":50},"bundles":["Summer Crops"],"recipes":["Pepper Poppers","Spicy Eel"],"quests":["Knee Therapy","Summer Help Wanted"],"notes":["Has a small chance for multiple fruit from each harvest"]}]`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchNotFound() error {
	expected := `"no matching crops found"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchNoValueSpecifiedMessage() error {
	expected := `"A value must be provided for each parameter specified"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

// remove below?

func (sc *ScenarioContext) MatchAllCropsBySeason() error {
	expected := utils.Open("./test_data/seasons/all.json")
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchSeasonCrops(season string) error {
	expected := utils.Open(fmt.Sprintf(`./test_data/seasons/%s.json`, season))
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}

func (sc *ScenarioContext) MatchUnknownCropMessage() error {
	expected := `"Unknown season for breakfast"`
	actual := strings.TrimSpace(sc.STDOut)

	return utils.AssertMatch(expected, actual)
}
