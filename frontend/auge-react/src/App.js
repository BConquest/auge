import './App.css';
import React from 'react';

class Signin extends React.Component {
    constructor(props){
        super(props);

        this.state = {
            username: '',
            password: '',
            confirmPassword: ''
        };
        
        this.signup = this.signup.bind(this);
        this.login = this.login.bind(this);
    }

    async postData(url = '', data = {}) {
        const response = await fetch(url, {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        return response.json();
    }

    signup() {
        var f = this.postData("http://174.138.52.72:1234/signup", this.state);
        console.log(this.state);
        console.log(f);
    }
    
    async login() {
        var f = await this.postData("http://174.138.52.72:1234/login", this.state);
        localStorage.setItem('token', f.tokens);
        window.location.reload();
    }

    render() {
        return (
            <div>
            <div>
                <label> Username: 
                    <input name="username"
                        onChange={e => this.setState({username: e.target.value})}/>
                </label>
                <br />
                <label> Password: 
                    <input name="password" 
                        type="password"
                        onChange={e => this.setState({password: e.target.value})}/>
                </label>
                <br />
                <label> Confirm Password: 
                    <input name="confirmPassword"
                        type="password"
                        onChange={e => this.setState({confirmPassword: e.target.value})}/>
                </label>
                <br />
                <button onClick={this.signup}> Signup </button>
            </div>
            <div>
                <label> Username: 
                    <input name="username"
                        onChange={e => this.setState({username: e.target.value})}/>
                </label>
                <br />
                <label> Password: 
                    <input name="password" 
                        type="password"
                        onChange={e => this.setState({password: e.target.value})}/>
                </label>
                <br />
                <button onClick={this.login}> Login </button>
            </div>
            </div>
        );
    }
}

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
