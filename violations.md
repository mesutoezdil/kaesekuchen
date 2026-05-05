# Violations

A violation is not a preference. It is a decision that places the result outside the boundary of what can be called Kaesekuchen without misrepresenting it. Some violations are more severe than others. All of them are documented here.

```cuda
typedef enum {
    QUARK_AS_PRIMARY_DAIRY  = 0x01,
    FRUIT_PRESENT           = 0x02,
    NOT_BAKED               = 0x04,
    BISCUIT_BASE            = 0x08,
    SERVED_WARM             = 0x10,
    WRONG_NAME              = 0x20
} ViolationFlag;

__global__ void audit(const Cake* cake, uint32_t* flags) {
    *flags = 0;

    if (cake->primary_dairy == QUARK)             *flags |= QUARK_AS_PRIMARY_DAIRY;
    if (cake->has_fruit)                          *flags |= FRUIT_PRESENT;
    if (!cake->was_baked)                         *flags |= NOT_BAKED;
    if (cake->base == GRAHAM_CRACKER
     || cake->base == DIGESTIVE_BISCUIT)          *flags |= BISCUIT_BASE;
    if (cake->serve_temp_celsius > 8.0f)          *flags |= SERVED_WARM;
    if (!cake->name_is_accurate)                  *flags |= WRONG_NAME;
}

// if (*flags != 0) the product is not Kaesekuchen
// if (*flags == QUARK_AS_PRIMARY_DAIRY) it is not even close
// the severity of each flag is described in the sections below
```


## Using Quark Instead of Cheese

The most common violation and the one with the most defenders. Quark has been used in cakes called Kaesekuchen for long enough that the substitution feels legitimate. It is not.

The name of the cake specifies the ingredient. When that ingredient is replaced with something else, the name no longer applies. A baker who makes a Quark-filled cake and calls it Kaesekuchen is not wrong about the recipe. They are wrong about the name. The cake they made is a Quarkkuchen. It may be a good Quarkkuchen. It is not a Kaesekuchen.

This is the foundational violation. Everything else in this document is secondary to this one.


## Adding Fruit

Fruit does not belong in a Kaesekuchen or on top of one. This includes strawberries (Erdbeeren), cherries (Kirschen), blueberries (Heidelbeeren), raspberries (Himbeeren), and any other fruit in any form including fresh, cooked, pureed, jellied, or candied.

The reasoning is not about flavor compatibility. Some fruits do pair reasonably well with cheese. The reasoning is about what the addition communicates. Adding fruit to a Kaesekuchen filling is an implicit statement that the cheese filling alone is insufficient. If that statement is true, the solution is better cheese, not more ingredients. If the statement is false, then the fruit is unnecessary and should be removed.

A Kaesekuchen that needs a strawberry topping to be interesting has not been made correctly. A Kaesekuchen that has been made correctly does not need anything on top of it.


## No-Bake Preparation

A Kaesekuchen must be baked. This is not a stylistic preference. It is definitional.

Products that use gelatine or agar to set a cold filling are not Kaesekuchen. They are cold-set desserts. The baking process changes the filling at a structural level: proteins in the cheese and eggs denature under heat, the filling sets from the inside out, and the result has a texture and density that cannot be achieved by chilling alone. A chilled, gelatine-set filling is a different product with different properties. It may be good. It is not Kaesekuchen.

No amount of labeling or presentation changes this. If it was not baked, it is not Kaesekuchen.


## Using a Graham Cracker or Biscuit Base

A crust made from crushed biscuits, Graham crackers, Digestives, or any other processed crumb product is not a valid base for Kaesekuchen. These substrates behave differently from a proper Mürbeteig. They absorb moisture from the filling in unpredictable ways, they introduce flavors that are foreign to the Kaesekuchen tradition, and they produce a texture at the base that is incompatible with what the filling expects to rest on.

If a base is used at all, it must be a Mürbeteig or, as a personal alternative, a Karottenkuchen base as described in the recipe. The crumb-crust approach belongs to American cheesecake traditions and should remain there.


## Serving Warm

Kaesekuchen is served cold. Not cool. Not at room temperature. Cold, having spent a minimum of four hours in the refrigerator after cooling completely from the bake.

The structure of the filling is designed to be eaten cold. When warm, the filling is softer than intended, the flavors are less distinct, and the texture is closer to a baked pudding than to Kaesekuchen. A baker who serves Kaesekuchen warm has completed approximately eighty percent of the work correctly and then stopped too early.

Serving warm is also a practical error. A Kaesekuchen removed from the oven before it has fully cooled will often fall apart when cut, because the filling has not fully set. Patience is not optional here.


## Adding Chocolate, Spices, or Flavor Extracts

Chocolate swirls, cinnamon, cardamom, or artificial flavor extracts do not belong in a Kaesekuchen. The filling has a flavor. That flavor is the point. Adding other flavors does not improve it. It competes with it.

Vanille and Zitronenschale are the only flavor additions that belong in the filling, and they are there not to flavor the cake but to support the cheese, which is different. Everything else is a distraction.

A Kaesekuchen with a chocolate swirl is a cake that was almost a Kaesekuchen and then a decision was made. The decision was wrong.


## Adding Food Coloring

There is no version of this that is acceptable. Kaesekuchen does not need color added to it. Its color comes from its ingredients: the cream-yellow of the cheese, the slight browning on the surface from the oven. These are correct colors. Any other color is a sign that something has gone wrong, either in the recipe or in the judgment of the person making it.


## Calling the Result Kaesekuchen When It Is Not

This is the violation that enables all the others. A product made with Quark, or with a biscuit base, or with fruit on top, or without baking, can exist in the world without causing harm as long as it is called what it is. The problem arises when it is called Kaesekuchen.

Misnamed food is more than a labeling problem. It displaces the original. When enough people experience a Quark-filled, fruit-topped, biscuit-based cold-set product labeled Kaesekuchen, the word Kaesekuchen begins to mean that product. The actual Kaesekuchen becomes the exception, then the obscure variant, then the thing people have to specify by adding qualifiers. This process is already underway.

Naming a thing correctly is not pedantry. It is the minimum required to preserve the thing.
