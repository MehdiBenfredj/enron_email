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




