document.addEventListener('DOMContentLoaded', async () => {
    try {
        const token = document.cookie.split('; ').find(row => row.startsWith('jwt=')).split('=')[1];
        console.log(token)
        const response = await fetch('/api/authent/isLogin', {
            method: 'GET',
            headers: {
                'Authorization': token,
                'Content-Type': 'application/json'
            },
        });
        if (response.ok) {
            console.log("test")
            document.getElementById('createArticle').style.display = 'block';
            document.getElementById('login').style.display = 'none';
            document.getElementById('logout').style.display = 'block';
        } else {
            console.error('Error:', data.message);
            document.getElementById('createArticle').style.display = 'none';
        }
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('createArticle').style.display = 'none';
    }
});

const svgNotification = function () {
    const notificationContainer = document.getElementById('notification-container');
    const imageContainer = document.getElementById('image-container');
    let img = document.createElement('img');
    
    if (notificationContainer.children.length === 0) {
      img.src = '../public/img/svg/notification-bell.svg';
    } else {
      img.src = '../public/img/svg/notification-alert.svg';
    }
    img.className ="w-12 h-12";
    imageContainer.appendChild(img);
}

document.addEventListener('DOMContentLoaded', svgNotification);

const socket = new WebSocket('ws://localhost:3000/ws');

socket.onmessage = function(event) {
    const message = event.data;
    console.log('Notification received:', message);
    // Affiche la notification sur la page
    const notificationElement = document.createElement('div');
    notificationElement.innerText = message;
    document.body.appendChild(notificationElement);
};

socket.onopen = function(event) { console.log('WebSocket connection established.'); };

socket.onclose = function(event) { console.log('WebSocket connection closed.'); };

async function logout(){
    try {
        const response = await fetch("http://127.0.0.1:3000/api/authent/logout",{
            method: 'GET'
        });
        if (response.ok){
            document.getElementById('login').style.display = 'block';
            document.getElementById('logout').style.display = 'none';
        }
    }catch (error) {
        console.error("Erreur lors de la déconnexion:", error);
        alert("Une erreur est survenue. Veuillez réessayer plus tard.");
    }
}

async function searchBar() {
    const searching = document.getElementById('search').value;
    console.log("searching bar say:", searching);

    try {
        const response = await fetch('/api/article/by-name', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ 'name': searching })
        });

        const data = await response.json();

        if (response.ok) {
            if(data.article.StatusID == 1){
                alert("l'article du même nom n'a pas encore etais verifier")    
            }else{
                window.location.href = "/article/" + data.article.ID;
            }
        } else {
            alert(data.error);
        }
    } catch (error) {
        console.error("Erreur lors de la recherche de l'article:", error);
        alert("Une erreur est survenue. Veuillez réessayer plus tard.");
    }
}
