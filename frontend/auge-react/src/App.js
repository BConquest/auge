import './css/App.css';
import React from 'react';

import Signin from './components/Signin.js';

class AddBookmark extends React.Component {
    constructor(props){
        super(props);

        this.state = {
            link: '',
        };
        
        this.addBookmark = this.addBookmark.bind(this);
    }

    async postData(url = '', data = {}, token = '') {
        const response = await fetch(url, {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });
        return response.json();
    }
    
    async addBookmark() {
        var token = 'Bearer ' + localStorage.getItem('token');
        var f = this.postData("http://174.138.52.72:1234/bookmark", this.state, token);
        console.log(f)
        localStorage.setItem('token', f.tokens);
        window.location.reload();
    }

    render() {
        return (
            <div>
               <label> Link
                    <input name="link"
                        onChange={e => this.setState({link: e.target.value})}/>
                </label>
                <button onClick={this.addBookmark}> Add Bookmark </button>
            </div>
        );
    }
}

class Bookmarks extends React.Component {
    logout() {
        localStorage.removeItem('token');
        window.location.reload();
    }

    render() {
        return (
            <div>
            <button onClick={this.logout}> Logout </button>
            <AddBookmark />
            </div>
        )
    }
}

class Auge extends React.Component {
    state = {
        token: ''
    };

    componentDidMount() {
        const token = localStorage.getItem('token') || '';
        this.setState({ token });
    }

    render() {
        if (this.state.token === "") {
            return (<Signin />);
        } else {
            return (<Bookmarks />);
        }
    }
}

export default function App() {
    return (
        <Auge />
    );
}
