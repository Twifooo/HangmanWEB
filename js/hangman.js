let motATrouver = "";
let lettresDevinees = [];
let lettresUtilisees = [];
let essaisRestants = 6;
let erreurs = 0;
let lettresCorrectes = [];
let lettresAffichees = 0;

const mots = {
    verbes: "verbes.txt",
    mots: "mots.txt",
    noms: "nomP.txt"
};

function startGame() {
    const category = localStorage.getItem("categories");
    const pseudo = localStorage.getItem("pseudo");

    fetch(`/mots/${mots[category]}`)
        .then(response => response.text())
        .then(data => {
            const words = data.split("\n");
            motATrouver = words[Math.floor(Math.random() * words.length)];
            lettresDevinees = Array(motATrouver.length).fill("_");
            lettresUtilisees = [];
            erreurs = 0;
            essaisRestants = 6;
            lettresCorrectes = [];
            lettresAffichees = localStorage.getItem("lettresAffichees");

            if (lettresAffichees > 0) {
                let positions = [];
                while (positions.length < lettresAffichees) {
                    let pos = Math.floor(Math.random() * motATrouver.length);
                    if (!positions.includes(pos)) {
                        positions.push(pos);
                    }
                }
                positions.forEach(pos => {
                    lettresDevinees[pos] = motATrouver[pos];
                });
            }

            localStorage.setItem("mot", motATrouver);
            localStorage.setItem("lettresAffichees", lettresAffichees);

            updateGameDisplay();
            updateHangman();
        })
        .catch(error => {
            console.error("Erreur lors du chargement des mots :", error);
        });
}

function updateGameDisplay() {
    document.getElementById("mot").innerText = lettresDevinees.join(" ");
    document.getElementById("lettresUtilisees").innerText = lettresUtilisees.join(", ");
    document.getElementById("essaisRestants").innerHTML = getHearts(essaisRestants);
}

function getHearts(remaining) {
    let hearts = '';
    for (let i = 0; i < remaining; i++) {
        hearts += '❤️ ';
    }
    return hearts.trim();
}

function makeGuess() {
    const lettre = document.getElementById("lettre").value.toLowerCase();
    const feedbackElement = document.getElementById("feedback");
    if (lettre && !lettresUtilisees.includes(lettre) && lettre.length === 1) {
        lettresUtilisees.push(lettre);
        if (motATrouver.includes(lettre)) {
            for (let i = 0; i < motATrouver.length; i++) {
                if (motATrouver[i] === lettre) {
                    lettresDevinees[i] = lettre;
                }
            }
            lettresCorrectes.push(lettre);
            feedbackElement.innerText = `Bien joué ! La lettre "${lettre}" est dans le mot.`;
            feedbackElement.style.color = "green";
        } else {
            erreurs++;
            essaisRestants--;
            feedbackElement.innerText = `Dommage ! La lettre "${lettre}" n'est pas dans le mot.`;
            feedbackElement.style.color = "red";
        }
        updateGameDisplay();
        updateHangman();
        checkGameStatus();
    } else if (lettresUtilisees.includes(lettre)) {
        feedbackElement.innerText = `La lettre "${lettre}" a déjà été utilisée.`;
        feedbackElement.style.color = "orange";
    } else {
        feedbackElement.innerText = "Veuillez entrer une lettre valide.";
        feedbackElement.style.color = "black";
    }
    document.getElementById("lettre").value = "";
}

function updateHangman() {
    const hangmanCanvas = document.getElementById('hangmanCanvas');
    const ctx = hangmanCanvas.getContext('2d');
    ctx.clearRect(0, 0, hangmanCanvas.width, hangmanCanvas.height);

    switch (erreurs) {
        case 1:
            ctx.beginPath();
            ctx.arc(30, 50, 20, 0, Math.PI * 2, true);
            ctx.stroke();
            break;
        case 2:
            ctx.beginPath();
            ctx.moveTo(30, 70);
            ctx.lineTo(30, 120);
            ctx.stroke();
            break;
        case 3:
            ctx.beginPath();
            ctx.moveTo(30, 80);
            ctx.lineTo(0, 90);
            ctx.stroke();
            break;
        case 4:
            ctx.beginPath();
            ctx.moveTo(30, 80);
            ctx.lineTo(60, 90);
            ctx.stroke();
            break;
        case 5:
            ctx.beginPath();
            ctx.moveTo(30, 120);
            ctx.lineTo(0, 150);
            ctx.stroke();
            break;
        case 6:
            ctx.beginPath();
            ctx.moveTo(30, 120);
            ctx.lineTo(60, 150);
            ctx.stroke();
            break;
    }
}

function checkGameStatus() {
    if (lettresDevinees.join('') === motATrouver) {
        localStorage.setItem('gameStatus', 'win');
        localStorage.setItem('mot', motATrouver);
        localStorage.setItem('essaisRestants', essaisRestants);
        window.location.href = "end_game.html";
    } else if (erreurs >= 6) {
        localStorage.setItem('gameStatus', 'lose');
        localStorage.setItem('mot', motATrouver);
        localStorage.setItem('essaisRestants', essaisRestants);
        window.location.href = "end_game.html";
    }
}

window.onload = function() {
    const pseudo = localStorage.getItem("pseudo");
    const category = localStorage.getItem("categories");
    const lettersDisplayed = localStorage.getItem("lettresAffichees");

    document.getElementById("pseudoDisplay").innerText = pseudo || "Inconnu";
    document.getElementById("categoryDisplay").innerText = category || "Inconnu";

    startGame();
};
