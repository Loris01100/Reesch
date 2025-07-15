# 🃏 Reesch - Application de gestion de cartes

**Reesch** est une application complète permettant de gérer une collection de cartes dans un univers fictif personnalisé. Elle est composée de deux parties : un frontend Angular et un backend Node.js/Express.

---

## 📦 Technologies utilisées

- **Frontend** : [Angular](https://angular.io/)
- **Backend** : [Node.js](https://nodejs.org/) + [Express](https://expressjs.com/)
- **CORS** : pour l'accès API
- **MongoDB / PostgreSQL** (à adapter selon ta stack)
- **TypeScript** : utilisé sur les deux côtés

---

## 📁 Structure du projet

/
├── frontend/ # Application Angular
├── backend/ # API Express
├── .gitignore
├── LICENSE
└── README.md

---

## 🚀 Installation

### 1. Cloner le projet
```bash
git clone https://github.com/ton-user/reesch-universe.git
cd reesch-universe


# Frontend
cd frontend
npm install

# Backend
cd ../backend
npm install

## Démarrage du back-end
cd backend
npm run dev

## Démarrage du Front
cd ../frontend
ng serve

📚 Fonctionnalités prévues
Classement par rareté, type et univers

Historique et lore de chaque carte

Système de collection par utilisateur

Ouverture de boosters

