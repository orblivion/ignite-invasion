This program was created as an exercise as part of an interview process.

# Alien Invasion Simulator

This program simulates an alien invasion. It accepts a file with a map of cities with roads connecting them. It also accepts a number of aliens that will land at random cities on the map.

If any city has two or more aliens, it gets destroyed along with the aliens, and you get an output letting you know what happened. Any remaining aliens can take a road to an adjacent city. These two steps are repeated 10,000 times. Afterwards, the remaining map is printed out.

## Building

```
go build .
```

## Creating a map

Maps have the following format:

```
Austin north=Cambridge
Cambridge east=Austin west=Chicago
Chicago south=Cambridge
```

Each city has one line. There can be up to four roads that connect to another city: `north` `east` `south` and `west`.

Roads can curve, so (for instance) a `north` road coming out of one city need not connect to the `south` road of the connected city. However, it does need to connect to _some_ road in the connected city. For instance, we see going east from Cambridge gets you to Austin, and going north from Austin gets you back to Cambridge.

## Running

(On Linux)

```
./ignite-invasion my-map.txt 20
```

This will load the map located at my-map.txt, and add 20 aliens in starting positions on the map.
