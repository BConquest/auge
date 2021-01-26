import './css/App.css';
import React, { useState, useEffect } from 'react';

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
        var f = this.postData("https://api.paxavis.dev/signup", this.state);
        console.log(this.state);
        console.log(f);
    }

    async login() {
        var f = await this.postData("https://api.paxavis.dev/login", this.state);
        console.log(f)
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
            headers: {
                'Authorization': token,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });
        return response.json();
    }

    async addBookmark() {
        const token = await 'Bearer ' + localStorage.getItem('token');
        console.log("Hey " + token);
        var f = await this.postData("https://api.paxavis.dev/addbookmark", this.state, token);
        console.log(f)
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

function UserBookmarks(){
    const [bookmarks, setBookmarks] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchBookmarks = async (token) => {
            const response = await fetch("https://api.paxavis.dev/bookmark", {
                "method": "GET",
                "headers": {
                    "Accept": "application/json",
                    "Authorization": token,
                    "Content-Type": "application/json"
                },
            });
            const bookmarks = await response.json();

            setBookmarks(bookmarks);
            setLoading(false);
        };

        fetchBookmarks('Bearer ' + localStorage.getItem('token'));
    }, []);

        return (
            <div>
                {loading ? "loading..." : <pre>{JSON.stringify(bookmarks)}</pre>}
            </div>
        );
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
            <UserBookmarks />
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
