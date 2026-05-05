# Verstöße

Ein Verstoß ist keine Präferenz. Es ist eine Entscheidung, die das Ergebnis außerhalb der Grenze dessen platziert, was man Käsekuchen nennen kann, ohne es falsch darzustellen. Einige Verstöße sind schwerwiegender als andere. Alle sind hier dokumentiert.

```cuda
typedef enum {
    QUARK_ALS_HAUPTZUTAT   = 0x01,
    OBST_VORHANDEN         = 0x02,
    NICHT_GEBACKEN         = 0x04,
    KEKSBODEN              = 0x08,
    WARM_SERVIERT          = 0x10,
    FALSCH_BEZEICHNET      = 0x20
} VerstossFlag;

__global__ void pruefen(const Kuchen* kuchen, uint32_t* flags) {
    *flags = 0;

    if (kuchen->hauptzutat == QUARK)              *flags |= QUARK_ALS_HAUPTZUTAT;
    if (kuchen->hat_obst)                         *flags |= OBST_VORHANDEN;
    if (!kuchen->wurde_gebacken)                  *flags |= NICHT_GEBACKEN;
    if (kuchen->boden == GRAHAM_CRACKER
     || kuchen->boden == DIGESTIVE_KEKS)          *flags |= KEKSBODEN;
    if (kuchen->serviertemperatur_celsius > 8.0f) *flags |= WARM_SERVIERT;
    if (!kuchen->bezeichnung_korrekt)             *flags |= FALSCH_BEZEICHNET;
}

// wenn (*flags != 0) ist das Produkt kein Käsekuchen
// wenn (*flags == QUARK_ALS_HAUPTZUTAT) ist es nicht einmal annähernd einer
// der Schweregrad jedes Flags wird in den folgenden Abschnitten beschrieben
```


## Quark statt Käse

Der häufigste Verstoß und der mit den meisten Verteidigern. Quark wird seit so langer Zeit in Kuchen verwendet, die Käsekuchen heißen, dass die Substitution legitim wirkt. Das ist sie nicht.

Der Name des Kuchens gibt die Zutat vor. Wenn diese Zutat durch etwas anderes ersetzt wird, gilt der Name nicht mehr. Ein Bäcker, der einen Quark-gefüllten Kuchen macht und ihn Käsekuchen nennt, irrt sich nicht beim Rezept. Er irrt sich beim Namen. Der Kuchen, den er gemacht hat, ist ein Quarkkuchen. Es kann ein guter Quarkkuchen sein. Es ist kein Käsekuchen.

Das ist der grundlegende Verstoß. Alles andere in diesem Dokument ist ihm untergeordnet.


## Obst hinzufügen

Obst gehört nicht in einen Käsekuchen und nicht obendrauf. Das gilt für Erdbeeren, Kirschen, Heidelbeeren, Himbeeren und jedes andere Obst in jeder Form, ob frisch, gekocht, püriert, geliert oder kandiert.

Die Begründung hat nichts mit Geschmackskompatibilität zu tun. Einige Früchte passen durchaus zu Käse. Die Begründung betrifft das, was die Hinzufügung kommuniziert. Obst zur Füllung hinzuzufügen ist ein implizites Statement, dass die Käsefüllung allein unzureichend ist. Wenn dieses Statement zutrifft, lautet die Lösung besserer Käse, nicht mehr Zutaten. Wenn es nicht zutrifft, ist das Obst überflüssig und sollte entfernt werden.

Ein Käsekuchen, der eine Erdbeerdeko braucht, um interessant zu sein, wurde nicht richtig gemacht. Ein Käsekuchen, der richtig gemacht wurde, braucht nichts obendrauf.


## Ohne Backen

Ein Käsekuchen muss gebacken werden. Das ist keine stilistische Präferenz. Es ist definitorisch.

Produkte, die Gelatine oder Agar verwenden, um eine kalte Füllung zu setzen, sind kein Käsekuchen. Sie sind kaltgestellte Desserts. Der Backprozess verändert die Füllung auf struktureller Ebene: Proteine in Käse und Eiern denaturieren unter Hitze, die Füllung setzt von innen nach außen, und das Ergebnis hat eine Textur und Dichte, die durch Kühlen allein nicht erreichbar ist. Eine kaltgestellte, gelatinegebundene Füllung ist ein anderes Produkt mit anderen Eigenschaften. Es kann gut sein. Es ist kein Käsekuchen.

