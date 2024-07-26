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
            document.getElementById('createArticle').style.display = 'block';
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