# 🐍🎲 Snake and Ladder – Production-Style LLD Implementation in Go

A modular and extensible implementation of the classic **Snake and Ladder** game written in **Go**, designed using clean architecture and strong Low-Level Design (LLD) principles.

This project focuses on building a **scalable**, **testable**, and **maintainable** game engine by properly separating responsibilities across packages.

---

## 🏗 Architecture Overview

Game (Orchestrator)
│  
├── Board   (Manages snakes, ladders, and position resolution)  
├── Dice    (Interface-based randomness abstraction)  
└── Player  (Encapsulated player state)  

---

## 📦 Component Breakdown

### 🎮 Game
- Controls turn rotation
- Handles win conditions
- Applies overshoot rules
- Orchestrates interactions between components

### 🗺 Board
- Maintains snakes & ladders mapping
- Resolves player positions in **O(1)** time using maps
- Validates board configuration defensively

### 🎲 Dice
- Interface-based design
- Supports multiple implementations
- Easily mockable for deterministic testing

### 👤 Player
- Fully encapsulated state
- Controlled mutations via getters/setters
- Clean domain modeling

---

## ✨ Key Highlights

- ✅ Fully encapsulated player state  
- ✅ Interface-based `Dice` for testability  
- ✅ Defensive input validation for board setup  
- ✅ O(1) snake & ladder resolution  
- ✅ Clean turn rotation logic  
- ✅ Overshoot rule handling  
- ✅ Extendable architecture  
- ✅ No global state usage  
- ✅ Production-style folder structure  

---

## 📌 Design Principles Applied

- Single Responsibility Principle (SRP)
- Dependency Inversion Principle (DIP)
- Encapsulation
- Clean Package Layering
- Interface-Driven Design
- Deterministic Testing Support

---

## 🧠 Why This Project?

This implementation is not just a game — it demonstrates:

- Real-world backend architecture thinking
- Clean Go package structuring
- Testability via dependency injection
- Production-level code separation
- LLD interview readiness

---

## 🔧 Future Enhancements

- 🎲 Multiple dice support
- 📝 Event logging system
- 👀 Observer pattern for game events
- ⚡ Concurrency support (goroutines for simulation mode)
- 🧪 Unit testing with mock dice
- 🎮 CLI / REST API wrapper

---

## 🚀 How to Run

```bash
go run cmd/main/main.go
