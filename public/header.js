let userId
document.addEventListener('DOMContentLoaded', async () => {
    try {
        var token = null
        const jwtCookie = document.cookie.split('; ').find(row => row.startsWith('jwt-'));
        const linkElement = document.getElementById('profiLink');
        if (jwtCookie) {
            token = jwtCookie.split('=')[1];
        }
        console.log("token",token)
        if (token){
            const response = await fetch('/api/authent/isLogin', {
                method: 'GET',
                headers: {
                    'Authorization': token,
                    'Content-Type': 'application/json'
                },
            });

            if (response.ok) {
                console.log("test login")
                const jdata = await response.json();
                document.getElementById('action').style.display = 'block';
                document.getElementById('login').style.display = 'none';
                document.getElementById('logout').style.display = 'block';
                userId = document.cookie.split('; ').find(row => row.startsWith('user_id=')).split('=')[1];
                linkElement.href = "/profil/"+jdata;
                document.getElementById("messages").href = "/message-liste/"+userId;
                document.getElementById("MyListeArticles").href = "/mes-article/"+userId;
            } else {
                console.error('Error:', data.message);
                document.getElementById('action').style.display = 'none';
            }
        }
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('action').style.display = 'none';
    }
});

let ws;
function connectWebSocket() {
    ws = new WebSocket('ws://127.0.0.1:3000/ws');

    ws.onopen = function() {
        console.log("WebSocket connection established");
    };

    ws.onmessage = function(event) {
        const message = JSON.parse(event.data);
        if (message.type === 'notification') {
            console.log("muid=",message.user_id,"suid=",userId)
            if(message.user_id == userId){
                console.log(`content: ${message.content}`);
                addNotification(message);
                svgNotification();
            }
        }
    };

    ws.onerror = function(error) {
        console.error("WebSocket error observed:", error);
    };

    ws.onclose = function(event) {
        console.log("WebSocket connection closed:", event);
        if (event.code === 1006) {
            console.log("Attempting to reconnect in 5 seconds...");
            setTimeout(connectWebSocket, 5000);
        }
    };
}

connectWebSocket();
function addNotification(message) {
    const notificationContainer = document.getElementById('notification-container');
    const notification = document.createElement('div');
    const numberOfChildren = notificationContainer.childElementCount;
    let id = `child-${numberOfChildren + 1}`
    console.log(id);
    notification.id = id;
    notification.className = 'bg-gray-500 hover:bg-gray-400 p-4 rounded-lg mt-5';
    notification.innerHTML = `
    <div class="last-content">
        <p class="text-black">${message.content}</p>
        <p class="text-black"><strong>Price:</strong> ${message.price}</p>
        <div class="flex justify-center space-x-2 mt-2">
            <button onclick="decideOffer('accept','${message.user_id}')" class="bg-green-300 px-3 py-1 rounded">Accept</button>
            <button onclick="decideOffer('${id}',null)" class="bg-red-300 px-3 py-1 rounded">Refuse</button>
        </div>
    </div>
    `;
    notificationContainer.appendChild(notification);
}
function svgNotification() {
    const notificationContainer = document.getElementById('notification-container');
    const imageContainer = document.getElementById('image-container');
    const svg1 = `
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0"/>
        </svg>
        
    `;
    const svg2 = `
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0M3.124 7.5A8.969 8.969 0 0 1 5.292 3m13.416 0a8.969 8.969 0 0 1 2.168 4.5"/>
        </svg>
    `;

    let svgContent;

    if (notificationContainer.children.length <= 1 ) {
        svgContent = svg1;
    } else {
        svgContent = svg2;
    }

    imageContainer.innerHTML = svgContent;
}

document.addEventListener('DOMContentLoaded', svgNotification);

function searchKey(event) {
    if (event.key === 'Enter') {
        searchBar();
    }
}

async function logout() {
    const userId = document.cookie.split('; ').find(row => row.startsWith('user_id=')).split('=')[1];
    console.log(userId);

    const response = await fetch("/api/authent/logout", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "userID": userId })
    });

    if (response.ok) {
        document.getElementById('login').style.display = 'block';
        document.getElementById('logout').style.display = 'none';
    } else {
        console.error('Logout failed:', response.statusText);
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
            if(data.articles.length === 0){
                alert("Aucun article trouvé.");
            }else{
                sessionStorage.setItem('searchResults', JSON.stringify(data.articles));
                window.location.href = "/article/search";
            }
        } else {
            alert(data.error);
        }
    } catch (error) {
        console.error("Erreur lors de la recherche de l'article:", error);
        alert("Une erreur est survenue. Veuillez réessayer plus tard.");
    }
}

function displaynotif() {
    const notification = document.getElementById('image-container');
    const notificationContainer = document.getElementById('notification-container');
    if(notificationContainer.classList.contains('opacity-0')){
        notificationContainer.classList.remove('opacity-0', 'invisible');
        notificationContainer.classList.add('opacity-100', 'visible');
    }else{
        notificationContainer.classList.add('opacity-0', 'invisible');
        notificationContainer.classList.remove('opacity-100', 'visible');
    }
}

function decideOffer(action, userId) {
    const element = document.getElementById(action);
    if (action == "accept") {
        console.log("create message");
        //ajouter la notif de messagerie ouvert
        //ajouter la creation de messagerie
        NotificationCreateMessage(userId)
        //remove notification
        if (element) {
            element.remove();
        }
        //redirection to messenger
    } else {
        if (element) {
            element.remove();
            //ajouter la notif de refus
            
        }
    }
}

async function NotificationCreateMessage(id){
    const bid = userId;
    const sid = id;
    const data = {
        "idBuyer": bid,
        "idSeller": sid,
    };
    console.log(data)
    try{
        const response = await fetch('/api/messenger/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body:  JSON.stringify(data),
        });
        
        if(response.ok){
            console.log("responce ok")

        }else{
            alert("Un problème serveur est survenu");
        }
    }catch{
        alert("Un problème serveur est survenu");
    }
}