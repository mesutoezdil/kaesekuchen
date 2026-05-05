# Käsekuchen

```
K   K    A    EEEEE SSSS EEEEE K   K U   U  CCC  H   H EEEEE N   N
K  K    A A   E    S     E     K  K  U   U C     H   H E     NN  N
KKK    AAAAA  EEE   SSS  EEE   KKK   U   U C     HHHHH EEE   N N N
K  K  A   A   E       S  E     K  K  U   U C     H   H E     N  NN
K   K A   A   EEEEE SSSS EEEEE K   K  UUU   CCC  H   H EEEEE N   N
```

Dieses Repository dokumentiert den Käsekuchen auf die einzig angemessene Weise: vollständig, ohne Kompromisse und mit dem Ernst, den das Thema verdient.

Der Käsekuchen ist kein einfacher Kuchen. Er ist eine Haltung. Eine Antwort auf die Frage, wie ein Dessert aussieht, das nichts mehr beweisen muss.

> ≪ Wer sich selbst nicht achtet, wird dem, was er im Supermarkt als sogenannten Käsekuchen kauft, natürlich diesen Namen geben und ihn mit Appetit verspeisen. ≫
> Platon, Käsekuchen: Eine Apologie

Vier Dokumente bilden die Grundlage dieses Projekts.

[Das Rezept](rezept.md) beschreibt den kanonischen Käsekuchen: was hineinkommt, in welcher Form und warum. Es richtet sich an jemanden, der backen kann und keine Anleitung braucht.

[Was kein Käsekuchen ist](kein-kaesekuchen.md) behandelt die Produkte, die fälschlicherweise dem Käsekuchen zugerechnet werden. Der Basque Burnt Cheesecake. Der New York Cheesecake. Der japanische Baumwollkäsekuchen. Jedes wird untersucht und in seine richtige Kategorie eingeordnet, die nicht Käsekuchen lautet.

[Verstöße](verstoesse.md) beschreibt die Verhaltensweisen, Zutaten und Entscheidungen, die mit einem ernst zu nehmenden Käsekuchen unvereinbar sind. Einige davon sind weit verbreitet. Das macht sie nicht akzeptabel.

[Das stille Zentrum](das-stille-zentrum.md) ist das vierte Dokument. Es enthält weder ein Rezept noch eine Regelliste. Es fragt, was es bedeutet, dass der Käsekuchen, richtig gemacht, nichts Zusätzliches braucht, und folgt dieser Frage bis zu ihrer Antwort.

In beliebiger Reihenfolge lesen. Die Schlussfolgerungen sind dieselben, unabhängig davon, wo man anfängt.


## kk

Ein CLI für Menschen, die einem Käsekuchen durch ein Terminal begegnen wollen. Das Werkzeug kann nicht backen. Es bringt dich bis zu dem Punkt, an dem du es musst.

Der Quellcode liegt in `cli/`. Benötigt Go 1.22 oder neuer.


### Installation

```sh
git clone https://github.com/mesutoezdil/kaesekuchen
cd kaesekuchen
go build -o kk ./cli/
```

Um `kk` systemweit verfügbar zu machen:

```sh
sudo mv kk /usr/local/bin/kk
```

Ohne sudo-Zugriff, in ein Verzeichnis im PATH verschieben:

```sh
mkdir -p ~/bin
mv kk ~/bin/kk
```

Dann sicherstellen, dass `~/bin` im PATH ist. In `~/.zshrc` oder `~/.bashrc` eintragen, falls noch nicht vorhanden:

```sh
export PATH="$HOME/bin:$PATH"
```


### Befehle

`kk make` geht das Rezept Schritt für Schritt durch. Enter drücken, wenn jeder Schritt abgeschlossen ist. Wer "fruit" oder "quark" eintippt, beendet die Sitzung sofort.

`kk audit` stellt fünf Fragen über das, was gemacht wurde. Gibt entweder Bestanden zurück oder eine Liste von Verstößen mit Codes und Erklärungen.

`kk classify` stellt Ja/Nein-Fragen und gibt eine Klassifizierung zurück. Beendet sich, sobald die Kategorie feststeht.

Mögliche Ausgaben: Kaesekuchen, Quarkkuchen, New York Cheesecake, Basque Burnt Cheesecake, Japanese Cheesecake, undefined.
