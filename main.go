package main

import "fmt"

//librairy (entrée sortie principalement)

func main() {
	println("Au blackjack, la valeur total de votre main ne dois pas dépacer 21\nLà personne qui a la main la plus forte gagne")
	// Création des joueurs
	joueurs := []Player{
		NewPlayer(1, "1st"),
		NewPlayer(2, "2nd"),
	}

	croupier := NewPlayer(0, "Croupier")

	//initialisation et mélange du deck
	deck := createDeck()
	shuffle(deck)

	//Creation du code du jeu de blackjack

	//On donne les cartes aux joueurs
	for i := 0; i < 2; i++ {
		for _, joueur := range joueurs {
			giveRandomCard(&joueur, &deck)
		}
	}

	//On donne les cartes au croupier
	for i := 0; i < 2; i++ {
		giveRandomCard(&croupier, &deck)
	}

	//On commence le jeu

	//Affichage de la main du croupier
	fmt.Println("Main du croupier:")
	for _, card := range croupier.hand {
		fmt.Printf(" - %s de %s\n", card.value, card.symbol)
	}

	//Tour des joueurs
	for _, joueur := range joueurs {
		for {
			joueur.PrintDetails()
			total := calculateHand(joueur.hand)
			fmt.Printf("Total des points : %d\n", total)

			if total >= 21 {
				break
			}

			var action string
			fmt.Println("Voulez vous tirer une autre carte? (oui/non)")
			fmt.Scanln(&action)

			if action == "oui" {
				giveRandomCard(&joueur, &deck)
			} else {
				break
			}
		}
	}

	//Tour du croupier
	for calculateHand(croupier.hand) < 17 {
		giveRandomCard(&croupier, &deck)
	}

	//Afficher a nouveau la main du croupier
	fmt.Println("Main du croupier:")
	for _, card := range croupier.hand {
		fmt.Printf(" - %s de %s\n", card.value, card.symbol)
	}

	//Comparaison des mains pour déclarer les gagnants
	croupierTotal := calculateHand(croupier.hand)
	for _, joueur := range joueurs {
		joueurTotal := calculateHand(joueur.hand)
		if (joueurTotal > croupierTotal) && (joueurTotal < 21) || (croupierTotal > 21) {
			fmt.Printf("Le joueur %s a la main la plus forte!\n", joueur.name)
			fmt.Printf("La valeur de la main du croupier est: %d\n", croupierTotal)
		} else {
			fmt.Printf("Le croupier avait une main plus forte que le joueur %s!\n", joueur.name)
			fmt.Printf("La valeur de la main du croupier est: %d\n", croupierTotal)
		}
	}
}
