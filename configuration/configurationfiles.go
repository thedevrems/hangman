package configuration

// ConfigFiles initializes and returns a pointer to a Structure_files struct

// Return :
// The .txt file that corresponds to the difficulty chosen by the player if he decided to do it in the language chosen

func ConfigFiles() *Structure_files {
	return &Structure_files{
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
		Path:                   "./managefiles/files/",

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
			"mountain", "street", "road", "boat", "bus", "motorcycle", "bicycle", "train", "party",
			"cake", "chocolate", "candy", "coffee", "tea", "water", "juice", "bread", "butter", "cheese",
			"salad", "soup", "fruit", "vegetable", "carrot", "chicken", "fish", "pork", "lamb", "rice",
			"pasta", "pizza", "sandwich", "ice cream", "sorbet", "cereal", "bar", "honey", "oil", "vinegar",
			"salt", "pepper", "sugar", "spice", "cinnamon", "vanilla", "lemon", "orange", "strawberry", "banana",
			"cherry", "kiwi", "mango", "watermelon", "grape", "peach", "melon", "grapefruit", "nut", "kitten",
			"puppy", "bird", "fish", "snake", "turtle", "guinea pig", "rat", "rabbit", "mouse", "crocodile",
			"elephant", "lion", "tiger", "zebra", "giraffe", "hippopotamus", "bear", "monkey", "shark", "cow",
			"sheep", "pig", "hen", "duck", "turkey", "bee", "butterfly", "insect", "spider", "squirrel",
			"raccoon", "fox", "wolf", "frog", "grasshopper", "dragon", "unicorn",
		},
		EnglishWordsEasy: []string{
			"cat", "dog", "house", "apple", "sun", "moon", "flower", "tree", "garden", "car",
			"table", "chair", "lamp", "book", "door", "window", "wall", "ceiling", "rain", "snow",
			"mountain", "street", "road", "boat", "bus", "motorcycle", "bicycle", "train", "party",
			"cake", "chocolate", "candy", "coffee", "tea", "water", "juice", "bread", "butter", "cheese",
			"salad", "soup", "fruit", "vegetable", "carrot", "chicken", "fish", "pork", "lamb", "rice",
			"pasta", "pizza", "sandwich", "ice cream", "sorbet", "cereal", "bar", "honey", "oil", "vinegar",
			"salt", "pepper", "sugar", "spice", "cinnamon", "vanilla", "lemon", "orange", "strawberry", "banana",
			"cherry", "kiwi", "mango", "watermelon", "grape", "peach", "melon", "grapefruit", "nut", "kitten",
			"puppy", "bird", "fish", "snake", "turtle", "guinea pig", "rat", "rabbit", "mouse", "crocodile",
			"elephant", "lion", "tiger", "zebra", "giraffe", "hippopotamus", "bear", "monkey", "shark", "cow",
			"sheep", "pig", "hen", "duck", "turkey", "bee", "butterfly", "insect", "spider", "squirrel",
			"raccoon", "fox", "wolf", "frog", "grasshopper", "dragon", "unicorn",
		},
		EnglishWordsMiddle: []string{
			"grocery", "library", "supermarket", "restaurant", "school", "university", "hospital", "museum", "cinema", "theater",
			"office", "bank", "post office", "gardener", "pharmacist", "hairdresser", "computer scientist", "architect", "engineer",
			"teacher", "doctor", "nurse", "police officer", "firefighter", "salesperson", "driver", "musician", "painter", "sculptor",
			"photographer", "writer", "journalist", "actor", "director", "host", "manager", "assistant", "electrician", "plumber",
			"gardening", "electricity", "plumbing", "mechanics", "engineering", "surgery", "pediatrics", "dentistry", "pharmacy", "psychology",
			"mathematics", "physics", "chemistry", "biology", "astronomy", "geology", "history", "geography", "philosophy", "sociology",
			"politics", "economics", "anthropology", "linguistics", "musicology", "visual arts", "game theory", "statistics", "strategy",
			"entrepreneurship", "innovation", "management", "marketing", "human resources", "finance", "accounting", "law", "legal", "contractual",
			"international", "national", "local", "comparable", "evaluation", "analysis", "study", "synthesis", "proposal", "conclusion",
			"presentation", "conference", "seminar", "workshop", "training", "course", "supervision", "mentorship",
		},
		EnglishWordsExpert: []string{
			"transcendental", "epistemological", "sociocultural", "unconstitutional", "intrinsically", "hypothetical", "multidimensional", "incommensurable", "democratization", "infrastructure",
			"biocompatibility", "reification", "transubstantiation", "epiphenomenon", "antimatter", "intertextuality", "deconstruction", "representation", "metalanguage", "oxymoron",
			"paradox", "subconscious", "systematization", "heteronormativity", "interpretability", "infinitesimal", "cybernetics", "cognitivism", "intersubjectivity",
			"hypothesis", "axiomatic", "metaphysics", "energetics", "thermodynamics", "cytoplasmic", "revolutionary", "functionality", "paradigm", "dialectics",
			"inherent", "resonance", "cognitive", "transmutation", "vitalism", "interconnection", "intentionality", "autonomy", "influencer", "mediatization",
			"globalization", "teleology", "transdisciplinarity", "dynamic", "singularity", "ubiquity", "telematics", "telescope", "uberization", "multidisciplinarity",
			"intangible", "dissociable", "differentiation", "cumulativity", "perspective", "apprehension", "entropy", "unionism", "recurrence", "hyperconnectivity",
			"multiculturalism", "transnationalism", "antipode", "reflection", "abstraction", "eudaimonism", "heterogeneity", "culturalism", "macroeconomics", "microeconomics",
		},
		EnglishWordsHacker: []string{
			"uncharacteristically", "incomprehensibility", "misinterpretation", "disproportionately", "uncontrollability", "interdisciplinary", "anticonstitutionally", "counterproductive",
			"psychophysiological", "deindustrialization", "unconstitutionalism", "hypercholesterolemia", "characteristically", "overgeneralization", "disenfranchisement", "intercontinental",
			"disproportionateness", "hyperthyroidism", "uncharacteristic", "overcompensation", "dichlorodifluoromethane", "photosynthesis", "individualistically", "overindustrialization",
			"overcommercialization", "interpersonal", "unconventionality", "counterintuitive", "interconnectivity", "overpopulation", "pseudointernational", "unquestionable",
			"autobiographical", "hyperventilation", "institutionalization", "microencapsulation", "uncompartmentalized", "antiestablishment", "transcontinental", "multidimensional",
			"individualization", "psychoneuroimmunology", "reindustrialization", "nonrepresentational", "thermodynamically", "interoperability", "counterproductive", "industrialization", "disproportionate", "intellectualization", "miscommunication", "nontraditional",
			"interdepartmental", "transformation", "overestimation", "dissemination", "antimaterialism", "heterogeneousness", "misrepresentation", "underestimation", "unrepresentative", "decontamination", "nonconformity", "counterarguments",
			"uncompromisingly", "incompatibility", "institutionalize", "unconventionalism", "psychopathological", "transgenderism", "dissonance", "nonparticipation",
			"unconventional", "unprofessionalism", "overcomplicated", "semiautomatic", "reconceptualization", "immunohistochemistry", "interdependence",
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
