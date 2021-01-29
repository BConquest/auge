import React, { useEffect, useState } from 'react';

import { login } from '../services/api.js';

function Login() {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const user = {
            "username": username,
            "password": password
        };
        const data = await login(user);
        if (data.tokens != null) {
            localStorage.setItem('token', data.tokens);
            window.location.reload(true);
        } 
    }

    return (
        <div>
            <h1> Auge </h1>
            <h3> Login </h3>
            <form onSubmit={e => handleSubmit(e)}>
                <label>
                    Name:
                    <input
                        type="text"
                        placeholder="Username"
                        onChange={e => setUsername(e.target.value)}/>
                </label>
                <br/>
                <label>
                    Password:
                    <input
                        type="password"
                        placeholder="Password"
                        onChange={e => setPassword(e.target.value)}/>
                </label>
                <br/>
                <input type="submit" value="Submit" />
            </form>
        </div>
    )
}

export default Login;
