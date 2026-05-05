# Was kein Käsekuchen ist

Mehrere Produkte existieren in der Welt, die gewisse Merkmale mit dem Käsekuchen teilen, aber kein Käsekuchen sind. Das ist keine Kritik an diesen Produkten. Die meisten von ihnen sind gut. Einige sind hervorragend. Der Punkt ist schlicht, dass Benennung eine Rolle spielt, und etwas Käsekuchen zu nennen, das kein Käsekuchen ist, erzeugt Verwirrung, die sich ansammelt, bis das Original nicht mehr erkennbar ist.

Die folgenden Einträge sind keine Angriffe. Sie sind Klassifizierungen.

```cuda
typedef enum {
    KAESEKUCHEN,
    QUARKKUCHEN,
    NEW_YORK_CHEESECAKE,
    BASQUE_BURNT_CHEESECAKE,
    JAPANISCHER_KAESEKUCHEN,
    UNDEFINIERT
} Kuchenkategorie;

__global__ void klassifizieren(const Kuchen* eingabe, Kuchenkategorie* ausgabe) {
    if (!eingabe->wurde_gebacken)
        { *ausgabe = UNDEFINIERT;              return; }  // kein Kuchen

    if (eingabe->oberflaeche_bewusst_verbrannt && eingabe->mitte_fluessig_beim_servieren)
        { *ausgabe = BASQUE_BURNT_CHEESECAKE;  return; }  // gut. kein Käsekuchen.

    if (eingabe->textur == WOLKENWEICH && eingabe->wackelt_beim_beruehren)
        { *ausgabe = JAPANISCHER_KAESEKUCHEN;  return; }  // beeindruckend. kein Käsekuchen.

    if (eingabe->boden == GRAHAM_CRACKER || eingabe->boden == DIGESTIVE)
        { *ausgabe = NEW_YORK_CHEESECAKE;      return; }  // gut. kein Käsekuchen.

    if (eingabe->hauptzutat == QUARK)
        { *ausgabe = QUARKKUCHEN;              return; }  // verbreitet. kein Käsekuchen.

    if (eingabe->hat_obst_oben || eingabe->hat_obst_in_fuellung)
        { *ausgabe = UNDEFINIERT;              return; }  // was hat man sich dabei gedacht.

    *ausgabe = KAESEKUCHEN;
    // geschafft
    // diese Ausgabe ist selten
    // mit Sorgfalt behandeln
}
```

```
    Klassifizierungsübersicht

                              gebacken   Boden             Füllung        Kern beim Servieren   Urteil
    Käsekuchen                     ja   keiner od. Karotte  echter Käse    gesetzt · kalt         korrekt
    Quarkkuchen                    ja   Mürbeteig           Quark          gesetzt · kalt         falscher Name
    New York Cheesecake            ja   Graham Cracker      Frischkäse     gesetzt · kalt         anderes Produkt
    Basque Burnt Cheesecake        ja   keiner              Frischkäse     flüssig · warm         anderes Produkt
    Japanischer Käsekuchen         ja   keiner              Frischkäse     gesetzt · kalt · luftig anderes Produkt
    No-Bake-Variante             nein   Kekskrümel          beliebig       nie gesetzt            kein Kuchen
```


## Basque Burnt Cheesecake

Der Basque Burnt Cheesecake stammt aus La Viña, einem Restaurant in der Altstadt von San Sebastián. Es ist eine Frischkäsefüllung ohne Boden, bei sehr hoher Hitze gebacken, bis die Oberseite tief und absichtlich verbrannt ist, mit einer Mitte, die bei Serviertemperatur bewusst flüssig bleibt. Es ist eines der interessanteren Desserts der letzten dreißig Jahre, und die Technik dahinter ist bemerkenswert durchdacht.

Es ist kein Käsekuchen.

Die Designziele sind grundlegend verschieden. Der Basque Burnt Cheesecake lebt von Kontrast: dem bitteren, karamellisierten Äußeren gegen das kühle, flüssige Innere. Der Reiz liegt in der Spannung zwischen diesen beiden Zuständen. Ein Käsekuchen hat keine solche Spannung. Er soll vollständig durchgebacken, vollständig gesetzt, gleichmäßig von Rand bis Mitte und kalt serviert werden. Das Ziel ist Konsistenz, nicht Kontrast. Das Ziel ist, dass jeder Bissen dem letzten gleicht, weil die Füllung der Punkt ist und die Füllung sich nicht aufführen muss.

Die verbrannte Oberfläche des Basque Burnt Cheesecake ist sein definierendes Merkmal. Die verbrannte Oberfläche eines Käsekuchens ist ein Backfehler. Das sind nicht dieselben Produkte unter verschiedenen Ästhetiken. Sie sind verschiedene Produkte.


## New York Cheesecake

New York Cheesecake verwendet Frischkäse als primäre Milchkomponente, hat einen Graham-Cracker-Boden und wird fast immer im Wasserbad gebacken, um die Oberfläche vor direkter Hitze zu schützen. Er wird viele Stunden vor dem Servieren gekühlt. Er ist dicht, reichhaltig und sehr gut darin, das zu sein, was er ist.

Er ist kein Käsekuchen.

