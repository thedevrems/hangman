package configuration

// ConfigFiles initializes and returns a pointer to a Structure_files struct

// Return :
// The .txt file that corresponds to the difficulty chosen by the player if he decided to do it in the language chosen

func ConfigFiles() *DataFiles {
	return &DataFiles{
		NameFilesConfigFrenchWordsDefault:  "frenchwordsdefault.txt",
		NameFilesConfigFrenchWordsVeryEasy: "frenchwordsveryeasy.txt",
		NameFilesConfigFrenchWordsEasy:     "frenchwordseasy.txt",
		NameFilesConfigFrenchWordsMiddle:   "frenchwordsmiddle.txt",
		NameFilesConfigFrenchWordsHard:     "frenchwordshard.txt",
		NameFilesConfigFrenchWordsExpert:   "frenchwordsexpert.txt",
		NameFilesConfigFrenchWordsHacker:   "frenchwordshacker.txt",

		NameFilesConfigEnglishWordsDefault:  "englishwordsdefault.txt",
		NameFilesConfigEnglishWordsVeryEasy: "englishwordsveryeasy.txt",
		NameFilesConfigEnglishWordsEasy:     "englishwordseasy.txt",
		NameFilesConfigEnglishWordsMiddle:   "englishwordsmiddle.txt",
		NameFilesConfigEnglishWordsHard:     "englishwordshard.txt",
		NameFilesConfigEnglishWordsExpert:   "englishwordsexpert.txt",
		NameFilesConfigEnglishWordsHacker:   "englishwordshacker.txt",

		NameFilesDisplayResult: "hangman.txt",
		Path:                   "./files/",

		FrenchWordsVeryEasy: []string{
			"chat", "chien", "maison", "pomme", "soleil", "lune", "fleur", "arbre", "jardin", "voiture",
			"table", "chaise", "lampe", "livre", "porte", "fenêtre", "mur", "plafond", "pluie", "neige",
			"montagne", "rue", "avenue", "bateau", "bus", "moto", "vélo", "train", "car", "fête",
			"gâteau", "chocolat", "bonbon", "café", "thé", "eau", "jus", "pain", "beurre", "fromage",
			"salade", "soupe", "fruit", "légume", "carotte", "poulet", "poisson", "porc", "agneau", "riz",
			"pasta", "pizza", "sandwich", "glace", "sorbet", "céréale", "barre", "miel", "huile", "vinaigre",
			"sel", "poivre", "sucre", "épice", "cannelle", "vanille", "citron", "orange", "fraise", "banane",
			"cerise", "kiwi", "mangue", "pastèque", "raisin", "pêche", "melon", "pamplemousse", "fruit", "noix",
			"chaton", "chiot", "oiseau", "poisson", "serpent", "tortue", "cobaye", "rat", "lapin", "souris",
			"crocodile", "éléphant", "lion", "tigre", "zèbre", "girafe", "hippopotame", "ours", "singes", "requin",
			"vache", "mouton", "cochon", "poule", "canard", "dinde", "abeille", "papillon", "insecte", "araignée",
			"écureuil", "raton", "laveur", "renard", "loup", "serpent", "grenouille", "sauterelle", "dragon", "licorne",
		},
		FrenchWordsEasy: []string{
			"chat", "chien", "maison", "pomme", "soleil", "lune", "fleur", "arbre", "jardin", "voiture",
			"table", "chaise", "lampe", "livre", "porte", "fenêtre", "mur", "plafond", "pluie", "neige",
			"montagne", "rue", "avenue", "bateau", "bus", "moto", "vélo", "train", "car", "fête",
			"gâteau", "chocolat", "bonbon", "café", "thé", "eau", "jus", "pain", "beurre", "fromage",
			"salade", "soupe", "fruit", "légume", "carotte", "poulet", "poisson", "porc", "agneau", "riz",
			"pasta", "pizza", "sandwich", "glace", "sorbet", "céréale", "barre", "miel", "huile", "vinaigre",
			"sel", "poivre", "sucre", "épice", "cannelle", "vanille", "citron", "orange", "fraise", "banane",
			"cerise", "kiwi", "mangue", "pastèque", "raisin", "pêche", "melon", "pamplemousse", "fruit", "noix",
			"chaton", "chiot", "oiseau", "poisson", "serpent", "tortue", "cobaye", "rat", "lapin", "souris",
			"crocodile", "éléphant", "lion", "tigre", "zèbre", "girafe", "hippopotame", "ours", "singes", "requin",
			"vache", "mouton", "cochon", "poule", "canard", "dinde", "abeille", "papillon", "insecte", "araignée",
			"écureuil", "raton", "laveur", "renard", "loup", "serpent", "grenouille", "sauterelle", "dragon", "licorne",
		},
		FrenchWordsMiddle: []string{
			"épicerie", "bibliothèque", "supermarché", "restaurant", "école", "université", "hôpital", "musée", "cinéma", "théâtre",
			"bureau", "bureau", "banque", "poste", "jardinier", "pharmacien", "coiffeur", "informaticien", "architecte", "ingénieur",
			"enseignant", "médecin", "infirmière", "policier", "pompiers", "vendeur", "chauffeur", "musicien", "peintre", "sculpteur",
			"photographe", "écrivain", "journaliste", "acteur", "réalisateur", "animateur", "directeur", "assistant", "électricien", "plombier",
			"jardinage", "électricité", "plomberie", "mécanique", "génie", "chirurgie", "pédiatrie", "dentisterie", "pharmacie", "psychologie",
			"mathématiques", "physique", "chimie", "biologie", "astronomie", "géologie", "histoire", "géographie", "philosophie", "sociologie",
			"politique", "économie", "anthropologie", "linguistique", "musicologie", "arts plastiques", "arts visuels", "théorie des jeux", "statistiques", "stratégie",
			"entrepreneuriat", "innovation", "gestion", "marketing", "ressources humaines", "finance", "comptabilité", "droit", "juridique", "contractuel",
			"international", "national", "local", "comparable", "évaluation", "analyse", "étude", "synthèse", "proposition", "conclusion",
			"exposé", "présentation", "conférence", "séminaire", "atelier", "formation", "cours", "atelier", "encadrement", "mentorat",
		},
		FrenchWordsHard: []string{
			"quasar", "phosphorescent", "transcendant", "épistémologie", "sociopolitique", "systémique", "biotechnologie", "nanotechnologie", "anthropocentrique", "néolibéralisme",
			"surconsommation", "durabilité", "biodiversité", "développement durable", "télédétection", "génomique", "microbiome", "phénomène", "accompagnement", "réflexivité",
			"périphérie", "génération", "innovation", "interdisciplinarité", "contemporanéité", "cybersécurité", "cryptographie", "blockchain", "virtualisation", "intelligence",
			"artificielle", "machine", "apprentissage", "algorithme", "robotique", "automatisation", "numérisation", "digitalisation", "architecture", "infrastructure",
			"conception", "élaboration", "réglementation", "standardisation", "protocoles", "opérabilité", "interopérabilité", "flexibilité", "résilience", "sustainability",
			"agilité", "pluridisciplinaire", "cognition", "réhabilitation", "équité", "réciprocité", "systèmes", "complexité", "expérience", "connaissance",
			"expertise", "influence", "interaction", "collaboration", "communication", "réseautage", "intégration", "cohésion", "partenariat",
			"convergence", "cohérence", "synergie", "sustainability", "équilibre", "inclusivité", "écosystème", "société", "environnement", "humanité",
			"philosophie", "éthique", "moralité", "valeurs", "croyance", "spiritualité", "tolerance", "respect", "diversité", "égalité",
		},
		FrenchWordsExpert: []string{
			"transcendantal", "épistémologique", "socioculturel", "anticonstitutionnel", "intrinsèquement", "hypothétique", "multidimensionnel", "incommensurable", "démocratisation", "infrastructure",
			"biocompatibilité", "réification", "transsubstantiation", "épiphénomène", "antimatière", "intertextualité", "déconstruction", "représentation", "métalangage", "oxymore",
			"paradoxe", "subconscient", "systématisation", "hétéronormativité", "interprétabilité", "infinitésimal", "cybernétique", "cognitivisme", "intersubjectivité",
			"hypothèse", "axiomatique", "métaphysique", "énergétique", "thermodynamique", "cytoplasmique", "révolutionnaire", "fonctionnalité", "paradigme", "dialectique",
			"inhérente", "résonance", "cognitif", "transmutation", "vitalisme", "interconnexion", "intentionnalité", "autonomie", "influenceur", "médiatisation",
			"globalisation", "téléologie", "transdisciplinarité", "dynamique", "singularité", "ubiquité", "télématique", "télescopage", "ubérisation", "pluridisciplinarité",
			"intangible", "dissociable", "différenciation", "cumulativité", "perspective", "appréhension", "entropie", "syndicalisme", "récurrence", "hyperconnectivité",
			"multiculturalisme", "transnationalisme", "antipode", "réflexion", "abstraction", "eudémonisme", "hétérogénéité", "culturalisme", "macroéconomie", "microéconomie",
		},
		FrenchWordsHacker: []string{"inconstitutionnel", "démocratisation", "anticonstitutionnel", "interdisciplinarité",
			"surconsommation", "sociopolitiquement", "hyperparamétrisation", "incommensurablement", "transsubstantiation", "pluridimensionnel", "développementaliste", "inappropriabilité", "antidémocratique", "révolutionnairement",
			"hyperconnectivité", "propriétaires privés", "interrelationnel", "réintroduction", "multidimensionnelle", "énergétiquement", "pseudointernational", "désinstitutionnalisation",
			"restructuration", "systématiquement", "antimatière", "télécommunication", "phosphorescence", "démocratiquement", "socioculturellement", "pragmatiquement", "inaccessibilité",
			"transdisciplinarité", "maladministration", "hypersensibilité", "indéfectiblement", "ultraconservateur", "expérimentaliste", "intergouvernemental", "psychopathologie",
			"photodégradation", "antipathiquement", "intersubjectivité", "prétentieux", "scientifiquement", "philosophiquement", "autobiographique", "recontextualisation", "interconnectivité", "démystification",
			"prolétarisation", "désintéressement", "sociohistorique", "réinterprétation", "électroniquement", "anesthésiste", "interculturel", "antimésopotamien",
			"internationalisation", "sociolinguistique", "homogénéisation", "surinformatisation", "antipersonnel", "excommunication", "détenteur de droits", "élaboration", "réductionnisme", "décolonisation",
			"éthiquement", "indivisibilité", "intergalactique", "sociétalement", "élitiste", "inintelligibilité", "précarisation", "médiatisation", "objectivation", "prolétariat", "hypothétisation",
			"pseudoscientifique", "ubiquité", "réflexivité",
		},

		EnglishWordsVeryEasy: []string{
			"cat", "dog", "house", "apple", "sun", "moon", "flower", "tree", "garden", "car",
			"table", "chair", "lamp", "book", "door", "window", "wall", "ceiling", "rain", "snow",
			"mountain", "street", "road", "boat", "bus", "train", "party", "cake", "chocolate",
			"candy", "coffee", "tea", "water", "juice", "bread", "butter", "cheese", "salad",
			"soup", "fruit", "vegetable", "carrot", "chicken", "fish", "pork", "lamb", "rice",
			"pasta", "pizza", "sandwich", "cereal", "honey", "oil", "vinegar", "salt", "pepper",
			"sugar", "spice", "cinnamon", "vanilla", "lemon", "orange", "banana", "cherry",
			"kiwi", "mango", "grape", "peach", "melon", "kitten", "puppy", "bird", "snake",
			"turtle", "rabbit", "mouse", "elephant", "lion", "tiger", "zebra", "giraffe", "bear",
			"cow", "sheep", "pig", "hen", "duck", "bee", "frog", "dragon", "unicorn",
		},

		EnglishWordsEasy: []string{
			"garden", "school", "family", "holiday", "island", "castle", "airport", "balloon",
			"friend", "vacation", "forest", "desert", "mountain", "ocean", "planet", "light", "cloud",
			"breeze", "storm", "bicycle", "rocket", "helmet", "garden", "animal", "flower", "doctor",
			"teacher", "window", "bottle", "pencil", "paper", "music", "piano", "violin", "guitar",
			"drum", "actor", "dancer", "singer", "author", "journal", "bridge", "river", "lake",
			"valley", "journey", "garden", "meadow", "cowboy", "fairy", "queen", "princess", "robot",
			"wizard", "giant", "knight", "chess", "riddle", "puzzle", "shell", "stone", "globe",
		},

		EnglishWordsMiddle: []string{
			"library", "restaurant", "grocery", "supermarket", "university", "hospital", "museum",
			"theater", "office", "bank", "gardener", "architect", "engineer", "teacher", "doctor",
			"nurse", "firefighter", "musician", "painter", "sculptor", "writer", "journalist", "actor",
			"electrician", "mechanics", "surgery", "dentistry", "psychology", "mathematics", "physics",
			"chemistry", "biology", "geology", "philosophy", "sociology", "economics", "linguistics",
			"visualarts", "gametheory", "innovation", "management", "marketing", "finance", "accounting",
			"evaluation", "analysis", "synthesis", "proposal", "seminar", "workshop", "training", "course",
			"mentorship", "conclusion", "conference", "presentation",
		},

		EnglishWordsHard: []string{
			"transcendental", "epistemology", "sociocultural", "unconstitutional", "intrinsically",
			"hypothetical", "multidimensional", "democratization", "infrastructure", "biocompatible",
			"transubstantiation", "epiphenomenon", "antimatter", "intertextual", "deconstruction",
			"metalanguage", "oxymoron", "subconscious", "systematic", "heteronormative", "infinitesimal",
			"cybernetics", "cognitivism", "intersubjectivity", "axiomatic", "metaphysics", "thermodynamics",
			"cytoplasmic", "revolutionary", "functionality", "paradigm", "dialectics", "resonance",
			"intentionality", "autonomy", "teleology", "singularity", "dissociable", "cumulative",
			"apprehension", "entropy", "recurrence", "macroeconomics", "microeconomics",
		},
		EnglishWordsExpert: []string{
			"psychophysiology", "incomprehensibility", "counterproductive", "psychoneuroimmunology",
			"unquestionable", "overgeneralization", "deindustrialization", "overindustrialization",
			"disenfranchisement", "intercontinental", "hyperthyroidism", "overpopulation",
			"pseudointernational", "overcommercialization", "hyperventilation", "institutionalization",
			"microencapsulation", "anticonstitutionally", "photosynthesis", "individualistically",
			"overcompensation", "intellectualization", "uncharacteristically", "counterarguments",
			"uncompartmentalized", "autobiographical", "nonrepresentational", "thermodynamically",
			"counterintuitive", "industrialization", "transformation", "psychopathological",
			"misrepresentation", "immunohistochemistry", "decontamination",
		},
		EnglishWordsHacker: []string{
			"uncharacteristically", "incomprehensibility", "misinterpretation", "disproportionately",
			"uncontrollability", "interdisciplinary", "anticonstitutionally", "counterproductive",
			"psychophysiological", "deindustrialization", "unconstitutionalism", "hypercholesterolemia",
			"overgeneralization", "disenfranchisement", "intercontinental", "overcommercialization",
			"dichlorodifluoromethane", "photosynthesis", "individualistically", "overindustrialization",
			"overpopulation", "pseudointernational", "unquestionable", "autobiographical", "hyperventilation",
			"institutionalization", "microencapsulation", "uncompartmentalized", "antiestablishment",
			"reconceptualization", "counterintuitive", "overcompensation", "interconnectivity",
			"overestimation", "miscommunication", "disproportionate", "psychoneuroimmunology",
			"nontraditional", "transcontinental", "multidimensional", "decontamination",
			"uncompromisingly", "incompatibility", "institutionalize", "unconventionalism",
			"transdisciplinarity", "nonrepresentational", "thermodynamically", "disproportionateness",
			"hyperthyroidism", "photosynthesis", "intellectualization", "counterarguments",
		},

		ContentDisplayResult: `// HangMan Pos 1

=========
		
// HangMan Pos 2

	  | 
	  |
	  | 
	  |
	  |
=========
		
// HangMan Pos 3

  +---+ 
	  |
	  | 
	  |  
	  |
	  | 
=========
		
// HangMan Pos 4

  +---+
  |   | 
	  | 
	  | 
	  | 
	  | 
=========
		
// HangMan Pos 5

  +---+  
  |   |  
  O   |  
	  |  
	  |  
	  |  
=========
		
// HangMan Pos 6

  +---+  
  |   |  
  O   |  
  |   |  
	  |  
	  |  
=========
		
// HangMan Pos 7

  +---+  
  |   |  
  O   |  
 /|   |  
	  |  
	  |  
=========
		
// HangMan Pos 8

  +---+  
  |   |  
  O   |  
 /|\  |  
	  |  
	  |  
=========
		
// HangMan Pos 9

  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
	  |  
=========
		
// HangMan Pos 10

  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
	  |  
=========`,
	}
}
