import React, { useEffect, useState } from 'react';

import { signup } from '../services/api.js';

function newP() {
    return <p> h </p>
}

function Signup() {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();
    const [confirmP, setConfirmP] = useState();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const user = {
            "username": username,
            "password": password
        };
        const data = await signup(user);
        if (data === "success") {
            console.log("here")
            newP();
        }
    }

    return (
        <div>
            <h1> Auge </h1>
            <h3> Signup </h3>
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
                <label>
                    Confirm Password:
                    <input
                        type="password"
                        placeholder="Confirm Password"
                        onChange={e => setConfirmP(e.target.value)}/>
                </label>
                <br/>
                <input type="submit" value="Submit" />
            </form>
        </div>
    )
}

export default Signup;
