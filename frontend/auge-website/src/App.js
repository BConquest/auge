import React, { useEffect, useState } from 'react';

import Login from './components/login.js';
import Signup from './components/signup.js';

function CheckLocalStorage() {
    return localStorage.getItem('token');
}

function App() {
    const token = CheckLocalStorage();

    useEffect(() => {
        console.log("Hey");
    });

    if (token == null){
        return (<div> <Login /> <Signup/> </div>);
    } else {
        return (<p> Hey </p>);
    }
}

export default App;
