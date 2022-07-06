type Card struct {
    ID      int
	numeral string
    name    string
    art     string
    desc    []string
    rDesc   []string
}

fool := new(Card){
    ID: 0,
	numeral: "O"
    name: "The Fool",
    art: "",
    desc: []{"beginnings", "new journeys", "trust", "hope", "optimism", "innocence", "spontaneity", "eccentricity"}
    rDesc: []{"naiveté", "inexperienced", "unwordly", "easily duped", "risky", "chaos", "bad judgement"}
}

magician := new(Card){
    ID: 1,
	numeral: "I"
    name: "The Magician",
    art: "",
    desc: []{"creative power", "vision", "mastery", "energy" , "focus", "control", "organization", "self-discipline", "transformation of dreams into reality"}
    rDesc: []{"indecisive", "poor with follow-through", "lacking commitment", "misuse of power", "lacking energy", "cheating", "conning"}
}

highpriestess := new(Card){
    ID: 2,
	numeral: "II"
    name: "The High Priestess",
    art: "",
    desc: []{"intuition", "instinct", "dreams", "mysteries", "deep emotions", "greater powers at work"}
    rDesc: []{"repressed intuition", "blocks", "false allies", "hidden enemies", "hidden agendas", "answers locked away"}
}

empress := new(Card){
    ID: 3,
	numeral: "III",
    name: "The Empress",
    art: "",
    desc: []{"fruitfulness", "fertility", "fulfillment", "sensuality", "beauty", "nurturing", "motherhood"}
    rDesc: []{"result of energy wasted", "abilities obstructed", "laziness", "poverty", "feeling impoverished", "problems at home"}
}

emperor := new(Card){
    ID: 4,
	numeral: "IV",
    name: "The Emperor",
    art: "",
    desc: []{"power", "authority", "assertion", "control", "order", "structure", "logic", "father figure ", "worldly success ", "possible promotions"}
    rDesc: []{"controlling", "inflexibility", "taxing", "abuse of power", "authority lost", "coldness", "lacking emotion", "feelings repressed"}
}

hierophant := new(Card){
    ID: 5,
	numeral "V"
    name: "The Hierophant",
    art: "",
    desc: []{"spirituality", "education", "higher learning", "mentor / teacher", "seeking guidance", "social groups", "belonging", "roles defined"}
    rDesc: []{"breaking rules", "anarchy", "rebellion", "cult mentality", "outmoded beliefs", "rejecting old beliefs", "poor advice", "disappointments"}
}

lovers := new(Card){
    ID: 6,
	numeral: "VI",
    name: "The Lovers",
    art: "",
    desc: []{"love", "passion", "relationships", "shared values", "balance", "choices", "decisions", "love for oneself", "staying true to oneself"}
    rDesc: []{"infidelity", "imbalance", "jealousy", "arguments", "broken hearts", "emotionless sex", "confusion", "indecision", "wrong choice"}
}

chariot := new(Card){
    ID: 7,
	numeral: "VII",
    name: "The Chariot",
    art: "",
    desc: []{"energy", "drive", "ambition", "moving forward", "victory", "triumph against the odds", "reward", "travel", "news", "opportunities"}
    rDesc: []{"unfocused", "energy scattered", "cravenness", "negativity", "lacking self-control", "journeys called off"}
}

strength := new(Card){
    ID: 8,
	numeral: "VII",
    name: "Strength",
    art: "",
    desc: []{"strength from within", "faith", "bravery", "patience", "ability", "self-control", "animalistic", "negativity overpowered", "befriending an animal"}
    rDesc: []{"weakness", "lack of self-control", "lack of courage", "self doubt", "negative blocks", "paralyzing fear"}
}

hermit := new(Card){
    ID: 9,
	numeral: "IX",
    name: "The Hermit",
    art: "",
    desc: []{"introspection", "meditation", "solitary endeavors", "withdrawing from society", "hunting for knowledge", "wisdom found", "study and learning", "teacher"}
    rDesc: []{"isolation", "loneliness", "fear of being alone", "misfit", "timidity", "distractions", "stubbornness"}
}

wheeloffortune := new(Card){
    ID: 10,
	numeral: "X",
    name: "The Wheel of Fortune",
    art: "",
    desc: []{"good luck", "destiny", "karma", "otherwordly assistance", "serendipity", "chance meetings", "new relationships"}
    rDesc: []{"bad luck", "bad news", "problems", "delays in new pursuits", "time of hardship", "unexpected changes", "self-delusion"}
}

