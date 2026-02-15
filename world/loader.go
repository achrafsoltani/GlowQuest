package world

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"

	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/data"
	"github.com/AchrafSoltani/GlowQuest/entity"
)

// --- JSON structures for deserialization ---

type jsonOverworldMeta struct {
	Width         int                      `json:"width"`
	Height        int                      `json:"height"`
	StartScreen   struct{ X, Y int }       `json:"start_screen"`
	StartPos      struct{ X, Y float64 }   `json:"start_pos"`
	StartInterior string                   `json:"start_interior"`
	Regions       map[string]jsonRegion     `json:"regions"`
}

type jsonRegion struct {
	Name    string   `json:"name"`
	Screens [][2]int `json:"screens"`
}

type jsonRow struct {
	Row     int          `json:"row"`
	Screens []jsonScreen `json:"screens"`
}

type jsonScreen struct {
	Col     int              `json:"col"`
	Tiles   [][]int          `json:"tiles"`
	Enemies []jsonEnemy      `json:"enemies"`
	Items   []jsonItem       `json:"items"`
	NPCs    []jsonNPC        `json:"npcs"`
	Warps   []jsonWarp       `json:"warps"`
}

type jsonEnemy struct {
	Type interface{} `json:"type"` // int or string
	X    int         `json:"x"`
	Y    int         `json:"y"`
}

type jsonItem struct {
	Type      interface{} `json:"type"` // int or string
	X         int         `json:"x"`
	Y         int         `json:"y"`
	Condition string      `json:"condition,omitempty"`
}

type jsonNPC struct {
	ID          string             `json:"id"`
	X           int                `json:"x"`
	Y           int                `json:"y"`
	Dir         int                `json:"dir"`
	Name        string             `json:"name"`
	DialogueKey string             `json:"dialogue_key"`
	Dialogues   []jsonDialogueOpt  `json:"dialogues,omitempty"`
	Condition   string             `json:"condition,omitempty"`
}

