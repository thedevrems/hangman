# Nantes Hangman Projet

## Description

Ce projet est une implémentation du jeu **Hangman** (ou Pendu en français) en ligne de commande. Le but du jeu est de deviner un mot aléatoire en proposant des lettres avant que José ne soit totalement pendu. Vous avez **10 tentatives** pour deviner le mot dans le mode de difficulté par default. Le jeu se termine lorsque toutes les lettres du mot sont révélées, quand le mot est trouvé ou lorsque le nombre de tentatives atteint zéro.

## Fonctionnalités

### Configuration
1. Par défaut, le jeu sera configuré en anglais, sans extension ajoutée.
2. Une fois le jeu démarré, vous aurez accès à un panneau de configuration permettant d'ajuster les paramètres du jeu.

### Paramètres du Jeu
Dans le panneau de configuration du jeu, vous aurez accès aux paramètres suivants :
- Language : Ce paramètre vous permet de changer la langue du jeu. Vous pouvez choisir "en" pour l'anglais ou "fr" pour le français.
- VictoryCounter : Active le compteur de victoires et de défaites.
EnableDifficulty : Active les différents niveaux de difficulté (veryeasy, easy, middle, hard, expert, hacker). Par exemple :
    - En mode veryeasy, vous bénéficiez de plus de tentatives. 
    - En mode hacker, aucun indice n'est donné pour deviner le mot, mais le programme révèle un nombre aléatoire de lettres égales à **(longueur du mot / 2) - 1.**. Également, les tolérances d'erreurs vous feront perdre une tantavive.
- AddWordAfterGame : Permet d'ajouter un mot supplémentaire dans le jeu, en fonction du mode de difficulté sélectionné.
- Compatibilité : Toutes ces fonctionnalités peuvent être activées simultanément pour enrichir l'expérience de jeu.

### Paramètres de Base
Le programme Hangman fonctionne de la manière suivante :

1. Fichier de mots : Le programme utilise un fichier .txt (voir la configuration des fichiers dans configuration/configurationfiles.go). Ces fichiers seront automatiquement créés dans managefiles/files s'ils n'existent pas. Si un fichier est vide, il sera rempli avec des valeurs par défaut. Cependant, si le fichier contient déjà des mots, le jeu ne le modifiera pas, vous permettant ainsi d'ajouter des mots manuellement selon vos besoins. Chaque mot est séparé par un saut de ligne (\n).

2. Lancement du jeu :
    - Le programme choisit aléatoirement un mot dans le fichier. Si le jeu détecte un mot invalide, il passera automatiquement à un autre mot, si tous les mots sont invalides le jeu vous le dira automatiquement.
    - Un mot est correct s'il est composé de ces caractères uniquement :
        - Majuscules : A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z
        - Minuscules : a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z
        - Caractères Accentués : À, Á, Â, Ã, Ä, Å, Ç, È, É, Ê, Ë, Ì, Í, Î, Ò, Ó, Ô, Õ, Ö, Ù, Ú, Û, Ü, Ý, à, á, â, ã, ä, å, è, é, ê, ë, ò, ó, ô, õ, ö, ù, ú, û, ü
    - Il révèle un certain nombre de lettres aléatoires du mot, où n = (longueur du mot / 2) - 1 (sauf dans le mode hacker, s'il est activé).
    - Vous disposez de 10 tentatives pour deviner le mot. (sauf dans le mode supereasy, s'il est activé).
    - Vous proposez une lettre en saisissant une entrée au clavier. Si la lettre n'est pas présente, vous perdez une tentative et un message d'erreur s'affiche.
    - Vous avez la possibilité de saisir un mot, cependant si la réponse est incorrect, alors vous perdez deux tentatives, et un message d'erreur s'affiche
    - Si la lettre est correcte, toutes ses occurrences dans le mot sont révélées.
    - Le jeu se poursuit jusqu'à ce que le mot soit entièrement deviné ou que le nombre de tentatives soit épuisé.
    - De plus, des messages vous indiqueront si un mot ou une lettre a déjà été proposé(e).

### Graphisme de José

Un aspect unique de ce projet est José, le personnage qui sera pendu si vous ne parvenez pas à deviner le mot. Un fichier hangman.txt vous est fourni, contenant les 10 étapes graphiques représentant la progression de sa pendaison. Pour le mode de difficulté supereasy, 5 étapes supplémentaires ont été ajoutées, prolongeant ainsi la partie.
 À chaque erreur, la position de José change progressivement jusqu'à ce qu'il soit complètement pendu. Voici un exemple de la septième position de José :
    
      +---+  
      |   |  
      O   |  
     /|   |  
          |  
          |  
    =========

## Prerequisites

Before running the project, ensure that the following are installed on your machine:

- ![Go Version](https://img.shields.io/badge/Go-1.23-blue)
- **Git** to clone the repository

## Installation

To install and run the project, follow these steps:

1. Clonez le dépôt du projet en utilisant Git :
   ```bash
   git clone https://ytrack.learn.ynov.com/git/gremy/hangman.git
   ```
2. Accédez au répertoire du projet :
   ```bash
   cd hangman
   ```
3. Vérifiez que tous les fichiers nécessaires sont présents, notamment cmd/cmd.go, avant de lancer le jeu.

4. Lancez le jeu en exécutant la commande :
   ```bash
   go run cmd/cmd.go
   ```
## Comment tester le projet ?

Pour pouvoir tester le projet il suffit de faire la commande suivante :

```bash
    go run cmd/cmd.go
```
Cette commande doit réaliser à la racine du projet.

## Développeurs du projet

Ce projet a été développé par :

- **[Guibert Rémy](https://github.com/thedevrems)** - Responsable du Projet  - Développeur
- **Laurent Monnier** - Développeur

Si vous souhaitez contribuer ou discuter des idées, n'hésitez pas à nous contacter !

### Remerciements

Un grand merci à tous ceux qui ont contribué au projet et à la communauté open-source pour leur soutien et leurs ressources.


Merci d'avoir tester notre projet !

![Go Version](https://img.shields.io/badge/Go-1.23-blue)
![Licence](https://img.shields.io/badge/License-MIT-green)