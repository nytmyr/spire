package models

import (
	"github.com/volatiletech/null/v8"
)

type Spawnentry struct {
	SpawngroupID           int         `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	NpcID                  int         `json:"npc_id" gorm:"Column:npcID"`
	Chance                 int16       `json:"chance" gorm:"Column:chance"`
	ConditionValueFilter   int32       `json:"condition_value_filter" gorm:"Column:condition_value_filter"`
	MinTime                int16       `json:"min_time" gorm:"Column:min_time"`
	MaxTime                int16       `json:"max_time" gorm:"Column:max_time"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Spawngroup             *Spawngroup `json:"spawngroup,omitempty" gorm:"foreignKey:spawngroupID;references:id"`
	NpcType                *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:npcID;references:id"`
}

func (Spawnentry) TableName() string {
    return "spawnentry"
}

func (Spawnentry) Relationships() []string {
    return []string{
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.AlternateCurrency.Item",
		"NpcType.AlternateCurrency.Item.AlternateCurrencies",
		"NpcType.AlternateCurrency.Item.CharacterCorpseItems",
		"NpcType.AlternateCurrency.Item.DiscoveredItems",
		"NpcType.AlternateCurrency.Item.Doors",
		"NpcType.AlternateCurrency.Item.Doors.Item",
		"NpcType.AlternateCurrency.Item.Fishings",
		"NpcType.AlternateCurrency.Item.Fishings.Item",
		"NpcType.AlternateCurrency.Item.Fishings.NpcType",
		"NpcType.AlternateCurrency.Item.Fishings.Zone",
		"NpcType.AlternateCurrency.Item.Forages",
		"NpcType.AlternateCurrency.Item.Forages.Item",
		"NpcType.AlternateCurrency.Item.Forages.Zone",
		"NpcType.AlternateCurrency.Item.GroundSpawns",
		"NpcType.AlternateCurrency.Item.GroundSpawns.Zone",
		"NpcType.AlternateCurrency.Item.ItemTicks",
		"NpcType.AlternateCurrency.Item.Keyrings",
		"NpcType.AlternateCurrency.Item.LootdropEntries",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Item",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcType.AlternateCurrency.Item.Merchantlists",
		"NpcType.AlternateCurrency.Item.Merchantlists.Items",
		"NpcType.AlternateCurrency.Item.Merchantlists.NpcTypes",
		"NpcType.AlternateCurrency.Item.ObjectContents",
		"NpcType.AlternateCurrency.Item.Objects",
		"NpcType.AlternateCurrency.Item.Objects.Item",
		"NpcType.AlternateCurrency.Item.Objects.Zone",
		"NpcType.AlternateCurrency.Item.TradeskillRecipeEntries",
		"NpcType.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.AlternateCurrency.Item.TributeLevels",
		"NpcType.Loottable",
		"NpcType.Loottable.LoottableEntries",
		"NpcType.Loottable.LoottableEntries.Lootdrop",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"NpcType.Loottable.LoottableEntries.Loottable",
		"NpcType.Loottable.NpcTypes",
		"NpcType.Merchantlists",
		"NpcType.Merchantlists.Items",
		"NpcType.Merchantlists.Items.AlternateCurrencies",
		"NpcType.Merchantlists.Items.AlternateCurrencies.Item",
		"NpcType.Merchantlists.Items.CharacterCorpseItems",
		"NpcType.Merchantlists.Items.DiscoveredItems",
		"NpcType.Merchantlists.Items.Doors",
		"NpcType.Merchantlists.Items.Doors.Item",
		"NpcType.Merchantlists.Items.Fishings",
		"NpcType.Merchantlists.Items.Fishings.Item",
		"NpcType.Merchantlists.Items.Fishings.NpcType",
		"NpcType.Merchantlists.Items.Fishings.Zone",
		"NpcType.Merchantlists.Items.Forages",
		"NpcType.Merchantlists.Items.Forages.Item",
		"NpcType.Merchantlists.Items.Forages.Zone",
		"NpcType.Merchantlists.Items.GroundSpawns",
		"NpcType.Merchantlists.Items.GroundSpawns.Zone",
		"NpcType.Merchantlists.Items.ItemTicks",
		"NpcType.Merchantlists.Items.Keyrings",
		"NpcType.Merchantlists.Items.LootdropEntries",
		"NpcType.Merchantlists.Items.LootdropEntries.Item",
		"NpcType.Merchantlists.Items.LootdropEntries.Lootdrop",
		"NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcType.Merchantlists.Items.Merchantlists",
		"NpcType.Merchantlists.Items.ObjectContents",
		"NpcType.Merchantlists.Items.Objects",
		"NpcType.Merchantlists.Items.Objects.Item",
		"NpcType.Merchantlists.Items.Objects.Zone",
		"NpcType.Merchantlists.Items.TradeskillRecipeEntries",
		"NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.Merchantlists.Items.TributeLevels",
		"NpcType.Merchantlists.NpcTypes",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"NpcType.NpcSpell",
		"NpcType.NpcSpell.BotSpellsEntries",
		"NpcType.NpcSpell.BotSpellsEntries.NpcSpell",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors.Item",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Item",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Item",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Zone",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.ItemTicks",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Keyrings",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.Items",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.ObjectContents",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Item",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Zone",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TributeLevels",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets",
		"NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals",
		"NpcType.NpcSpell.NpcSpell",
		"NpcType.NpcSpell.NpcSpellsEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"Spawngroup",
		"Spawngroup.Spawn2",
		"Spawngroup.Spawn2.Spawnentries",
		"Spawngroup.Spawn2.Spawngroup",
	}
}

func (Spawnentry) Connection() string {
    return "eqemu_content"
}
