# Kaesekuchen

This repository documents the Kaesekuchen in the only way it deserves to be documented: completely, without compromise, and with the seriousness the subject demands.

The Kaesekuchen is not a simple cake. It is a position. A stance. An answer to the question of what dessert looks like when it has nothing left to prove.

> ≪ Wer sich selbst nicht achtet, wird dem, was er im Supermarkt als sogenannten Kaesekuchen kauft, natürlich diesen Namen geben und ihn mit Appetit verspeisen. ≫
> Platon, Kaesekuchen: Eine Apologie

> ❝ If a man has no respect for himself, he will of course give that name to the so-called Kaesekuchen he bought at the supermarket and eat it heartily. ❞
> Plato, Kaesekuchen: An Apology

> 『 若人不自重，自然会把超市里买来的所谓Kaesekuchen冠以此名，欣然食之。』
> 柏拉图，《Kaesekuchen申辩篇》

Three documents form the foundation of this project.

[The Recipe](recipe.md) describes the canonical Kaesekuchen: what goes into it, in what form, and why. It is written for someone who already knows how to bake and does not need their hand held.

[What Kaesekuchen Is Not](not-kaesekuchen.md) addresses the products that have been mistakenly grouped with Kaesekuchen over the years. San Sebastián cheesecake. New York cheesecake. The Japanese cotton cheesecake. Each is examined and placed in its correct category, which is not Kaesekuchen.

[Violations](violations.md) covers the behaviors, ingredients, and decisions that are incompatible with a Kaesekuchen that can be taken seriously. Some of these are common. That does not make them acceptable.

[The Still Center](the-still-center.md) is the fourth document. It does not contain a recipe or a list of rules. It asks what it means that the Kaesekuchen, made correctly, needs nothing added to it, and follows that question to where it leads.

Read in any order. The conclusions are the same regardless of where you start.


## kk

A CLI for people who want to encounter a Kaesekuchen through a terminal. The tool cannot bake for you. It can get you to the point where you have to.

Source is in `cli/`. Requires Go 1.22 or later.


### Install

```sh
git clone https://github.com/mesutoezdil/kaesekuchen
cd kaesekuchen
go build -o kk ./cli/
mv kk /usr/local/bin/kk
```


### Commands


#### kk help

```sh
kk help
```

```
kk  Kaesekuchen CLI

commands:
    make        walk through the recipe, step by step
    audit       determine whether what you made is a Kaesekuchen
    classify    identify what kind of cheesecake you have

this tool cannot bake for you.
that part is yours.
```


#### kk make

Walks through the recipe one step at a time. Press enter when each step is done. Typing "fruit" or "quark" at any prompt ends the session immediately.

```sh
kk make
```

```
Kaesekuchen. Step by step.
Press enter when each step is done.

[1/12] Take your eggs out of the refrigerator. Wait until they reach room temperature.
       note: cold eggs will affect the texture. this step is not optional.
       [enter to continue]

[2/12] Prepare your Springform. Butter the interior thoroughly. Set aside.
       [enter to continue]

[3/12] Grate the Zitronenschale from an unwaxed lemon. Set aside.
       note: use an unwaxed lemon. the zest is functional, not decorative.
       [enter to continue]

...

[12/12] Serve cold. No toppings. No fruit. Nothing on top.
        [enter to continue]

if you followed the steps, you have a Kaesekuchen.
if you added fruit at any point, you do not.
```


#### kk audit

Asks five questions about what you made. Returns a pass or a list of violations with codes and explanations.

```sh
kk audit
```

Failing case:

```
Kaesekuchen audit.
Answer honestly. The result reflects your cake, not your intentions.

Primary dairy ingredient (cheese, quark, or other): quark
Was it baked? no
Base type (none, Mürbeteig, Karottenkuchen, Graham cracker, biscuit, or other): graham cracker
Does it contain fruit, or is there fruit on top? yes
Serving temperature (cold, room temperature, or warm): warm

audit failed. 5 violation(s):

  [QUARK_AS_PRIMARY_DAIRY]
  the name is Kaesekuchen. Käse means cheese. this is not a preference.

  [NOT_BAKED]
  a Kaesekuchen must be baked. gelatine is not heat.

  [BISCUIT_BASE]
  processed crumb substrates belong to American cheesecake traditions. they should remain there.

  [FRUIT_PRESENT]
  fruit does not belong in or on a Kaesekuchen. if the cheese needed help, find better cheese.

  [WRONG_TEMPERATURE]
  Kaesekuchen is served cold. the structure of the filling is designed for cold.

this is not a Kaesekuchen.
```

Passing case:

```
Kaesekuchen audit.
Answer honestly. The result reflects your cake, not your intentions.

Primary dairy ingredient (cheese, quark, or other): cheese
Was it baked? yes
Base type (none, Mürbeteig, Karottenkuchen, Graham cracker, biscuit, or other): none
Does it contain fruit, or is there fruit on top? no
Serving temperature (cold, room temperature, or warm): cold

audit passed.

this is a Kaesekuchen.
```


#### kk classify

Asks yes/no questions and returns a classification. Exits as soon as the category is determined.

```sh
kk classify
```

```
Kaesekuchen classifier.
Answer yes or no.

  Was it baked? [yes/no] yes
  Is the top deliberately burnt and the center liquid at serving temperature? [yes/no] no
  Is the texture cloud-soft, and does the cake wobble noticeably when touched? [yes/no] no
  Is the base made from Graham crackers, Digestives, or any crushed biscuit? [yes/no] no
  Is the primary dairy ingredient Quark? [yes/no] no
  Does the cake have fruit in the filling or on top? [yes/no] no
  Is the primary dairy ingredient actual cheese, not Quark, not cream cheese, not a spreadable substitute? [yes/no] yes

classification: Kaesekuchen

this is a Kaesekuchen. this output is rare. handle with care.
```

Other possible outputs: Quarkkuchen, New York Cheesecake, Basque Burnt Cheesecake, Japanese Cheesecake, undefined.


## Auf Deutsch

Die vollständige deutsche Fassung dieses Projekts befindet sich im Ordner [de](de/README.md). Alle Dokumente, einschließlich Rezept, Verstöße, Klassifizierungen und das stille Zentrum, sind dort vollständig auf Deutsch verfasst.