type jsonDialogueOpt struct {
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

type jsonWarp struct {
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Target string  `json:"target"`
	SX     float64 `json:"sx"`
	SY     float64 `json:"sy"`
	EX     float64 `json:"ex"`
	EY     float64 `json:"ey"`
}

// --- Interior JSON structures ---

type jsonInteriorFile struct {
	Interiors []jsonInterior `json:"interiors"`
}

type jsonInterior struct {
	ID      string       `json:"id"`
	Tiles   [][]int      `json:"tiles"`
	Enemies []jsonEnemy  `json:"enemies"`
	Items   []jsonItem   `json:"items"`
	NPCs    []jsonNPC    `json:"npcs"`
	Warps   []jsonWarp   `json:"warps"`
}

// --- Dialogue table ---

// DialogueTable maps dialogue keys to dialogue lines.
var DialogueTable = map[string][]string{
	// Original NPCs
	"old_man_intro": {
		"It's dangerous to go",
		"alone! Take the sword",
		"by the well.",
	},
	"merchant_welcome": {
		"Welcome to our village.",
		"The ruins to the south-",
		"east hold great treasure.",
	},
	"traveller_beware": {
		"Beware the mountain.",
		"Many Moblins lurk",
		"there.",
	},
	"ghost_ruins": {
		"You've found the",
		"ancient ruins. The",
		"stairs lead deeper...",
	},
	"villager_house_1": {
		"Please make yourself",
		"at home. The village",
		"is peaceful... for now.",
	},
	"scholar_house_2": {
		"I've been studying the",
		"ruins. Ancient power",
		"sleeps beneath them.",
	},

	// Tarin
	"tarin_intro": {
		"Yer finally awake!",
		"I'm Tarin. Found ye",
		"washed up on shore.",
	},
	"tarin_has_sword": {
		"Ah, ye found yer",
		"sword! The beach",
		"can be dangerous...",
	},
	"tarin_shield": {
		"Here, take this",
		"shield. Ye'll need",
		"it out there.",
	},

	// Marin
	"marin_singing": {
		"The wind fish in",
		"name only, for it",
		"is neither...",
	},
	"marin_met": {
		"You remind me of",
		"someone... Please be",
		"careful out there.",
	},

	// Madam MeowMeow
	"meowmeow_intro": {
		"My BowWow is the",
		"best! Don't get too",
		"close though!",
	},

	// Library
	"librarian_lore": {
		"The Wind Fish sleeps",
		"in the Egg atop the",
		"mountains...",
	},

	// Shop
	"shopkeeper_hello": {
		"Welcome! Take a look",
		"around. I've got",
		"supplies for sale.",
	},

	// Village NPCs
	"kid_village": {
		"I wanna be an",
		"adventurer when I",
		"grow up!",
	},
	"villager_east": {
		"The library has old",
		"books about this",
		"island's secrets.",
	},
	"owl_statue_village": {
		"Head south to find",
		"what the sea washed",
		"ashore...",
	},

	// Beach
	"beach_hermit": {
		"This shore is called",
		"Toronbo. Many things",
		"wash up here...",
	},

	// Old Man (moved to interior)
	"old_man_cave": {
		"This cave is safe.",
		"Rest here before you",
		"venture further.",
	},

	// Phone booth
	"phone_hint": {
		"Ring ring! The path",
		"south leads to the",
		"Toronbo Shores.",
	},
}

// --- Loading functions ---

// OverworldMeta holds the parsed overworld metadata.
type OverworldMeta struct {
	Width         int
	Height        int
	StartScreen   [2]int
	StartPos      [2]float64
	StartInterior string
}

// LoadOverworldMeta loads the overworld.json metadata.
func LoadOverworldMeta() *OverworldMeta {
	raw, err := data.MapsFS.ReadFile("maps/overworld.json")
	if err != nil {
		log.Printf("loader: failed to read overworld.json: %v", err)
		return &OverworldMeta{Width: 16, Height: 16, StartScreen: [2]int{8, 8}, StartPos: [2]float64{113, 113}}
	}
	var meta jsonOverworldMeta
	if err := json.Unmarshal(raw, &meta); err != nil {
		log.Printf("loader: failed to parse overworld.json: %v", err)
		return &OverworldMeta{Width: 16, Height: 16, StartScreen: [2]int{8, 8}, StartPos: [2]float64{113, 113}}
	}
	return &OverworldMeta{
		Width:         meta.Width,
		Height:        meta.Height,
		StartScreen:   [2]int{meta.StartScreen.X, meta.StartScreen.Y},
		StartPos:      [2]float64{meta.StartPos.X, meta.StartPos.Y},
		StartInterior: meta.StartInterior,
	}
}

// LoadOverworldScreens loads all overworld screens from JSON row files.
func LoadOverworldScreens() map[[2]int]*Screen {
	screens := make(map[[2]int]*Screen)

	entries, err := fs.ReadDir(data.MapsFS, "maps/overworld")
	if err != nil {
		log.Printf("loader: failed to read overworld dir: %v", err)
		return screens
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		raw, err := data.MapsFS.ReadFile("maps/overworld/" + entry.Name())
		if err != nil {
			log.Printf("loader: failed to read %s: %v", entry.Name(), err)
			continue
		}
		var row jsonRow
		if err := json.Unmarshal(raw, &row); err != nil {
			log.Printf("loader: failed to parse %s: %v", entry.Name(), err)
			continue
		}
		for _, js := range row.Screens {
			screen := convertJSONScreen(&js)
			screens[[2]int{js.Col, row.Row}] = screen
		}
	}

	return screens
}

// LoadInteriors loads all interior definitions from JSON files.
func LoadInteriors() (map[string]*InteriorDef, []DoorLink) {
	interiors := make(map[string]*InteriorDef)
	var allDoorLinks []DoorLink

	entries, err := fs.ReadDir(data.MapsFS, "maps/interiors")
	if err != nil {
		log.Printf("loader: failed to read interiors dir: %v", err)
		return interiors, allDoorLinks
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		raw, err := data.MapsFS.ReadFile("maps/interiors/" + entry.Name())
		if err != nil {
			log.Printf("loader: failed to read %s: %v", entry.Name(), err)
			continue
		}
		var file jsonInteriorFile
		if err := json.Unmarshal(raw, &file); err != nil {
			log.Printf("loader: failed to parse %s: %v", entry.Name(), err)
			continue
		}
		for _, ji := range file.Interiors {
			def := convertJSONInterior(&ji)
			interiors[def.ID] = def
		}
	}

	return interiors, allDoorLinks
}

// --- Conversion helpers ---

func convertJSONScreen(js *jsonScreen) *Screen {
	s := &Screen{}

	// Load tiles
	for y := 0; y < config.ScreenGridH && y < len(js.Tiles); y++ {
		for x := 0; x < config.ScreenGridW && x < len(js.Tiles[y]); x++ {
			s.Tiles[y][x] = TileType(js.Tiles[y][x])
		}
	}

	// Load enemies
	for _, je := range js.Enemies {
		s.EnemySpawns = append(s.EnemySpawns, EnemySpawn{
			Type:  resolveEnemyType(je.Type),
			TileX: je.X,
			TileY: je.Y,
		})
	}

	// Load items
	for _, ji := range js.Items {
		s.ItemSpawns = append(s.ItemSpawns, ItemSpawn{
			Type:  resolveItemType(ji.Type),
			TileX: ji.X,
			TileY: ji.Y,
		})
	}

	// Load NPCs
	for _, jn := range js.NPCs {
		spawn := convertJSONNPC(&jn)
		s.NPCSpawns = append(s.NPCSpawns, spawn)
	}

	// Load warps into screen
	for _, jw := range js.Warps {
		s.Warps = append(s.Warps, ScreenWarp{
			TileX:  jw.X,
			TileY:  jw.Y,
			Target: jw.Target,
			SpawnX: jw.SX,
			SpawnY: jw.SY,
			ExitX:  jw.EX,
			ExitY:  jw.EY,
		})
	}

	return s
}

func convertJSONInterior(ji *jsonInterior) *InteriorDef {
	s := &Screen{}
	for y := 0; y < config.ScreenGridH && y < len(ji.Tiles); y++ {
		for x := 0; x < config.ScreenGridW && x < len(ji.Tiles[y]); x++ {
			s.Tiles[y][x] = TileType(ji.Tiles[y][x])
		}
	}

	def := &InteriorDef{
		ID:     ji.ID,
		Screen: s,
	}

	for _, je := range ji.Enemies {
		def.EnemySpawns = append(def.EnemySpawns, EnemySpawn{
			Type:  resolveEnemyType(je.Type),
			TileX: je.X,
			TileY: je.Y,
		})
	}

	for _, jit := range ji.Items {
		def.ItemSpawns = append(def.ItemSpawns, ItemSpawn{
			Type:  resolveItemType(jit.Type),
			TileX: jit.X,
			TileY: jit.Y,
		})
	}

	for _, jn := range ji.NPCs {
		spawn := convertJSONNPC(&jn)
		def.NPCSpawns = append(def.NPCSpawns, spawn)
	}

	for _, jw := range ji.Warps {
		def.DoorLinks = append(def.DoorLinks, DoorLink{
			DoorTileX:  jw.X,
			DoorTileY:  jw.Y,
			InteriorID: jw.Target,
			SpawnX:     jw.SX,
			SpawnY:     jw.SY,
			ExitX:      jw.EX,
			ExitY:      jw.EY,
		})
	}

	return def
}

func convertJSONNPC(jn *jsonNPC) NPCSpawn {
	dialogue := DialogueTable[jn.DialogueKey]
	if dialogue == nil {
		dialogue = []string{"..."}
	}

	spawn := NPCSpawn{
		ID:       jn.ID,
		TileX:    jn.X,
		TileY:    jn.Y,
		Dir:      jn.Dir,
		Name:     jn.Name,
		Dialogue: dialogue,
	}

	// Build conditional dialogues from "dialogues" array
	for _, d := range jn.Dialogues {
		lines := DialogueTable[d.Key]
		if lines == nil {
			continue
		}
		spawn.ConditionalDialogues = append(spawn.ConditionalDialogues, entity.DialogueOption{
			Condition: d.Condition,
			Lines:     lines,
		})
	}

	return spawn
}

func resolveEnemyType(v interface{}) int {
	switch t := v.(type) {
	case float64:
		return int(t)
	case string:
		switch t {
		case "octorok":
			return 0
		case "moblin":
			return 1
		case "stalfos":
			return 2
		case "boss":
			return 3
		default:
			return 0
		}
	}
	return 0
}

func resolveItemType(v interface{}) int {
	switch t := v.(type) {
	case float64:
		return int(t)
	case string:
		switch t {
		case "heart":
			return 0
		case "rupee":
			return 1
		case "key":
			return 2
		case "sword":
			return 3
		case "heart_container":
			return 4
		default:
			return 0
		}
	}
	return 0
}

// --- Overworld door link builder from screen warps ---

// BuildDoorLinksFromScreens builds the overworld door links by scanning all screen warps.
func BuildDoorLinksFromScreens(screens map[[2]int]*Screen) []DoorLink {
	var links []DoorLink
	for pos, screen := range screens {
		for _, w := range screen.Warps {
			if len(w.Target) > 9 && w.Target[:9] == "interior:" {
				interiorID := w.Target[9:]
				links = append(links, DoorLink{
					ScreenX:    pos[0],
					ScreenY:    pos[1],
					DoorTileX:  w.TileX,
					DoorTileY:  w.TileY,
					InteriorID: interiorID,
					SpawnX:     w.SpawnX,
					SpawnY:     w.SpawnY,
					ExitX:      w.ExitX,
					ExitY:      w.ExitY,
				})
			}
		}
	}
	return links
}

// LoadDungeon loads a dungeon definition from JSON. Placeholder for Phase 8+.
func LoadDungeon(id string) *Dungeon {
	path := fmt.Sprintf("maps/dungeons/%s.json", id)
	raw, err := data.MapsFS.ReadFile(path)
	if err != nil {
		log.Printf("loader: dungeon %s not found: %v", id, err)
		return nil
	}
	_ = raw // TODO: parse dungeon JSON in Phase 8
	return nil
}
