# Projet de PR72

L'objectif du projet est d'analyser un corpus d'e-mails pour en extraire des informations. L'analyse doit être parallélisée en limitant le nombre de tâches à un une valeur raisonnable dépendante du nombre de processeurs/threads de la machine.

## Le corpus d'e-mails Enron

Suite à la faillite de la société américaine Enron en 2001, un corpus de 517401 e-mails de la société a été publié. Il est disponible sur [ce site](https://www.cs.cmu.edu/~./enron/enron_mail_20150507.tar.gz). Le jeu de données est constitué d'un répertoire par utilisateur (150 au total). Chacun de ces répertoires contient des sous répertoires (les dossiers d'e-mails), contenant chacun des e-mails nommés `X.` où `X` est un nombre positif. Remarquez bien le `.` à la fin du nom.

## L'objectif de l'analyse

L'objectif est de produire un fichier qui recense les communications entre adresses e-mails sous la forme suivante :

```
expediteur1@domain1.ext1: occ1_1:destinataire1_1@domain1_1.ext1_1 occ1_2:destinataire1_2@domain1_2.ext1_2 ... occ1_n1:destinataire1_n1@domain1_n1.ext1_n1
expediteur2@domain2.ext2: occ2_1:destinataire2_1@domain2_1.ext2_1 occ2_2:destinataire2_2@domain2_2.ext2_2 ... occ2_n2:destinataire2_n2@domain2_n2.ext2_n2
...
expediteurN@domainN.extN: occN_1:destinataireN_1@domainN_1.extN_1 occN_2:destinataireN_2@domainN_2.extN_2 ... occN_nN:destinataireN_nN@domainN_nN.extN_nN
```

Par exemple :

```
toto@enron.com: 2:bob@enron.com 5:bill@enron.com
bob@enron.com: 8:toto.enron.com
bill@enron.com: 7:bob@enron.com 4:toto@enron.com 2:contact@sec.gov
```

Chaque ligne commence par un expéditeur unique, suivi par `:` (sans espace avant). Puis, séparés par des espaces, viennent une suite de destinataires précédés par le nombre d'occurrences de mails depuis l'expéditeur vers ce destinataire. Les occurrences sont séparées du destinataire par le caractère `:`.

## Procédé de l'analyse

Pour faire cette analyse et produire le résultat, il faut créer des threads ou des processus, dits _workers_, qui traiteront des e-mails un par un. Le processus/thread principal (qui exécute le `main`), nommé _task dispatcher_, sera chargé d'alimenter les _workers_ en tâches. Chaque _worker_ reçoit une tâche, la traite, stocke le résultat dans un fichier de résultat intermédiaire, puis informe le _task dispatcher_ de la fin de sa tâche, avant d'en recevoir éventuellement une nouvelle. Chaque _worker_ produit un fichier de sortie comme celui défini ci dessus. À la fin, le _task dispatcher_ fusionne les sorties de tous ses _workers_.

## Analyse d'un e-mail

Un e-mail peut contenir parmi ses champs, les champs suivants :

- From : expéditeur
- To : destinataires directs
- Bcc : destinataires en copie cachée
- Cc : destinataires en copie

Ne confondez par ces champs avec les champs X-From, X-To, etc. que nous ne traiterons pas dans ce projet.

X-From étant présent dans tout e-mail, et placé après les 4 champs qui nous intéressent, le traitement d'un e-mail s'arrête dès ce champ atteint. Il est possible que des champs soient manquant. Dans ce cas, si From est vide, ou si aucun destinataire (direct, copie ou copie cachée) n'est trouvé, le mail ne donnera lieu à aucune sortie.

## Tâches du projet

Il vous sera nécessaire pour réaliser le projet :

- de proposer une architecture logicielle tirant parti du parallélisme de votre PC
- qui permette l'analyse telle qu'elle est décrite
- d'implémenter votre architecture et de l'exécuter sur le jeu de données complet d'Enron

Les contraintes de réalisation sont les suivantes :

- Vous réaliserez le projet par groupes de 3 étudiants.
- Le programme doit être implémenté en langage C++ (pouvant contenir du C)
- Le programme ne doit pas faire appel à des bibliothèques/frameworks autres que les bibliothèques standard C (dont `getopt` fait partie) et C++ (vous utiliserez en particulier les conteneurs de la STL). En particulier, vous n'utiliserez ni boost, ni Poco, ni Qt.
- Le programme doit compiler et fonctionner sous Linux. Un Makefile est fourni pour permettre la compilation avec une simple commande `make`.
- L'analyse ne doit pas prendre plus d'une minute
- Le programme ne doit pas fuire de mémoire (vérifier avec `valgrind`)

## Architecture

L'architecture doit permettre l'évolution du programme, par exemple sur de nouvelles tâches. Pour cela, vous vous appuierez sur les propriétés des langages orientés objets et de C++.

Une attention particulière sera apportée au choix du canal de communication entre les _workers_ et le _task dispatcher_. Dans ce projet, les _workers_ doivent en effet être persistants (i.e. la fin d'une tâche ne termine pas le _worker_, mais ce dernier communique avec le _task dispatcher_ pour obtenir une nouvelle tâche). L'arrêt des _workers_ ne se fait qu'avant la fusion des fichiers de sortie intermédiaires.

L'architecture devra être validée avec l'enseignant du TP, et fera l'objet d'une mise en forme en UML à intégrer dans le rapport.

## Analyse des e-mails

L'analyse des e-mails est parallélisée. Vous pourrez utiliser soit des processus (`fork` en C), soit des threads. Les tâches envoyées doivent pouvoir changer si l'on souhaite étendre le programme à d'autres types d'analyses.

## Implémentation

L'implémentation en C++ doit permettre l'application des bonnes pratiques vues sur les premières séances de TP. Pour celà, vous serez amenés à montrer régulièrement le code produit à l'intervenant de TP.

# Livrables

À la fin du projet, vous produirez :

- le code, compilable et utilisable. Il doit être commenté pour aider à sa compréhension (rôle des classes, des variables, algorithmes complexes, etc.)
- un rapport décrivant les étapes de la conception à la réalisation, ainsi qu'un retour d'expérience (RETEX) consistant à réponse à la question "_si je refaisais le projet maintenant avec l'expérience acquise, qu'est-ce que je changerais ?_"
- une présentation avec support de présentation (ppt, odp, beamer, etc.). La présentation reprend les points clé du rapport, et ne comporte pas de démo (l'examinateur doit pouvoir compiler et faire fonctionner le programme sans problème lui-même). La présentation dure 10 minutes, suivies de questions par l(es)'examinateur(s).

# Dates

Les dates suivantes sont les jalons à suivre dans le développement :

- 19 janvier : constitution des groupes, début de la réflexion sur l'architecture
- 25 et 26 janvier : proposition et validation de l'architecture du programme à l'intervenant de TP.
- 2 février : prototype du programme, séquentiel uniquement
- 8 et 9 février : prototype du programme, traitement des e-mails parallèle
- 14 février : envoi du rapport, et du code
- 16 février : présentations des projets

# Conseils

Il est fortement conseillé de commencer le projet en y travaillant beaucoup au début pour mieux cerner le rythme nécessaire pour le faire aboutir et éviter d'entrer en concurrence avec les épreuves de fin de période.

Il est également fortement conseillé de poser des questions à l'intervenant de TP : le sujet est complexe et il vaut mieux déjouer le plus vite possible les incompréhensions/zones floues pour vous.

Enfin, il vous sera nécessaire de travailler ensemble plutôt que chacun de son côté, mettez donc en oeuvre une méthode de gestion de projet pour vous assurer du bon déroulement (répartition des tâches, suivi, transmission des acquis à vos camarades du groupes, etc.).
