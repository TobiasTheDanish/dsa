package main

import (
	"bufio"
	"log"
	"os"
	"time"

	levdist "github.com/tobiasthedanish/lev-dist/lev-dist"
)

var (
	words = []string{
		"apple", "brain", "crane", "dance", "earth", "flame", "giant", "house", "image", "juice",
		"knife", "lemon", "money", "night", "ocean", "plane", "quest", "river", "snake", "tiger",
		"under", "vivid", "whale", "zebra", "angel", "block", "charm", "drift", "enter", "frost",
		"glare", "happy", "ideal", "joker", "karma", "layer", "magic", "north", "orbit", "patch",
		"quiet", "round", "stone", "trend", "unity", "vital", "winds", "xerox", "yield", "zebra",
		"actor", "basic", "chase", "drink", "early", "focus", "grape", "heart", "index", "jelly",
		"knock", "leapt", "metal", "noble", "other", "proud", "queen", "rapid", "shift", "toast",
		"upper", "vital", "waste", "xerox", "yacht", "amber", "brisk", "clear", "depth", "eager",
		"flour", "grasp", "harsh", "input", "joint", "kneel", "latch", "march", "nerve", "orbit",
		"plant", "quick", "reign", "sharp", "toast", "usher", "value", "weary", "youth", "adopt",
		"blink", "creek", "draft", "enact", "field", "gloom", "hinge", "inner", "jolly", "knees",
		"laser", "miner", "naval", "often", "piano", "quote", "resin", "smile", "tease", "urban",
		"vault", "woven", "adapt", "blend", "catch", "dream", "elite", "frank", "grill", "hover",
		"inner", "jumbo", "knead", "large", "maple", "nicer", "outer", "point", "quiet", "round",
		"stack", "taste", "ultra", "vivid", "write", "angry", "blind", "craft", "dodge", "equal",
		"fight", "glide", "hoist", "irony", "jumbo", "lodge", "magic", "never", "ocean", "pound",
		"query", "rider", "stack", "trust", "upset", "voter", "wiser", "adapt", "blend", "crisp",
		"drain", "event", "faith", "gloom", "habit", "index", "jolly", "kneel", "laser", "march",
		"nerve", "orbit", "plant", "quiet", "rapid", "shine", "tiger", "upper", "vault", "woven",
		"amber", "blend", "catch", "draft", "elite", "focus", "grape", "harsh", "index", "joint",
		"kneel", "layer", "march", "noble", "outer", "place", "quick", "rider", "sharp", "trend",
		"upper", "value", "wiser", "apple", "blend", "charm", "draft", "event", "focus", "grape",
		"harsh", "index", "jolly", "kneel", "laser", "march", "nerve", "orbit", "plant", "rapid",
		"shine", "tiger", "upper", "vault", "woven", "angle", "blend", "crisp", "drain", "elite",
		"focus", "glide", "harsh", "index", "jolly", "kneel", "large", "march", "nerve", "orbit",
		"piano", "quiet", "rapid", "shine", "tiger", "upper", "vault", "woven", "adapt", "blink",
		"catch", "drain", "elite", "faith", "glide", "habit", "index", "jolly", "kneel", "laser",
		"march", "nerve", "orbit", "place", "quiet", "rapid", "shine", "trust", "upper", "value",
		"woven", "apple", "blend", "charm", "draft", "elite", "focus", "gloom", "harsh", "index",
		"jolly", "kneel", "laser", "march", "nerve", "orbit", "place", "quiet", "rapid", "shine",
		"stack", "trust", "upper", "value", "woven", "angle", "blend", "crisp", "draft", "elite",
		"faith", "glide", "habit", "index", "jolly", "kneel", "large", "march", "nerve", "orbit",
		"piano", "quiet", "rapid", "shine", "trust", "upper", "value", "woven", "adapt", "blink",
		"charm", "draft", "elite", "faith", "focus", "glide", "habit", "index", "jolly", "kneel",
		"laser", "march", "nerve", "orbit", "place", "quiet", "rapid", "shine", "stack", "trust",
		"upper", "value", "woven", "angle", "blink", "crisp", "draft", "elite", "faith", "focus",
		"glide", "habit", "index", "jolly", "kneel", "laser", "march", "nerve", "orbit", "piano",
		"quiet", "rapid", "shine", "stack", "trust", "upper", "value", "woven", "adapt", "blink",
		"charm", "draft", "elite", "focus", "gloom", "habit", "index", "jolly", "kneel", "large",
		"march", "nerve", "orbit", "piano", "quiet", "rapid", "shine", "stack", "trust", "upper",
		"value", "woven", "angle", "blink", "crisp", "drain", "elite", "faith", "glide", "habit",
		"index", "jolly", "kneel", "laser", "march", "nerve", "orbit", "piano", "quiet", "rapid",
		"shine", "stack", "trust", "upper", "value", "woven", "adapt", "blink", "catch", "drain",
		"elite", "focus", "gloom", "habit", "index", "jolly", "kneel", "large", "march", "nerve",
		"orbit", "piano", "quiet", "rapid", "shine", "stack", "trust", "upper", "value", "woven",
		"a", "an", "be", "by", "do", "go", "if", "in", "is", "it",
		"me", "no", "of", "on", "or", "to", "up", "us", "we", "am",
		"as", "at", "he", "hi", "my", "so", "ox", "ax", "pi", "we",
		"ace", "act", "air", "ant", "ape", "art", "ask", "bat", "bee",
		"bid", "bin", "bit", "bow", "bug", "cat", "cap", "cop", "cry",
		"cut", "dam", "den", "die", "dim", "dip", "dot", "dry", "dug",
		"eat", "egg", "era", "fan", "fat", "fig", "fit", "fix", "fly",
		"fog", "gap", "gas", "get", "god", "gum", "gut", "ham", "hat",
		"hit", "hop", "hot", "hug", "ice", "ill", "ink", "inn", "jar",
		"jet", "joy", "key", "kid", "kit", "law", "leg", "log", "mad",
		"man", "map", "mix", "mud", "net", "nod", "now", "oak", "odd",
		"oil", "owl", "pad", "pal", "pan", "pen", "pet", "pig", "pop",
		"pot", "put", "ram", "rat", "raw", "red", "rod", "row", "run",
		"sad", "sea", "see", "set", "sit", "sky", "son", "sun", "tap",
		"tax", "tie", "tip", "top", "toy", "try", "war", "wax", "web",
		"win", "zip", "zero", "alive", "above", "after", "again", "agree",
		"alert", "allow", "along", "angry", "aside", "below", "blast",
		"bring", "broil", "charm", "cling", "close", "crisp", "dream",
		"drive", "eager", "eight", "enter", "extra", "fairy", "fleet",
		"globe", "happy", "haste", "honor", "inner", "jolly", "knot",
		"large", "latch", "lively", "marry", "melon", "neon", "north",
		"object", "orange", "quiet", "rapid", "shout", "trend", "ultra",
		"vivid", "wagon", "whale", "xerox", "yellow", "zeal", "actor",
		"acorn", "blaze", "blend", "clown", "crowd", "delight", "engine",
		"galaxy", "hunger", "inside", "kitten", "leader", "market",
		"number", "oceanic", "picture", "quality", "reason", "silent",
		"tornado", "umbrella", "voltage", "whisper", "xylophone",
		"yesterday", "zodiac", "balloon", "carpet", "diamond", "flavor",
		"guitar", "history", "internet", "jungle", "kitchen", "laughter",
		"marathon", "notebook", "operation", "question", "resource",
		"treasure", "vacation", "whistle", "wonder", "xenophobia",
		"zookeeper", "alignment", "beautiful", "creature", "decoration",
		"elevation", "foundation", "generation", "imagination", "knowledge",
		"lightning", "meditation", "navigation", "perception", "situation",
		"transformation", "understanding", "vibration", "wonderful",
		"championship", "determination", "entertainment", "friendship",
		"inspiration", "justification", "magnification", "negotiation",
		"organization", "revolution", "stimulation", "transportation",
		"verification", "celebration", "determination", "exaggeration",
	}
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatalln("Please provide a word for me to calculate the levensthein edit distance of.")
	}

	input := args[0]

	for {
		if input == "quit" {
			break
		}

		start := time.Now().UnixMicro()
		smallest := 65000
		closestWord := ""
		for _, word := range words {
			dist := levdist.Calculate(input, word)

			if dist < smallest {
				smallest = dist
				closestWord = word
			}
		}
		end := time.Now().UnixMicro()

		log.Printf("The closest word to '%s' was '%s' with a distance of %d\n", input, closestWord, smallest)
		log.Printf("Calculation took: %d microseconds\n", end-start)

		r := bufio.NewReader(os.Stdin)
		newInput, err := r.ReadString('\n')
		if err != nil {
			log.Fatal("Could not read from stdin: ", err)
		}

		input = newInput[0 : len(newInput)-1]
	}
}
