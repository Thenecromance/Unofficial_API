package wow

import (
	"Unofficial_API/internal"
	"Unofficial_API/model"
	"Unofficial_API/utils"
	"context"
	"encoding/json"
)

type Equipment struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		EquippedItemSets []struct {
			DisplayString string `json:"display_string"`
			Effects       []struct {
				DisplayString string `json:"display_string"`
				IsActive      bool   `json:"is_active"`
				RequiredCount int    `json:"required_count"`
			} `json:"effects"`
			ItemSet struct {
				Name string `json:"name"`
			} `json:"item_set"`
			Items []struct {
				Item struct {
					Name string `json:"name"`
				} `json:"item"`
				IsEquipped bool `json:"is_equipped,omitempty"`
			} `json:"items"`
		} `json:"equipped_item_sets"`
		EquippedItems []struct {
			Armor struct {
				Display struct {
					Color *struct {
						A int `json:"a"`
						B int `json:"b"`
						G int `json:"g"`
						R int `json:"r"`
					} `json:"color"`
					DisplayString string `json:"display_string"`
				} `json:"display"`
				Value int `json:"value"`
			} `json:"armor"`
			ArtifactXp     int `json:"artifact_xp"`
			AzeriteDetails struct {
				Level struct {
					DisplayString string `json:"display_string"`
					Value         int    `json:"value"`
				} `json:"level"`
				PercentageToNextLevel int         `json:"percentage_to_next_level"`
				SelectedEssences      interface{} `json:"selected_essences"`
				SelectedPowers        interface{} `json:"selected_powers"`
				SelectedPowersString  string      `json:"selected_powers_string"`
			} `json:"azerite_details"`
			Binding struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"binding"`
			BonusList  []int `json:"bonus_list"`
			Context    int   `json:"context"`
			Durability struct {
				DisplayString string `json:"display_string"`
				Value         int    `json:"value"`
			} `json:"durability"`
			InventoryType struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"inventory_type"`
			ItemClass struct {
				Name string `json:"name"`
			} `json:"item_class"`
			ItemSubclass struct {
				Name string `json:"name"`
			} `json:"item_subclass"`
			Legacy string `json:"legacy"`
			Level  struct {
				DisplayString string `json:"display_string"`
				Value         int    `json:"value"`
			} `json:"level"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"Key"`
				Icon string `json:"icon"`
				Id   int    `json:"id"`
			} `json:"media"`
			ModifiedAppearanceId int    `json:"modified_appearance_id,omitempty"`
			Name                 string `json:"name"`
			NameDescription      struct {
				Color *struct {
					A int `json:"a"`
					B int `json:"b"`
					G int `json:"g"`
					R int `json:"r"`
				} `json:"color"`
				DisplayString string `json:"display_string"`
			} `json:"name_description"`
			Quality struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"quality"`
			Quantity     int `json:"quantity"`
			Requirements struct {
				Level struct {
					DisplayString string `json:"display_string"`
					Value         int    `json:"value"`
				} `json:"level"`
				PlayableClasses struct {
					DisplayString string `json:"display_string"`
					Links         []struct {
						Name string `json:"name"`
					} `json:"links"`
				} `json:"playable_classes"`
				PlayableSpecializations interface{} `json:"playable_specializations"`
			} `json:"requirements"`
			SellPrice struct {
				DisplayStrings struct {
					Copper string `json:"copper"`
					Gold   string `json:"gold"`
					Header string `json:"header"`
					Silver string `json:"silver"`
				} `json:"display_strings"`
				Value int `json:"value"`
			} `json:"sell_price"`
			Set struct {
				DisplayString string `json:"display_string"`
				Effects       []struct {
					DisplayString string `json:"display_string"`
					IsActive      bool   `json:"is_active"`
					RequiredCount int    `json:"required_count"`
				} `json:"effects"`
				ItemSet struct {
					Name string `json:"name"`
				} `json:"item_set"`
				Items []struct {
					IsEquipped bool `json:"is_equipped,omitempty"`
					Item       struct {
						Name string `json:"name"`
					} `json:"item"`
				} `json:"items"`
			} `json:"set"`
			Slot struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"slot"`
			Stats []struct {
				Display struct {
					Color struct {
						A int `json:"a"`
						B int `json:"b"`
						G int `json:"g"`
						R int `json:"r"`
					} `json:"color"`
					DisplayString string `json:"display_string"`
				} `json:"display"`
				Type struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"type"`
				Value        int  `json:"value"`
				IsNegated    bool `json:"is_negated,omitempty"`
				IsEquipBonus bool `json:"is_equip_bonus,omitempty"`
			} `json:"stats"`
			Transmog struct {
				DisplayString string `json:"display_string"`
				Item          struct {
					Id  int `json:"id"`
					Key struct {
					} `json:"key"`
					Name string `json:"name"`
				} `json:"item"`
				ItemModifiedAppearanceId int `json:"item_modified_appearance_id"`
			} `json:"transmog"`
			Weapon struct {
				AttackSpeed struct {
					DisplayString string `json:"display_string"`
					Value         int    `json:"value"`
				} `json:"attack_speed"`
				Damage struct {
					DamageClass struct {
						Name string `json:"name"`
						Type string `json:"type"`
					} `json:"damage_class"`
					DisplayString string `json:"display_string"`
					MaxValue      int    `json:"max_value"`
					MinValue      int    `json:"min_value"`
				} `json:"damage"`
				Dps struct {
					DisplayString string  `json:"display_string"`
					Value         float64 `json:"value"`
				} `json:"dps"`
			} `json:"weapon"`
			IsSubclassHidden bool `json:"is_subclass_hidden,omitempty"`
			Sockets          []struct {
				DisplayString string `json:"display_string"`
				Item          struct {
					Name string `json:"name"`
				} `json:"item"`
				Media struct {
					Icon string `json:"icon"`
					Id   int    `json:"id"`
				} `json:"media"`
				SocketType struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"socket_type"`
			} `json:"sockets,omitempty"`
			Enchantments []struct {
				DisplayString   string `json:"display_string"`
				EnchantmentId   int    `json:"enchantment_id"`
				EnchantmentSlot struct {
					Id   int    `json:"id"`
					Type string `json:"type"`
				} `json:"enchantment_slot"`
				SourceItem struct {
					Id  int `json:"id"`
					Key struct {
					} `json:"key"`
					Name string `json:"name"`
				} `json:"source_item"`
				Spell struct {
					Description string `json:"description"`
					Spell       struct {
						Name string `json:"name"`
					} `json:"spell"`
				} `json:"spell"`
			} `json:"enchantments,omitempty"`
			Spells []struct {
				Description string `json:"description"`
				Spell       struct {
					Name string `json:"name"`
				} `json:"spell"`
			} `json:"spells,omitempty"`
			UniqueEquipped       string `json:"unique_equipped,omitempty"`
			LimitCategory        string `json:"limit_category,omitempty"`
			ModifiedCraftingStat []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"modified_crafting_stat,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"equipped_items"`
	} `json:"data"`
}

