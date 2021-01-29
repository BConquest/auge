const API_URL = "https://api.paxavis.dev:1235"

async function signup(data = {}) {
    const response = await fetch(API_URL+"/signup", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    return response.json();
}


async function login(data = {}) {
    const response = await fetch(API_URL+"/login", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    return response.json();
}

export {
    signup,
    login
}
