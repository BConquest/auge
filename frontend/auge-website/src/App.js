import React from 'react';

import SLForm from './containers/slForm.js';

function CheckLocalStorage() {
    return localStorage.getItem('token');
}

function App() {
    const token = CheckLocalStorage();

    if (token == null){
        return <SLForm />;
    } else {
        return (<p> Hey </p>);
    }
}

export default App;