justice := new(Card){
    ID: 11,
	numeral: "XI",
    name: "Justice",
    art: "",
    desc: []{"universal equilibrium", "justice", "fairness", "honor", "resolution of a dispute", "reconciliation", "legal dealings"}
    rDesc: []{"disequilibrium", "injustice", "long delays", "resolutions not in your favor", "separation", "resentment", "dishonesty"}
}

hangedman := new(Card){
    ID: 12,
	numeral: "XII",
    name: "The Hanged Man",
    art: "",
    desc: []{"sacrifice", "understanding", "letting go", "detachment", "suspension", "a crucible"}
    rDesc: []{"selfishness", "pain", "cannot let go", "materialism", "unaware", "lacking perspective"}
}

death := new(Card){
    ID: 13,
	numeral: "XII",
    name: "Death",
    art: "",
    desc: []{"change", "transformation", "transition", "era coming to a close", "end", "new beginning", "renewal"}
    rDesc: []{"boredom", "sloth", "letharhy", "depression", "laziness", "ready to change", "unable to move on"}
}

temperance := new(Card){
    ID: 14,
	numeral: "XIV",
    name: "Temperance",
    art: "",
    desc: []{"moderation", "balance", "blending", "harmony", "resolution", "parnership", "excellent cooperation", "new ideas", "creative work"}
    rDesc: []{"discordance", "disagreements", "outbursts", "arguments", "out of balance", " restlessness", "dissatisfaction", "overworked", "no vision", "lack of focus"}
}

devil := new(Card){
    ID: 15,
	numeral: "XV"
    name: "The Devil",
    art: "",
    desc: []{"bondage", "unable to sever ties", "obsession", "money / power / sex", "materialism", "imbalanced relationships", "debauchery", "lust", "kinkiness", "your own personal dark side"}
    rDesc: []{"chains broken", "severing ties", "exhaustion", "shifting values", "change within a relationship"}
}

tower := new(Card){
    ID: 16,
	numeral: "XVI",
    name: "The Tower",
    art: "",
    desc: []{"disruptive change", "extreme changes", "unexpected thoughts and feelings", "anger", "disaster", "ruin", "becoming humbled"}
    rDesc: []{"upheaval", "imprisonment", "freedom, for a price", "clinging to outmoded values", "volatility"}
}

star := new(Card){
    ID: 17,
	numeral: "XVII",
    name: "The Star",
    art: "",
    desc: []{"hope", "goals within reach", "inspiration", "a new approach", "healing", "spending time with nature", "working for the benefit of all"}
    rDesc: []{"pessimism", "faithlessness", "hopelessness", "self-doubt", "feeling discouraged"}
}

moon := new(Card){
    ID: 18,
	numeral: "XVIII",
    name: "The Moon",
    art: "",
    desc: []{"concealed changes", "dreams / nightmares", "intuition", "portent", "omen", "growth of mental faculties", "secrets", "trickery"}
    rDesc: []{"insomnia", "unhappiness", "morbid thoughts", "irrational action", "superstition", "cruelty", "paranoid"}
}

sun := new(Card){
    ID: 19,
	numeral: "XIX",
    name: "The Sun",
    art: "",
    desc: []{"positivity", "happiness", "vitality", "lifeblood", "warmth", "success", "vacation", "children"}
    rDesc: []{"half success, half failure", "temporary difficulties", "delayed success", "diminished vitality", "things not in quite the right order"}
}

judgement := new(Card){
    ID: 20,
	numeral: "XX",
    name: "Judgement",
    art: "",
    desc: []{"rebirth", "awakening", "fresh start", "replenishment", "renewal", "forgiveness"}
    rDesc: []{"uninspired", "boredom", "attempts to avoid the inevitable", "refusal to self-assess", "inability to commit", "ambivalence", "fear of change"}
}

world := new(Card){
    ID: 21,
	numeral: "XXI"
    name: "The World",
    art: "",
    desc: []{"completion", "conclusion", "full circle", "success", "next level", "new era in life", "recognition", "atonement", "new beginnings"}
    rDesc: []{"lesson not learned", "completion delayed", "something left unifinished", "stagnation", "unprepared for the end", "unpleasant loops", "need to try again"}
}

// minor arcana: wands

aceWands := new(Card){
    ID: 22,
	numeral: "I",
    name: "Ace of Wands",
    art: "",
    desc: []{"creation", "origin", "venture", "invention", "source", "inheritance", "beginning", "birth / transformation"}
    rDesc: []{"fall", "delays", "ruin", "spiritual blocks", "misunderstandings", "disappointments"}
}

