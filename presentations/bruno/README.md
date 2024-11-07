# We don't talk about Bruno

Below are the notes of my presentation, pay not attention.

## Pre-requisite


1. `docker compose up -d`
2. `cd slides && yarn run dev`
3. Change the password in the collection to `password`

## Intro (2min)

Follow the slides.

## Demo (10min)

Before running the scenario, make sure some products are inserted in the database.

1. Show Bruno UI
2. Show Collection configuration
3. Show environment configuration
4. Show public routes
   * get all products -> show variables / secrets, scripts, assertions and tests
   * get 1 product by id
5. Show product creation route that fails because of authentication
6. Show login route
7. Show auth configuration in the auth pre-script (as well as other auth methods)
8. Inherit auth in the create request and run it.
9. Show Bruno file structure in the IDE 
10. Run the whole collection via CLI
    * `cd collection && bru run --env local --env-var password=password`

## Comparison (3min)

* Reopen the slides
