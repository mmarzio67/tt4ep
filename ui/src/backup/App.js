import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      articles: []
    };
  }
  componentDidMount() {
    let myHeaders = new Headers();

    let myInit = {
      method: "GET",
      headers: myHeaders,
      mode: "cors",
      cache: "default"
    };

    fetch("http://localhost:10000/sel", myInit)
      .then(res => res.json()) //response type
      .then(data => {
        console.log(data); //log the data
        this.setState({ articles: data });
      });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Fetch Example</h1>
        </header>
        {this.state.articles.map(article => {
          return (
            <ul key={article.descr}>
              <li>{article.project}</li>
              <li>{article.task}</li>
            </ul>
          );
        })}
      </div>
    );
  }
}

export default App;