Der Graham-Cracker-Boden ist das erste Problem. Graham Cracker sind ein verarbeitetes Keksprodukt mit einer spezifischen Süße und Krümelstruktur, die in der deutschen Backtradition kein Äquivalent hat. Zu einem Keksboden gepresst, verhält er sich in jeder relevanten Dimension anders als ein Mürbeteig: wie er Feuchtigkeit aus der Füllung aufnimmt, wie er gegen den Käse schmeckt, wie er seine Form über Zeit behält.

Der Frischkäse ist das zweite Problem, aber nicht auf die erwartete Weise. Frischkäse ist als Backzutat nicht von Grund auf falsch. Das Problem liegt darin, dass er ein spezifisches Geschmacksprofil hat, eine bestimmte Schärfe und Reichhaltigkeit, die eine bestimmte Art von Kuchen erzeugt. Dieser Kuchen ist New York Cheesecake. Er ist kein Käsekuchen. Die Geschmacksdistanz zwischen beiden ist so groß, dass keine Anpassung der Technik sie überbrücken kann.

New York Cheesecake ist sein eigenes Ding. Er hat sich zu seinen eigenen Bedingungen durchgesetzt. Er braucht das Käsekuchen-Label nicht, und der Käsekuchen profitiert nicht von der Verbindung.


## Japanischer Käsekuchen

Der japanische Käsekuchen, manchmal auch Soufflé-Käsekuchen oder Baumwollkäsekuchen genannt, wird durch das Unterfalten von steif geschlagenem Eiweiß in einen Frischkäseteig und Backen im Wasserbad bei niedriger Temperatur hergestellt. Das Ergebnis ist ein extrem leichter, wolkenweicher Kuchen, der beim Berühren zittert und beim Abkühlen leicht zusammenfällt. Die Textur ist einem Chiffon-Kuchen näher als allem in der Käsekuchentradition.

Er ist kein Käsekuchen.

Allein die Textur ist disqualifizierend. Käsekuchen hat Dichte. Wenn man eine Gabel hineindrückt, sollte die Füllung sauber nachgeben und die Schnittkante halten. Der japanische Käsekuchen gibt unter Druck nach und springt zurück. Das sind entgegengesetzte Verhaltensweisen aus entgegengesetzten Absichten. Der japanische Käsekuchen versucht, so leicht wie möglich zu sein. Der Käsekuchen versucht nichts außer zu sein, was er ist, was einschließt, dass er substanziell ist.

Die Technik, diese Textur konsistent herzustellen, ist anspruchsvoller, als die meisten Menschen erkennen. Das Produkt verdient Anerkennung zu seinen eigenen Bedingungen. Diese Bedingungen sind nicht Käsekuchen.


## Quarkkuchen

Dieser Punkt ist heikel, weil Quarkkuchen lange genug als Käsekuchen verkauft und beschrieben wurde, dass viele Menschen glauben, sie seien dasselbe. Das sind sie nicht.

Quark ist ein Frischmilchprodukt. Er ist kein Käse in dem Sinne, den der Name Käsekuchen impliziert. Die Textur, die er beim Backen erzeugt, ist leichter und zarter als eine Füllung aus echtem Käse. Der Geschmack ist milder. Das Ergebnis ist angenehm, aber es ist kein Käsekuchen.

Das Wort Käse in Käsekuchen ist kein historischer Zufall und keine lockere Annäherung. Es bedeutet Käse. Wenn der Kuchen Quark statt Käse enthält, ist er ein Quarkkuchen. Dieser Unterschied ist wichtig, weil beide Kuchen anders schmecken, sich beim Backen anders verhalten und vom Esser unterschiedliche Erwartungen erfordern.

Einen Quarkkuchen als Käsekuchen zu bezeichnen ist keine Lüge im Sinne einer Täuschungsabsicht. Es ist eine Lüge im Sinne eines Kategorienfehlers, der so oft wiederholt wurde, dass er wahr zu sein begann. Kategorienfehler werden durch Wiederholung nicht korrekt.


## Käsekuchen mit Obstbelag

Das ist keine separate Käsekuchentradition, sondern eine Modifikation, die auf verschiedene Käsekuchen angewendet wird, einschließlich solcher, die sonst als Käsekuchen gelten würden. Die Modifikation besteht darin, Obst, ob frisch, gekocht oder in Geleeform, auf den fertigen Kuchen zu legen.

Ein Käsekuchen mit Erdbeerobendrauf ist kein Käsekuchen. Er ist ein anderes Produkt. Der Belag verändert das Esserlebnis auf eine Weise, die nicht additiv ist. Das Obst bringt Süße mit, die mit dem Käse konkurriert, Feuchtigkeit, die die Oberfläche der Füllung beeinflusst, und visuelle Betonung, die die Aufmerksamkeit von der Füllung selbst abzieht. Das Ergebnis ist ein Kuchen, der versucht, mehr Menschen zu gefallen, indem er mehr Dinge in sich hat, auf Kosten der Klarheit darüber, was er eigentlich ist.

Ein Käsekuchen hat keinen Belag. Wenn er einen hat, ist er kein Käsekuchen mehr. Das ist definitorisch, nicht präferentiell.
