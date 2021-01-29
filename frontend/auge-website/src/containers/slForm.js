import React, { useState } from 'react';

import Login from '../components/login.js';
import Signup from '../components/signup.js';

export default function SLForm() {
    const [isNewUser, setNewUser] = useState(false);

    return (
        <div>
        {isNewUser
         ? <Signup />
         : <Login />
        }
        <button onClick={e => setNewUser(!isNewUser)}>
            {isNewUser ? 'Login' : 'Create New User'}
        </button>
        </div>
    )
}