twoWands := new(Card){
    ID: 23,
	numeral: "II",
    name: "Two of Wands",
    art: "",
    desc: []{"magnificence", "dominion", "productive relationships", "profitable contracts", "work friendships", "employment mentors", "successful real estate transactions"}
    rDesc: []{"trouble", "fears", "hollow success", "end of partnerships", "unexpected expenses", "stubbornness and pride", "miscommunication"}
}

threeWands := new(Card){
    ID: 24,
	numeral: "III",
    name: "Three of Wands",
    art: "",
    desc: []{"manifesting dreams into reality", "fruitful effort", "success", "trade", "commerce", "discovery", "co-operation with another business", "old contacts", "visionary"}
    rDesc: []{"halt in difficulties", "need for practical planning", "unrealistic expectations", "hard work and disappointment", "wanting a free ride"}
}

fourWands := new(Card){
    ID: 25,
	numeral: "IV",
    name: "Four of Wands",
    art: "",
    desc: []{"new home", "rest", "refuge", "harvest", "alignment", "harmony", "prosperity", "work completed"}
    rDesc: []{"growth", "prosperity", "happiness", "beauty", "decoration / ornamentation"}
}

fiveWands := new(Card){
    ID: 26,
	numeral: "V",
    name: "Five of Wands",
    art: "",
    desc: []{"arguments", "conflicts", "testing circumstances", "forging new ideas", "creating new products", "power struggles", "mastering yourself", "sharing viewpoints"}
    rDesc: []{"litigation", "court battles", "negotiation", "inner conflict", "anger", "possible resolution"}
}

sixWands := new(Card){
    ID: 27,
	numeral: "VI",
    name: "Six of Wands",
    art: "",
    desc: []{"victory", "hard won battles at work / creative endeavors", "legal triumphs", "good news", "success"}
    rDesc: []{"enemy at the gate", "troubles caused by others", "arguments at work", "lack of reward", "indecision", "unable to move ahead", "failure to recieve promotion"}
}

sevenWands := new(Card){
    ID: 28,
	numeral: "VII",
    name: "Seven of Wands",
    art: "",
    desc: []{"standing up for yourself", "tenacity", "struggle", "you against the world", "faith in future success"}
    rDesc: []{"fear of failure", "insecurity", "risk", "embarrassments", "anxiety", "make a decision", "encouragement needed"}
}

eightWands := new(Card){
    ID: 29,
	numeral: "VIII",
    name: "Eight of Wands",
    art: "",
    desc: []{"swiftness", "haste", "action", "travel", "invitations", "education", "change in career", "increase in social life", "freedom", "extra energy"}
    rDesc: []{"delays", "dispute", "less speed", "hurrying creates mistakes", "miscommunication", "cancelled travel plans", "strikes", "reassess choices"}
}

nineWands := new(Card){
    ID: 30,
	numeral: "IX",
    name: "Nine of Wands",
    art: "",
    desc: []{"strength", "resilience", "one last challenge ", "persistence", "inner reserves", "prudence", "success is in sight"}
    rDesc: []{"obstacles", "insecurities", "loss of will", "tiredness", "stress", "meditstion", "delays", "overwork", "mental exhaustion", "doubting your own strength"}
}

tenWands := new(Card){
    ID: 31,
	numeral: "X",
    name: "Ten of Wands",
    art: "",
    desc: []{"oppression", "hard work", "ambition", "burden", "drive", "dream realized", "vitality", "workaholic"}
    rDesc: []{"schemes", "difficulties", "biting off more than you can chew", "someone without moral compuction"}
}

pageWands := new(Card){
    ID: 32,
	numeral: "p",
    name: "Page of Wands",
    art: "",
    desc: []{"child / teenager", "creative ideas", "education", "enthusiasm", "surge in energy", "hyperactive"}
    rDesc: []{"bad news", "setbacks", "invitations", "conversation", "tales", "announcements"}
}

knightWands := new(Card){
    ID: 33,
	numeral: "k",
    name: "Knight of Wands",
    art: "",
    desc: []{"young man", "charming", "bright", "creative", "adventure", "easily distracred", "failure to follow through"}
    rDesc: []{"con artist", "fast-talker", "selfish", "rogue", "scattered", "social whirlwind", "invitations"}
}

queenWands := new(Card){
    ID: 34,
	numeral: "Q",
    name: "Queen of Wands",
    art: "",
    desc: []{"exuberant woman", "enthusiastic", "independent", "passionate", "inspirational", "warm", "sense of humor"}
    rDesc: []{"jealousy", "breaking promises", "vengeful", "drama queen", "infidelity", "malicious", "flakiness"}
}

