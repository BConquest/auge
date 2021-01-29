import React from 'react';

import SLForm from './containers/slForm.js';
import Auge from './containers/auge.js';

function CheckLocalStorage() {
    return localStorage.getItem('token');
}

function App() {
    const token = CheckLocalStorage();

    if (token == null){
        return <SLForm />;
    } else {
        return <Auge />;
    }
}

export default App;