Keine Menge an Beschriftung oder Präsentation ändert das. Wenn es nicht gebacken wurde, ist es kein Käsekuchen.


## Graham-Cracker- oder Keksboden

Ein Boden aus zerdrückten Keksen, Graham Crackern, Digestives oder anderen verarbeiteten Krümelprodukten ist kein gültiger Boden für Käsekuchen. Diese Substrate verhalten sich anders als ein ordentlicher Mürbeteig. Sie absorbieren Feuchtigkeit aus der Füllung auf unvorhersehbare Weise, bringen Aromen mit, die der Käsekuchentradition fremd sind, und erzeugen am Boden eine Textur, die mit dem unvereinbar ist, was die Füllung erwartet.

Wenn ein Boden verwendet wird, muss es ein Mürbeteig sein oder, als persönliche Alternative, ein Karottenkuchen-Boden wie im Rezept beschrieben. Der Keksboden-Ansatz gehört zu amerikanischen Cheesecake-Traditionen und sollte dort bleiben.


## Warm Servieren

Käsekuchen wird kalt serviert. Nicht lauwarm. Nicht bei Zimmertemperatur. Kalt, nachdem er mindestens vier Stunden im Kühlschrank war, nachdem er vollständig aus dem Backofen abgekühlt ist.

Die Struktur der Füllung ist darauf ausgelegt, kalt gegessen zu werden. Warm ist die Füllung weicher als vorgesehen, die Aromen sind weniger distinkt, und die Textur kommt einem Pudding näher als einem Käsekuchen. Ein Bäcker, der Käsekuchen warm serviert, hat ungefähr achtzig Prozent der Arbeit richtig gemacht und dann zu früh aufgehört.

Warm servieren ist auch ein praktischer Fehler. Ein Käsekuchen, der aus dem Ofen genommen wird, bevor er vollständig abgekühlt ist, fällt beim Schneiden oft auseinander, weil die Füllung noch nicht vollständig gesetzt hat. Geduld ist hier keine Option.


## Schokolade, Gewürze oder Aromen hinzufügen

Schokoladenwirbel, Zimt, Kardamom oder künstliche Aromen gehören nicht in einen Käsekuchen. Die Füllung hat einen Geschmack. Dieser Geschmack ist der Punkt. Andere Aromen hinzuzufügen verbessert ihn nicht. Sie konkurrieren damit.

Vanille und Zitronenschale sind die einzigen Aromaten, die in die Füllung gehören, und sie sind nicht da, um den Kuchen zu aromatisieren, sondern um den Käsegeschmack zu unterstützen. Das ist ein Unterschied. Alles andere ist Ablenkung.

Ein Käsekuchen mit Schokoladenwirbel ist ein Kuchen, der fast ein Käsekuchen war, und dann wurde eine Entscheidung getroffen. Die Entscheidung war falsch.


## Lebensmittelfarbe hinzufügen

Es gibt keine Version davon, die akzeptabel ist. Käsekuchen braucht keine hinzugefügte Farbe. Seine Farbe kommt von seinen Zutaten: das Cremgelb des Käses, das leichte Bräunen der Oberfläche aus dem Ofen. Das sind die richtigen Farben. Jede andere Farbe ist ein Zeichen, dass etwas schiefgelaufen ist, entweder im Rezept oder im Urteilsvermögen der Person, die ihn gemacht hat.


## Das Ergebnis Käsekuchen nennen, wenn es keiner ist

Das ist der Verstoß, der alle anderen ermöglicht. Ein Produkt aus Quark, mit Keksboden, mit Obst obendrauf oder ohne Backen kann in der Welt existieren, ohne Schaden anzurichten, solange es heißt, was es ist. Das Problem entsteht, wenn es Käsekuchen genannt wird.

Falsch benanntes Essen ist mehr als ein Etikettierungsproblem. Es verdrängt das Original. Wenn genug Menschen ein Quark-gefülltes, obst-belegtes, keks-bodenbasiertes kaltgestelltes Produkt namens Käsekuchen erleben, beginnt das Wort Käsekuchen, dieses Produkt zu bedeuten. Der eigentliche Käsekuchen wird zur Ausnahme, dann zur obskuren Variante, dann zu dem Ding, das man durch Zusätze spezifizieren muss. Dieser Prozess ist bereits im Gange.

Ein Ding richtig zu benennen ist keine Pedanterie. Es ist das Minimum, das erforderlich ist, um das Ding zu erhalten.