func (e Equipment) ToBNet() *model.Equipment {
	equip := &model.Equipment{}
	equip.EquippedItems = make([]model.EquippedItem, len(e.Data.EquippedItems))
	equip.EquippedItemSets = make([]model.EquippedItemSet, len(e.Data.EquippedItemSets))

	for i, item := range e.Data.EquippedItems {
		equip.EquippedItems[i].Name = item.Name
		{
			equip.EquippedItems[i].Level.Value = item.Level.Value
			equip.EquippedItems[i].Level.DisplayString = item.Level.DisplayString
		}
		{
			equip.EquippedItems[i].Quality.Name = item.Quality.Name
			equip.EquippedItems[i].Quality.Type = item.Quality.Type
		}
		{
			equip.EquippedItems[i].Binding.Name = item.Binding.Name
			equip.EquippedItems[i].Binding.Type = item.Binding.Type
		}
		{
			equip.EquippedItems[i].ItemClass.Name = item.ItemClass.Name
			equip.EquippedItems[i].ItemClass.Id = 0
		}
		{
			equip.EquippedItems[i].ItemSubclass.Name = item.ItemSubclass.Name
			//equip.EquippedItems[i].ItemSubclass.Key
			equip.EquippedItems[i].ItemSubclass.Id = 0
		}
		{
			equip.EquippedItems[i].InventoryType.Name = item.InventoryType.Name
			equip.EquippedItems[i].InventoryType.Type = item.InventoryType.Type
		}
		{
			equip.EquippedItems[i].Binding.Name = item.Binding.Name
			equip.EquippedItems[i].Binding.Type = item.Binding.Type
		}
		{
			equip.EquippedItems[i].Armor.Value = item.Armor.Value
			{
				equip.EquippedItems[i].Armor.Display.DisplayString = item.Armor.Display.DisplayString
				{
					equip.EquippedItems[i].Armor.Display.Color.R = item.Armor.Display.Color.R
					equip.EquippedItems[i].Armor.Display.Color.G = item.Armor.Display.Color.G
					equip.EquippedItems[i].Armor.Display.Color.B = item.Armor.Display.Color.B
					equip.EquippedItems[i].Armor.Display.Color.A = item.Armor.Display.Color.A
				}
			}
		}
		{
			for statIndex, stat := range item.Stats {
				equip.EquippedItems[i].Stats = append(equip.EquippedItems[i].Stats, model.Stat{})
				equip.EquippedItems[i].Stats[statIndex].Value = stat.Value
				equip.EquippedItems[i].Stats[statIndex].IsNegated = stat.IsNegated
				equip.EquippedItems[i].Stats[statIndex].IsEquipBonus = stat.IsEquipBonus
				{
					equip.EquippedItems[i].Stats[statIndex].Type.Name = stat.Type.Name
					equip.EquippedItems[i].Stats[statIndex].Type.Type = stat.Type.Type

				}
				{
					equip.EquippedItems[i].Stats[statIndex].Display.DisplayString = stat.Display.DisplayString
					equip.EquippedItems[i].Stats[statIndex].Display.Color.R = stat.Display.Color.R
					equip.EquippedItems[i].Stats[statIndex].Display.Color.G = stat.Display.Color.G
					equip.EquippedItems[i].Stats[statIndex].Display.Color.B = stat.Display.Color.B
					equip.EquippedItems[i].Stats[statIndex].Display.Color.A = stat.Display.Color.A
				}
			}
		}
		{
			equip.EquippedItems[i].SellPrice.Value = item.SellPrice.Value
			equip.EquippedItems[i].SellPrice.DisplayStrings.Copper = item.SellPrice.DisplayStrings.Copper
			equip.EquippedItems[i].SellPrice.DisplayStrings.Gold = item.SellPrice.DisplayStrings.Gold
			equip.EquippedItems[i].SellPrice.DisplayStrings.Header = item.SellPrice.DisplayStrings.Header
			equip.EquippedItems[i].SellPrice.DisplayStrings.Silver = item.SellPrice.DisplayStrings.Silver
		}
		{
			{
				equip.EquippedItems[i].Requirements.Level.Value = item.Requirements.Level.Value
				equip.EquippedItems[i].Requirements.Level.DisplayString = item.Requirements.Level.DisplayString
			}
			{
				equip.EquippedItems[i].Requirements.PlayableClasses.DisplayString = item.Requirements.PlayableClasses.DisplayString
				for _, link := range item.Requirements.PlayableClasses.Links {
					equip.EquippedItems[i].Requirements.PlayableClasses.Links = append(equip.EquippedItems[i].Requirements.PlayableClasses.Links, model.PlayableClassesLink{
						Name: link.Name,
						Id:   0,
					})
				}
			}
		}
		{
			equip.EquippedItems[i].Set.ItemSet.Name = item.Set.ItemSet.Name
			equip.EquippedItems[i].Set.DisplayString = item.Set.DisplayString
			for itemIdx, itm := range item.Set.Items {
				equip.EquippedItems[i].Set.Items = append(equip.EquippedItems[i].Set.Items, model.Item{})
				equip.EquippedItems[i].Set.Items[itemIdx].Item.Name = itm.Item.Name
				equip.EquippedItems[i].Set.Items[itemIdx].IsEquipped = itm.IsEquipped
				equip.EquippedItems[i].Set.Items[itemIdx].Item.Id = e.Data.EquippedItems[i].Transmog.Item.Id
			}

			equip.EquippedItems[i].Set.Effects = make([]model.SetEffect, len(item.Set.Effects))
			for effectIdx, effect := range item.Set.Effects {
				equip.EquippedItems[i].Set.Effects[effectIdx].DisplayString = effect.DisplayString
				equip.EquippedItems[i].Set.Effects[effectIdx].IsActive = effect.IsActive
				equip.EquippedItems[i].Set.Effects[effectIdx].RequiredCount = effect.RequiredCount
			}
		}
		{
			equip.EquippedItems[i].Level.Value = item.Level.Value
			equip.EquippedItems[i].Level.DisplayString = item.Level.DisplayString
		}
		{
			equip.EquippedItems[i].Transmog.Item.Id = item.Transmog.Item.Id
			equip.EquippedItems[i].Transmog.Item.Name = item.Transmog.Item.Name
			equip.EquippedItems[i].Transmog.DisplayString = item.Transmog.DisplayString
			equip.EquippedItems[i].Transmog.ItemModifiedAppearanceId = item.Transmog.ItemModifiedAppearanceId
		}
		{
			equip.EquippedItems[i].Durability.Value = item.Durability.Value
			equip.EquippedItems[i].Durability.DisplayString = item.Durability.DisplayString
		}
		{
			equip.EquippedItems[i].NameDescription.DisplayString = item.NameDescription.DisplayString
			{
				equip.EquippedItems[i].NameDescription.Color.R = item.NameDescription.Color.R
				equip.EquippedItems[i].NameDescription.Color.G = item.NameDescription.Color.G
				equip.EquippedItems[i].NameDescription.Color.B = item.NameDescription.Color.B
				equip.EquippedItems[i].NameDescription.Color.A = item.NameDescription.Color.A
			}
		}
		equip.EquippedItems[i].IsSubclassHidden = item.IsSubclassHidden
		{
			equip.EquippedItems[i].ModifiedCraftingStat = make([]model.ModifiedCraftingStat, len(item.ModifiedCraftingStat))
			for statIdx, stat := range item.ModifiedCraftingStat {
				equip.EquippedItems[i].ModifiedCraftingStat[statIdx].Id = stat.Id
				equip.EquippedItems[i].ModifiedCraftingStat[statIdx].Name = stat.Name
				equip.EquippedItems[i].ModifiedCraftingStat[statIdx].Type = stat.Type
			}
		}
		{
			equip.EquippedItems[i].Enchantments = make([]model.Enchantment, len(item.Enchantments))
			for enchIdx, ench := range item.Enchantments {
				equip.EquippedItems[i].Enchantments[enchIdx].DisplayString = ench.DisplayString
				equip.EquippedItems[i].Enchantments[enchIdx].EnchantmentId = ench.EnchantmentId

				equip.EquippedItems[i].Enchantments[enchIdx].EnchantmentSlot.Id = ench.EnchantmentSlot.Id
				equip.EquippedItems[i].Enchantments[enchIdx].EnchantmentSlot.Type = ench.EnchantmentSlot.Type

				equip.EquippedItems[i].Enchantments[enchIdx].SourceItem.Id = ench.SourceItem.Id
				equip.EquippedItems[i].Enchantments[enchIdx].SourceItem.Name = ench.SourceItem.Name

				equip.EquippedItems[i].Enchantments[enchIdx].Spell.Description = ench.Spell.Description
				equip.EquippedItems[i].Enchantments[enchIdx].Spell.Spell.Name = ench.Spell.Spell.Name
				equip.EquippedItems[i].Enchantments[enchIdx].Spell.Spell.Id = 0
			}
		}

		equip.EquippedItems[i].LimitCategory = item.LimitCategory

		{
			equip.EquippedItems[i].Spells = make([]model.Spell, len(item.Spells))
			for spellIdx, spell := range item.Spells {
				equip.EquippedItems[i].Spells[spellIdx].Description = spell.Description
				equip.EquippedItems[i].Spells[spellIdx].Spell.Name = spell.Spell.Name
				equip.EquippedItems[i].Spells[spellIdx].Spell.Id = 0

				equip.EquippedItems[i].Spells[spellIdx].DisplayColor.R = 255
				equip.EquippedItems[i].Spells[spellIdx].DisplayColor.G = 255
				equip.EquippedItems[i].Spells[spellIdx].DisplayColor.B = 255
				equip.EquippedItems[i].Spells[spellIdx].DisplayColor.A = 255
			}
		}

		equip.EquippedItems[i].UniqueEquipped = item.UniqueEquipped
		{
			equip.EquippedItems[i].Weapon.Damage.MaxValue = item.Weapon.Damage.MaxValue
			equip.EquippedItems[i].Weapon.Damage.MinValue = item.Weapon.Damage.MinValue
			equip.EquippedItems[i].Weapon.Damage.DisplayString = item.Weapon.Damage.DisplayString
			equip.EquippedItems[i].Weapon.Damage.DamageClass.Name = item.Weapon.Damage.DamageClass.Name
			equip.EquippedItems[i].Weapon.Damage.DamageClass.Type = item.Weapon.Damage.DamageClass.Type

			equip.EquippedItems[i].Weapon.AttackSpeed.Value = item.Weapon.AttackSpeed.Value
			equip.EquippedItems[i].Weapon.AttackSpeed.DisplayString = item.Weapon.AttackSpeed.DisplayString
		}
		equip.EquippedItems[i].Media.Id = item.Media.Id
		equip.EquippedItems[i].Durability.Value = item.Durability.Value
		equip.EquippedItems[i].Durability.DisplayString = item.Durability.DisplayString
	}

	for i, itemSet := range e.Data.EquippedItemSets {
		equip.EquippedItemSets[i].ItemSet.Name = itemSet.ItemSet.Name
		equip.EquippedItemSets[i].ItemSet.Id = 0
		for itemIdx, itm := range itemSet.Items {
			equip.EquippedItemSets[i].Items = append(equip.EquippedItems[i].Set.Items, model.Item{})
			equip.EquippedItemSets[i].Items[itemIdx].Item.Name = itm.Item.Name
			equip.EquippedItemSets[i].Items[itemIdx].IsEquipped = itm.IsEquipped
			equip.EquippedItemSets[i].Items[itemIdx].Item.Id = e.Data.EquippedItems[i].Transmog.Item.Id
		}

		for effectIdx, effect := range itemSet.Effects {
			equip.EquippedItemSets[i].Effects = append(equip.EquippedItemSets[i].Effects, model.SetEffect{})
			equip.EquippedItemSets[i].Effects[effectIdx].DisplayString = effect.DisplayString
			equip.EquippedItemSets[i].Effects[effectIdx].IsActive = effect.IsActive
			equip.EquippedItemSets[i].Effects[effectIdx].RequiredCount = effect.RequiredCount
		}
		equip.EquippedItemSets[i].DisplayString = itemSet.DisplayString
	}

	return equip
}

func StringEquipmentSummary(ctx context.Context, name string, realmSlug string) (string, error) {
	client := utils.NewRequest()
	token := internal.TryToGetToken(name, realmSlug)
	if token == "" {
		summary := CNPlayerSummary(ctx, name, realmSlug)
		if summary == nil {
			return "", nil
		}
		token = internal.TryToGetToken(name, realmSlug)
	}

	return client.GET("https://webapi.blizzard.cn/wow-armory-server/api/do", "api", "equipment", "token", token)
}

func CNEquipmentSummary(ctx context.Context, name string, realmSlug string) *Equipment {
	obj, err := StringEquipmentSummary(ctx, name, realmSlug)
	if err != nil {
		return nil
	}
	equip := &Equipment{}
	err = json.Unmarshal([]byte(obj), equip)
	if err != nil {
		return nil
	}
	return equip
}

func BNetEquipmentSummary(ctx context.Context, name string, realmSlug string) *model.Equipment {
	obj := CNEquipmentSummary(ctx, name, realmSlug)
	if obj == nil {
		return nil
	}
	return obj.ToBNet()
}