kingWands := new(Card){
    ID: 35,
	numeral: "K",
    name: "King of Wands",
    art: "",
    desc: []{"honest man", "fair", "natural leader", "successful in business", "humorous", "hasty", "passionate"}
    rDesc: []{"sober", "dogmatic", "violent temper", "impulsive", "severe", "charitable"}
}

// minor arcana: cups

aceCups := new(Card){
    ID: 36,
	numeral: "I",
    name: "Ace of Cups",
    art: "",
    desc: []{"heart's home", "abundance", "love", "passion", "growth in relationships", "birth of new ideas or children", "intuition", "fruitfulness"}
    rDesc: []{"chaos", "need for replenishment", "one-sided relationships", "need to withdraw", "mutation", "revolution"}
}

twoCups := new(Card){
    ID: 37,
	numeral: "II",
    name: "Two of Cups",
    art: "",
    desc: []{"union", "love", "connections", "harmony", "romantic attachments", "working partnerships", "fun friendships", "social butterfly"}
    rDesc: []{"emotional battles", "misunderstandings", "separations", "over-dependence", "need for a break", "dissolution", "acting rashly", "looking for common ground"}
}

threeCups := new(Card){
    ID: 38,
	numeral: "III",
    name: "Three of Cups",
    art: "",
    desc: []{"celebration", "parties", "social events", "weddings", "babies", "family gatherings", "fulfillment"}
    rDesc: []{"excess", "lack of sleep", "hedonism", "overspending", "overindulgence", "forgetting about responsibilities", "need for rest and quiet relaxation"}
}

fourCups := new(Card){
    ID: 39,
	numeral: "IV",
    name: "Four of Cups",
    art: "",
    desc: []{"emotional boredom", "depression", "stuck in a rut", "need for new interests", "ignoring the obvious", "short holiday", "fear of love", "emotional scars"}
    rDesc: []{"fear of being alone", "harbinger", "new challenges", "training for a new career", "new hobby"}
}

fiveCups := new(Card){
    ID: 40,
	numeral: "V",
    name: "Five of Cups",
    art: "",
    desc: []{"loss", "destruction of ideals", "broken dreams", "old wounds", "unaware why things went wrong", "emotional loneliness", "inheritance", "transmission"}
    rDesc: []{"reserved hope", "new chances", "rebuilding old relationships", "family relationships", "emotional wisdom", "return"}
}

sixCups := new(Card){
    ID: 41,
	numeral: "VI",
    name: "Six of Cups",
    art: "",
    desc: []{"happy memories", "nostalgia", "old relationships", "looking to the past to understand the future", "rewards for past efforts", "return of past lovers and friends"}
    rDesc: []{"clinging to the past", "blaming the past", "concern for the future", "refusal to budge", "self-doubt", "needing a fresh start"}
}

sevenCups := new(Card){
    ID: 42,
	numeral: "VII",
    name: "Seven of Cups",
    art: "",
    desc: []{"illusions", "wishful thinking", "imagination", "fool's gold", "temporary success", "wait and see", "emotional choices"}
    rDesc: []{"desire", "deceptions", "seeing only what you want to see", "losing touch with reality", "drugs, alcohol", "self-creating fantasies"}
}

eightCups := new(Card){
    ID: 43,
	numeral: "VIII",
    name: "Eight of Cups",
    art: "",
    desc: []{"abandoning relationships", "abandoning enterprise", "moving away emotionally", "travel", "lack of substance", "futile success", "rejection", "restraint"}
    rDesc: []{"hopelessness", "depression", "search for meaning", "emotional confusion", "exhaustion", "lack of energy", "perfectionism"}
}

nineCups := new(Card){
    ID: 44,
	numeral: "IX",
    name: "Nine of Cups",
    art: "",
    desc: []{"joy", "happiness", "fulfillment", "well-being", "living with ease", "creative bounty", "harmonious relationships"}
    rDesc: []{"greedy", "smug", "taking things for granted", "ignoring work", "no maintenance", "dissatisfaction"}
}

tenCups := new(Card){
    ID: 45,
	numeral: "X",
    name: "Ten of Cups",
    art: "",
    desc: []{"joy in abundance", "lasting happiness", "friendships blossoming", "lengthy romance", "happy homes", "success in all aspects of life", "dreams come true"}
    rDesc: []{"broken home", "moving", "quarrels", "discord", "loss of friendship", "emotional distance"}
}

pageCups := new(Card){
    ID: 46,
	numeral: "p",
    name: "Page of Cups",
    art: "",
    desc: []{"thoughtful", "imaginative", "meditative", "loving", "awareness", "emotional", "dreams", "education"}
    rDesc: []{"emotionally immature", "seduction", "illusion", "creativity blocked"}
}

