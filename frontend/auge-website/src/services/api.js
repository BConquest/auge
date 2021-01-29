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

async function addBookmark(data = {}, token = "") {
    const response = await fetch(API_URL+"/addbookmark", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWRpZW5jZSI6InVzZXIiLCJleHAiOjE2MTIxNTE5MzgsImlkIjoiYmx1In0.dpZ9WTNBBewqO0jY2a5thla1T54x2aEfrYrobEQv-Fg"
        },
        body: JSON.stringify(data)
    });
    return response.json();
}

export {
    signup,
    login,
    addBookmark
}
