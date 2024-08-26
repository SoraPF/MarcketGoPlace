document.addEventListener('DOMContentLoaded', async () => {
    try {
        var token = null
        const jwtCookie = document.cookie.split('; ').find(row => row.startsWith('jwt-'));
        const linkElement = document.getElementById('profiLink');
        console.log(jwtCookie)
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
                const jdata = await response.json();
                document.getElementById('createArticle').style.display = 'block';
                document.getElementById('login').style.display = 'none';
                document.getElementById('logout').style.display = 'block';
                linkElement.href = "/profil/"+jdata;
            } else {
                console.error('Error:', data.message);
                document.getElementById('createArticle').style.display = 'none';
            }
        }
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('createArticle').style.display = 'none';
    }
});

const ws = new WebSocket('ws://127.0.0.1:3000/ws'); // Connecter au serveur WebSocket

ws.onmessage = function(event) {
    const message = JSON.parse(event.data);

    if (message.Type === 'notification') {
        addNotification(message);
        svgNotification();
    }
};

function addNotification(message) {
    const notificationContainer = document.getElementById('notification-container');
    const notification = document.createElement('div');
    notification.className = 'bg-white border border-gray-300 p-4 rounded-lg shadow-md mb-2';
    notification.innerHTML = `
        <p><strong>Notification:</strong> ${message.Content}</p>
        <p><strong>Price:</strong> ${message.Price}</p>
    `;
    notificationContainer.appendChild(notification);
}
function svgNotification() {
    const notificationContainer = document.getElementById('notification-container');
    const imageContainer = document.getElementById('image-container');
    const svg1 = `
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0M3.124 7.5A8.969 8.969 0 0 1 5.292 3m13.416 0a8.969 8.969 0 0 1 2.168 4.5"/>
        </svg>
    `;
    const svg2 = `
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0"/>
        </svg>
    `;

    let svgContent;

    if (notificationContainer.children.length === 0) {
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


async function logout(){
        const userId = document.cookie.split('; ').find(row => row.startsWith('user_id=')).split('=')[1];
        const response = await fetch("http://127.0.0.1:3000/api/authent/logout",{
            method: 'POST',
            body: JSON.stringify({"userID": userId })
        });
        
        if (response.ok){
            document.getElementById('login').style.display = 'block';
            document.getElementById('logout').style.display = 'none';
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
