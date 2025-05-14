# Pokédex CLI 🕹️

A command-line Pokédex written in Go that lets you explore the Pokémon world right from your terminal. Built using the PokéAPI, this tool allows you to catch Pokémon, inspect them, and explore different in-game locations.

---

## 📦 Features

- 🔍 `pokedex` — List all caught Pokémon
- 📍 `explore` — Explore next location area and available Pokémon
- ⚔️ `catch <pokemon_name>` — Try to catch a Pokémon
- 🧾 `inspect <pokemon_name>` — View stats and info about a Pokémon
- 🗺️ `map` — Display current location path
- 🔄 `help` — Show all available commands
- ❌ `exit` — Quit the Pokédex

---

## 🛠️ Technologies Used

- [Go](https://golang.org/)
- [PokéAPI](https://pokeapi.co/)
- Go modules, structs, slices, maps, interfaces

---

## 🚀 Getting Started

### 📋 Prerequisites

- Go 1.18 or above installed

### 🧪 Run Locally

```bash
# Clone the repository
git clone https://github.com/keshav78-78/pokedexcli.git

# Navigate into the project
cd pokedexcli

# Run the app
go run main.go