knightCups := new(Card){
    ID: 47,
	numeral: "k",
    name: "Knight of Cups",
    art: "",
    desc: []{"romantic dreamer", "spiritual", "creative", "playful", "idealistic", "love drawing near", "someone / something drawing near"}
    rDesc: []{"talented but unrealistic", "lazy", "living in a daze", "thinking instead of doing", "false relationships", "narcotics dealer"}
}

queenCups := new(Card){
    ID: 48,
	numeral: "Q",
    name: "Queen of Cups",
    art: "",
    desc: []{"devoted woman", "fantasy and creation", "profound intution", "pleasure", "wisdom", "artistic", "spiritual", "intuition", "sympathetic"}
    rDesc: []{"perverse", "overly emotional", "manipulative", "emotional pain", "depravity", "wickedness"}
}

kingCups := new(Card){
    ID: 49,
	numeral: "K",
    name: "King of Cups",
    art: "",
    desc: []{"emotionally balanced", "quet confidence", "fascination with arts and sciences", "creative", "intuitive", "successful", "moody"}
    rDesc: []{"intense", "difficult to be around", "two-faced", "unfaithful", "anxious or depressed", "hiding from reality", "addiction", "affairs", "overwhelmed", "scandal"}
}

// minor arcana: swords

aceSwords := new(Card){
    ID: 50,
	numeral: "I",
    name: "Ace of Swords",
    art: "",
    desc: []{""}
    rDesc: []{}
}

twoSwords := new(Card){
    ID: 51,
	numeral: "II",
    name: "Two of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

threeSwords := new(Card){
    ID: 52,
	numeral: "III",
    name: "Three of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

fourSwords := new(Card){
    ID: 53,
	numeral: "IV",
    name: "Four of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

fiveSwords := new(Card){
    ID: 54,
	numeral: "V",
    name: "Five of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

sixSwords := new(Card){
    ID: 55,
	numeral: "VI",
    name: "Six of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

sevenSwords := new(Card){
    ID: 56,
	numeral: "VII",
    name: "Seven of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

eightSwords := new(Card){
    ID: 57,
	numeral: "VIII",
    name: "Eight of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

nineSwords := new(Card){
    ID: 58,
	numeral: "IX",
    name: "Nine of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

tenSwords := new(Card){
    ID: 59,
	numeral: "X",
    name: "Ten of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

pageSwords := new(Card){
    ID: 60,
	numeral: "p",
    name: "Page of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

knightSwords := new(Card){
    ID: 61,
	numeral: "k",
    name: "Knight of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

queenSwords := new(Card){
    ID: 62,
	numeral: "Q",
    name: "Queen of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

kingSwords := new(Card){
    ID: 63,
	numeral: "K",
    name: "King of Swords",
    art: "",
    desc: []{}
    rDesc: []{}
}

// minor arcana: pentacles
/*
    .—'‘’’—.
   /        \
  │     $    │
   \        /
    ‘—.,.-—‘    
*/

acePentacles := new(Card){
    ID: 64,
	numeral: "I",
    name: "Ace of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

twoPentacles := new(Card){
    ID: 65,
	numeral: "II",
    name: "Two of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

threePentacles := new(Card){
    ID: 66,
	numeral: "III",
    name: "Three of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

fourPentacles := new(Card){
    ID: 67,
	numeral: "IV",
    name: "Four of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

fivePentacles := new(Card){
    ID: 68,
	numeral: "V",
    name: "Five of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

sixPentacles := new(Card){
    ID: 69,
	numeral: "VI",
    name: "Six of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

sevenPentacles := new(Card){
    ID: 70,
	numeral: "VII",
    name: "Seven of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

eightPentacles := new(Card){
    ID: 71,
	numeral: "VIII",
    name: "Eight of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

ninePentacles := new(Card){
    ID: 72,
	numeral: "IX",
    name: "Nine of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

tenPentacles := new(Card){
    ID: 73,
	numeral: "X",
    name: "Ten of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

pagePentacles := new(Card){
    ID: 74,
	numeral: "p",
    name: "Page of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

knightPentacles := new(Card){
    ID: 75,
	numeral: "k",
    name: "Knight of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

queenPentacles := new(Card){
    ID: 76,
	numeral: "Q",
    name: "Queen of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

kingPentacles := new(Card){
    ID: 77,
	numeral: "K",
    name: "King of Pentacles",
    art: "",
    desc: []{}
    rDesc: []{}
}

// block template
title := new(Card){
    ID: 0,
	numeral:
    name: "",
    art: "",
    desc: []{}
    rDesc: []{}
}